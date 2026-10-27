-- name: CreateCompanyForVerifiedUser :one
WITH new_company AS (
  INSERT INTO companies (name)
  SELECT sqlc.arg(name)
  WHERE EXISTS (
    SELECT 1 FROM users WHERE users.id = sqlc.arg(user_id) AND users.email_verified = TRUE
  )
  RETURNING id, name
), new_membership AS (
  INSERT INTO user_companies (user_id, company_id, role)
  SELECT sqlc.arg(user_id), id, 'owner' FROM new_company
)
SELECT id, name, 'owner'::text AS role FROM new_company;

-- name: ListCompaniesByUser :many
SELECT c.id, c.name, uc.role
FROM companies AS c
JOIN user_companies AS uc ON uc.company_id = c.id
WHERE uc.user_id = $1
ORDER BY c.created_at DESC;

-- name: FindCompanyForUser :one
SELECT c.id, c.name, uc.role
FROM companies AS c
JOIN user_companies AS uc ON uc.company_id = c.id
WHERE c.id = $1 AND uc.user_id = $2;
