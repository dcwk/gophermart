package auth

import (
	"context"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

type Claims struct {
	jwt.RegisteredClaims
	UserID int64
}

type UserIDKey string

// TODO: добавить возможность управлять протуханием из конфига
const TokenExp = time.Hour * 3

// TODO: убрать ключ в конфиг
const SecretKey = "supersecretkey"

func BuildJWTString(userID int64) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, Claims{
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(TokenExp)),
		},
		UserID: userID,
	})

	tokenString, err := token.SignedString([]byte(SecretKey))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func GetUserID(tokenString string) (int64, error) {
	claims := &Claims{}
	token, err := jwt.ParseWithClaims(tokenString, claims,
		func(t *jwt.Token) (interface{}, error) {
			if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", t.Header["alg"])
			}
			return []byte(SecretKey), nil
		})
	if err != nil {
		return 0, fmt.Errorf("failed to parse token: %v", err)
	}

	if !token.Valid {
		fmt.Println("Token is not valid")
		return 0, fmt.Errorf("invalid token")
	}

	fmt.Println("Token os valid")
	return claims.UserID, nil
}

func GetUserIDFromCtx(ctx context.Context) int64 {
	userKey := UserIDKey("userID")
	return ctx.Value(userKey).(int64)
}
