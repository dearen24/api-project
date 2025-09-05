package models

type User struct{
	Id   int    `json:"id"`
	Username string `json:"username"`
	Email string `json:"email"`
	Password_hash  string    `json:"password_hash"`
	First_name    string    `json:"first_name"`
	Last_name     string    `json:"last_name"`
	Is_active    bool      `json:"is_active"`
	Created_at   string    `json:"created_at"`
	Updated_at   string    `json:"updated_at"`
}