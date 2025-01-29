-- name: CreateProject :one
INSERT INTO projects (
    project_name,
    client_id
) VALUES (
    $1, $2
) RETURNING *;

-- name: UpdateProjectByID :exec
UPDATE projects SET project_name = $2,client_id = $3
WHERE id = $1;

-- name: DeleteProject :exec
DELETE FROM projects WHERE id = $1;

-- name: GetProjectByID :one
SELECT * FROM projects
WHERE id = $1;

-- name: ListProjects :many
SELECT * FROM projects
ORDER BY id;