package helpers

import (
	"github.com/gofiber/fiber/v2"
	jtoken "github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
	userModels "golang/api/models/user"
	"os"
	"time"
)

func HashPassword(password string) string {
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	return string(hashedPassword)
}

func CompareHashAndPassword(hashedPassword string, password string) (bool, error) {
	matchErr := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))

	if matchErr != nil {
		return false, matchErr
	}

	return true, nil
}

func GenerateJWT(user *userModels.User) (string, error) {
	jwtSecret := os.Getenv("JWT_SECRET")

	day := time.Hour * 24
	claims := jtoken.MapClaims{
		"ID":    user.ID,
		"email": user.Email,
		"exp":   time.Now().Add(day * 1).Unix(),
	}
	token := jtoken.NewWithClaims(jtoken.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(jwtSecret))

	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func ParseToken(c *fiber.Ctx) map[string]interface{} {
	user := c.Locals("user").(*jtoken.Token)
	claims := user.Claims.(jtoken.MapClaims)

	return claims
}

func GetUserId(c *fiber.Ctx) int {
	claims := ParseToken(c)
	userIdFloat := claims["ID"].(float64)
	userId := int(userIdFloat)

	return userId
}
