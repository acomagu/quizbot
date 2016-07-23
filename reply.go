package main

import (
	"math/rand"
)

type QA struct {
	question []string
	answer string
}

type State struct {
	qa QA
	stage int
}

var states map[string]State = make(map[string]State)

func reply(text string, userID string) error {
	var err error

	userIDs := []string{userID}

	state, ok := states[userID]
	if !ok {
		state = initialState()
	}

	switch state.stage {
	case 0:
		state.qa = oneQA()
		err = sendTexts(userIDs, state.qa.question)
		state.stage = 1
	case 1:
		if text == state.qa.answer {
			_, err = bot.SendText(userIDs, "なんで知ってるの...?")
		} else {
			sendTexts(userIDs, []string{
				"やーいやーーいwwwwwwwwwwwwwwwwwww",
				"せぃかぃゎ"+state.qa.answer,
			})
		}
		state.stage = 0
	}

	states[userID] = state

	return err
}

func initialState() State {
	return State{
		qa: oneQA(),
		stage: 0,
	}
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
				"こんにちは。僕のラッキーカラーは何でしょう?",
				"1. Blue",
				"2. イエロー☆",
				"3. Red",
			},
			answer: "1",
		},
		QA{
			question: []string{
				"やっほー",
				"1. うっほー",
				"2. ごっほー",
				"3. サッポー",
			},
			answer: "2",
		},
	}
	return qas[rand.Intn(len(qas))]
}
