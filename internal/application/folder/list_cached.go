package folder

import (
	"context"
	"fmt"

	"github.com/Nikhiliitg/atlasdrive/internal/ports/repository"
)


type CachedFolderQuery struct {
	dbQuery repository.FolderQuery
	cache   FolderCache
}


type FolderCache interface {
	Get(ctx context.Context, key string, dest interface{}) (bool, error)
	Set(ctx context.Context, key string, value interface{}) error
}

func NewCachedFolderQuery(
	dbQuery repository.FolderQuery,
	cache FolderCache,
) *CachedFolderQuery {
	return &CachedFolderQuery{
		dbQuery: dbQuery,
		cache:   cache,
	}
}

func (q *CachedFolderQuery) ListChildFolders(
	ctx context.Context,
	parentID string,
	ownerID string,
) ([]repository.FolderSummary, error) {
	return q.dbQuery.ListChildFolders(ctx, parentID, ownerID)
}

// func (q *CachedFolderQuery) ListFilesInFolder(
// 	ctx context.Context,
// 	folderID string,
// 	ownerID string,
// ) ([]repository.FileSummary, error) {

// 	key := fmt.Sprintf("folder:%s:user:%s", folderID, ownerID)

// 	var result []repository.FileSummary
// 	ok, _ := q.cache.Get(ctx, key, &result)
// 	if ok {
// 		return result, nil
// 	}

// 	result, err := q.dbQuery.ListFilesInFolder(ctx, folderID, ownerID)
// 	if err != nil {
// 		return nil, err
// 	}

// 	_ = q.cache.Set(ctx, key, result)
// 	return result, nil
// }
func (q *CachedFolderQuery) ListFolderContents(
	ctx context.Context,
	folderID string,
	ownerID string,
) ([]repository.FolderSummary, []repository.FileSummary, error) {

	key := fmt.Sprintf("folder:%s:user:%s", folderID, ownerID)

	var cached FolderContents
	ok, _ := q.cache.Get(ctx, key, &cached)
	if ok {
		return cached.Folders, cached.Files, nil
	}

	folders, err := q.dbQuery.ListChildFolders(ctx, folderID, ownerID)
	if err != nil {
		return nil, nil, err
	}

	files, err := q.dbQuery.ListFilesInFolder(ctx, folderID, ownerID)
	if err != nil {
		return nil, nil, err
	}

	_ = q.cache.Set(ctx, key, FolderContents{
		Folders: folders,
		Files:   files,
	})

	return folders, files, nil
}
