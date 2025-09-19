-- name: CreateNote :one
INSERT INTO notes (description)
VALUES ($1)
    RETURNING *;

-- name: UpdateNoteByID :one
UPDATE notes
SET description = $2, updated = now()
WHERE id = $1
    RETURNING id, description, created, updated;


-- name: GetAllNotes :many
SELECT id, description, created, updated
FROM notes
ORDER BY created DESC;


-- name: DeleteNoteByID :exec
DELETE FROM notes WHERE id = $1;


-- name: GetNoteByID :one
SELECT id, description, created, updated FROM notes WHERE id = $1;
