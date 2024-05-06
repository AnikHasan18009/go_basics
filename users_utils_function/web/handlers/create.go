package handler

import (
	"encoding/json"
	"net/http"
	"passwordList/web/utils"
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
			utils.SendError(w, http.StatusBadRequest, "Error decoding body")
		} else {
			Users = append(Users, user)
			utils.SendJson(w, http.StatusOK, map[string]interface{}{
				"status":  true,
				"message": "Success",
			})
		}
	} else {
		utils.SendError(w, http.StatusBadRequest, "Request method not allowed")
	}

}
