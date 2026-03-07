-- name: CreateCategory :exec
INSERT INTO course.categories(name) VALUES(@category)
ON CONFLICT DO NOTHING;
