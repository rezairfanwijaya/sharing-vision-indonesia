package main

import (
	"log"
	"svid/database"
	"svid/middleware"
	"svid/router"

	"github.com/gin-gonic/gin"
)

func main() {
	// koneksi ke database
	dbConnection, err := database.NewConnection("./.env")
	if err != nil {
		log.Fatal(err.Error())
	}

	// router
	route := gin.Default()
	route.Use(middleware.CorsMiddleware())
	router.NewRouter(route, dbConnection)

	// start server
	if err := route.Run(":6868"); err != nil {
		log.Fatal("tidak dapat memulai server : ", err.Error())
	}
}
