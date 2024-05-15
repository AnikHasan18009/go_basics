package types

type User struct {
	Id        int    `json:"id"  db:"id"`
	FirstName string `json:"first_ame" db:"first_name"`
	LastName  string `json:"last_name" db:"last_name"`
	Password  string `json:"password" db:"password"`
}
