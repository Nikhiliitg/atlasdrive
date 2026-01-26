package memory

import (
	"context"
	"sync"

	"github.com/Nikhiliitg/atlasdrive/internal/domain/file"
	"github.com/Nikhiliitg/atlasdrive/internal/ports/repository"
)

type FileRepo struct {
	mu    sync.RWMutex
	files map[string]*file.File
}

func NewFileRepo() *FileRepo {
	return &FileRepo{
		files: make(map[string]*file.File),
	}
}

func (r *FileRepo) Save(_ context.Context, f *file.File) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	r.files[f.ID] = f
	return nil
}

func (r *FileRepo) ListByFolder(
	_ context.Context,
	folderID string,
	_ string,
) ([]repository.FileSummary, error) {

	r.mu.RLock()
	defer r.mu.RUnlock()

	var result []repository.FileSummary

	for _, f := range r.files {
		if f.FolderID == folderID {
			result = append(result, repository.FileSummary{
				ID:   f.ID,
				Name: f.Name,
			})
		}
	}

	return result, nil
}
func (r *FileRepo) SaveWithFolderCheck(
    ctx context.Context,
    folderID string,
    f *file.File,
) error {
    // No transaction in memory, just call Save
    return r.Save(ctx, f)
}
