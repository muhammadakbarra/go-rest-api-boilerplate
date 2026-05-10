package posts

import (
	"context"
	"errors"
	"fmt"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

var ErrPostNotFound = errors.New("post not found")

type Repository struct {
	db *pgxpool.Pool
}

func NewRepository(db *pgxpool.Pool) *Repository {
	return &Repository{db: db}
}

func (r *Repository) Create(ctx context.Context, req CreatePostRequest) (*Post, error) {
	query := `
		INSERT INTO posts (title, content)
		VALUES ($1, $2)
		RETURNING id, title, content, created_at, updated_at
	`

	var post Post

	err := r.db.QueryRow(ctx, query, req.Title, req.Content).Scan(
		&post.ID,
		&post.Title,
		&post.Content,
		&post.CreatedAt,
		&post.UpdatedAt,
	)
	if err != nil {
		return nil, fmt.Errorf("create post: %w", err)
	}

	return &post, nil
}

func (r *Repository) List(ctx context.Context) ([]Post, error) {
	query := `
		SELECT id, title, content, created_at, updated_at
		FROM posts
		ORDER BY created_at DESC
	`

	rows, err := r.db.Query(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("list posts: %w", err)
	}
	defer rows.Close()

	var result []Post

	for rows.Next() {
		var post Post

		err := rows.Scan(
			&post.ID,
			&post.Title,
			&post.Content,
			&post.CreatedAt,
			&post.UpdatedAt,
		)
		if err != nil {
			return nil, fmt.Errorf("scan post: %w", err)
		}

		result = append(result, post)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("rows error: %w", err)
	}

	return result, nil
}

func (r *Repository) GetByID(ctx context.Context, id string) (*Post, error) {
	query := `
		SELECT id, title, content, created_at, updated_at
		FROM posts
		WHERE id = $1
	`

	var post Post

	err := r.db.QueryRow(ctx, query, id).Scan(
		&post.ID,
		&post.Title,
		&post.Content,
		&post.CreatedAt,
		&post.UpdatedAt,
	)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, ErrPostNotFound
		}

		return nil, fmt.Errorf("get post by id: %w", err)
	}

	return &post, nil
}

func (r *Repository) Update(ctx context.Context, id string, req UpdatePostRequest) (*Post, error) {
	query := `
		UPDATE posts
		SET title = $1,
		    content = $2,
		    updated_at = now()
		WHERE id = $3
		RETURNING id, title, content, created_at, updated_at
	`

	var post Post

	err := r.db.QueryRow(ctx, query, req.Title, req.Content, id).Scan(
		&post.ID,
		&post.Title,
		&post.Content,
		&post.CreatedAt,
		&post.UpdatedAt,
	)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, ErrPostNotFound
		}

		return nil, fmt.Errorf("update post: %w", err)
	}

	return &post, nil
}

func (r *Repository) Delete(ctx context.Context, id string) error {
	query := `
		DELETE FROM posts
		WHERE id = $1
	`

	tag, err := r.db.Exec(ctx, query, id)
	if err != nil {
		return fmt.Errorf("delete post: %w", err)
	}

	if tag.RowsAffected() == 0 {
		return ErrPostNotFound
	}

	return nil
}
