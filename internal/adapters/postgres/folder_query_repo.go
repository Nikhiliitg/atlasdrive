package postgres

import (
	"context"
	"database/sql"

	"github.com/Nikhiliitg/atlasdrive/internal/ports/repository"
)

type FolderQueryRepo struct {
	db *sql.DB
}

func NewFolderQueryRepo(db *sql.DB) *FolderQueryRepo {
	return &FolderQueryRepo{db: db}
}

func (r *FolderQueryRepo) ListChildFolders(
	ctx context.Context,
	parentID string,
	ownerID string,
) ([]repository.FolderSummary, error) {

	query := `
		SELECT id, name
		FROM folders
		WHERE parent_id = $1 AND owner_id = $2
	`

	rows, err := r.db.QueryContext(ctx, query, parentID, ownerID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var result []repository.FolderSummary
	for rows.Next() {
		var f repository.FolderSummary
		if err := rows.Scan(&f.ID, &f.Name); err != nil {
			return nil, err
		}
		result = append(result, f)
	}

	return result, nil
}

func (r *FolderQueryRepo) ListFilesInFolder(
	ctx context.Context,
	folderID string,
	ownerID string,
) ([]repository.FileSummary, error) {
	// Files still in memory for now
	return []repository.FileSummary{}, nil
}
