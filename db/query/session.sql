-- name: CreateSession :one
INSERT INTO sessions
(
	id,
	user_id,
	token,
	blocked,
	expired_at
) VALUES (
	$1,
	$2,
	$3,
	$4,
	$5
) RETURNING *;
