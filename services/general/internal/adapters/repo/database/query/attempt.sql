-- name: CreateAttempt :one
INSERT INTO general.attempts(material_id, student_id)
VALUES(@material_id, @student_id)
RETURNING id;

-- name: CreateAttemptDetail :exec
INSERT INTO general.attempt_details(attempt_id, question_id, answer)
VALUES(@id, @question_id, @answer);

-- name: GetAttemptsByMaterial :many
SELECT *
FROM general.attempts
WHERE material_id = @material_id AND student_id = @student_id;

-- name: GetAttempt :one
SELECT *
FROM general.attempts
WHERE id = @id;

-- name: GetAttemptDetails :many
SELECT *
FROM general.attempt_details
WHERE attempt_id = @id;

-- name: UpdateAttempt :exec
UPDATE general.attempts
SET score = @score
WHERE id = @id;
