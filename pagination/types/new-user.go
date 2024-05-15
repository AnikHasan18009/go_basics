package types

type NewUser struct {
	FirstName string `json:"first_ame" db:"first_name"`
	LastName  string `json:"last_name" db:"last_name"`
	Password  string `json:"password" db:"password"`
}
