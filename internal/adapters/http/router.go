package http
import (
	"net/http"
	"github.com/Nikhiliitg/atlasdrive/internal/adapters/http/middleware"
)

func NewRouter(handler *Handler, authHandler *AuthHandler) http.Handler {
	mux := http.NewServeMux()

	// Public auth routes
	mux.HandleFunc("/auth/register", authHandler.Register)
	mux.HandleFunc("/auth/login", authHandler.Login)

	// Protected routes
	mux.Handle("/folders", middleware.AuthMiddleware((http.HandlerFunc(handler.CreateFolder))))
	mux.Handle("/folders/", middleware.AuthMiddleware(http.HandlerFunc(handler.ListFolderContents)))
	mux.Handle("/files", middleware.AuthMiddleware(http.HandlerFunc(handler.CreateFile)))

	return mux
}
