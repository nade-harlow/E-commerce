package models

type PasswordRecovery struct {
	UserID    string `json:"user_id"`
	UserEmail string `json:"user_email"`
}
