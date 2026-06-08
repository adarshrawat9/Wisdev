package main

import (
	"Wisdev/internals/database"
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
}