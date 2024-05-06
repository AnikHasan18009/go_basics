package handler

import (
	"net/http"
	"passwordList/web/utils"
)

func ReturnList(w http.ResponseWriter, r *http.Request) {

	if r.Method == http.MethodGet {
		utils.SendData(w, Users)
	} else {

		utils.SendError(w, http.StatusInternalServerError, "Request method not allowed")
	}

}
