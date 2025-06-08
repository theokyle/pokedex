package main

import (
	"strings"
)

func cleanInput(text string) []string {
	split_text := strings.Split(text, " ")
	var clean_text []string
	for _, word := range split_text {
		if word != "" {
			clean_text = append(clean_text, strings.ToLower(strings.TrimSpace(word)))
		}
	}
	return clean_text
}
