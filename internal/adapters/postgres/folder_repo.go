package postgres

import (
	"context"
	"database/sql"

	"github.com/Nikhiliitg/atlasdrive/internal/domain/folder"
)

type FolderRepo struct {
	db *sql.DB
}

func NewFolderRepo(db *sql.DB) *FolderRepo {
	return &FolderRepo{db: db}
}

func (r *FolderRepo) Save(ctx context.Context, f *folder.Folder) error {
	query := `
		INSERT INTO folders (id, name, owner_id, parent_id, created_at)
		VALUES ($1, $2, $3, $4, $5)
	`

	_, err := r.db.ExecContext(
		ctx,
		query,
		f.ID,
		f.Name,
		f.OwnerID,
		f.ParentID,
		f.CreatedAt,
	)

	return err
}

func (r *FolderRepo) GetByID(ctx context.Context, id string) (*folder.Folder, error) {
	query := `
		SELECT id, name, owner_id, parent_id, created_at
		FROM folders
		WHERE id = $1
	`

	row := r.db.QueryRowContext(ctx, query, id)

	var f folder.Folder
	err := row.Scan(
		&f.ID,
		&f.Name,
		&f.OwnerID,
		&f.ParentID,
		&f.CreatedAt,
	)

	if err != nil {
		return nil, err
	}

	return &f, nil
}
