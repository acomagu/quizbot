package main

import ()

func reply(text string, userID string) error {
	var err error
	userIDs := []string{userID}

	question := []string{
		"もんだぃ。ゎたしゎなんさぃ?",
		"1. 14さぃ",
		"2. 24さぃ",
		"3. 64さぃ",
	}

	err = sendTexts(userIDs, question)
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
