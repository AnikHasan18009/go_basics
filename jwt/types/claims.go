package types

type Claims struct {
	Name     string `json:"name"`
	Password string `json:"password"`
	Exp      int64  `json:"exp"`
}
