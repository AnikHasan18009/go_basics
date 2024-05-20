package handler

import (
	"encoding/json"
	"net/http"
	"user-service/db"
	"user-service/types"
	"user-service/web/utils"
)

func CreateNewUser(w http.ResponseWriter, r *http.Request) {
	var user types.NewUser
	_ = json.NewDecoder(r.Body).Decode(&user)

	err := db.InsertUserIntoDB(user)
	if err == nil {
		utils.SendJson(w, http.StatusOK, map[string]interface{}{
			"status":  true,
			"message": "Success",
		})
	} else {
		utils.SendError(w, http.StatusInternalServerError, err)
	}

}
