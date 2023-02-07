package dto

type UserUpdateDTO struct {
	Email    string `json:"email"`
	Name     string `json:"name"`
	Password string `json:"password"`
}
