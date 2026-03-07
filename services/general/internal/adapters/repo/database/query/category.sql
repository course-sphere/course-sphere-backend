-- name: CreateCategory :exec
INSERT INTO general.categories(name) VALUES(@category)
ON CONFLICT DO NOTHING;
