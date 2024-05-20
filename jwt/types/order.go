package types

type Order struct {
	Id     int    `json:"id" db:"id"`
	UserId int    `json:"user_id" db:"user_id"`
	Name   string `json:"name" db:"name"`
	Price  int    `json:"price" db:"price"`
}
