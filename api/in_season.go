package api

import (
	"database/sql"
	"log"
	"os"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

type crops struct {
	Time    string   `json:"time"`
	Produce []string `json:"produce"`
}

func queryInSeason() crops {
	result := crops{}
	sqlQuery := `
	SELECT crop_name 
	FROM season
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
		crop := make([]string, 0)
		if err != nil {
			log.Fatal(err)
		}
		defer rows.Close()

		for rows.Next() {
			var crop_name string
			rows.Scan(&crop_name)
			crop = append(crop, crop_name)
		}
		return crops{
			Time:    time.Now().String(),
			Produce: crop,
		}
	}
	return result
}
