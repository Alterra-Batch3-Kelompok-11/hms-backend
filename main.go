package main

import (
	"fmt"
	"hms-backend/configs"
	"hms-backend/databases"
	"hms-backend/routes"
)

func main() {
	databases.InitDB()
	configs.InitConfig()

	e := routes.New(databases.DB)
	e.Logger.Fatal(e.Start(fmt.Sprintf("%s%s", ":", configs.Cfg.ApiPort)))
}
