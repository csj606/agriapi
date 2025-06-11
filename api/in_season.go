package api

// "os"

type crops struct {
	Time    string   `json:"time"`
	Produce []string `json:"produce"`
}

func queryInSeason() crops {
	result := crops{}
	return result
}
