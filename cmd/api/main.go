package main

import (
	"fd-test/application/config"
	"fd-test/application/database"
	"fd-test/application/infra"
)

func main() {
	config.LoadConfig("cmd/api/.env")

	db := database.NewDB().ConnectPostgres()

	infra := infra.NewInfraFactory()
	infraHttp, err := infra.CreateInfraHttp(config.GetString(config.CFG_APP_PORT, ""), &db)
	if err != nil {
		panic(err)
	}

	infraHttp.Run()

}
