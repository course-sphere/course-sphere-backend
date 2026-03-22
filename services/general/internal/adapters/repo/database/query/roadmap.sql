-- name: CreateRoadmap :one
INSERT INTO general.roadmaps(
    author_id,
    title,
    description,
    position
) VALUES(
    @author_id,
    @title,
    @description,
    COALESCE((SELECT MAX(position) FROM general.roadmaps), 0) + 1000
)
RETURNING id;

-- name: AddRoadmapCourse :exec
INSERT INTO general.roadmap_courses(roadmap_id, course_id)
VALUES(@id, @course_id);

-- name: ApplyRoadmap :exec
INSERT INTO general.student_roadmaps(student_id, roadmap_id)
VALUES(@student_id, @id);

-- name: GetAllRoadmaps :many
SELECT id FROM general.roadmaps;

-- name: GetRoadmapsByStudent :many
SELECT id FROM general.roadmaps
WHERE id IN (
    SELECT roadmap_id
    FROM general.student_roadmaps
    WHERE student_id = @student_id
);

-- name: GetRoadmap :one
SELECT * FROM general.roadmaps WHERE id = @id;

-- name: GetRoadmapCourse :many
SELECT course_id FROM general.roadmap_courses WHERE roadmap_id = @id;

-- name: UpdateRoadmap :exec
UPDATE general.roadmaps
SET
    position = COALESCE(sqlc.narg('position'), position),
    title = COALESCE(sqlc.narg('title'), title),
    description = COALESCE(sqlc.narg('description'), description)
WHERE id = @id;
