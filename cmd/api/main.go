package main

import (
	"Wisdev/internal/database"
	"Wisdev/internal/server"
	"log"

	"github.com/joho/godotenv"
)

func main(){
	err := godotenv.Load()
	if err != nil{
		log.Fatal("error loading .env file ", err)
	}

	err = database.ConnectDB()
	if err != nil{
		log.Fatal("error connecting to database ", err)
	}

	log.Println("successfully connected to postgres")

	server := server.New()

	if err := server.Run(":8080"); err != nil{
		log.Fatalf("error starting server: %v", err)
	}
}