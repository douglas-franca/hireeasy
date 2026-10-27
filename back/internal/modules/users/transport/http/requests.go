package http

type registerUserRequest struct {
	Name     string `json:"name" binding:"required,max=255"`
	Email    string `json:"email" binding:"required,email,max=255"`
	Password string `json:"password" binding:"required,min=8"`
}
type updateCurrentUserRequest struct {
	Name string `json:"name" binding:"required,max=255"`
}
type confirmEmailVerificationRequest struct {
	Token string `json:"token" binding:"required"`
}
