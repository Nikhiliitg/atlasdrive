package repository

import "context"

type FolderSummary struct {
	ID   string
	Name string
}

type FileSummary struct {
	ID   string
	Name string
}

type FolderQuery interface {
	ListChildFolders(
		ctx context.Context,
		parentID string,
		ownerID string,
	) ([]FolderSummary, error)

	ListFilesInFolder(
		ctx context.Context,
		folderID string,
		ownerID string,
	) ([]FileSummary, error)
}
