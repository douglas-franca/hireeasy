package http

import (
	"github.com/douglas-franca/hireeasy/internal/modules/companies/domain"
	"github.com/google/uuid"
)

type companyResponse struct {
	ID   uuid.UUID `json:"id"`
	Name string    `json:"name"`
	Role string    `json:"role"`
}

func companyResponseFrom(company *domain.Company) companyResponse {
	return companyResponse{ID: company.ID, Name: company.Name, Role: company.Role}
}
func companyResponsesFrom(companies []domain.Company) []companyResponse {
	responses := make([]companyResponse, 0, len(companies))
	for index := range companies {
		responses = append(responses, companyResponseFrom(&companies[index]))
	}
	return responses
}
