package jwt

import (
	"fmt"
	"github.com/WHUPRJ/woj-server/internal/e"
	"github.com/WHUPRJ/woj-server/pkg/utils"
	"github.com/golang-jwt/jwt/v4"
	"go.uber.org/zap"
	"time"
)

func (s *service) ParseToken(tokenText string) (*Claim, e.Err) {
	if tokenText == "" {
		return nil, e.TokenEmpty
	}

	token, err := jwt.ParseWithClaims(
		tokenText,
		&Claim{},
		func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}
			return s.SigningKey, nil
		})

	if ve, ok := err.(*jwt.ValidationError); ok {
		if ve.Errors&jwt.ValidationErrorMalformed != 0 {
			return nil, e.TokenMalformed
		} else if ve.Errors&(jwt.ValidationErrorExpired|jwt.ValidationErrorNotValidYet) != 0 {
			// Token is either expired or not active yet
			return nil, e.TokenTimeError
		} else {
			return nil, e.TokenInvalid
		}
	} else if err != nil {
		s.log.Warn("JWT Token Parse Error", zap.Error(err))
		return nil, e.TokenUnknown
	}

	if token.Valid {
		c := token.Claims.(*Claim)
		return c, e.Success
	}

	return nil, e.TokenInvalid
}

func (s *service) SignClaim(claim *Claim) (string, e.Err) {
	now := time.Now()

	claim.IssuedAt = jwt.NewNumericDate(now)
	claim.ExpiresAt = jwt.NewNumericDate(now.Add(time.Duration(s.ExpireHour) * time.Hour))
	claim.ID = utils.RandomString(16)
	// TODO: use per-user claim.Version to tracker invalidation
	claim.NotBefore = jwt.NewNumericDate(time.Now())

	token := jwt.NewWithClaims(jwt.SigningMethodHS512, claim)
	ss, err := token.SignedString(s.SigningKey)
	if err != nil {
		s.log.Warn("jwt.SignedString error", zap.Error(err))
		return "", e.TokenSignError
	}
	return ss, e.Success
}
