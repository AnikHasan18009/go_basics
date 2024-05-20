package utils

import "net/http"

func SendError(w http.ResponseWriter, status int, err error) {
	if err != nil {
		SendJson(w, status, map[string]any{
			"status":  false,
			"message": err.Error(),
		})
	} else {
		SendJson(w, http.StatusOK, map[string]any{
			"status":  true,
			"message": "Success",
		})
	}

}
