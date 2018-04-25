package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

// RmDup remove duplicate element in a sorted array
// [3, 3, 3, 2, 1, 1] -> [3, 2, 1]
// [3] -> [3]
func rmDup(ar []int32) []int32 {
	n := len(ar)
	if n == 0 || n == 1 {
		return ar
	}
	var ar2 []int32
	ar2 = append(ar2, ar[0])
	for i := 1; i < n; i++ {
		if ar[i] != ar[i-1] {
			ar2 = append(ar2, ar[i])
		}
	}
	return ar2
}

// find a minimum index ar[idx] <= el
// if all values in array bigger than el, return len(ar)
// ([100, 50, 40, 20, 10], 5)   -> 5
// ([100, 50, 40, 20, 10], 25)  -> 3
// ([100, 50, 40, 20, 10], 50)  -> 1
// ([100, 50, 40, 20, 10], 120) -> 0
func binaryInsert(ar []int32, el int32) int32 {
	n := len(ar)
	if n == 1 {
		if ar[0] <= el {
			return 0
		}
		return 1
	}

	if n == 2 {
		if ar[0] <= el {
			return 0
		}
		if ar[1] <= el {
			return 1
		}
		return 2
	}

	leftIdx, midIdx, rightIdx := 0, n/2, n-1
	leftVal, midVal, rightVal := ar[leftIdx], ar[midIdx], ar[rightIdx]

	if el >= leftVal {
		return int32(leftIdx)
	}
	if el < rightVal {
		return int32(rightIdx + 1)
	}
	if el == rightVal {
		return int32(rightIdx)
	}
	if el == midVal {
		return int32(midIdx)
	}

	if el > midVal {
		return binaryInsert(ar[leftIdx:midIdx+1], el)
	}
	return int32(midIdx+1) + binaryInsert(ar[midIdx+1:rightIdx+1], el)
}

/*
 * Complete the climbingLeaderboard function below.
 */
func climbingLeaderboard(scores []int32, alice []int32) []int32 {
	scores = rmDup(scores)
	dp := make(map[int32]int32)
	result := make([]int32, 0, len(alice))

	for _, a := range alice {
		loc, ok := dp[a]
		if !ok {
			loc = binaryInsert(scores, a) + 1
			dp[a] = loc
		}
		result = append(result, loc)
	}
	return result
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 1024*1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 1024*1024)

	scoresCount, err := strconv.ParseInt(readLine(reader), 10, 64)
	checkError(err)

	scoresTemp := strings.Split(readLine(reader), " ")

	var scores []int32

	for scoresItr := 0; scoresItr < int(scoresCount); scoresItr++ {
		scoresItemTemp, err := strconv.ParseInt(scoresTemp[scoresItr], 10, 64)
		checkError(err)
		scoresItem := int32(scoresItemTemp)
		scores = append(scores, scoresItem)
	}

	aliceCount, err := strconv.ParseInt(readLine(reader), 10, 64)
	checkError(err)

	aliceTemp := strings.Split(readLine(reader), " ")

	var alice []int32

	for aliceItr := 0; aliceItr < int(aliceCount); aliceItr++ {
		aliceItemTemp, err := strconv.ParseInt(aliceTemp[aliceItr], 10, 64)
		checkError(err)
		aliceItem := int32(aliceItemTemp)
		alice = append(alice, aliceItem)
	}

	result := climbingLeaderboard(scores, alice)

	for resultItr, resultItem := range result {
		fmt.Fprintf(writer, "%d", resultItem)

		if resultItr != len(result)-1 {
			fmt.Fprintf(writer, "\n")
		}
	}

	fmt.Fprintf(writer, "\n")

	writer.Flush()
}

func readLine(reader *bufio.Reader) string {
	str, _, err := reader.ReadLine()
	if err == io.EOF {
		return ""
	}

	return strings.TrimRight(string(str), "\r\n")
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
