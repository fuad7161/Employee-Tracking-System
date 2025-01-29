-- name: CreateTask :one
INSERT INTO tasks (
    task_title,
    progress,
    project_id,
    assigned_user_id
) VALUES (
    $1, $2, $3, $4
) RETURNING *;

-- name: UpdateTaskTile :exec
UPDATE tasks SET task_title = $2
WHERE id = $1;

-- name: UpdateAssignedUserID :exec
UPDATE tasks SET assigned_user_id = $2
WHERE id = $1;

-- name: UpdateTaskProgress :exec
UPDATE tasks SET progress = $2
WHERE id = $1;

-- name: UpdateTaskProjectID :exec
UPDATE tasks SET project_id = $2
WHERE id = $1;

-- name: DeleteTask :exec
DELETE FROM tasks WHERE id = $1;

-- name: GetTaskByTaskID :one
SELECT * FROM tasks
WHERE id = $1;

-- name: GetTaskByID :one
SELECT * FROM tasks
WHERE id = $1;

-- name: ListTasksByUserID :many
SELECT * FROM tasks
WHERE assigned_user_id = $1;

-- name: ListTasksByProjectID :many
SELECT * FROM tasks
ORDER BY project_id;