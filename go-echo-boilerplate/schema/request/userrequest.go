package request

type UpdateUserRequest struct {
	Name  string `json:"name,omitempty" validate:"omitempty,min=2"`
	Email string `json:"email,omitempty" validate:"omitempty,email"`
	Role  string `json:"role,omitempty"`
}
