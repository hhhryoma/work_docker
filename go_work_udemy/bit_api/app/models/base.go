package models

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	"../../config"

	_ "github.com/mattn/go-sqlite3"
)

const (
	tableNameSignalEvents = "signal_events"
)

var DbConnection *sql.DB

func GetCandleTableName(product_code string, duration time.Duration) string {
	return fmt.Sprintf("%s_%s", product_code, duration)
}

func init() {
	var err error
	DbConnection, err = sql.Open(config.Config.SQLDriver, config.Config.DbName)
	if err != nil {
		log.Fatalln(err)
	}
	cmd := fmt.Sprintf(`
        CREATE TABLE IF NOT EXISTS %s (
            time DATETIME PRIMARY KEY NOT NULL,
            product_code STRING,
            side STRING,
            price FLOAT,
            size FLOAT)`, tableNameSignalEvents)
	DbConnection.Exec(cmd)

	for _, duration := range config.Config.Durations {
		tableName := GetCandleTableName(config.Config.ProductCode, duration)
		c := fmt.Sprintf(`
		CREATE TABLE IF NOT EXISTS %s (
			time DATETIME PRIMARY KEY NOT NULL,
			open FLOAT,
			close FLOAT,
			high FLOAT,
			low opne FLOAT, 
			volume FLOAT)`, tableName)
		DbConnection.Exec(c)
	}
}
