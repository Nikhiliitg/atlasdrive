package folder

import (
	"context"
	"testing"

	"github.com/Nikhiliitg/atlasdrive/internal/adapters/memory"
)

func TestListFolderContents(t *testing.T) {
	ctx := context.WithValue(context.Background(), "user_id", "test-user")


	folderRepo := memory.NewFolderRepo()
	fileRepo := memory.NewFileRepo()

	folderQuery := memory.NewFolderQueryRepo(folderRepo, fileRepo)

	createHandler := NewCreateFolderHandler(folderRepo)
	// listHandler := NewListFolderContentsHandler(folderQuery)
	cache := newFakeCache()
	listHandler := NewListFolderContentsHandler(folderQuery, cache)


	rootID := "root"

	_, err := createHandler.Handle(ctx, CreateFolderCommand{
		ID:      rootID,
		Name:    "root",
	})
	if err != nil {
		t.Fatalf("failed to create root folder: %v", err)
	}

	_, _ = createHandler.Handle(ctx, CreateFolderCommand{
		ID:       "child-1",
		Name:     "docs",
		ParentID: &rootID,
	})

	_, _ = createHandler.Handle(ctx, CreateFolderCommand{
		ID:       "child-2",
		Name:     "images",
		ParentID: &rootID,
	})

	folders, files, err := listHandler.Handle(ctx, ListFolderContentsQuery{
		FolderID: rootID,
	})
	if err != nil {
		t.Fatalf("list failed: %v", err)
	}

	if len(folders) != 2 {
		t.Fatalf("expected 2 child folders, got %d", len(folders))
	}

	if len(files) != 0 {
		t.Fatalf("expected 0 files, got %d", len(files))
	}
}
