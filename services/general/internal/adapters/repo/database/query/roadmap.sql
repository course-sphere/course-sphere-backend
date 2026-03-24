-- name: CreateRoadmap :one
INSERT INTO general.roadmaps(author_id, title, description)
VALUES(@author_id, @title, @description)
RETURNING id;

-- name: AddRoadmapCourse :exec
INSERT INTO general.roadmap_courses(
    roadmap_id,
    course_id,
    position
)
VALUES(
    @id,
    @course_id,
    COALESCE((SELECT MAX(position) FROM general.roadmap_courses WHERE roadmap_id = @id), 0) + 1000
);

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
SELECT course_id
FROM general.roadmap_courses
WHERE roadmap_id = @id
ORDER BY position;

-- name: GetRoadmapCoursePosition :one
SELECT position FROM general.roadmap_courses WHERE roadmap_id = @id AND course_id = @course_id;

-- name: UpdateRoadmap :exec
UPDATE general.roadmaps
SET
    title = COALESCE(sqlc.narg('title'), title),
    description = COALESCE(sqlc.narg('description'), description)
WHERE id = @id;

-- name: UpdateRoadmapCourse :exec
UPDATE general.roadmap_courses
SET position = @posistion
WHERE roadmap_id = @id AND course_id = @course_id;
