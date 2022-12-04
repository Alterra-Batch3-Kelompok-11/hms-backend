package main

import (
	"fmt"
	echoSwagger "github.com/swaggo/echo-swagger"
	"hms-backend/configs"
	"hms-backend/databases"
	_ "hms-backend/docs"
	"hms-backend/routes"
)

func main() {
	fmt.Println("Hello World")
	configs.InitConfig()
	databases.InitDB()
	//
	e := routes.New(databases.DB, echoSwagger.WrapHandler)
	e.Logger.Fatal(e.Start(fmt.Sprintf("%s%s", ":", configs.Cfg.ApiPort)))
}
