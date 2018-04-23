package main

import "fmt"

func abs(x int) int {
	if x > 0 {
		return x
	}
	return -x
}

func rotate(m [9]int) [9]int {
	m[0], m[1], m[2], m[3], m[5], m[6], m[7], m[8] = m[6], m[3], m[0], m[7], m[1], m[8], m[5], m[2]
	return m
}

func mirror(m [9]int) [9]int {
	m[0], m[2], m[3], m[5], m[6], m[8] = m[2], m[0], m[5], m[3], m[8], m[6]
	return m
}

func sub(m1, m2 [9]int) (sum int) {
	for i := 0; i < 9; i++ {
		sum += abs(m1[i] - m2[i])
	}
	return
}

func main() {
	correctMat := [9]int{8, 3, 4, 1, 5, 9, 6, 7, 2}
	var inputMat [9]int

	for i := 0; i < 9; i++ {
		fmt.Scanf("%d", &inputMat[i])
	}

	minResult := 81
	for i := 0; i < 4; i++ {
		correctMat = rotate(correctMat)
		for j := 0; j < 2; j++ {
			correctMat = mirror(correctMat)
			result := sub(correctMat, inputMat)
			if minResult > result {
				minResult = result
			}
		}
	}

	fmt.Println(minResult)
}

// Ref: https://baike.baidu.com/item/%E4%B8%89%E9%98%B6%E5%B9%BB%E6%96%B9
