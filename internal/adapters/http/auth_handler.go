package http

import (
	"encoding/json"
	"net/http"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"database/sql"

	"github.com/Nikhiliitg/atlasdrive/internal/auth"
)

type AuthHandler struct {
	db *sql.DB
}

func NewAuthHandler(db *sql.DB) *AuthHandler {
	return &AuthHandler{db: db}
}

func (h *AuthHandler) Register(w http.ResponseWriter, r *http.Request) {
	var req struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	json.NewDecoder(r.Body).Decode(&req)

	hash, _ := bcrypt.GenerateFromPassword([]byte(req.Password), 10)
	id := uuid.New().String()

	_, err := h.db.Exec(
		`INSERT INTO users VALUES ($1,$2,$3,now())`,
		id, req.Email, string(hash),
	)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}

	w.Write([]byte("registered"))
}

func (h *AuthHandler) Login(w http.ResponseWriter, r *http.Request) {
	var req struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	json.NewDecoder(r.Body).Decode(&req)

	var id, hash string
	err := h.db.QueryRow(
		`SELECT id, password_hash FROM users WHERE email=$1`,
		req.Email,
	).Scan(&id, &hash)

	if err != nil {
		http.Error(w, "user not found", 401)
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(hash), []byte(req.Password))
	if err != nil {
		http.Error(w, "wrong password", 401)
		return
	}

	token, _ := auth.GenerateToken(id)
	json.NewEncoder(w).Encode(map[string]string{
		"token": token,
	})
}
