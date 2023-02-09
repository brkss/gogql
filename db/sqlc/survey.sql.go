// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.16.0
// source: survey.sql

package db

import (
	"context"
)

const createQuestionAnswer = `-- name: CreateQuestionAnswer :one
INSERT INTO "answers"
( id, ans, val, question_id )
VALUES ( $1, $2, $3, $4)
RETURNING id, ans, question_id, val
`

type CreateQuestionAnswerParams struct {
	ID         string `json:"id"`
	Ans        string `json:"ans"`
	Val        int32  `json:"val"`
	QuestionID string `json:"question_id"`
}

func (q *Queries) CreateQuestionAnswer(ctx context.Context, arg CreateQuestionAnswerParams) (Answer, error) {
	row := q.db.QueryRowContext(ctx, createQuestionAnswer,
		arg.ID,
		arg.Ans,
		arg.Val,
		arg.QuestionID,
	)
	var i Answer
	err := row.Scan(
		&i.ID,
		&i.Ans,
		&i.QuestionID,
		&i.Val,
	)
	return i, err
}

const createSurvey = `-- name: CreateSurvey :one
INSERT INTO "surveys"
( id, name )
VALUES ( $1, $2 )
RETURNING id, name
`

type CreateSurveyParams struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

func (q *Queries) CreateSurvey(ctx context.Context, arg CreateSurveyParams) (Survey, error) {
	row := q.db.QueryRowContext(ctx, createSurvey, arg.ID, arg.Name)
	var i Survey
	err := row.Scan(&i.ID, &i.Name)
	return i, err
}

const createSurveyQuestion = `-- name: CreateSurveyQuestion :one
INSERT INTO "questions"
( id, qst, survey_id )
VALUES ( $1, $2, $3 )
RETURNING id, qst, survey_id
`

type CreateSurveyQuestionParams struct {
	ID       string `json:"id"`
	Qst      string `json:"qst"`
	SurveyID string `json:"survey_id"`
}

func (q *Queries) CreateSurveyQuestion(ctx context.Context, arg CreateSurveyQuestionParams) (Question, error) {
	row := q.db.QueryRowContext(ctx, createSurveyQuestion, arg.ID, arg.Qst, arg.SurveyID)
	var i Question
	err := row.Scan(&i.ID, &i.Qst, &i.SurveyID)
	return i, err
}

const createSurveyResult = `-- name: CreateSurveyResult :one
INSERT INTO "results"
( id, min, max, comment, surver_id )
VALUES ( $1, $2, $3, $4, $5 )
RETURNING id, min, max, comment, surver_id
`

type CreateSurveyResultParams struct {
	ID       string `json:"id"`
	Min      int32  `json:"min"`
	Max      int32  `json:"max"`
	Comment  string `json:"comment"`
	SurverID string `json:"surver_id"`
}

func (q *Queries) CreateSurveyResult(ctx context.Context, arg CreateSurveyResultParams) (Result, error) {
	row := q.db.QueryRowContext(ctx, createSurveyResult,
		arg.ID,
		arg.Min,
		arg.Max,
		arg.Comment,
		arg.SurverID,
	)
	var i Result
	err := row.Scan(
		&i.ID,
		&i.Min,
		&i.Max,
		&i.Comment,
		&i.SurverID,
	)
	return i, err
}

const getSurvies = `-- name: GetSurvies :many
SELECT id, name FROM "surveys"
ORDER BY id
`

func (q *Queries) GetSurvies(ctx context.Context) ([]Survey, error) {
	rows, err := q.db.QueryContext(ctx, getSurvies)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Survey{}
	for rows.Next() {
		var i Survey
		if err := rows.Scan(&i.ID, &i.Name); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}