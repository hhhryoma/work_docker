package main

import (
	"fmt"

	"./app/controllers"
	"./app/models"
	"./config"
	"./utils"
)

func main() {
	utils.LogingSettings(config.Config.LogFile)
	fmt.Println(models.DbConnection)
	controllers.StreamIngestionData()
	controllers.StartWebServer()

}
