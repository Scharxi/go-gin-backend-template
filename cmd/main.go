package main

import (
	"database/sql"
	"github.com/BxfferOverflow/go-gin-template/routes"
	"github.com/BxfferOverflow/go-gin-template/util"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"log"
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
	conn, err := sql.Open("postgres", "postgresql://postgres:lms@localhost:5432/postgres?sslmode=disable")

	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}

	util.InitStore(conn)
}
