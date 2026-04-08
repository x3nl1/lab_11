package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func register(c *gin.Context) {
	var user User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, user)
}

func calculate(c *gin.Context) {
	var numbers Numbers
	if err := c.ShouldBindJSON(&numbers); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	sum := 0
	for _, n := range numbers.Values {
		sum += n
	}

	c.JSON(http.StatusOK, gin.H{"sum": sum})
}

func token(c *gin.Context) {
	t, err := generateToken()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"token": t})
}

func protected(c *gin.Context) {
	auth := c.GetHeader("Authorization")
	if auth == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "no token"})
		return
	}

	const prefix = "Bearer "
	if len(auth) <= len(prefix) || auth[:len(prefix)] != prefix {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid format"})
		return
	}

	tokenStr := auth[len(prefix):]

	if err := validateToken(tokenStr); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "ok"})
}