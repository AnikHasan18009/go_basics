package handler

import (
	"net/http"
	"user-service/db"
	"user-service/types"
	"user-service/web/utils"
)

func GetUsers(w http.ResponseWriter, r *http.Request) {
	var users []types.User

	users, err := db.RetriveUsersFromDB(users)

	if err != nil {
		utils.SendError(w, http.StatusInternalServerError, err)
	} else {

		utils.SendData(w, users)
	}

}
