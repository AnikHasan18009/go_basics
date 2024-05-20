package types

type Claims struct {
	Name  string `json:"name"`
	Email string `json:"email"`
	Exp   int64  `json:"exp"`
}
