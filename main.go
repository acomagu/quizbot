package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("Hello, world!")
	fmt.Println("I'm in heroku now!")

	lineChannelID := os.Getenv("LINE_CHANNEL_ID")
	lineChannelSecret := os.Getenv("LINE_CHANNEL_SECRET")
	lineMID := os.Getenv("LINE_MID")

	fmt.Println(lineChannelID)
	fmt.Println(lineChannelSecret)
	fmt.Println(lineMID)
}
