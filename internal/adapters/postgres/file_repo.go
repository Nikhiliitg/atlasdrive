package postgres

import (
	"context"
	"database/sql"
	"github.com/Nikhiliitg/atlasdrive/internal/domain/file"
	"github.com/Nikhiliitg/atlasdrive/internal/ports/repository"
)


type FileRepo struct {
	db *sql.DB
}

func NewFileRepo(db *sql.DB) *FileRepo {
	return &FileRepo{db: db}
}

func (r *FileRepo) Save(ctx context.Context, f *file.File) error {
	query := `
	INSERT INTO files (id, name, folder_id, owner_id, created_at)
	VALUES ($1, $2, $3, $4, $5)
	`
	_, err := r.db.ExecContext(ctx, query, f.ID, f.Name, f.FolderID, f.OwnerID, f.CreatedAt)
	return err
}

func (r *FileRepo) ListByFolder(
	ctx context.Context,
	folderID string,
	ownerID string,
) ([]repository.FileSummary, error) {
	query := `
		SELECT id, name
		FROM files
		WHERE folder_id = $1 AND owner_id = $2
		`
	rows, err := r.db.QueryContext(ctx, query, folderID, ownerID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var result []repository.FileSummary
	for rows.Next() {
		var f repository.FileSummary
		if err := rows.Scan(&f.ID, &f.Name); err != nil {
			return nil, err
		}
		result = append(result, f)
	}
	return result, nil
}