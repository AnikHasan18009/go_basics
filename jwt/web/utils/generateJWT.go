package utils

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"log"
	"time"
	configuration "user-service/config"
	"user-service/types"
)

func GenetateJWTForUser(user types.NewUser) string {

	claims := types.Claims{
		Name:  user.Name,
		Email: user.Email,
		Exp:   time.Now().Add(time.Minute * 5).Unix(),
	}
	header := `{"alg":"HS256","typ":"JWT"}`
	encodedHeader := base64.RawURLEncoding.EncodeToString([]byte(header))
	jsonClaims, err := json.Marshal(claims)
	if err != nil {
		log.Fatal("Error encoding claims:", err)

	}
	encodedClaims := base64.RawURLEncoding.EncodeToString(jsonClaims)

	signature := CreateSignature(encodedHeader, encodedClaims, configuration.GetConfig().SECRET_KEY)

	token := encodedHeader + "." + encodedClaims + "." + signature

	fmt.Println("JWT token:", token)
	return token

}
