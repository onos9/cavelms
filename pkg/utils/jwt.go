package utils

import (
	"log"
	"time"

	"github.com/cavelms/config"
	"github.com/cavelms/internal/model"
	jwt "github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
)

// RefreshClaims represents refresh token JWT claims
type RefreshClaims struct {
	RefreshTokenID string `json:"refreshTokenID"`
	ExternalID     string `json:"userId"`
	Role           string `json:"role"`
	jwt.RegisteredClaims
}

// AccessClaims represents access token JWT claims
type AccessClaims struct {
	AccessTokenID string `json:"accessTokenID"`
	ExternalID    string `json:"userID"`
	Role          string `json:"role"`
	jwt.RegisteredClaims
}

type Token struct {
	TokenID          string    `json:"accessTokenID"`
	AccessToken      string    `json:"accessToken"`
	RefreshToken     string    `json:"refreshToken"`
	AccessExpiresAt  time.Time `json:"accessExpiresAt"`
	RefreshExpiresAt time.Time `json:"refreshExpiresAt"`
	ExternalID       string    `json:"userID"`
	Role             string    `json:"role"`
	jwt.RegisteredClaims
}

// IssueAccessToken generate access tokens used for authentication
func IssueToken(u model.User) (*Token, error) {
	t := new(Token)
	t.AccessExpiresAt = time.Now().Add(time.Hour)              // 1 hour
	t.RefreshExpiresAt = time.Now().Add((24 * time.Hour) * 14) // 14 days

	// Generate encoded token
	claims := AccessClaims{
		uuid.New().String(),
		u.ID,
		u.Role,
		jwt.RegisteredClaims{
			Issuer:    config.GetConfig().JWTIssuer,
		},
	}

	// IssueRefreshToken generate refresh tokens used for refreshing authentication
	claims.ExpiresAt = jwt.NewNumericDate(t.AccessExpiresAt)
	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tk, err := tokenClaims.SignedString([]byte(config.GetConfig().JWTAccessSecret))
	t.AccessToken = tk
	if err != nil {
		return nil, err
	}

	// IssueRefreshToken generate refresh tokens used for refreshing authentication
	claims.ExpiresAt = jwt.NewNumericDate(t.RefreshExpiresAt)
	tokenClaims = jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tk, err = tokenClaims.SignedString([]byte(config.GetConfig().JWTAccessSecret))
	t.RefreshToken = tk
	if err != nil {
		return nil, err
	}

	return t, nil
}

// IssueAccessToken generate access tokens used for authentication
func IssueAccessToken(u model.User) (string, error) {
	expireTime := time.Now().Add(time.Hour) // 1 hour

	// Generate encoded token
	claims := AccessClaims{
		uuid.New().String(),
		u.ID,
		u.Role,
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expireTime),
			Issuer:    config.GetConfig().JWTIssuer,
		},
	}
	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return tokenClaims.SignedString([]byte(config.GetConfig().JWTAccessSecret))
}

// IssueRefreshToken generate refresh tokens used for refreshing authentication
func IssueRefreshToken(u model.User) (string, error) {
	expireTime := time.Now().Add((24 * time.Hour) * 14) // 14 days

	// Generate encoded token
	claims := RefreshClaims{
		uuid.New().String(),
		u.ID,
		u.Role,

		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expireTime),
			Issuer:    config.GetConfig().JWTIssuer,
		},
	}

	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return tokenClaims.SignedString([]byte(config.GetConfig().JWTRefreshSecret))
}

func ParseToken(tokenString string) (jwt.MapClaims, error) {

	hmacSecret := []byte(config.GetConfig().JWTRefreshSecret)
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// check token signing method etc
		return hmacSecret, nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims, nil
	} else {
		log.Printf("Invalid JWT Token")
		return nil, ErrInvalidAuthToken
	}
}
