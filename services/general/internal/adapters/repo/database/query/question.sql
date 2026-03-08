-- name: CreateQuestion :one
INSERT INTO general.questions(material_id, position, question) VALUES(
    @material_id,
    COALESCE((SELECT MAX(position) FROM general.questions WHERE material_id = @material_id), 0) + 1000,
    @question
)
RETURNING id;

-- name: CreateQuestionPossibleAnswer :one
INSERT INTO general.question_possible_answers(question_id, answer)
VALUES(@question_id, @answer)
RETURNING id;

-- name: CreateQuestionCriterion :one
INSERT INTO general.question_criteria(question_id, criterion, score)
VALUES(@question_id, @criterion, @score)
RETURNING id;

-- name: GetQuestionsByMaterial :many
SELECT * FROM general.questions WHERE material_id = @material_id;

-- name: GetQuestionPossibleAnswers :many
SELECT * FROM general.question_possible_answers WHERE question_id = @question_id;

-- name: GetQuestionCriteria :many
SELECT * FROM general.question_criteria WHERE question_id = @question_id;

-- name: UpdateQuestion :exec
UPDATE general.questions
SET
    question = COALESCE(sqlc.narg('question'), question),
    position = COALESCE(sqlc.narg('position'), position)
WHERE id = @id;
