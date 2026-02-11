package main

import (
	"strings"
)

func cleanInput(text string) []string {
	cleanedTextSlice := strings.Fields(strings.ToLower(text))
	return cleanedTextSlice
}

