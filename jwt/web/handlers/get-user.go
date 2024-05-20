package handler

import (
	"fmt"
	"net/http"
	"user-service/db"
	"user-service/web/utils"
)

func GetRequestedUser(w http.ResponseWriter, r *http.Request) {

	userId := utils.GetUserIdFromPath(r)
	fmt.Print(userId)

	user, err := db.RetriveUserByIdFromDB(userId)
	if err != nil {
		utils.SendError(w, http.StatusInternalServerError, err)
	} else {
		utils.SendData(w, user)
	}

}
