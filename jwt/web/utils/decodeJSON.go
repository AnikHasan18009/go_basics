package utils

import (
	"encoding/json"
	"io"
	"log"
)

func DecodeJSON(requestBody io.ReadCloser, userDefinedType interface{}) {

	err := json.NewDecoder(requestBody).Decode(userDefinedType)
	log.Fatal(err)
}
