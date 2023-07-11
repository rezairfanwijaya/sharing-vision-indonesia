package main

import (
	"log"
	"svid/database"
)

func main() {
	// koneksi ke database
	dbConnection, err := database.NewConnection("./.env")
	if err != nil {
		log.Fatal(err.Error())
	}

	log.Println(dbConnection)
}
