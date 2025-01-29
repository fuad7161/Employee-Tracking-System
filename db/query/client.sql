-- name: CreateClient :one
INSERT INTO clients (
    client_name,
    status
) VALUES (
    $1, $2
) RETURNING *;

-- name: UpdateClientStatus :exec
UPDATE clients SET status = $2
WHERE id = $1;

-- name: UpdateClient :exec
UPDATE clients SET status = $2, client_name = $3
WHERE id = $1;

-- name: DeleteClient :exec
DELETE FROM clients WHERE id = $1;

-- name: GetClientByID :one
SELECT * FROM clients
WHERE id = $1;

-- name: ListClients :many
SELECT * FROM clients
ORDER BY id;