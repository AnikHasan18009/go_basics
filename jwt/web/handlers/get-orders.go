package handler

import (
	"fmt"

	"net/http"
	"user-service/db"
	"user-service/types"
	"user-service/web/utils"
)

func GetOrders(w http.ResponseWriter, r *http.Request) {
	var err error
	var orders []types.Order
	if verified := utils.VerifyJWT(r.Header.Get("jwt")); verified {
		userId := utils.GetUserIdFromPath(r)
		if orders, err = db.GetUserOrdersFromDB(userId); err == nil {

			utils.SendJson(w, http.StatusOK, map[string]any{
				"data":    orders,
				"status":  true,
				"message": "Success",
			})
			return
		}

		utils.SendError(w, http.StatusInternalServerError, err)

	} else {
		utils.SendError(w, http.StatusBadRequest, fmt.Errorf("failed authentication"))
	}

}
