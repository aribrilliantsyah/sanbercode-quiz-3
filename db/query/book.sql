-- name: CreateBook :one
INSERT INTO books (title, description, image_url, release_year, price, total_page, thickness, category_id, created_at, created_by)
VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)
RETURNING *;

-- name: GetBook :one
SELECT * FROM books
WHERE id = $1;

-- name: ListBooks :many
SELECT * FROM books
ORDER BY id
LIMIT $1 OFFSET $2;

-- name: UpdateBook :one
UPDATE books
SET title = $2, description = $3, image_url = $4, release_year = $5, price = $6, total_page = $7, thickness = $8, category_id = $9, modified_at = $10, modified_by = $11
WHERE id = $1
RETURNING *;

-- name: DeleteBook :exec
DELETE FROM books
WHERE id = $1;


-- name: GetBooksByCategory :many
SELECT * FROM books
WHERE category_id = $1
ORDER BY id
LIMIT $2 OFFSET $3;