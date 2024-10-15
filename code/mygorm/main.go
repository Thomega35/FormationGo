package main

import (
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

type User struct {
	ID       uint `gorm:"primaryKey"`
	Name     string
	Email    string `gorm:"unique"`
	Password string
}

func initDatabase() {
	dsn := "host=127.0.0.1 user=postgres password=root dbname=postgres port=5432"
	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database: ", err)
	}
	fmt.Println("Database connected successfully!")
}

func migrate() {
	DB.AutoMigrate(&User{})
	fmt.Println("Database migrated!")
}

func createUser(name string, email string, password string) {
	user := User{Name: name, Email: email, Password: password}
	DB.Create(&user)
	fmt.Printf("User %s created successfully!\n", user.Name)
}

func getUsers() {
	var users []User
	DB.Find(&users)
	for _, user := range users {
		fmt.Printf("User: %s, Email: %s\n", user.Name, user.Email)
	}
}

func updateUserEmail(id uint, newEmail string) {
	var user User
	DB.First(&user, id)
	user.Email = newEmail
	DB.Save(&user)
	fmt.Printf("User %s updated with new email: %s\n", user.Name, user.Email)
}

func deleteUser(id uint) {
	var user User
	DB.Delete(&user, id)
	fmt.Printf("User with ID %d deleted successfully!\n", id)
}

func main() {
	initDatabase()
	migrate()
	createUser("John Doe", "john@example.com", "securepassword")
	getUsers()
	updateUserEmail(1, "john.doe@newemail.com")
	getUsers()
	deleteUser(1)
}
