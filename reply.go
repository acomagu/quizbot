package main

import ()

var stage int = 0

func reply(text string, userID string) error {
	userIDs := []string{userID}
	_, err := bot.SendText(userIDs, "Hello, "+text)
	if err != nil {
		return err
	}
	question := []string{
		"もんだぃ。ゎたしゎなんさぃ?",
		"1. 14さぃ",
		"2. 24さぃ",
		"3. 64さぃ",
	}
	answer := "3"

	switch stage {
	case 0:
		err = sendTexts(userIDs, question)
		stage = 1
	case 1:
		if text == answer {
			bot.SendText(userIDs, "どうしてしってるの?")
		} else {
			err = sendTexts(userIDs, []string{
				"やーいやーいwwwwwww",
				"せぃかぃは"+answer,
			})
		}
		stage = 0
	}
	return err
}

func sendTexts(userIDs []string, texts []string) error {
	for _, text := range texts {
		_, err := bot.SendText(userIDs, text)
		if err != nil {
			return err
		}
	}
	return nil
}
