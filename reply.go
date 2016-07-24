package main

import ()

func reply(text string, userID string) error {
	userIDs := []string{userID}
	state := stateBy(userID)
	replies, newState := generateReplies(text, state)
	err := sendTexts(userIDs, replies)
	if err != nil {
		return err
	}
	setStateBy(userID, newState)

	return nil
}

func generateReplies(text string, state State) ([]string, State) {
	var replies []string
	newState := state

	switch state.stage {
	case 0:
		qa := oneQA()
		replies = qa.question
		newState.qa = qa
		newState.stage = 1
	case 1:
		if isCorrectAnswer(text, state.qa) {
			replies = []string{"なんで知ってるの...?"}
		} else {
			replies = []string{
				"やーいやーーいwwwwwwwwwwwwwwwwwww",
				"せぃかぃゎ"+state.qa.answer,
			}
		}
		newState.stage = 0
	}
	return replies, newState
}

func isCorrectAnswer(text string, qa QA) bool {
	return text == qa.answer
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
