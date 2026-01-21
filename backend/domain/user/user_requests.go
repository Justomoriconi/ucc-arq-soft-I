package user

type SignupRequest struct {
	Name     string `json:"name"`
	LastName string `json:"last_name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginResponse struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
