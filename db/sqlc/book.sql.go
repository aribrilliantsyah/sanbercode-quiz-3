// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: book.sql

package db

import (
	"context"
	"database/sql"
)

const createBook = `-- name: CreateBook :one
INSERT INTO books (title, description, image_url, release_year, price, total_page, thickness, category_id, created_at, created_by)
VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)
RETURNING id, title, description, image_url, release_year, price, total_page, thickness, category_id, created_at, created_by, modified_at, modified_by
`

type CreateBookParams struct {
	Title       string         `json:"title"`
	Description sql.NullString `json:"description"`
	ImageUrl    sql.NullString `json:"image_url"`
	ReleaseYear sql.NullInt32  `json:"release_year"`
	Price       sql.NullInt32  `json:"price"`
	TotalPage   sql.NullInt32  `json:"total_page"`
	Thickness   sql.NullString `json:"thickness"`
	CategoryID  int64          `json:"category_id"`
	CreatedAt   sql.NullTime   `json:"created_at"`
	CreatedBy   sql.NullString `json:"created_by"`
}

func (q *Queries) CreateBook(ctx context.Context, arg CreateBookParams) (Book, error) {
	row := q.queryRow(ctx, q.createBookStmt, createBook,
		arg.Title,
		arg.Description,
		arg.ImageUrl,
		arg.ReleaseYear,
		arg.Price,
		arg.TotalPage,
		arg.Thickness,
		arg.CategoryID,
		arg.CreatedAt,
		arg.CreatedBy,
	)
	var i Book
	err := row.Scan(
		&i.ID,
		&i.Title,
		&i.Description,
		&i.ImageUrl,
		&i.ReleaseYear,
		&i.Price,
		&i.TotalPage,
		&i.Thickness,
		&i.CategoryID,
		&i.CreatedAt,
		&i.CreatedBy,
		&i.ModifiedAt,
		&i.ModifiedBy,
	)
	return i, err
}

const deleteBook = `-- name: DeleteBook :exec
DELETE FROM books
WHERE id = $1
`

func (q *Queries) DeleteBook(ctx context.Context, id int64) error {
	_, err := q.exec(ctx, q.deleteBookStmt, deleteBook, id)
	return err
}

const getBook = `-- name: GetBook :one
SELECT id, title, description, image_url, release_year, price, total_page, thickness, category_id, created_at, created_by, modified_at, modified_by FROM books
WHERE id = $1
`

func (q *Queries) GetBook(ctx context.Context, id int64) (Book, error) {
	row := q.queryRow(ctx, q.getBookStmt, getBook, id)
	var i Book
	err := row.Scan(
		&i.ID,
		&i.Title,
		&i.Description,
		&i.ImageUrl,
		&i.ReleaseYear,
		&i.Price,
		&i.TotalPage,
		&i.Thickness,
		&i.CategoryID,
		&i.CreatedAt,
		&i.CreatedBy,
		&i.ModifiedAt,
		&i.ModifiedBy,
	)
	return i, err
}

const getBooksByCategory = `-- name: GetBooksByCategory :many
SELECT id, title, description, image_url, release_year, price, total_page, thickness, category_id, created_at, created_by, modified_at, modified_by FROM books
WHERE category_id = $1
ORDER BY id
LIMIT $2 OFFSET $3
`

type GetBooksByCategoryParams struct {
	CategoryID int64 `json:"category_id"`
	Limit      int32 `json:"limit"`
	Offset     int32 `json:"offset"`
}

func (q *Queries) GetBooksByCategory(ctx context.Context, arg GetBooksByCategoryParams) ([]Book, error) {
	rows, err := q.query(ctx, q.getBooksByCategoryStmt, getBooksByCategory, arg.CategoryID, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Book{}
	for rows.Next() {
		var i Book
		if err := rows.Scan(
			&i.ID,
			&i.Title,
			&i.Description,
			&i.ImageUrl,
			&i.ReleaseYear,
			&i.Price,
			&i.TotalPage,
			&i.Thickness,
			&i.CategoryID,
			&i.CreatedAt,
			&i.CreatedBy,
			&i.ModifiedAt,
			&i.ModifiedBy,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const listBooks = `-- name: ListBooks :many
SELECT id, title, description, image_url, release_year, price, total_page, thickness, category_id, created_at, created_by, modified_at, modified_by FROM books
ORDER BY id
LIMIT $1 OFFSET $2
`

type ListBooksParams struct {
	Limit  int32 `json:"limit"`
	Offset int32 `json:"offset"`
}

func (q *Queries) ListBooks(ctx context.Context, arg ListBooksParams) ([]Book, error) {
	rows, err := q.query(ctx, q.listBooksStmt, listBooks, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Book{}
	for rows.Next() {
		var i Book
		if err := rows.Scan(
			&i.ID,
			&i.Title,
			&i.Description,
			&i.ImageUrl,
			&i.ReleaseYear,
			&i.Price,
			&i.TotalPage,
			&i.Thickness,
			&i.CategoryID,
			&i.CreatedAt,
			&i.CreatedBy,
			&i.ModifiedAt,
			&i.ModifiedBy,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const updateBook = `-- name: UpdateBook :one
UPDATE books
SET title = $2, description = $3, image_url = $4, release_year = $5, price = $6, total_page = $7, thickness = $8, category_id = $9, modified_at = $10, modified_by = $11
WHERE id = $1
RETURNING id, title, description, image_url, release_year, price, total_page, thickness, category_id, created_at, created_by, modified_at, modified_by
`

type UpdateBookParams struct {
	ID          int64          `json:"id"`
	Title       string         `json:"title"`
	Description sql.NullString `json:"description"`
	ImageUrl    sql.NullString `json:"image_url"`
	ReleaseYear sql.NullInt32  `json:"release_year"`
	Price       sql.NullInt32  `json:"price"`
	TotalPage   sql.NullInt32  `json:"total_page"`
	Thickness   sql.NullString `json:"thickness"`
	CategoryID  int64          `json:"category_id"`
	ModifiedAt  sql.NullTime   `json:"modified_at"`
	ModifiedBy  sql.NullString `json:"modified_by"`
}

func (q *Queries) UpdateBook(ctx context.Context, arg UpdateBookParams) (Book, error) {
	row := q.queryRow(ctx, q.updateBookStmt, updateBook,
		arg.ID,
		arg.Title,
		arg.Description,
		arg.ImageUrl,
		arg.ReleaseYear,
		arg.Price,
		arg.TotalPage,
		arg.Thickness,
		arg.CategoryID,
		arg.ModifiedAt,
		arg.ModifiedBy,
	)
	var i Book
	err := row.Scan(
		&i.ID,
		&i.Title,
		&i.Description,
		&i.ImageUrl,
		&i.ReleaseYear,
		&i.Price,
		&i.TotalPage,
		&i.Thickness,
		&i.CategoryID,
		&i.CreatedAt,
		&i.CreatedBy,
		&i.ModifiedAt,
		&i.ModifiedBy,
	)
	return i, err
}
