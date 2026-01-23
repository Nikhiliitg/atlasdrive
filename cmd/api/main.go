package main

import (
	"log"
	"net/http"

	httpadapter "github.com/Nikhiliitg/atlasdrive/internal/adapters/http"
	"github.com/Nikhiliitg/atlasdrive/internal/adapters/memory"
	fileapp "github.com/Nikhiliitg/atlasdrive/internal/application/file"
	folderapp "github.com/Nikhiliitg/atlasdrive/internal/application/folder"
)

func main() {
	// Repositories
	folderRepo := memory.NewFolderRepo()
	fileRepo := memory.NewFileRepo()

	// Composed query (THIS is the key)
	folderQuery := memory.NewFolderQueryRepo(folderRepo, fileRepo)

	// Application handlers
	createFolderHandler := folderapp.NewCreateFolderHandler(folderRepo)
	listFolderHandler := folderapp.NewListFolderContentsHandler(folderQuery)
	createFileHandler := fileapp.NewCreateFileHandler(fileRepo)

	// HTTP handler
	handler := httpadapter.NewHandler(
		createFolderHandler,
		listFolderHandler,
		createFileHandler,
	)

	router := httpadapter.NewRouter(handler)

	log.Println("Server running on :8080")
	http.ListenAndServe(":8080", router)
}
