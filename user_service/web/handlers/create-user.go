package handler

import (
	"encoding/json"
	"net/http"
	"user-service/web/utils"
)

type NewUser struct {
	Name     string `json:"name"`
	Password string `json:"password"`
}
type User struct {
	Id       int    `json:"id"`
	Name     string `json:"name"`
	Password string `json:"password"`
}

var Users []User
var ID int = 1

func CreateUser(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		var user NewUser

		err := json.NewDecoder(r.Body).Decode(&user)
		if err != nil {
			utils.SendError(w, http.StatusBadRequest, "Error decoding body")
		} else {
			newuser := User{Id: ID, Name: user.Name, Password: user.Password}
			Users = append(Users, newuser)
			ID++
			utils.SendJson(w, http.StatusOK, map[string]interface{}{
				"status":  true,
				"message": "Success",
			})
		}
	} else {
		utils.SendError(w, http.StatusBadRequest, "Request method not allowed")
	}

}
