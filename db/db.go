package db

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq" //PostgreSQL Driver
	migrate "github.com/rubenv/sql-migrate"
	log "github.com/sirupsen/logrus"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func ConnectDB() *gorm.DB {
	DSN := getConnectionDSN()
	db, err := gorm.Open(postgres.Open(DSN), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})
	if err != nil {
		panic("Exception.ConnectDB: Failed to connect")
	}
	return db
}

func MigrateDB() {
	log.Debug("MigrateDB.Start")

	DSN := getConnectionDSN()
	db, err := sql.Open("postgres", DSN)
	if err != nil {
		log.Fatal("Exception.MigrateDb.OpenConnection:", err)
	}
	defer db.Close()

	migrations := &migrate.FileMigrationSource{
		Dir: "db/migrations",
	}
	_, err = migrate.Exec(db, "postgres", migrations, migrate.Up)
	if err != nil {
		log.Fatal("Exception.MigrateDb.Apply:", err)
	}

	log.Debug("MigrateDB.End")
}

func getConnectionDSN() string {
	return fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=5433 sslmode=disable",
		config.Props.DBAutoURL, config.Props.DBAutoUser, config.Props.DBAutoPass, config.Props.DBAutoName)
}
