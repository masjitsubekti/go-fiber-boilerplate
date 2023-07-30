package database

import (
	"os"

	"github.com/jmoiron/sqlx"
	"gorm.io/gorm"
)

type DBConn struct {
	DB   *sqlx.DB
	Gorm *gorm.DB
}

// OpenDBConnection func for opening database connection.
func CreateDBConnection() (dbConn DBConn, err error) {
	// Define Database connection variables.
	var (
		db   *sqlx.DB
		gorm *gorm.DB
	)

	// Get DB_TYPE value from .env file.
	dbType := os.Getenv("DB_TYPE")

	// Define a new Database connection with right DB type.
	switch dbType {
	case "pgx":
		db, err = PostgreSQLConnection()
		if err != nil {
			os.Exit(1)
		}

		gorm, err = GormPostgreSQLConnection()
		if err != nil {
			os.Exit(1)
		}

	case "mysql":
		db, err = MysqlConnection()
		if err != nil {
			os.Exit(1)
		}

		gorm, err = GormMysqlConnection()
		if err != nil {
			os.Exit(1)
		}
	}

	dbConn = DBConn{
		DB:   db,
		Gorm: gorm,
	}

	return dbConn, nil
}
