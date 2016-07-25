package main

import (
	"math/rand"
)

type QA struct {
	question []Message
	answer   string
}

var qas []QA = []QA{
	QA{
		question: []Message{
			"もんだぃ。ゎたしゎなんさぃ?",
			"1. 14さぃ",
			"2. 24さぃ",
			"3. 64さぃ",
		},
		answer: "3",
	},
	QA{
		question: []Message{
			"こんにちは。僕のラッキーカラーは何でしょう?",
			"1. Blue",
			"2. イエロー☆",
			"3. Red",
		},
		answer: "1",
	},
	QA{
		question: []Message{
			"やっほー",
			"1. うっほー",
			"2. ごっほー",
			"3. サッポー",
		},
		answer: "2",
	},
}

func oneQA() QA {
	return qas[rand.Intn(len(qas))]
}
