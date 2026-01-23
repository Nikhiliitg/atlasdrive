package repository

import (
	"context"
	"github.com/Nikhiliitg/atlasdrive/internal/domain/file"
)

type FileRepository interface {
	Save(ctx context.Context, f *file.File) error
	ListByFolder(
		ctx context.Context,
		folderID string,
		ownerID string,
	)([]FileSummary, error)
}