package config

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"time"
)

func NewPostgresSQL(configuration *Config) (*sql.DB, error) {
	username := configuration.DBUsername
	password := configuration.DBPassword
	host := configuration.DBHost
	port := configuration.DBPort
	database := configuration.DBDatabase
	sslMode := configuration.DBSSLMode

	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s", host, port, username, password, database, sslMode)
	db, err := sql.Open(configuration.DBConnection, dsn)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	db, err = databasePooling(configuration, db)
	if err != nil {
		return nil, err
	}

	return db, nil
}

func databasePooling(configuration *Config, db *sql.DB) (*sql.DB, error) {
	// Limit connection with db pooling
	setMaxIdleConns := configuration.DBPoolMin
	setMaxOpenConns := configuration.DBPoolMax
	setConnMaxIdleTime := configuration.DBMaxIdleTimeSecond
	setConnMaxLifetime := configuration.DBMaxLifeTimeSecond

	db.SetMaxIdleConns(setMaxIdleConns)                                    // minimal connection
	db.SetMaxOpenConns(setMaxOpenConns)                                    // maximal connection
	db.SetConnMaxLifetime(time.Duration(setConnMaxIdleTime) * time.Second) // unused connections will be deleted
	db.SetConnMaxIdleTime(time.Duration(setConnMaxLifetime) * time.Second) // connection that can be used

	return db, nil
}
