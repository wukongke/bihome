package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func init() {
	mode := flag.String("mode", "dev", "app run mode")
	flag.Parse()
	if *mode == "dev" || *mode == "test" {
		err := godotenv.Load("./.env")
		if err != nil {
			log.Fatal("Error loading .env file")
		}
		mode := os.Getenv("APP_MODE")
		fmt.Println("mode: ", mode)
	}
}
