package service

import (
	"time"

	"github.com/cavelms/config"
	"github.com/cavelms/internal/model"
	"github.com/cavelms/pkg/utils"
	jwt "github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
)

// jwt service
type JWTService interface {
	GenerateToken(u model.User, refresh bool) (*Token, error)
	ValidateAccessToken(tokenString string) (jwt.MapClaims, error)
	ValidateRefreshToken(tokenString string) (jwt.MapClaims, error)
}

type TokenClaims struct {
	TokenID string `json:"tokenId"`
	UserID  string `json:"userId"`
	Email   string `json:"email"`
	Role    string `json:"role"`
	jwt.RegisteredClaims
}

type Token struct {
	AccessToken      string `json:"accessToken"`
	RefreshToken     string `json:"refreshToken"`
	AccessExpiresAt  int64  `json:"accessExpiresAt"`
	RefreshExpiresAt int64  `json:"refreshExpiresAt"`
}

type jwtService struct {
	accessSecret     string
	refreshSecret    string
	accessExpiresAt  int64
	refreshExpiresAt int64
	issure           string
}

// auth-jwt
func JWTAuthService() JWTService {
	return &jwtService{
		accessSecret:  config.GetConfig().JWTAccessSecret,
		refreshSecret: config.GetConfig().JWTRefreshSecret,
		issure:        config.GetConfig().JWTIssuer,

		accessExpiresAt:  int64(time.Hour),             // 1 hour
		refreshExpiresAt: int64((24 * time.Hour) * 14), // 14 days
	}
}

// IssueAccessToken generate access tokens used for authentication
func (j *jwtService) GenerateToken(u model.User, refresh bool) (*Token, error) {
	token := &Token{}

	// Generate encoded token
	claims := TokenClaims{
		TokenID:          uuid.New().String(),
		UserID:           u.ID,
		Email:            u.Email,
		Role:             u.Role,
		RegisteredClaims: jwt.RegisteredClaims{Issuer: j.issure},
	}

	// IssueRefreshToken generate refresh tokens used for refreshing authentication
	if refresh {
		claims.TokenID = uuid.New().String()
		t := time.Now().Add(time.Duration(j.refreshExpiresAt))
		claims.ExpiresAt = jwt.NewNumericDate(t)
		tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
		rt, err := tokenClaims.SignedString([]byte(j.refreshSecret))
		if err != nil {
			return nil, err
		}
		token.RefreshToken = rt
	}

	// IssueAccessToken generate accees token used for authentication
	claims.TokenID = uuid.New().String()
	t := time.Now().Add(time.Duration(j.accessExpiresAt))
	claims.ExpiresAt = jwt.NewNumericDate(t)
	tc := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	at, err := tc.SignedString([]byte(j.accessSecret))
	if err != nil {
		return nil, err
	}
	token.AccessToken = at

	token.AccessExpiresAt = j.accessExpiresAt
	token.RefreshExpiresAt = j.refreshExpiresAt
	return token, nil

}

func (j *jwtService) ValidateRefreshToken(tokenString string) (jwt.MapClaims, error) {
	hmacSecret := []byte(j.refreshSecret)
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return hmacSecret, nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims, nil
	}

	return nil, utils.ErrInvalidAuthToken
}

func (j *jwtService) ValidateAccessToken(tokenString string) (jwt.MapClaims, error) {
	hmacSecret := []byte(j.accessSecret)
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return hmacSecret, nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims, nil
	}

	return nil, utils.ErrInvalidAuthToken
}
