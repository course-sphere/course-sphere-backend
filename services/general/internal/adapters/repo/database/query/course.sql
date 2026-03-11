-- name: CreateCourse :one
INSERT INTO general.courses(
    instructor_id,
    title,
    description,
    level,
    price,
    learning_objectives
) VALUES(
    @instructor_id,
    @title,
    @description,
    @level,
    @price,
    @learning_objectives
)
RETURNING id;

-- name: AddCourseCategory :exec
INSERT INTO general.course_categories(course_id, category_id) VALUES (
    @id,
    (SELECT id FROM general.categories WHERE name = @category)
);

-- name: AddCoursePrerequisite :exec
INSERT INTO general.course_prerequisites(course_id, other_id) VALUES (@id, @other_id);

-- name: EnrollCourse :exec
INSERT INTO general.enrolls(course_id, student_id) VALUES(@id, @student_id);

-- name: GetAllCourses :many
SELECT id
FROM general.courses;

-- name: GetEnrolledCourses :many
SELECT id
FROM general.courses
WHERE status = 'approved' AND id IN (SELECT course_id FROM general.enrolls WHERE student_id = @student_id);

-- name: GetCourse :one
SELECT 
    id,
    instructor_id,
    title,
    subtitle,
    description,
    level,
    thumbnail_url,
    promo_video_url,
    price,
    requirements,
    learning_objectives,
    target_audiences,
    status,
    (SELECT COUNT(id) FROM general.materials m WHERE m.course_id = @id) as total,
    (SELECT COUNT(id) FROM general.materials m WHERE m.course_id = @id AND is_required = true) as total_required
FROM general.courses
WHERE id = @id;

-- name: GetCourseCategories :many
SELECT name
FROM general.categories
WHERE id IN (
    SELECT category_id
    FROM general.course_categories c
    WHERE c.course_id = @id
);

-- name: GetCoursePrerequisites :many
SELECT other_id
FROM general.course_prerequisites
WHERE course_id = @id;

-- name: UpdateCourse :exec
UPDATE general.courses
SET
    title = COALESCE(sqlc.narg('title'), title),
    subtitle = COALESCE(sqlc.narg('subtitle'), subtitle),
    description = COALESCE(sqlc.narg('description'), description),
    level = COALESCE(sqlc.narg('level'), level),
    thumbnail_url = COALESCE(sqlc.narg('thumbnail_url'), thumbnail_url),
    promo_video_url = COALESCE(sqlc.narg('promo_video_url'), promo_video_url),
    price = COALESCE(sqlc.narg('price'), price),
    requirements = COALESCE(sqlc.narg('requirements'), requirements),
    learning_objectives = COALESCE(sqlc.narg('learning_objectives'), learning_objectives),
    target_audiences = COALESCE(sqlc.narg('target_audiences'), target_audiences),
    status = COALESCE(sqlc.narg('status'), status)
WHERE id = @id AND instructor_id = @instructor_id;

-- name: DeleteCourseCategories :exec
DELETE FROM general.course_categories
WHERE course_id = @id;

-- name: DeleteCoursePrerequisites :exec
DELETE FROM general.course_prerequisites
WHERE course_id = @id;
