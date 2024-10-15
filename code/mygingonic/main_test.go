package main

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestGetUsers(t *testing.T) {
	// Configurer Gin en mode test
	r := gin.Default()

	// Route de test
	r.GET("/users", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"users": users,
		})
	})

	// Créer une requête HTTP
	req, _ := http.NewRequest("GET", "/users", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	// Assertions
	assert.Equal(t, http.StatusOK, w.Code)
	assert.Contains(t, w.Body.String(), "Alice")
}

func TestPostUser(t *testing.T) {
	// Configurer Gin en mode test
	r := gin.Default()

	// Route de test
	r.POST("/users", func(c *gin.Context) {
		var newUser string
		if err := c.BindJSON(&newUser); err != nil || newUser == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user data"})
			return
		}
		users = append(users, newUser)
		c.JSON(http.StatusCreated, gin.H{"user": newUser})
	})

	// Créer une requête HTTP
	user := `"David"`
	req, _ := http.NewRequest("POST", "/users", bytes.NewBufferString(user))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	// Assertions
	assert.Equal(t, http.StatusCreated, w.Code)
	assert.Contains(t, w.Body.String(), "David")
}
