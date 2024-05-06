package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type User struct {
	Name     string `json:"name"`
	Password string `json:"password"`
}

var Users []User

func Create(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		var user User

		err := json.NewDecoder(r.Body).Decode(&user)
		if err != nil {
			http.Error(w, "Error decoding body", http.StatusBadRequest)
			return
		} else {
			Users = append(Users, user)
			fmt.Fprintf(w, "Successfully Added")
		}
	} else {
		http.Error(w, "Request method not allowed", http.StatusBadRequest)
	}

}
