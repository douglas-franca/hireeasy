package postgres

import (
	"errors"

	"github.com/jackc/pgx/v5/pgconn"
)

func isEmailAlreadyExists(err error) bool {
	var postgresError *pgconn.PgError
	return errors.As(err, &postgresError) && postgresError.Code == "23505" && postgresError.ConstraintName == "users_email_key"
}
