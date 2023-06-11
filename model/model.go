package model

type User struct {
	Password string
	Role     string
}

type Credentials struct {
	Password string `json:"password"`
	Username string `json:"username"`
}
