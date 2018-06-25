package main

import (
	"fmt"
	"os"
)

var greeting string

func init() {
	glang := os.Getenv("G4O_LANG")
	switch glang {
	case "de":
		greeting = "Hallo und willkommen!"
	case "nl":
		greeting = "Hallo en welkom!"
	default:
		greeting = "Let's Go!"
	}
}

func main() {
	fmt.Println(greeting)
	fmt.Println("👋")
}
