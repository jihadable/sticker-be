package validators

import (
	"errors"
	"os"
	"strings"

	"github.com/golang-jwt/jwt/v5"
	"github.com/jihadable/sticker-be/services"
)

type AuthHeaderType string

var AuthHeader AuthHeaderType = "AuthHeader"

func AuthValidator(authHeader string, userService services.UserService) (map[string]string, error) {
	if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
		return nil, errors.New("token tidak ditemukan")
	}

	tokenStr := strings.TrimPrefix(authHeader, "Bearer ")
	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (any, error) {
		return []byte(os.Getenv("JWT_SECRET")), nil
	})

	if err != nil || !token.Valid {
		return nil, errors.New("token tidak valid")
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return nil, errors.New("gagal membaca token")
	}

	userId := claims["user_id"].(string)
	role := claims["role"].(string)

	_, err = userService.GetUserById(userId)
	if err != nil {
		return nil, errors.New("pengguna tidak terdaftar")
	}

	return map[string]string{
		"user_id": userId,
		"role":    role,
	}, nil
}

func RoleValidator(authHeader string, userService services.UserService, allowedRole string) (map[string]string, error) {
	credit, err := AuthValidator(authHeader, userService)
	if err != nil {
		return nil, err
	}

	if credit["user_id"] != allowedRole {
		return nil, errors.New("peran pengguna tidak diizinkan")
	}

	return credit, nil
}
