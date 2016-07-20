package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("Hello, world!")
	fmt.Println("I'm in heroku now!")

	hoge := os.Getenv("HOGE")
	fmt.Println(hoge)
}
