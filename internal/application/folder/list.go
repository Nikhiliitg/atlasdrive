package folder

import (
	"context"
	"github.com/Nikhiliitg/atlasdrive/internal/ports/repository"
)

type ListFolderContentsQuery struct {
	FolderID string
	OwnerID  string
}
type ListFolderContentsHandler struct {
	query repository.FolderQuery
}

func NewListFolderContentsHandler(
	query repository.FolderQuery,
) *ListFolderContentsHandler {
	return &ListFolderContentsHandler{query: query}
}

func (h *ListFolderContentsHandler) Handle(
	ctx context.Context,
	q ListFolderContentsQuery,
) ([]repository.FolderSummary, []repository.FileSummary, error) {

	folders, err := h.query.ListChildFolders(ctx, q.FolderID, q.OwnerID)
	if err != nil {
		return nil, nil, err
	}

	files, err := h.query.ListFilesInFolder(ctx, q.FolderID, q.OwnerID)
	if err != nil {
		return nil, nil, err
	}

	return folders, files, nil
}