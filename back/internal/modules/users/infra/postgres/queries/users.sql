-- name: CreateUser :exec
INSERT INTO users (
  id,
  name,
  email,
  password_hash
) VALUES (
  sqlc.arg(id),
  sqlc.arg(name),
  sqlc.arg(email),
  sqlc.arg(password_hash)
);

-- name: FindUserByEmail :one
SELECT * FROM users WHERE email = $1 LIMIT 1;
