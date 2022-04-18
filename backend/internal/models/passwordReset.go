package models
const TablePassReset = "password_resets"
type PasswordReset struct {
	ID    int `json:"id"`
	Email string `json:"email"`
	Token string `json:"token"`
}