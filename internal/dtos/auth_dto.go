package dtos

type GetProfile struct {
	ID       string `json:"user_id" validate:"required"`
	Username string `json:"user_name" validate:"required"`
	Email    string `json:"user_email" validate:"required"`
}
