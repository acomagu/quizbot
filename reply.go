package main

import ()

func reply(text string, userID string) error {
	_, err := bot.SendText([]string{userID}, "Hello, "+text)
	return err
}
