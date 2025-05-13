package main

import (
	Server "Eventify-API/internal/api/server"
	PostgresDb "Eventify-API/repository/db/postgres"
)

func main() {
	PostgresDb.InitDB()

	Server.SetupApi()
}
