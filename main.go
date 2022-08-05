package main

import (
	"github.com/Fajar-Islami/ais_code_test/infrastructure/container"
	"github.com/Fajar-Islami/ais_code_test/infrastructure/db"
)

func main() {
	cont := container.New()
	db.SetupDatabaseConnection(*cont.Pgsql)

}
