package posts

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/go-chi/chi/v5"
)

type Handler struct {
	repo *Repository
}

func NewHandler(repo *Repository) *Handler {
	return &Handler{repo: repo}
}

func (h *Handler) RegisterRoutes(r chi.Router) {
	r.Route("/posts", func(r chi.Router) {
		r.Post("/", h.CreatePost)
		r.Get("/", h.ListPosts)
		r.Get("/{id}", h.GetPost)
		r.Put("/{id}", h.UpdatePost)
		r.Delete("/{id}", h.DeletePost)
	})
}

// CreatePost godoc
// @Summary Create a new post
// @Description Create a new post with the provided title and content
// @Tags posts
// @Accept json
// @Produce json
// @Param post body CreatePostRequest true "Post to create"
// @Success 201 {object} Post
// @Failure 400 {object} ErrorResponse "Invalid request body or missing required fields"
// @Failure 500 {object} ErrorResponse "Internal server error"
// @Router /posts [post]
func (h *Handler) CreatePost(w http.ResponseWriter, r *http.Request) {
	var req CreatePostRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		writeError(w, http.StatusBadRequest, "invalid JSON body")
		return
	}

	if req.Title == "" || req.Content == "" {
		writeError(w, http.StatusBadRequest, "title and content are required")
		return
	}

	post, err := h.repo.Create(r.Context(), req)
	if err != nil {
		writeError(w, http.StatusInternalServerError, "failed to create post")
		return
	}

	writeJSON(w, http.StatusCreated, post)
}

// ListPosts godoc
// @Summary List all posts
// @Description Get a list of all posts
// @Tags posts
// @Produce json
// @Success 200 {array} Post
// @Failure 500 {object} ErrorResponse "Internal server error"
// @Router /posts [get]
func (h *Handler) ListPosts(w http.ResponseWriter, r *http.Request) {
	result, err := h.repo.List(r.Context())
	if err != nil {
		writeError(w, http.StatusInternalServerError, "failed to list posts")
		return
	}

	writeJSON(w, http.StatusOK, result)
}

// GetPost godoc
// @Summary Get a post by ID
// @Description Get a single post by its ID
// @Tags posts
// @Produce json
// @Param id path string true "Post ID"
// @Success 200 {object} Post
// @Failure 404 {object} ErrorResponse "Post not found"
// @Failure 500 {object} ErrorResponse "Internal server error"
// @Router /posts/{id} [get]
func (h *Handler) GetPost(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	post, err := h.repo.GetByID(r.Context(), id)
	if err != nil {
		if errors.Is(err, ErrPostNotFound) {
			writeError(w, http.StatusNotFound, "post not found")
			return
		}

		writeError(w, http.StatusInternalServerError, "failed to get post")
		return
	}

	writeJSON(w, http.StatusOK, post)
}

// UpdatePost godoc
// @Summary Update a post
// @Description Update an existing post by its ID
// @Tags posts
// @Accept json
// @Produce json
// @Param id path string true "Post ID"
// @Param post body UpdatePostRequest true "Post update data"
// @Success 200 {object} Post
// @Failure 400 {object} ErrorResponse "Invalid request body or missing required fields"
// @Failure 404 {object} ErrorResponse "Post not found"
// @Failure 500 {object} ErrorResponse "Internal server error"
// @Router /posts/{id} [put]
func (h *Handler) UpdatePost(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	var req UpdatePostRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		writeError(w, http.StatusBadRequest, "invalid JSON body")
		return
	}

	if req.Title == "" || req.Content == "" {
		writeError(w, http.StatusBadRequest, "title and content are required")
		return
	}

	post, err := h.repo.Update(r.Context(), id, req)
	if err != nil {
		if errors.Is(err, ErrPostNotFound) {
			writeError(w, http.StatusNotFound, "post not found")
			return
		}

		writeError(w, http.StatusInternalServerError, "failed to update post")
		return
	}

	writeJSON(w, http.StatusOK, post)
}

// DeletePost godoc
// @Summary Delete a post
// @Description Delete an existing post by its ID
// @Tags posts
// @Param id path string true "Post ID"
// @Success 204 "No Content"
// @Failure 404 {object} ErrorResponse "Post not found"
// @Failure 500 {object} ErrorResponse "Internal server error"
// @Router /posts/{id} [delete]
func (h *Handler) DeletePost(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	err := h.repo.Delete(r.Context(), id)
	if err != nil {
		if errors.Is(err, ErrPostNotFound) {
			writeError(w, http.StatusNotFound, "post not found")
			return
		}

		writeError(w, http.StatusInternalServerError, "failed to delete post")
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func writeJSON(w http.ResponseWriter, status int, data any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	if data != nil {
		_ = json.NewEncoder(w).Encode(data)
	}
}

func writeError(w http.ResponseWriter, status int, message string) {
	writeJSON(w, status, ErrorResponse{
		Error: message,
	})
}
