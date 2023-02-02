package models

type Credentials struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type User struct {
	UUID     string `json:"uuid"`
	Email    string `json:"email"`
	Password string `json:"password"`
}
