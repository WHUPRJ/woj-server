package jwt

import "github.com/golang-jwt/jwt/v4"

type Claim struct {
	UID      uint   `json:"id"`
	UserName string `json:"user_name"`
	NickName string `json:"nick_name"`
	Version  int    `json:"version"`
	jwt.RegisteredClaims
}
