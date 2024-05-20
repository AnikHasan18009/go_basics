package handler

import (
	"net/http"
	"user-service/db"
	"user-service/types"
	"user-service/web/utils"
)

func UpdateRequestedUser(w http.ResponseWriter, r *http.Request) {

	updatedUserId := utils.GetUserIdFromPath(r)
	var (
		err    error
		exists bool
	)

	if exists, err = db.CheckIfUserExistsInDB(updatedUserId); exists {
		updatedUser := types.NewUser{}
		utils.DecodeJSON(r.Body, updatedUser)
		if err = db.UpdateUserInDB(updatedUserId, updatedUser); err == nil {
			utils.SendJson(w, http.StatusOK, map[string]interface{}{
				"status":  true,
				"message": "Success",
			})
			return
		}
	}

	utils.SendError(w, http.StatusInternalServerError, err)

}
