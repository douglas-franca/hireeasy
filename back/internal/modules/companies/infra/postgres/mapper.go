package postgres

import (
	"github.com/douglas-franca/hireeasy/internal/modules/companies/domain"
	companiessqlc "github.com/douglas-franca/hireeasy/internal/modules/companies/infra/postgres/sqlc"
	"github.com/google/uuid"
)

func companyFromCreateRow(row companiessqlc.CreateCompanyForVerifiedUserRow) *domain.Company {
	return &domain.Company{ID: uuid.UUID(row.ID.Bytes), Name: row.Name, Role: row.Role}
}
func companyFromListRow(row companiessqlc.ListCompaniesByUserRow) domain.Company {
	return domain.Company{ID: uuid.UUID(row.ID.Bytes), Name: row.Name, Role: row.Role}
}
func companyFromFindRow(row companiessqlc.FindCompanyForUserRow) domain.Company {
	return domain.Company{ID: uuid.UUID(row.ID.Bytes), Name: row.Name, Role: row.Role}
}
