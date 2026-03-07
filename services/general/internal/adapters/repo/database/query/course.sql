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
    (
        SELECT ARRAY(
            SELECT name from course.categories
            WHERE id IN (
                SELECT category_id
                FROM course.course_categories c
                WHERE c.course_id = @id
            )
        )
    ) as categories,
    level,
    thumbnail_url,
    promo_video_url,
    price,
    (
        SELECT ARRAY(
            SELECT other_id from course.course_prerequisites
            WHERE course_id = @id
        )
    ) as prerequisites,
    requirements,
    learning_objectives,
    target_audiences
FROM course.courses
WHERE id = @id;

-- name: AddCourseCategory :exec
INSERT INTO course.course_categories(course_id, category_id) VALUES (
    @id,
    (SELECT id FROM course.categories WHERE name = @category)
);

-- name: AddCoursePrerequisite :exec
INSERT INTO course.course_prerequisites(course_id, other_id) VALUES (@id, @other_id);
