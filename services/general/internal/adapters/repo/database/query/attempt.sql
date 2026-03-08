-- name: CreateAttempt :one
INSERT INTO general.material_attempts(material_id, student_id, score)
VALUES (@material_id, @student_id, @score)
RETURNING id;

-- name: GetAttemptsByMaterial :many
SELECT *
FROM general.material_attempts
WHERE material_id = @material_id AND student_id = @student_id;
