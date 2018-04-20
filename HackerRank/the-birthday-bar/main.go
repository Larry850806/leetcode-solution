package main

import "fmt"

func sum(arr []int) (res int) {
	for _, e := range arr {
		res += e
	}
	return
}

func solve(n int, s []int, d, m int) int {
	count := 0
	for i := 0; i < n-m+1; i++ {
		if sum(s[i:i+m]) == d {
			count++
		}
	}
	return count
}

func main() {
	var n, d, m int
	fmt.Scanf("%d", &n)

	s := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scanf("%d", &s[i])
	}

	fmt.Scanf("%d %d", &d, &m)

	fmt.Print(solve(n, s, d, m))
}
