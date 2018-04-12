// https://leetcode.com/problems/jewels-and-stones/description/

package main

func numJewelsInStones(J string, S string) int {
	m := make(map[rune]bool)

	for _, j := range J {
		m[j] = true
	}

	output := 0
	for _, s := range S {
		if m[s] {
			output++
		}
	}

	return output
}
