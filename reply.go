package main

import ()

type Message string
type UserID string

func reply(text Message, userID UserID) error {
	userIDs := []UserID{userID}
	state := stateBy(userID)
	replies, newState := generateReplies(text, state)
	err := sendTexts(userIDs, replies)
	if err != nil {
		return err
	}
	setStateBy(userID, newState)

	return nil
}

func generateReplies(text Message, state State) ([]Message, State) {
	var replies []Message
	newState := state

	switch state.stage {
	case _Ask:
		qa := oneQA()
		replies = qa.question
		newState.qa = qa
		newState.stage = _Answer
	case _Answer:
		if isCorrectAnswer(text, state.qa) {
			replies = []Message{"なんで知ってるの...?"}
		} else {
			replies = []Message{
				"やーいやーーいwwwwwwwwwwwwwwwwwww",
				Message("せぃかぃゎ"+state.qa.answer),
			}
		}
		newState.stage = _Ask
	}
	return replies, newState
}

func isCorrectAnswer(text Message, qa QA) bool {
	return text == Message(qa.answer)
}

func sendTexts(userIDs []UserID, texts []Message) error {
	stringUserIDs := userIDsToStrings(userIDs)
	for _, text := range texts {
		_, err := bot.SendText(stringUserIDs, string(text))
		if err != nil {
			return err
		}
	}
	return nil
}

func userIDsToStrings(orig []UserID) []string {
	var res []string
	for _, userID := range orig {
		res = append(res, string(userID))
	}
	return res
}
