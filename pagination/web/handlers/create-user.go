package handler

import (
	"net/http"
	"user-service/db"
	"user-service/types"
	"user-service/web/utils"
)

func CreateNewUser(w http.ResponseWriter, r *http.Request) {
	var user types.NewUser
	err := utils.DecodeJSON(r.Body, user)
	if err == nil {
		err = db.InsertUserIntoDB(user)
	}
	utils.SendError(w, http.StatusInternalServerError, err)

}
