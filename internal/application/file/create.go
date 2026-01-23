package file

import(
	"context"
	"github.com/Nikhiliitg/atlasdrive/internal/domain/file"
	"github.com/Nikhiliitg/atlasdrive/internal/ports/repository"
)

type CreateFileCommand struct{
	ID string
	Name string
	FolderID string
	OwnerID string
}

type CreateFileHandler struct{
	repo repository.FileRepository
}

func NewCreateFileHandler(repo repository.FileRepository) *CreateFileHandler{
	return &CreateFileHandler{repo: repo}
}

func (h *CreateFileHandler) Handle(
	ctx context.Context,
	cmd CreateFileCommand,
) (*file.File, error) {

	f, err := file.NewFile(
		cmd.ID,
		cmd.Name,
		cmd.FolderID,
		cmd.OwnerID,
	)
	if err != nil {
		return nil, err
	}

	if err := h.repo.Save(ctx, f); err != nil {
		return nil, err
	}

	return f, nil
}

