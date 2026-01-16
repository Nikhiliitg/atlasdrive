package folder

import (
	"context"
	"testing"

	"github.com/Nikhiliitg/atlasdrive/internal/adapters/memory"
	"github.com/Nikhiliitg/atlasdrive/internal/domain/folder"
)

func TestCreateFolder(t *testing.T) {
	ctx := context.Background()
	// In-memory repo (no DB, no excuses)
	repo := memory.NewFolderRepo()
	handler := NewCreateFolderHandler(repo)

	// 1. Create root folder
	rootID := "root"
	rootCmd := CreateFolderCommand{
		ID:      rootID,
		Name:    "root",
		OwnerID: "user-1",
	}

	root, err := handler.Handle(ctx, rootCmd)
	if err != nil {
		t.Fatalf("expected root folder creation to succeed, got error: %v", err)
	}

	if root.ParentID != nil {
		t.Fatalf("expected root folder to have no parent")
	}

	// 2. Create child folder
	childID := "child"
	childCmd := CreateFolderCommand{
		ID:       childID,
		Name:     "documents",
		OwnerID:  "user-1",
		ParentID: &rootID,
	}

	child, err := handler.Handle(ctx, childCmd)
	if err != nil {
		t.Fatalf("expected child folder creation to succeed, got error: %v", err)
	}

	if child.ParentID == nil || *child.ParentID != rootID {
		t.Fatalf("expected child parent to be %s", rootID)
	}

	// 3. Try to create a cycle (illegal)
	cycleCmd := CreateFolderCommand{
		ID:       rootID,
		Name:     "evil-root",
		OwnerID:  "user-1",
		ParentID: &rootID,
	}

	_, err = handler.Handle(ctx, cycleCmd)
	if err == nil {
		t.Fatalf("expected cycle creation to fail, but it succeeded")
	}

	if err != folder.ErrCycleDetected {
		t.Fatalf("expected ErrCycleDetected, got %v", err)
	}
}
