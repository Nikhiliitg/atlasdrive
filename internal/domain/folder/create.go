package folder

import (
	"context"
	"atlasdrive/internal/domain/folder"
	"atlasdrive/internal/ports/repository"
)

type CreateFolderCommand struct{
	ID string
	Name string
	OwnerID string
	ParentID *string
}

