package main

import (
	"carimurah/engines"
	"carimurah/postgres"
	"fmt"
)

func main() {

	// Init database
	db, err := postgres.New()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	// Init engines
	eng := engines.NewEngine(db)

	eng.DoSearchQuery("samsung")
}
