-- name: CreateCourse :one
INSERT INTO course.courses(
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
    target_audiences
FROM course.courses
WHERE id = @id;

-- name: GetCourseCategories :many
SELECT name
FROM course.categories
WHERE id IN (
    SELECT category_id
    FROM course.course_categories c
    WHERE c.course_id = @id
);

-- name: GetCoursePrerequisites :many
SELECT other_id
FROM course.course_prerequisites
WHERE course_id = @id;

-- name: AddCourseCategory :exec
INSERT INTO course.course_categories(course_id, category_id) VALUES (
    @id,
    (SELECT id FROM course.categories WHERE name = @category)
);

-- name: AddCoursePrerequisite :exec
INSERT INTO course.course_prerequisites(course_id, other_id) VALUES (@id, @other_id);
