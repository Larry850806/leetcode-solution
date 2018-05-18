package main

import (
	"bytes"
	"strings"
)

func reverse(s string) string {
	var reversedBuffer bytes.Buffer

	length := len(s)
	for i := range s {
		reversedBuffer.WriteByte(s[length-i-1])
	}

	return reversedBuffer.String()
}

// "Let's take LeetCode contest"
// -> ["Let's", "take", "LeetCode", "contest"]
// -> ["s'teL", "ekat", "edoCteeL", "tsetnoc"]
// -> "s'teL ekat edoCteeL tsetnoc"
func reverseWords(s string) string {
	words := strings.Split(s, " ")
	for i, word := range words {
		words[i] = reverse(word)
	}
	return strings.Join(words, " ")
}
