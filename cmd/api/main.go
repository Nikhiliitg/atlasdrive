package main

import (
	"log"
	"net/http"
	"database/sql"

	httpadapter "github.com/Nikhiliitg/atlasdrive/internal/adapters/http"
	"github.com/Nikhiliitg/atlasdrive/internal/adapters/postgres"
	fileapp "github.com/Nikhiliitg/atlasdrive/internal/application/file"
	folderapp "github.com/Nikhiliitg/atlasdrive/internal/application/folder"
	_ "github.com/lib/pq"
)

func main() {
	db, err := sql.Open("postgres", "postgres://localhost/atlasdrive?sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}

	// Write side
	folderRepo := postgres.NewFolderRepo(db)
	fileRepo := postgres.NewFileRepo(db)

	// Read side (Postgres)
	folderQuery := postgres.NewFolderQueryRepo(db)

	// Application handlers
	createFolderHandler := folderapp.NewCreateFolderHandler(folderRepo)
	listFolderHandler := folderapp.NewListFolderContentsHandler(folderQuery)
	createFileHandler := fileapp.NewCreateFileHandler(fileRepo)

	// HTTP
	handler := httpadapter.NewHandler(
		createFolderHandler,
		listFolderHandler,
		createFileHandler,
	)

	authHandler := httpadapter.NewAuthHandler(db)
	router := httpadapter.NewRouter(handler, authHandler)


	log.Println("Server running on :8080")
	http.ListenAndServe(":8080", router)
}
