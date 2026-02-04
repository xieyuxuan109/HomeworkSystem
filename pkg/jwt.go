package pkg

import (
	"errors"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var JWT_SECRET string = os.Getenv("JWT_SECRET")

// 自定义 Claims，明确指定类型
type Claims struct {
	UserID     uint   `json:"user_id"`
	Username   string `json:"username"`
	Department string `json:"department"`
	Role       string `json:"role"`
	Type       string `json:"type"`
	jwt.RegisteredClaims
}

// 生成双Token
func GenerateTokens(userID uint, username string, role string, department string) (map[string]string, error) {
	// Access Token
	accessClaims := Claims{
		UserID:     userID,
		Username:   username,
		Role:       role,
		Department: department,
		Type:       "access",
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(15 * time.Minute)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, accessClaims)
	accessTokenStr, err := accessToken.SignedString([]byte(JWT_SECRET))
	if err != nil {
		return nil, err
	}
	// Refresh Token (7天过期)
	refreshClaims := Claims{
		UserID:     userID,
		Username:   username,
		Role:       role,
		Department: department,
		Type:       "refresh",
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(7 * 24 * time.Hour)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}
	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshClaims)
	refreshTokenStr, err := refreshToken.SignedString([]byte(JWT_SECRET))
	if err != nil {
		return nil, err
	}
	return map[string]string{
		"access_token":  accessTokenStr,
		"refresh_token": refreshTokenStr,
	}, nil
}

// 通用验证函数
func VerifyToken(tokenStr, expectedType string) (*Claims, error) {
	token, err := jwt.ParseWithClaims(tokenStr, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(JWT_SECRET), nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(*Claims); ok && token.Valid {
		if claims.Type != expectedType {
			return nil, errors.New("token类型错误")
		}
		return claims, nil
	}
	return nil, errors.New("invalid token")
}
