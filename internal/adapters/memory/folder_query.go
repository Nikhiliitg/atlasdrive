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

func (r *FileRepo) ListFilesInFolder(
	ctx context.Context,
	folderID string,
	ownerID string,
) ([]repository.FileSummary, error) {
	return r.ListByFolder(ctx, folderID, ownerID)
}
type FolderQueryRepo struct {
	folderRepo *FolderRepository
	fileRepo   *FileRepo
}

func NewFolderQueryRepo(
	folderRepo *FolderRepository,
	fileRepo *FileRepo,
) *FolderQueryRepo {
	return &FolderQueryRepo{
		folderRepo: folderRepo,
		fileRepo:   fileRepo,
	}
}

func (r *FolderQueryRepo) ListChildFolders(
	ctx context.Context,
	parentID string,
	ownerID string,
) ([]repository.FolderSummary, error) {
	return r.folderRepo.ListChildFolders(ctx, parentID, ownerID)
}

func (r *FolderQueryRepo) ListFilesInFolder(
	ctx context.Context,
	folderID string,
	ownerID string,
) ([]repository.FileSummary, error) {
	return r.fileRepo.ListByFolder(ctx, folderID, ownerID)
}