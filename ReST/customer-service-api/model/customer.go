package model

type Customer struct {
	Id    int    `json:"id"`
	Name  string `json:"name"`
	City  string `json:"city"`
	Email string `json:"email"`
}
