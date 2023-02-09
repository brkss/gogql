
-- name: CreateSurvey :one
INSERT INTO "surveys"
( id, name )
VALUES ( $1, $2 )
RETURNING *;


-- name: CreateSurveyQuestion :one
INSERT INTO "questions"
( id, qst, survey_id )
VALUES ( $1, $2, $3 )
RETURNING *;


-- name: CreateQuestionAnswer :one
INSERT INTO "answers"
( id, ans, val, question_id )
VALUES ( $1, $2, $3, $4)
RETURNING *;


-- name: CreateSurveyResult :one 
INSERT INTO "results"
( id, min, max, comment, surver_id )
VALUES ( $1, $2, $3, $4, $5 )
RETURNING *;

-- name: GetSurvies :many 
SELECT * FROM "surveys"
ORDER BY id;
