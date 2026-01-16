package memory

import(
	"context"
	"sync"

	"github.com/Nikhiliitg/atlasdrive/internal/domain/folder"
)

type FolderRepository struct {
	folders map[string]*folder.Folder
	mu sync.Mutex
}
func NewFolderRepo() *FolderRepository {
	return &FolderRepository{
		folders: make(map[string]*folder.Folder),
	}
}

func (r *FolderRepository) Save(_ context.Context, f *folder.Folder) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.folders[f.ID] = f
	return nil
}

func (r *FolderRepository) GetByID(_ context.Context, id string) (*folder.Folder, error) {
	r.mu.Lock()
	defer r.mu.Unlock()
	return r.folders[id], nil
}

