// https://leetcode.com/problems/max-increase-to-keep-city-skyline/description/

package main

func maxIncreaseKeepingSkyline(grid [][]int) int {
	//  grid =
	// 	[ [3, 0, 8, 4],
	//    [2, 4, 5, 7],
	//    [9, 2, 6, 3],
	//    [0, 3, 1, 0] ]
	size := len(grid)

	// colMax = [9, 4, 8, 7]
	// rowMax = [8, 7, 9, 3]
	colMax := make([]int, size)
	rowMax := make([]int, size)
	for i, row := range grid {
		for j, elem := range row {
			if colMax[j] < elem {
				colMax[j] = elem
			}
			if rowMax[i] < elem {
				rowMax[i] = elem
			}
		}
	}

	// gridNew =
	// [ [8, 4, 8, 7],
	//   [7, 4, 7, 7],
	//   [9, 4, 8, 7],
	//   [3, 3, 3, 3] ]
	gridNew := make([][]int, size)
	for i := range gridNew {
		gridNew[i] = make([]int, size)
	}
	for i, row := range gridNew {
		for j := range row {
			if rowMax[i] > colMax[j] {
				gridNew[i][j] = colMax[j]
			} else {
				gridNew[i][j] = rowMax[i]
			}
		}
	}

	// result = sum(gridNew - grid)
	result := 0
	for i, row := range gridNew {
		for j := range row {
			result += gridNew[i][j] - grid[i][j]
		}
	}

	return result
}
