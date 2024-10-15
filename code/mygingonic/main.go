// Création de routes et gestion des requêtes
package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

var users = []string{"Alice", "Bob", "Charlie"}

func LoggerMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("Request received: ", c.Request.Method, c.Request.URL)
		c.Next()
	}
}

func main() {
	r := gin.Default()

	// Utiliser le middleware de journalisation
	r.Use(LoggerMiddleware())

	// Définir une route GET
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	// Route pour obtenir tous les utilisateurs
	r.GET("/users", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"users": users,
		})
	})

	// Route pour ajouter un utilisateur
	r.POST("/users", func(c *gin.Context) {
		var newUser string
		if err := c.BindJSON(&newUser); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid data"})
			return
		}
		users = append(users, newUser)
		c.JSON(http.StatusCreated, gin.H{"user": newUser})
	})

	r.GET("/error", func(c *gin.Context) {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Something went wrong"})
	})

	// Démarrer le serveur
	r.Run(":8080")
}
