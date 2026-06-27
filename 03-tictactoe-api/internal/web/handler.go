package web

import (
	"encoding/json"
	"net/http"
	"strings"
	"tictactoe/internal/datasource"
	"tictactoe/internal/domain"
)

type Handler struct {
	repo    datasource.Repository
	service domain.Service
}

func NewHandler(repo datasource.Repository, service domain.Service) *Handler {
	return &Handler{repo: repo, service: service}
}

func (h *Handler) PlayTarget(w http.ResponseWriter, r *http.Request) {
	parts := strings.Split(strings.Trim(r.URL.Path, "/"), "/")
	if len(parts) != 2 || parts[0] != "game" {
		http.Error(w, "not found", http.StatusNotFound)
		return
	}
	uuid := parts[1]

	var reqGame Game
	if err := json.NewDecoder(r.Body).Decode(&reqGame); err != nil {
		http.Error(w, "invalid json", http.StatusBadRequest)
		return
	}

	if reqGame.UUID != "" && reqGame.UUID != uuid {
		http.Error(w, "uuid mismatch", http.StatusBadRequest)
		return
	}
	reqGame.UUID = uuid

	newDGame := toDomainGame(reqGame)

	oldDGame, err := h.repo.Get(uuid)
	if err != nil {
		oldDGame = &domain.Game{UUID: uuid}
	}

	if err := h.service.ValidateField(oldDGame, newDGame); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := h.service.NextMove(newDGame); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := h.repo.Save(newDGame); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(fromDomainGame(newDGame))
}

func NewServeMux(h *Handler) *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/game/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
			return
		}
		h.PlayTarget(w, r)
	})
	return mux
}
