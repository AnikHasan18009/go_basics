package utils

import (
	"encoding/json"
	"io"
)

func DecodeJSON(requestBody io.ReadCloser, userDefinedType interface{}) error {

	err := json.NewDecoder(requestBody).Decode(userDefinedType)
	return err
}
