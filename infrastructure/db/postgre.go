package db

import (
	"fmt"
	"time"

	"github.com/Fajar-Islami/ais_code_test/entity"
	"github.com/Fajar-Islami/ais_code_test/infrastructure/container"
	"github.com/fatih/color"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// SetupDatabaseConnection is creating a new connection to our database
func SetupDatabaseConnection(cont container.Pgsql) *gorm.DB {
	// sslmode=disable for avoid error pq: SSL is not enabled on the server.
	dbUrl := fmt.Sprintf(`postgres://%s:%s@%s:%d/%s`,
		cont.Username,
		cont.Password,
		cont.Host,
		cont.Port,
		cont.DbName,
	)

	db, err := gorm.Open(postgres.Open(dbUrl), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	psqlDB, err := db.DB()
	if err != nil {
		panic(err.Error())
	}

	// SetMaxIdleConns sets the maximum number of connections in the idle connection pool.
	psqlDB.SetMaxIdleConns(cont.MinIdleConnections)

	// SetMaxOpenConns sets the maximum number of open connections to the database.
	psqlDB.SetMaxOpenConns(cont.MaxOpenConnections)

	// SetConnMaxLifetime sets the maximum amount of time a connection may be reused.
	maxLifeTime := time.Duration(cont.MaxLifetime) * time.Second
	psqlDB.SetConnMaxLifetime(maxLifeTime)

	if err != nil {
		panic("Failed to create a connection to your database")
	}

	color.Green("⇨ Postgre status is connected")

	db.AutoMigrate(&entity.Article{})

	return db

}

func CloseDatabaseConnection(db *gorm.DB) {
	dbSQL, err := db.DB()
	if err != nil {
		panic("Failed to close connection to database")
	}

	dbSQL.Close()

}
