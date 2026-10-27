package http

type createCompanyRequest struct {
	Name string `json:"name" binding:"required,max=255"`
}
