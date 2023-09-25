package main

import (
	"crypto/sha256"
	"lab-seguridad/util"
	"log"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	util.Hasher = sha256.New()
	util.DB_Connection = Connect_Database()

	server := Setup_Server()
	server.Logger.Fatal(server.Start(":1323"))
}
