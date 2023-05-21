package main

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
)


func healthHandler(c *gin.Context) {
	c.JSON(200, gin.H{
		"status": "running",
	})
}

func loginHandler(c *gin.Context, db *gorm.DB) {
	var user User
	var dbUser User
	c.BindJSON(&user)

	result := db.Where("email = ?", user.Email).First(&dbUser)

	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "E-mail ou senha inválidos"})
		return
	}

	err := bcrypt.CompareHashAndPassword([]byte(dbUser.Password), []byte(user.Password))
	if err != nil {
		if err == bcrypt.ErrMismatchedHashAndPassword {
			c.JSON(http.StatusNotFound, gin.H{"message": "E-mail ou senha inválidos"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"message": "Erro interno do servidor"})
		}
		return
	}


	token := generateSignedToken(dbUser.ID, time.Minute * 15, false)
	refreshToken := generateSignedToken(dbUser.ID, time.Hour * 24 * 30, true)

	c.JSON(http.StatusOK, gin.H{"token": token, "refreshToken": refreshToken})
}

func refreshHandler(c *gin.Context, db *gorm.DB) {
	authHeader := c.GetHeader("Authorization")
	var req RefreshTokenRequest
	c.BindJSON(&req)

	c.ShouldBindHeader(gin.H{"token_data": c.Request.Header["Token"],})

	accessTokenString := extractAuthHeaderToken(authHeader)
	isAccessTokenValid := validateTokenSignature(accessTokenString)

	if accessTokenString == "" || !isAccessTokenValid {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "Parâmetros inválidos"})
		return
	}

	refreshTokenString := req.RefreshToken
	tokenClaims := validateToken(refreshTokenString)

	if tokenClaims == nil {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "Parâmetros inválidos"})
		return
	}

	userId := string(tokenClaims["userId"].(string))

	newTokenString := generateSignedToken(userId, time.Minute * 15, false)

	c.JSON(http.StatusOK, gin.H{"token": newTokenString})
}