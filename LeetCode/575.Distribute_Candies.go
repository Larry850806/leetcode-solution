package main

func min(a int, b int) int {
	if a < b {
		return a
	}
	return b
}

// [1, 1, 2, 2, 3, 3] -> 3 kind -> min(3, 6/2) = 3
// [1, 1, 2, 3] -> 3 kind -> min(3, 4/2) = 2
// [1, 1, 1, 1, 1, 1] -> 1 kind -> min(1, 6/2) = 1
func distributeCandies(candies []int) int {
	m := map[int]bool{}
	for _, candy := range candies {
		m[candy] = true
	}
	return min(len(m), len(candies)/2)
}

// func distributeCandies(candies []int) int {
// 	arr := [200001]bool{false}
// 	k := 0
// 	for _, v := range candies {
// 		if !arr[100000+v] {
// 			k++
// 			arr[100000+v] = true
// 		}
// 		if k == len(candies)/2 {
// 			return k
// 		}
// 	}

// 	return k
// }
