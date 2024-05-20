package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"user-service/db"
	"user-service/types"
	"user-service/web/utils"
)

func LoginUser(w http.ResponseWriter, r *http.Request) {
	var user types.NewUser
	_ = json.NewDecoder(r.Body).Decode(&user)

	if count := db.CountOfRecordsBasedOnConditions(map[string]string{
		"name": user.Name, "email": user.Email, "password": user.Password,
	}); count == 1 {
		jwt := utils.GenetateJWTForUser(user)
		utils.SendData(w, map[string]any{
			"login-status": "logged in",
			"jwt-token":    jwt,
		})
	} else {
		utils.SendError(w, http.StatusBadRequest, fmt.Errorf("wrong credentials"))
	}

}
