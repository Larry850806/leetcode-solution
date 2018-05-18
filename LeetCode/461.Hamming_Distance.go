package main

func countOne(x int) int {
	var count = 0
	for x > 0 {
		if x%2 > 0 {
			count++
		}
		x /= 2
	}
	return count
}

func hammingDistance(x int, y int) int {
	return countOne(x ^ y)
}
