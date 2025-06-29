package helpers

import (
	"errors"
	"os"
	"strings"

	"github.com/golang-jwt/jwt/v5"
)

func Auth(authHeader string) (map[string]string, error) {
	if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
		return map[string]string{}, errors.New("token tidak ditemukan")
	}

	tokenStr := strings.TrimPrefix(authHeader, "Bearer ")
	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (any, error) {
		return []byte(os.Getenv("JWT_SECRET")), nil
	})

	if err != nil || !token.Valid {
		return map[string]string{}, errors.New("token tidak valid")
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return map[string]string{}, errors.New("gagal membaca token")
	}

	userId := claims["user_id"].(string)

	// _, err = userService.GetUserById(userId)
	// if err != nil {
	// 	return fiber.NewError(fiber.StatusUnauthorized, "Pengguna tidak terdaftar")
	// }
	return map[string]string{
		"user_id": userId,
	}, nil
}
