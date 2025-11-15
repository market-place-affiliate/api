package jwtPkg

import (
	"encoding/json"
	"fmt"
	"log"
	"reflect"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

const issuer = "radna-api"

type authCustomClaims struct {
	Device string `json:"device"`
	Userid string `json:"userid"`
	Issuer string `json:"issuer"`
	Role   string `json:"role"`
}

func GenerateToken(secret, id string, role, device string, durationMinute int) string {
	claims := &jwt.MapClaims{
		"userid": id,
		"device": device,
		"issuer": issuer,
	}

	if durationMinute > 0 {
		claims = &jwt.MapClaims{
			"userid": id,
			"device": device,
			"issuer": issuer,
			"exp":    time.Now().Add(time.Minute * time.Duration(durationMinute)).Unix(),
		}
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	//encoded string
	t, err := token.SignedString([]byte(secret)) //secret key
	if err != nil {
		panic(err)
	}
	return t
}

func ValidateToken(secret, encodedToken string) (*jwt.Token, error) {
	return jwt.Parse(encodedToken, func(token *jwt.Token) (interface{}, error) {
		if _, isvalid := token.Method.(*jwt.SigningMethodHMAC); !isvalid {
			return nil, fmt.Errorf("Invalid_token_%v", token.Header["alg"])
		}
		return []byte(secret), nil //secret key
	})

}

func GetClaims(tokenObj *jwt.Token) (*authCustomClaims, bool) {
	if claims, ok := tokenObj.Claims.(jwt.MapClaims); ok && tokenObj.Valid {
		var authClaim authCustomClaims
		b, err := json.Marshal(claims)
		if err != nil {
			return nil, false
		}
		err = json.Unmarshal(b, &authClaim)
		if err != nil {
			return nil, false
		}
		return &authClaim, true
	} else {
		fmt.Println(reflect.TypeOf(tokenObj.Claims))
		log.Printf("Invalid JWT Token")
		return nil, false
	}
}

func ValidAndGetClaims(secret, encodedToken string) (*authCustomClaims, bool) {
	tokenObj, err := ValidateToken(secret, encodedToken)
	if err != nil {
		return nil, false
	}
	return GetClaims(tokenObj)
}
