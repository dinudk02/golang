-- name: CreateUser :one
INSERT INTO users (id,username,created_at,password,email)
VALUES ($1,$2,$3,$4,$5)
RETURNING *;

