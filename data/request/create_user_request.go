package request

type CreateUserRequest struct {
	Email       string `validate:"required" json:"email"`
	Password    string `validate:"required" json:"password"`
	DateOfBirth string `validate:"required" json:"dateOfBirth"`
}

type LoginUserRequest struct {
	Email    string `validate:"required" json:"email"`
	Password string `validate:"required" json:"password"`
}
