package main

import (
	"fmt"
	"main/app/controllers"
	"main/app/models"
)

func main() {
	controllers.StartMainServer()

	user, _ := models.GetUserByEmail("test@gmail.com")
	fmt.Println(user)

}