package utils

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"log"
	"strings"
	"time"
	configuration "user-service/config"
	"user-service/types"
)

func VerifyJWT(token string) bool {
	claims := types.Claims{}
	parts := strings.Split(token, ".")

	signature := CreateSignature(parts[0], parts[1], configuration.GetConfig().SECRET_KEY)
	if signature != parts[2] {
		fmt.Println("Signature mismatch! Token is not valid.")
		return false
	}

	decodedPayload, err := base64.RawURLEncoding.DecodeString(parts[1])
	if err != nil {
		log.Fatal("error decoding payload")
	}

	_ = json.Unmarshal(decodedPayload, &claims)
	if claims.Exp == 0 {
		return false
	}

	if time.Now().Unix() > claims.Exp {

		return false
	}

	return true
}
