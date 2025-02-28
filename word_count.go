package main

import "strings"

type WordCountMap map[string]int

func WordCount(s string) WordCountMap {
	words := strings.Fields(s)
	wordCountMap := make(WordCountMap)

	for _, word := range words {
		count := wordCountMap[word]
		wordCountMap[word] = count + 1
	}

	return wordCountMap
}
