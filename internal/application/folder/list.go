// package folder

// import (
// 	"context"
// 	"fmt"

// 	"github.com/Nikhiliitg/atlasdrive/internal/ports/repository"
// )

// type FolderContents struct {
// 	Folders []repository.FolderSummary
// 	Files   []repository.FileSummary
// }

// type ListFolderContentsQuery struct {
// 	FolderID string
// }

// type ListFolderContentsHandler struct {
// 	query repository.FolderQuery
// 	cache FolderCache
// }

// func NewListFolderContentsHandler(
// 	query repository.FolderQuery,
// 	cache FolderCache,
// ) *ListFolderContentsHandler {
// 	return &ListFolderContentsHandler{
// 		query: query,
// 		cache: cache,
// 	}
// }

// func (h *ListFolderContentsHandler) Handle(
// 	ctx context.Context,
// 	q ListFolderContentsQuery,
// ) ([]repository.FolderSummary, []repository.FileSummary, error) {

// 	userID := ctx.Value("user_id").(string)
// 	key := fmt.Sprintf("folder:%s:user:%s", q.FolderID, userID)

// 	// 1️⃣ Try cache
// 	var cached FolderContents
// 	ok, _ := h.cache.Get(ctx, key, &cached)
// 	if ok {
// 		return cached.Folders, cached.Files, nil
// 	}

// 	// 2️⃣ Query DB
// 	folders, err := h.query.ListChildFolders(ctx, q.FolderID, userID)
// 	if err != nil {
// 		return nil, nil, err
// 	}

// 	files, err := h.query.ListFilesInFolder(ctx, q.FolderID, userID)
// 	if err != nil {
// 		return nil, nil, err
// 	}

// 	// 3️⃣ Store full response
// 	_ = h.cache.Set(ctx, key, FolderContents{
// 		Folders: folders,
// 		Files:   files,
// 	})

// 	return folders, files, nil
// }
package folder

import (
	"context"
	"fmt"

	"github.com/Nikhiliitg/atlasdrive/internal/ports/repository"
)

type FolderContents struct {
	Folders []repository.FolderSummary
	Files   []repository.FileSummary
}

type ListFolderContentsQuery struct {
	FolderID string
}

type ListFolderContentsHandler struct {
	query repository.FolderQuery
	cache FolderCache
}

func NewListFolderContentsHandler(
	query repository.FolderQuery,
	cache FolderCache,
) *ListFolderContentsHandler {
	return &ListFolderContentsHandler{
		query: query,
		cache: cache,
	}
}

func (h *ListFolderContentsHandler) Handle(
	ctx context.Context,
	q ListFolderContentsQuery,
) ([]repository.FolderSummary, []repository.FileSummary, error) {

	userID := ctx.Value("user_id").(string)
	key := fmt.Sprintf("folder:%s:user:%s", q.FolderID, userID)

	// 1. Try cache
	var cached FolderContents
	ok, _ := h.cache.Get(ctx, key, &cached)
	if ok {
		return cached.Folders, cached.Files, nil
	}

	// 2. Query DB
	folders, err := h.query.ListChildFolders(ctx, q.FolderID, userID)
	if err != nil {
		return nil, nil, err
	}

	files, err := h.query.ListFilesInFolder(ctx, q.FolderID, userID)
	if err != nil {
		return nil, nil, err
	}

	// 3. Cache full response
	_ = h.cache.Set(ctx, key, FolderContents{
		Folders: folders,
		Files:   files,
	})

	return folders, files, nil
}
