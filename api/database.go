package api

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

func CreateDatabase() {
	db, err := sql.Open("sqlite3", "./agriapi.db")
	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	plantingTableCreation := `
	CREATE TABLE IF NOT EXISTS planting (
		id INTEGER PRIMARY KEY,
		crop_name TEXT,
		start_month INTEGER,
		end_month INTEGER
	);
	`

	inSeasonTableCreation := `
	CREATE TABLE IF NOT EXISTS season (
		id INTEGER PRIMARY KEY,
		crop_name TEXT,
		start_month INTEGER,
		end_month INTEGER
	);
	`

	_, err = db.Exec(plantingTableCreation)
	if err != nil {
		log.Fatal(err)
	}

	_, err = db.Exec(inSeasonTableCreation)
	if err != nil {
		log.Fatal(err)
	}

	plant_stmt, err := db.Prepare("INSERT INTO planting(crop_name, start_month, end_month) VALUES (? , ? , ?)")
	if err != nil {
		log.Fatal(err)
	}
	defer plant_stmt.Close()

	season_stmt, err := db.Prepare("INSERT INTO season(crop_name, start_month, end_month) VALUES (? , ? , ?)")
	if err != nil {
		log.Fatal(err)
	}
	defer season_stmt.Close()

	crop_names := [19]string{
		"Green Beans", "Lima Beans", "Beets", "Brussel Sprouts", "Carrots", "Sweet Corn", "Cucumber",
		"Edamame", "Eggplant", "Garlic", "Baby Leaf Lettuce", "Honeydew Melon", "Muskmelon", "Okra",
		"Dry Onions", "Parsnips", "Potatoes", "Pumpkins", "Rutabaga"}

	plant_starts := [19]int{
		5, 5, 4, 5, 4, 4, 4, 6, 4, 10, 3, 4, 4, 4, 3, 4, 4, 4, 6}

	plant_ends := [19]int{
		7, 6, 7, 6, 7, 6, 7, 6, 5, 11, 9, 6, 6, 6, 4, 5, 5, 6, 7}

	season_starts := [19]int{
		7, 8, 6, 9, 6, 7, 8, 8, 8, 7, 4, 8, 8, 7, 8, 10, 7, 8, 10}

	season_ends := [19]int{
		10, 10, 11, 11, 10, 10, 10, 9, 10, 8, 10, 10, 10, 10, 8, 11, 9, 10, 11}

	for i := range 19 {
		_, err = plant_stmt.Exec(crop_names[i], plant_starts[i], plant_ends[i])
		if err != nil {
			log.Fatal(err)
		}

		_, err = season_stmt.Exec(crop_names[i], season_starts[i], season_ends[i])
		if err != nil {
			log.Fatal(err)
		}
	}
}
