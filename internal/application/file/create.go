package file

import(
	"context"
	"github.com/Nikhiliitg/atlasdrive/internal/domain/file"
	"github.com/Nikhiliitg/atlasdrive/internal/ports/repository"
	// "github.com/Nikhiliitg/atlasdrive/internal/adapters/postgres"
)

type CreateFileCommand struct{
	ID string
	Name string
	FolderID string
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
	userID := ctx.Value("user_id").(string)

	f, err := file.NewFile(
		cmd.ID,
		cmd.Name,
		cmd.FolderID,
		userID,
	)
	if err != nil {
		return nil, err
	}

	if err := h.repo.SaveWithFolderCheck(ctx, cmd.FolderID, f); err != nil {
	return nil, err
}


	return f, nil
}

