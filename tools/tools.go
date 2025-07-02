package tools

import (
	_ "github.com/99designs/gqlgen"
	_ "github.com/go-playground/validator/v10"
	_ "github.com/gofiber/adaptor/v2"
	_ "github.com/gofiber/fiber/v2"
	_ "github.com/golang-jwt/jwt/v5"
	_ "github.com/google/uuid"
	_ "github.com/joho/godotenv"
	_ "github.com/redis/go-redis/v9"
	_ "github.com/stretchr/testify"
	_ "github.com/stretchr/testify/assert"
	_ "golang.org/x/crypto/bcrypt"
	_ "gorm.io/driver/postgres"
	_ "gorm.io/gorm"
)
