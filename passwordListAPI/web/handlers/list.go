package handler

import (
	"encoding/json"

	"net/http"
)

func ReturnList(w http.ResponseWriter, r *http.Request) {

	if r.Method == http.MethodGet {
		jsonData, err := json.Marshal(Users)
		if err != nil {
			http.Error(w, "Error encoding JSON", http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")

		// Write the JSON response
		_, err = w.Write(jsonData)
		if err != nil {

			http.Error(w, "Error writing JSON response:", http.StatusInternalServerError)
		}
	} else {
		http.Error(w, "Request method not allowed", http.StatusBadRequest)
	}

}
