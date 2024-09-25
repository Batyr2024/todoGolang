-- name: Create :one
INSERT INTO tasks (
    text
)VALUES (
         $1
        )RETURNING *;

-- name: GetAll :many
SELECT * FROM tasks
ORDER BY id ASC;

-- name: ChangeCheckedByID :one
UPDATE  tasks
    set checked=$2
WHERE id=$1
RETURNING  *;

-- name: ChangeCheckedAll :many
UPDATE tasks
    set checked=$1
RETURNING  1;

-- name: ChangeTextByID :one
UPDATE  tasks
    set text=$2,
        checked=false
WHERE id=$1
RETURNING  *;

-- name: DeleteByID :one
DELETE FROM tasks
WHERE id=$1
RETURNING  *;

-- name: DeleteAll :many
DELETE FROM tasks
WHERE checked=true
RETURNING  1;
