

-- name: CreateAwarenessContent :one
INSERT INTO "awareness"
( id, content, image, survey_id )
VALUES ( $1, $2, $3, $4 )
RETURNING *;

-- name: GetAwarnesses :many
SELECT * FROM "awareness"
ORDER BY id;

-- name: GetAwa
