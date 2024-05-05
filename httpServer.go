package main

import (
	"fmt"
	"log"
	"net/http"
)

const INFO = `
<html> 
<body>

<table style="text-align: center; border: 2px solid blue;
display : flex; align-items: center; justify-content: center;">
<tr><td>Anik Hasan</td></tr>
<tr><td>Junior SWE</td></tr>
<tr><td>Golang</td></tr>
</table>
</body>
</html>
`

// Handler functions.
func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, INFO)
}

func main() {
	log.Println("Starting our simple http server.")

	http.HandleFunc("/index", Index)

	log.Println("Started on port", ":8080")
	fmt.Println("To close connection CTRL+C :-)")

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}
