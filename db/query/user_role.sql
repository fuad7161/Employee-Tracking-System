-- name: CreateUserRole :one
INSERT INTO user_roles (
    role_name
) VALUES (
    $1
) RETURNING *;

-- name: UpdateUserByRole :exec
UPDATE user_roles SET role_name = $2
WHERE id = $1;

-- name: DeleteRole :exec
DELETE FROM user_roles WHERE id = $1;

-- name: GetUserRoleByID :one
SELECT * FROM user_roles
WHERE id = $1;

-- name: ListUserRole :many
SELECT * FROM user_roles
ORDER BY id;