

-- name: CreateAwarenessContent :one
INSERT INTO "awarness"
( id, content, image, survey_id, title )
VALUES ( $1, $2, $3, $4, $5 )
RETURNING *;

-- name: GetAwarnesses :many
SELECT id, title, survey_id FROM "awarness"
ORDER BY id;

-- name: GetAwarness :one 
SELECT * FROM "awarness"
WHERE id = $1
LIMIT 1;
