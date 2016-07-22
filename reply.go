package main

import ()

var stage int = 0

func reply(text string, userID string) error {
	var err error
	userIDs := []string{userID}
	question := []string{
		"もんだぃだょ。ゎたしゎなんさぃ?",
		"1. 14さぃ",
		"2. 24さぃ",
		"3. 64さぃ",
	}
	answer := "3"
	switch stage {
	case 0:
		err = sendTexts(userIDs, question)
	case 1:
		if text == answer {
			_, err = bot.SendText(userIDs, "せぃかぃ")
		} else {
			err = sendTexts(userIDs, []string{
				"やーいやーいwwwwww",
				"正解は"+answer,
			})
		}
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
