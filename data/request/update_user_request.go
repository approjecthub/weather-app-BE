package request

type UpdateUserRequest struct {
	Id          uint   `validate:"required"`
	Email       string `validate:"required" json:"email"`
	Password    string `validate:"required" json:"password"`
	DateOfBirth string `validate:"required" json:"dateOfBirth"`
}
