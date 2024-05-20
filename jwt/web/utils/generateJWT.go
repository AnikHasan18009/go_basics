package utils

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"log"
	"time"
	configuration "user-service/config"
	"user-service/types"
)

func GenetateJWTForUser(user types.NewUser) string {

	hashedPassword := Sha1Hash(user.Password)
	claims := types.Claims{
		Name:     user.Name,
		Password: hashedPassword,
		Exp:      time.Now().Add(time.Minute * 3).Unix(),
	}
	header := `{"alg":"HS256","typ":"JWT"}`
	encodedHeader := base64.RawURLEncoding.EncodeToString([]byte(header))
	jsonClaims, err := json.Marshal(claims)
	if err != nil {
		log.Fatal("Error encoding claims:", err)

	}
	encodedClaims := base64.RawURLEncoding.EncodeToString(jsonClaims)

	signature := computeHMAC(encodedHeader, encodedClaims, configuration.GetConfig().SECRET_KEY)

	token := encodedHeader + "." + encodedClaims + "." + signature

	fmt.Println("JWT token:", token)
	return token

}

func computeHMAC(header, payload, key string) string {
	h := hmac.New(sha256.New, []byte(key))
	h.Write([]byte(header + "." + payload))
	return base64.RawURLEncoding.EncodeToString(h.Sum(nil))
}
