-- name: CreateMaterial :one
INSERT INTO general.materials(
    course_id,
    kind,
    lesson,
    title,
    content,
    required_score,
    required_peers,
    is_required,
    position
) VALUES(
    @course_id,
    @kind,
    @lesson,
    @title,
    @content,
    @required_score,
    @required_peers,
    @is_required,
    COALESCE((SELECT MAX(position) FROM general.materials WHERE course_id = @course_id), 0) + 1000
)
RETURNING id;

-- name: GetMaterialsByCourse :many
SELECT * FROM general.materials WHERE course_id = @course_id;

-- name: GetMaterialPosition :one
SELECT position FROM general.materials WHERE id = @id;

-- name: UpdateMaterial :exec
UPDATE general.materials
SET
    position = COALESCE(sqlc.narg('position'), position),
    lesson = COALESCE(sqlc.narg('lesson'), lesson),
    title = COALESCE(sqlc.narg('title'), title),
    required_score = COALESCE(sqlc.narg('required_score'), required_score),
    required_peers = COALESCE(sqlc.narg('required_peers'), required_peers),
    is_required = COALESCE(sqlc.narg('is_required'), is_required)
WHERE id = @id;
