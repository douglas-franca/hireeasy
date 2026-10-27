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

-- name: FindUserByID :one
SELECT * FROM users WHERE id = $1 LIMIT 1;

-- name: UpdateUserName :execrows
UPDATE users SET name = $2, updated_at = now() WHERE id = $1;

-- name: DeleteUser :execrows
DELETE FROM users WHERE id = $1;

-- name: HasOwnedCompanies :one
SELECT EXISTS (SELECT 1 FROM user_companies WHERE user_id = $1 AND role = 'owner');

-- name: SetEmailVerificationToken :execrows
UPDATE users
SET email_verification_token_hash = $2,
    email_verification_expires_at = $3,
    email_verification_sent_at = now()
WHERE id = $1;

-- name: VerifyEmail :one
UPDATE users
SET email_verified = TRUE,
    email_verified_at = now(),
    email_verification_token_hash = NULL,
    email_verification_expires_at = NULL
WHERE email_verification_token_hash = $1
  AND email_verification_expires_at > now()
RETURNING *;
