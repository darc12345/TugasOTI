package service

import (
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

var secretKey = []byte(os.Getenv("SECRETKEY"))

// CreateToken generates a signed JWT for a user with role and userid.
func CreateToken(role string, userid string, secretKey []byte) (string, error) {
	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"subject":    "user-credential",
		"role":       role,
		"userid":     userid,
		"expiration": time.Now().Add(time.Hour * 2).Unix(),
	})
	signedToken, err := claims.SignedString(secretKey)
	if err != nil {
		log.Printf("Error creating token: %v", err)
		return "", fmt.Errorf("failed to create token")
	}
	return signedToken, nil
}

// PostLoginService handles user login, validates credentials, and sets a JWT token as a cookie.
func (h *ServiceDB) PostLoginService(c *gin.Context) error {
	loginInfo := struct {
		Userid   string `json:"userid"`
		Password string `json:"password"`
	}{}
	if err := c.ShouldBindJSON(&loginInfo); err != nil {
		log.Printf("Error binding login data: %v", err)
		return fmt.Errorf("invalid input")
	}

	hash := sha256.New()
	hash.Write([]byte(loginInfo.Password))
	passwordhash := base64.StdEncoding.EncodeToString(hash.Sum(nil))

	userid, role, err := h.RepoDB.PostLoginRepo(loginInfo.Userid, passwordhash)
	if err != nil {
		log.Printf("Database error during login: %v", err)
		return fmt.Errorf("internal server error")
	}

	if userid == "" || role == "" {
		log.Println("Invalid credentials provided.")
		return fmt.Errorf("credential not found")
	}

	signedToken, err := CreateToken(role, userid, secretKey)
	if err != nil {
		log.Printf("Error creating JWT token: %v", err)
		return fmt.Errorf("failed to generate token")
	}

	c.SetCookie("login-token", signedToken, 60*60*2, "/", "localhost", false, true)
	return nil
}

// PostRegisterService handles user registration and stores their information in the database.
func (h *ServiceDB) PostRegisterService(c *gin.Context) error {
	regInfo := struct {
		Email    string `json:"email"`
		Address  string `json:"address"`
		Password string `json:"password"`
	}{}
	if err := c.ShouldBindJSON(&regInfo); err != nil {
		log.Printf("Error binding registration data: %v", err)
		return fmt.Errorf("invalid input")
	}

	hash := sha256.New()
	hash.Write([]byte(regInfo.Password))
	passwordhash := base64.StdEncoding.EncodeToString(hash.Sum(nil))

	if err := h.RepoDB.PostRegisterRepo(regInfo.Email, regInfo.Address, passwordhash); err != nil {
		log.Printf("Error during registration database operation: %v", err)
		return fmt.Errorf("failed to register user")
	}

	return nil
}
