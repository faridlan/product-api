package web

import (
	"time"

	"github.com/golang-jwt/jwt/v4"
)

var JwtSigningMEethod = jwt.SigningMethodHS256
var JwtSecret = []byte("Anj1ngAd4l4hH3w4nuuAAii8sdA73DFed7")
var ExpiredTime = time.Now().Add(time.Minute * 5)

type Claims struct {
	Id       int    `json:"id,omitempty"`
	Username string `json:"username,omitempty"`
	*jwt.RegisteredClaims
}
