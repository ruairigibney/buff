package main

import (
	"log"

	"github.com/ruairigibney/buff/internal/config"
	"github.com/ruairigibney/buff/internal/data/db"
	"github.com/ruairigibney/buff/internal/http"
)

func main() {
	v, err := config.Init(nil)
	if err != nil {
		log.Fatal(err)
	}

	newDB, err := db.NewDB(config.GetDSN(v))
	if err != nil {
		log.Fatal(err)
	}

	http.StartAndServe(newDB)
}
