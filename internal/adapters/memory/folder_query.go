package memory

import (
	"context"
	"github.com/Nikhiliitg/atlasdrive/internal/ports/repository"
)

func (r *FolderRepository) ListChildFolders(
	ctx context.Context,
	parentID string,
	ownerID string,
) ([]repository.FolderSummary, error) {
	r.mu.Lock()
	defer r.mu.Unlock()
	var result []repository.FolderSummary

	for _, f := range r.folders {
		if f.ParentID != nil && *f.ParentID == parentID {
			result = append(result, repository.FolderSummary{
				ID:   f.ID,
				Name: f.Name,
			})
		}
	}
	return result, nil
}

func (r *FolderRepository) ListFilesInFolder(
	_ context.Context,
	_ string,
	_ string,
) ([]repository.FileSummary, error) {

	// Files are not implemented yet.
	// Returning empty list is correct and honest.
	return []repository.FileSummary{}, nil
}
