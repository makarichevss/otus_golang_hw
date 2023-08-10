package hw03frequencyanalysis

import (
	"sort"
	"strings"
)

func Top10(text string) []string {
	words := strings.Fields(text)

	wordCounts := make(map[string]int, len(words))
	for _, word := range words {
		wordCounts[word]++
	}

	// Удаляем дубликаты.
	uniqueWords := make([]string, 0, len(wordCounts))
	for word := range wordCounts {
		uniqueWords = append(uniqueWords, word)
	}

	// Сортируем слайс.
	sort.Slice(uniqueWords, func(i, j int) bool {
		if wordCounts[uniqueWords[i]] == wordCounts[uniqueWords[j]] {
			return uniqueWords[i] < uniqueWords[j]
		}
		return wordCounts[uniqueWords[i]] > wordCounts[uniqueWords[j]]
	})

	// Если есть более 10 слов, обрезаем слайс.
	if len(uniqueWords) > 10 {
		uniqueWords = uniqueWords[:10]
	}

	return uniqueWords
}
