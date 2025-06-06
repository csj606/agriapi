package api

// "github.com/jackc/pgx/v5"
// "os"
// "github.com/joho/godotenv"

type crops struct {
	Time    string `json:"time"`
	Produce string `json:"produce"`
}

func queryInSeason() crops {
	result := crops{}
	return result
}
