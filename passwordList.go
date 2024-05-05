package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type User struct {
	Name     string `json:"name"`
	Password string `json:"password"`
}

var Users []User

// Handler functions.
func Create(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		var user User

		err := json.NewDecoder(r.Body).Decode(&user)
		if err != nil {
			http.Error(w, "Error decoding body", http.StatusBadRequest)
			return
		} else {
			Users = append(Users, user)
			fmt.Fprintf(w, "Successfully Added")
		}
	} else {
		http.Error(w, "Request method not allowed", http.StatusBadRequest)
	}

}

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

func main() {
	//routes
	http.HandleFunc("/create", Create)
	http.HandleFunc("/list", ReturnList)

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}
