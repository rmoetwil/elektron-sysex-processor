package main

import (
	"github.com/joho/godotenv"
	"github.com/rmoetwil/elektron-sysex-processor/cmd"
	"log"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}
	cmd.Execute()
}
