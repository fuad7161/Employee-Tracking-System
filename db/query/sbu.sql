-- name: CreateSBU :one
INSERT INTO sbus (
    sbu_name, sbu_head_user_id
) VALUES (
    $1, $2
) RETURNING *;

-- name: UpdateSbuByID :exec
UPDATE sbus SET sbu_name =$2, sbu_head_user_id = $3
WHERE id = $1;

-- name: UpdateSBUHeadUserID :exec
UPDATE sbus SET sbu_head_user_id = $2
WHERE id = $1;

-- name: DeleteSBU :exec
DELETE FROM sbus WHERE id = $1;

-- name: GetSBU :one
SELECT * FROM sbus
WHERE id = $1;

-- name: ListSBUs :many
SELECT * FROM sbus
ORDER BY id;