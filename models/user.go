package models

type User struct {
    Username  string `json:"username"`
    Password  string `json:"password"`
	Active    bool   `json:"active"`
}