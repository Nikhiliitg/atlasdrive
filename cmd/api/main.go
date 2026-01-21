package main

import (
	"log"
	"net/http"

	httpadapter "github.com/Nikhiliitg/atlasdrive/internal/adapters/http"
	"github.com/Nikhiliitg/atlasdrive/internal/adapters/memory"
	folderapp "github.com/Nikhiliitg/atlasdrive/internal/application/folder"
)

func main() {
	repo := memory.NewFolderRepo()

	createHandler := folderapp.NewCreateFolderHandler(repo)
	listHandler := folderapp.NewListFolderContentsHandler(repo)

	handler := httpadapter.NewHandler(createHandler, listHandler)
	router := httpadapter.NewRouter(handler)

	log.Println("Server running on :8080")
	http.ListenAndServe(":8080", router)
}
