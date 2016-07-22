package main

import (
	"fmt"
	"os"
	"strconv"
	"net/http"
	"github.com/line/line-bot-sdk-go/linebot"
)

var bot *linebot.Client

func main() {
	var err error
	configProxyURL()

	port := os.Getenv("PORT")

	bot, err = lineClient()
	if err != nil {
		fmt.Println(err)
		return
	}

	http.HandleFunc("/callback", handleRequest)
	http.ListenAndServe(":"+port, nil)
}

func handleRequest(w http.ResponseWriter, req *http.Request) {
	fmt.Println(req.Body)

	received, err := bot.ParseRequest(req)
	if err != nil {
		fmt.Println(err)
		return
	}

	for _, result := range received.Results {
		content := result.Content()
		if content != nil && content.IsMessage && content.ContentType == linebot.ContentTypeText {
			text, err := content.TextContent()
			if err != nil {
				fmt.Println(err)
				continue
			}
			fmt.Println(text.Text)

			_, err = bot.SendText([]string{content.From}, "Hello, "+text.Text)
			if err != nil {
				fmt.Println(err)
			}
		}
	}
}

func configProxyURL() {
	os.Setenv("HTTP_PROXY", os.Getenv("FIXIE_URL"))
	os.Setenv("HTTPS_PROXY", os.Getenv("FIXIE_URL"))
}

func lineClient() (*linebot.Client, error) {
	lineChannelID, err := strconv.Atoi(os.Getenv("LINE_CHANNEL_ID"))
	if err != nil {
		return nil, err
	}
	lineChannelSecret := os.Getenv("LINE_CHANNEL_SECRET")
	lineMID := os.Getenv("LINE_MID")

	bot, err := linebot.NewClient(int64(lineChannelID), lineChannelSecret, lineMID)
	return bot, err
}
