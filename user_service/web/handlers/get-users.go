package handler

import (
	"net/http"
	"user-service/web/utils"
)

func GetUsers(w http.ResponseWriter, r *http.Request) {

	if r.Method == http.MethodGet {
		utils.SendData(w, Users)
	} else {

		utils.SendError(w, http.StatusInternalServerError, "Request method not allowed")
	}

}
