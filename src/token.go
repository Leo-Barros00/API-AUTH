package main

import (
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func generateDynamicSecretKey(isRefreshToken bool) string {
	baseSecretKey := os.Getenv("BASE_SECRET_KEY")
	var secretKey string

	if isRefreshToken {
		secretKey = fmt.Sprintf("%v03b62516184fb6ef591f45bd4974b753", baseSecretKey)
	} else {
		secretKey = fmt.Sprintf("%vc21f969b5f03d33d43e04f8f136e7682", baseSecretKey)
	}

	return secretKey
}

func generateSignedToken(userId string, tokenExpiration time.Duration, isRefreshToken bool) gin.H {
	expiresIn := time.Now().Add(tokenExpiration).Unix()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userId": userId,
		"exp":    expiresIn,
	})

	secretKey := generateDynamicSecretKey(isRefreshToken)
	tokenString, _ := token.SignedString([]byte(secretKey))

	return gin.H{"value": tokenString, "expiresIn": expiresIn}
}

func parseToken(token string) (*jwt.Token, error) {
	secretKey := generateDynamicSecretKey(true)

	return jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		return []byte(secretKey), nil
	})
}

func validateTokenSignature(tokenString string) bool {
	secretKey := generateDynamicSecretKey(false)
	parser := jwt.Parser{
		SkipClaimsValidation: true,
	}

	token, err := parser.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(secretKey), nil
	})

	return err == nil && token.Valid
}

func validateToken(tokenString string) jwt.MapClaims {
	token, err := parseToken(tokenString)

	if err != nil || !token.Valid {
		return nil
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return nil
	}

	return claims
}

func extractAuthHeaderToken(authHeader string) string {
	if authHeader == "" {
		return  ""
	}

	authParts := strings.Split(authHeader, " ")
	if len(authParts) != 2 || authParts[0] != "Bearer" {
		return ""
	}

	return authParts[1]
}