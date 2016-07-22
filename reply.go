package main

import (
	"math/rand"
)

type QA struct {
	question []string
	answer string
}

var stage int = 0

func reply(text string, userID string) error {
	var err error

	userIDs := []string{userID}

	qa := oneQA()

	switch stage {
	case 0:
		err = sendTexts(userIDs, qa.question)
		stage = 1
	case 1:
		if text == qa.answer {
			_, err = bot.SendText(userIDs, "なんで知ってるの...?")
		} else {
			sendTexts(userIDs, []string{
				"やーいやーーいwwwwwwwwwwwwwwwwwww",
				"せぃかぃゎ"+qa.answer,
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

func oneQA() QA {
	qas := []QA{
		QA{
			question: []string{
				"もんだぃ。ゎたしゎなんさぃ?",
				"1. 14さぃ",
				"2. 24さぃ",
				"3. 64さぃ",
			},
			answer: "3",
		},
		QA{
			question: []string{
				"こんにちは。僕の今日のラッキーカラーは?",
				"1. Blue",
				"2. イエロー☆",
				"3. Red",
			},
			answer: "1",
		},
		QA{
			question: []string{
				"やっほー",
				"1. あっほー",
				"2. ぽっぽー",
				"3. ぴっぴー",
			},
			answer: "2",
		},
	}
	return qas[rand.Intn(len(qas))]
}
