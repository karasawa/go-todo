package main

import (
	"fmt"
	"log"
	"main/app/models"
	"main/config"
)

func main() {
	fmt.Println(config.Config.Port)
	fmt.Println(config.Config.SQLDriver)
	fmt.Println(config.Config.DbName)
	fmt.Println(config.Config.LogFile)

	log.Println("test")

	fmt.Println(models.Db)
}