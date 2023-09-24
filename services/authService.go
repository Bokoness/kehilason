package services

import (
	"errors"
	"fmt"
	"github.com/bokoness/lashon/models"
	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
	"os"
	"time"
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

func UnHashAuthJWT(hash string) (interface{}, error) {
	token, err := jwt.Parse(hash, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(jwtSecret), nil
	})

	if err != nil {
		return nil, errors.New("token not valid")
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims["user"], nil
	} else {
		return nil, errors.New("token is not valid")
	}
}

func GetUserByJWT(hash string) (*models.User, error) {
	result, err := UnHashAuthJWT(hash)

	if err != nil {
		return nil, errors.New("token is not valid")
	}

	uid, _ := result.(uint)

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
