package main

import (
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"log"
	"os"
)

func main() {
	godotenv.Load(".env")
	dumpenv()

	r := gin.Default()
	r.GET("/api/people", handleGetPeople)
	r.POST("/api/people", handleCreatePerson)
	r.GET("/api/visit", handleGetVisit)
	r.POST("/api/visit", handleNewVisit)
	r.Run()
}

func dumpenv() {
	log.Printf("Dumping env vars : ")
	for _, v := range os.Environ() {
		log.Printf("- %s", v)
	}
}
