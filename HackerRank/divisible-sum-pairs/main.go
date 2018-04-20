package main

import "fmt"

func solve(a []int, k int) (count int) {
	for j := range a {
		for i := 0; i < j; i++ {
			if (a[i]+a[j])%k == 0 {
				count++
			}
		}
	}
	return
}

func main() {
	var n, k int
	fmt.Scanf("%d %d", &n, &k)

	a := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scanf("%d", &a[i])
	}

	fmt.Print(solve(a, k))
}
