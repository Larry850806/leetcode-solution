// https://leetcode.com/problems/unique-morse-code-words/description/

package main

var mosMap = [...]string{".-", "-...", "-.-.", "-..", ".", "..-.", "--.", "....", "..", ".---", "-.-", ".-..", "--", "-.", "---", ".--.", "--.-", ".-.", "...", "-", "..-", "...-", ".--", "-..-", "-.--", "--.."}

func generateMorseCode(s string) string {
	result := ""
	for _, r := range s {
		result += mosMap[r-'a']
	}
	return result
}

func uniqueMorseRepresentations(words []string) int {
	m := make(map[string]bool)

	for _, w := range words {
		mos := generateMorseCode(w)
		m[mos] = true
	}

	return len(m)
}
