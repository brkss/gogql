// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.16.0
// source: session.sql

package db

import (
	"context"
	"time"
)

const createSession = `-- name: CreateSession :one
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
) RETURNING id, user_id, token, blocked, expired_at, created_at
`

type CreateSessionParams struct {
	ID        string    `json:"id"`
	UserID    string    `json:"user_id"`
	Token     string    `json:"token"`
	Blocked   bool      `json:"blocked"`
	ExpiredAt time.Time `json:"expired_at"`
}

func (q *Queries) CreateSession(ctx context.Context, arg CreateSessionParams) (Session, error) {
	row := q.db.QueryRowContext(ctx, createSession,
		arg.ID,
		arg.UserID,
		arg.Token,
		arg.Blocked,
		arg.ExpiredAt,
	)
	var i Session
	err := row.Scan(
		&i.ID,
		&i.UserID,
		&i.Token,
		&i.Blocked,
		&i.ExpiredAt,
		&i.CreatedAt,
	)
	return i, err
}
