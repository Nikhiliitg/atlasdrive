package http

import "net/http"

func NewRouter(handler *Handler) http.Handler {
	mux := http.NewServeMux()

	mux.HandleFunc("/folders", handler.CreateFolder)
	mux.HandleFunc("/folders/", handler.ListFolderContents)
	mux.HandleFunc("/files", handler.CreateFile)
	return mux
}
