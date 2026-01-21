package http

import (
	"encoding/json"
	"net/http"
	"strings"

	folderapp "github.com/Nikhiliitg/atlasdrive/internal/application/folder"
)

type Handler struct {
	createHandler *folderapp.CreateFolderHandler
	listHandler   *folderapp.ListFolderContentsHandler
}

func NewHandler(
	create *folderapp.CreateFolderHandler,
	list *folderapp.ListFolderContentsHandler,
) *Handler {
	return &Handler{
		createHandler: create,
		listHandler:   list,
	}
}
func (h *Handler) CreateFolder(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	var req struct {
		ID       string  `json:"id"`
		Name     string  `json:"name"`
		OwnerID  string  `json:"owner_id"`
		ParentID *string `json:"parent_id"`
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	folder, err := h.createHandler.Handle(r.Context(), folderapp.CreateFolderCommand{
		ID:       req.ID,
		Name:     req.Name,
		OwnerID:  req.OwnerID,
		ParentID: req.ParentID,
	})
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}

	json.NewEncoder(w).Encode(folder)
}

func (h *Handler) ListFolderContents(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	parts := strings.Split(r.URL.Path, "/")
	if len(parts) < 3 {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	folderID := parts[2]
	ownerID := r.URL.Query().Get("owner_id")

	folders, files, err := h.listHandler.Handle(r.Context(),
		folderapp.ListFolderContentsQuery{
			FolderID: folderID,
			OwnerID:  ownerID,
		},
	)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	resp := map[string]interface{}{
		"folders": folders,
		"files":   files,
	}

	json.NewEncoder(w).Encode(resp)
}
