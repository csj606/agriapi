package api

import (
	"database/sql"
	"log"
	"os"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

func queryPlantingSeason() crops {
	result := crops{}
	sqlQuery := `
	SELECT crop_name 
	FROM planting
	WHERE month >= (?) AND end_month <= (?)
	`
	if os.Getenv("CLOUD_ENV") != "true" {
		db, err := sql.Open("sqlite3", "./agriapi.db")
		if err != nil {
			log.Fatal(err)
		}
		defer db.Close()
		month := time.Time.Month(time.Now())
		rows, err := db.Query(sqlQuery, month)
		if err != nil {
			log.Fatal(err)
		}
		defer rows.Close()

		for rows.Next() {

		}
	}
	return result
}
