package main

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
)

// Auth structure will be used to handle the authenticated user data.
type Auth struct {
	ID       uint   `json:"id"`
	Name     string `json:"name"`
	Surname  string `json:"surname"`
	Username string `json:"username"`
	Password string `json:"password"`
}

// GetUser will parse incoming request and returns the user data.
func (c *Auth) GetUser(req *http.Request, key string) error {
	bearerSchema := "Bearer "
	tokenString := req.Header.Get("Authorization")

	claims := jwt.MapClaims{}
	_, _ = jwt.ParseWithClaims(tokenString[len(bearerSchema):], claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(key), nil
	})

	userData := claims["user"].(string)
	err := json.Unmarshal([]byte(userData), &c)

	if err != nil {
		return err
	}

	return nil
}

// NewToken will return a new JWT token
func (c *Auth) NewToken(key string) (string, bool) {
	c.Password = ""
	userDataString, _ := json.Marshal(c)
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user": string(userDataString),
		"exp":  time.Now().Add(time.Hour * time.Duration(2)).Unix(),
		"iat":  time.Now().Unix(),
	})

	tokenString, err := token.SignedString([]byte(key))

	if err != nil {
		return "", false
	}

	return tokenString, true
}

// RefreshToken will grefresh the a speficic token
func (c *Auth) RefreshToken(req http.ResponseWriter, key string) bool {
	expirationTime := time.Now().Add(5 * time.Minute)
	userDataString, _ := json.Marshal(c)
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user": string(userDataString),
		"exp":  expirationTime,
		"iat":  time.Now().Unix(),
	})

	tokenString, err := token.SignedString([]byte(key))
	req.Header().Set("refresh-token", tokenString)

	if err != nil {
		return false
	}

	return true
}
