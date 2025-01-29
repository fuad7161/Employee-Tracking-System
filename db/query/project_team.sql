-- name: CreateProjectTeam :one
INSERT INTO project_teams (
    project_id,
    user_id
) VALUES (
    $1, $2
) RETURNING *;

-- name: UpdateUserProjectIDByUserID :exec
UPDATE project_teams SET project_id = $2
WHERE user_id = $1;

-- name: UpdateUserProjectIDByID :exec
UPDATE project_teams SET project_id = $2
WHERE id = $1;

-- name: DeleteProjectTeamUser :exec
DELETE FROM project_teams WHERE id = $1;

-- name: GetProjectTeamUserByID :one
SELECT * FROM project_teams
WHERE id = $1;

-- name: ListProjectTeamUsers :many
SELECT * FROM project_teams
ORDER BY project_id;