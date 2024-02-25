package middleware

import (
	"encoding/json"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"github.com/nexryai/pgharmony/internal/core/config"
	"github.com/nexryai/pgharmony/internal/core/logger"
)

type apiRequest struct {
	Token string `json:"token"`
}

func Auth(ctx *fiber.Ctx) error {
	log := logger.GetLogger("AuthMiddleware")

	var tokenString string

	req := apiRequest{}

	err := json.Unmarshal(ctx.Body(), &req)
	if err != nil {
		log.Debug(err.Error())
		return ctx.Status(400).JSON(fiber.Map{
			"error": "Invalid JSON format",
		})
	}

	tokenString = req.Token

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(config.Secret), nil // Verify signing key
	})

	if err != nil || !token.Valid {
		return ctx.Status(403).JSON(fiber.Map{
			"error": "Invalid token. Please re-login.",
		})
	}

	// 色々チェック
	claims := token.Claims.(jwt.MapClaims)
	userId := claims["id"].(string)
	accountToken := claims["accountToken"].(string)

	if userId != "admin" || accountToken == "" {
		panic("JWT token was successfully decrypted but claims are invalid")
	}

	// Store the user information in locals for later use
	ctx.Locals("user", token)

	return ctx.Next()
}
