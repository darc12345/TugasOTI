package service

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func verifyToken(tokenString string) (*jwt.Token, error) {
	// Parse the token with the secret key
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return secretKey, nil
	})

	// Check for verification errors
	if err != nil {
		log.Printf("Failed to verify token. Original Error: %s", err)
		return nil, fmt.Errorf("token verification error")
	}

	// Check if the token is valid
	if !token.Valid {
		log.Printf("Invalid token detected.")
		return nil, fmt.Errorf("invalid token provided")
	}

	return token, nil
}

func CheckAdmin(tokenString string) error {
	// Verify the token
	token, err := verifyToken(tokenString)
	if err != nil {
		log.Printf("Token verification failed. Original Error: %s", err)
		return fmt.Errorf("unauthorized access")
	}

	// Extract claims
	claim := token.Claims.(jwt.MapClaims)
	role, roleOk := claim["role"].(string)
	expiration, expOk := claim["expiration"].(float64)

	// Validate claims
	if !roleOk || !expOk || role != "ADMIN" || expiration < float64(time.Now().Unix()) {
		log.Printf("Access denied. Role: %v, Expiration: %v", role, expiration)
		return fmt.Errorf("unauthorized access or expired token")
	}

	return nil
}

func (h *ServiceDB) PostProductService(c *gin.Context) error {
	// Retrieve token from cookies
	tokenString, err := c.Cookie("login-token")
	if err != nil {
		log.Printf("Failed to retrieve login token from cookies. Original Error: %s", err)
		return fmt.Errorf("error retrieving login token")
	}

	// Check admin authorization
	err = CheckAdmin(tokenString)
	if err != nil {
		log.Printf("Admin check failed during product posting. Original Error: %s", err)
		c.Redirect(http.StatusSeeOther, "login")
		return fmt.Errorf("unauthorized access")
	}

	return h.RepoDB.PostProductRepo(c)
}

func (h *ServiceDB) PutProductService(c *gin.Context) error {
	// Retrieve token from cookies
	tokenString, err := c.Cookie("login-token")
	if err != nil {
		log.Printf("Failed to retrieve login token from cookies. Original Error: %s", err)
		return fmt.Errorf("error retrieving login token")
	}

	// Check admin authorization
	err = CheckAdmin(tokenString)
	if err != nil {
		log.Printf("Admin check failed during product update. Original Error: %s", err)
		c.Redirect(http.StatusSeeOther, "login")
		return fmt.Errorf("unauthorized access")
	}

	return h.RepoDB.PutProductRepo(c)
}

func (h *ServiceDB) DeleteProductService(c *gin.Context) error {
	// Retrieve token from cookies
	tokenString, err := c.Cookie("login-token")
	if err != nil {
		log.Printf("Failed to retrieve login token from cookies. Original Error: %s", err)
		return fmt.Errorf("error retrieving login token")
	}

	// Check admin authorization
	err = CheckAdmin(tokenString)
	if err != nil {
		log.Printf("Admin check failed during product deletion. Original Error: %s", err)
		c.Redirect(http.StatusSeeOther, "login")
		return fmt.Errorf("unauthorized access")
	}

	return h.RepoDB.DeleteProductRepo(c)
}
