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
		jwt := utils.GenetateJWTForUser(user)
		user.Password = utils.Sha1Hash(user.Password)
		utils.SendData(w, map[string]any{
			"user-data": user,
			"jwt-token": jwt,
		})
	} else {
		utils.SendError(w, http.StatusInternalServerError, err)
	}

}
