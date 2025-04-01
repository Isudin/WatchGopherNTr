package main

import (
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		println("could not load the environmental variables")
		println("starting with default variables")
	}

}
