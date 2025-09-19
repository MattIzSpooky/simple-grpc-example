package service

import (
	"database/sql"
	"encoding/json"
	"errors"
	"github.com/mattizspooky/simple-grpc-example/v2/internal/db"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
)

type Service struct {
	q *db.Queries
}

func NewService(q *db.Queries) *Service { return &Service{q: q} }

func (s *Service) RegisterHTTP(r *chi.Mux) {
	r.Put("/notes/{id}", s.handleUpdateNote)
	r.Get("/notes", s.handleGetAll)
	r.Delete("/notes/{id}", s.handleDelete)
}

// Update: PUT /notes/{id}
// body: {"description": "..."}
func (s *Service) handleUpdateNote(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	id, err := uuid.Parse(idStr)
	if err != nil {
		http.Error(w, "invalid id", http.StatusBadRequest)
		return
	}

	var body struct {
		Description string `json:"description"`
	}
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		http.Error(w, "bad body", http.StatusBadRequest)
		return
	}

	note, err := s.q.UpdateNoteByID(r.Context(), db.UpdateNoteByIDParams{
		ID:          id,
		Description: body.Description,
	})
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			http.Error(w, "not found", http.StatusNotFound)
			return
		}
		http.Error(w, "internal", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(note)
}

func (s *Service) handleGetAll(w http.ResponseWriter, r *http.Request) {
	notes, err := s.q.GetAllNotes(r.Context())
	if err != nil {
		http.Error(w, "internal", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(notes)
}

func (s *Service) handleDelete(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	id, err := uuid.Parse(idStr)
	if err != nil {
		http.Error(w, "invalid id", http.StatusBadRequest)
		return
	}

	if err := s.q.DeleteNoteByID(r.Context(), id); err != nil {
		http.Error(w, "internal", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}
