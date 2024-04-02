package main

import (
	"database/sql"
	"github.com/BxfferOverflow/gogintemplate/routes"
	"github.com/BxfferOverflow/gogintemplate/util"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"log"
	"os"
)

func main() {
	initApp()
	router := gin.Default()
	v1 := router.Group("/api/v1")

	routes.AddSampleRoutes(v1)

	router.Run(":8080")
}

func initApp() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalln(err.Error())
		return
	}
	conn, err := sql.Open("postgres", os.Getenv("DB_URL"))

	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}

	util.InitStore(conn)
}
