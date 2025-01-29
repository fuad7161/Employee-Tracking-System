-- name: CreateUser :one
INSERT INTO users (
    firstname,
    lastname,
    email,
    password,
    user_role_id,
    sbu_id
) VALUES (
    $1, $2, $3, $4, $5, $6
) RETURNING *;

-- name: UpdateUserPassword :exec
UPDATE users SET password = $2
WHERE id = $1;

-- name: UpdateUserRole :exec
UPDATE users SET user_role_id = $2
WHERE id = $1;

-- name: UpdateUserSBU :exec
UPDATE users SET sbu_id = $2
WHERE id = $1;

-- name: DeleteUser :exec
DELETE FROM users WHERE id = $1;

-- name: GetUserByID :one
SELECT * FROM users
WHERE id = $1;

-- name: GetUserByEmail :one
SELECT * FROM users
WHERE email = $1;

-- name: ListAdmins :many
SELECT * FROM users
WHERE user_role_id = $1;

-- name: ListUsers :many
SELECT * FROM users
ORDER BY id;

-- name: UpdateUserInformation :exec
UPDATE users
SET    firstname = $1,    lastname = $2,    email = $3,    user_role_id = $4,    sbu_id = $5
WHERE id = $6;