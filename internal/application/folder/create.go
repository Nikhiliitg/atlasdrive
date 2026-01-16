package folder

import (
	"context"
	"github.com/Nikhiliitg/atlasdrive/internal/domain/folder"
	"github.com/Nikhiliitg/atlasdrive/internal/ports/repository"
)


type CreateFolderCommand struct{
	ID string
	Name string
	OwnerID string
	ParentID *string
}
type CreateFolderHandler struct {
	repo repository.FolderRepository
}

func NewCreateFolderHandler(repo repository.FolderRepository) *CreateFolderHandler {
	return &CreateFolderHandler{repo: repo}
}

func (h *CreateFolderHandler) Handle(
	ctx context.Context,
	cmd CreateFolderCommand,
) (*folder.Folder, error) {

	// Domain creation (rules enforced here)
	f, err := folder.NewFolder(
		cmd.ID,
		cmd.Name,
		cmd.OwnerID,
		cmd.ParentID,
	)
	if err != nil {
		return nil, err
	}

	// Persistence is a detail
	if err := h.repo.Save(ctx, f); err != nil {
		return nil, err
	}

	return f, nil
}