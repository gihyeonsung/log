package infrastructure

import (
	"database/sql"

	"github.com/gihyeonsung/log/internal/domain"
)

type SqlitePostRepository struct {
	db *sql.DB
}

var _ domain.PostRepository = (*SqlitePostRepository)(nil)

func NewSqlitePostRepository(db *sql.DB) *SqlitePostRepository {
	return &SqlitePostRepository{db: db}
}

func (r *SqlitePostRepository) Get(id domain.PostID) (*domain.Post, error) {
	row := r.db.QueryRow(`
		SELECT id, title, slug, created_at, updated_at, revision, content
		FROM post
		WHERE id = ? AND _deleted_at IS NULL
	`, id.String())

	var post domain.Post
	var idStr, slugStr string
	var err error
	if err := row.Scan(
		&idStr, &post.Title, &slugStr, &post.CreatedAt, &post.UpdatedAt, &post.Revision, &post.Content,
	); err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	post.ID, err = domain.PostIDFromString(idStr)
	if err != nil {
		return nil, err
	}
	post.Slug = domain.PostSlug(slugStr)
	return &post, nil
}

func (r *SqlitePostRepository) Find() ([]*domain.Post, error) {
	rows, err := r.db.Query(`
		SELECT id, title, slug, created_at, updated_at, revision, content
		FROM post
		WHERE _deleted_at IS NULL
		ORDER BY updated_at DESC
	`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var posts []*domain.Post
	for rows.Next() {
		var post domain.Post
		var idStr, slugStr string
		if err := rows.Scan(
			&idStr, &post.Title, &slugStr, &post.CreatedAt, &post.UpdatedAt, &post.Revision, &post.Content,
		); err != nil {
			return nil, err
		}
		post.ID, err = domain.PostIDFromString(idStr)
		if err != nil {
			return nil, err
		}
		post.Slug = domain.PostSlug(slugStr)
		posts = append(posts, &post)
	}
	return posts, nil
}

func (r *SqlitePostRepository) Save(post *domain.Post) error {
	_, err := r.db.Exec(`
		INSERT INTO post (
			id, title, slug, created_at, updated_at, revision, content,
			_created_at, _updated_at, _deleted_at
		) VALUES (?, ?, ?, ?, ?, ?, ?,
			CURRENT_TIMESTAMP, CURRENT_TIMESTAMP, NULL
		)
		ON CONFLICT(id) DO UPDATE SET
			title=excluded.title,
			slug=excluded.slug,
			created_at=excluded.created_at,
			updated_at=excluded.updated_at,
			revision=excluded.revision,
			content=excluded.content,
			_updated_at=CURRENT_TIMESTAMP,
			_deleted_at=NULL
	`,
		post.ID.String(), post.Title, string(post.Slug), post.CreatedAt, post.UpdatedAt, post.Revision, post.Content,
	)
	return err
}

func (r *SqlitePostRepository) Delete(id domain.PostID) error {
	_, err := r.db.Exec(`
		UPDATE post SET _deleted_at = CURRENT_TIMESTAMP WHERE id = ?
	`, id.String())
	return err
}
