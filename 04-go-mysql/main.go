package main

import (
	"fmt"
	"gomysql/db"
	"gomysql/models"
)

func main() {
	db.Connect()
	user := models.CreateUser("alejo", "alejo123", "alejo@gmail.com")
	// user := models.GetUser(2)
	fmt.Println(user)
	// user.Delete()

	// db.TruncateTable("users")
	fmt.Println(models.ListUsers())

	// fmt.Println(db.ExistTable("users"))
	// db.CreateTable(models.UserSchema, "users")
	// db.TruncateTable("users")
	db.Close()
}
