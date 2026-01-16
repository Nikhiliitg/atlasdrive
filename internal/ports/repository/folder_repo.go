package repository

import (
	"context"

	"github.com/Nikhiliitg/atlasdrive/internal/domain/folder"
)

type FolderRepository interface {
	Save(ctx context.Context, f *folder.Folder) error
	GetByID(ctx context.Context, id string) (*folder.Folder, error)
}
