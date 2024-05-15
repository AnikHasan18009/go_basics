package handler

import (
	"net/http"
	"user-service/db"
	utils "user-service/web/utils"
)

func RemoveRequestedUser(w http.ResponseWriter, r *http.Request) {
	var (
		err    error
		exists bool
	)

	id := utils.GetUserIdFromPath(r)
	if exists, err = db.CheckIfUserExistsInDB(id); exists {
		if err = db.DeleteUserByIdFromDB(id); err == nil {
			utils.SendJson(w, http.StatusOK, map[string]interface{}{
				"status":  true,
				"message": "Success",
			})
			return
		}
	}

	utils.SendError(w, http.StatusInternalServerError, err)

}
