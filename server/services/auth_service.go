package services

import (
	"crypto/rand"
	"encoding/hex"
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/bokoness/lashon/database"
	"github.com/bokoness/lashon/models"
	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
)

var jwtSecret = []byte(os.Getenv("JWT_SECRET"))

func GenerateAuthJWT(uid uint) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)

	claims["authorized"] = true
	claims["user"] = uid
	claims["exp"] = time.Now().Add(time.Minute * 30).Unix()

	tokenString, err := token.SignedString(jwtSecret)

	if err != nil {
		return "", fmt.Errorf("something went wrong: %s", err.Error())
	}

	return tokenString, nil
}

func UnHashAuthJWT(hash string) (uint, error) {
	token, err := jwt.Parse(hash, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(jwtSecret), nil
	})

	if err != nil {
		return 0, errors.New("token is not valid")
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		userFloat, ok := claims["user"].(float64)
		if !ok {
			return 0, errors.New("user claim is not a float64")
		}

		userId := uint(userFloat)
		return userId, nil
	} else {
		return 0, errors.New("token is not valid")
	}
}

func GetUserByJWT(hash string) (*models.User, error) {
	uid, err := UnHashAuthJWT(hash)

	if err != nil {
		return nil, errors.New("token is not valid")
	}

	user := GetUser(uid)

	return &user, nil
}

func SaveUserInSession(c *fiber.Ctx, user models.User) error {
	sess, err := GetStore(c)

	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError)
	}

	sess.Set("user", user)

	if err = sess.Save(); err != nil {
		return fiber.NewError(fiber.StatusInternalServerError)
	}

	return nil
}

func GetUserFromSession(c *fiber.Ctx) (*models.User, error) {
	sess, err := GetStore(c)

	if err != nil {
		return nil, errors.New("cant get session store")
	}

	uInterface := sess.Get("user")

	user, ok := uInterface.(*models.User)

	if !ok {
		return nil, errors.New("cant fetch user from session")
	}

	return user, nil
}

func GenerateRefreshToken(userID uint) (string, error) {
	tokenBytes := make([]byte, 32)
	if _, err := rand.Read(tokenBytes); err != nil {
		return "", err
	}
	token := hex.EncodeToString(tokenBytes)

	refreshToken := models.Sessions{
		UserID:       uint(userID),
		RefreshToken: token,
		ExpiresAt:    time.Now().Add(time.Hour * 24 * 30), // Set your desired expiration time
	}

	if err := database.DB.Create(&refreshToken).Error; err != nil {
		log.Error("Error saving refresh token", err)
		return "", fiber.NewError(fiber.StatusInternalServerError)
	}

	return token, nil

}
