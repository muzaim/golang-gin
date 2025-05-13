package dto

type StudentRequest struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Major     string `json:"major"`
	Email     string `json:"email"`
}
