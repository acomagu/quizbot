package main

import (
	"fmt"
	"os"
	"strconv"
	"net/http"
	"github.com/line/line-bot-sdk-go/linebot"
)

func main() {
	fmt.Println("Hello, world!")
	fmt.Println("I'm in heroku now!")

	lineChannelID, err := strconv.Atoi(os.Getenv("LINE_CHANNEL_ID"))
	if err != nil {
		fmt.Println(err)
	}
	lineChannelSecret := os.Getenv("LINE_CHANNEL_SECRET")
	lineMID := os.Getenv("LINE_MID")
	port := os.Getenv("PORT")

	_, err = linebot.NewClient(int64(lineChannelID), lineChannelSecret, lineMID)
	if err != nil {
		fmt.Println(err)
	}

	http.HandleFunc("/callback", func(w http.ResponseWriter, req *http.Request) {
		fmt.Println(req.Body)
	})
	http.ListenAndServe(":"+port, nil)
}
