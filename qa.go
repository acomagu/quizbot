package main

import (
	"math/rand"
)

type QA struct {
	question []string
	answer string
}

var qas []QA = []QA{
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

func oneQA() QA {
	return qas[rand.Intn(len(qas))]
}
