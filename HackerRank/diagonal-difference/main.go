package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

/*
 * Complete the diagonalDifference function below.
 */
func diagonalDifference(a [][]int32) int32 {
	n := len(a)

	var primarySum, secondarySum int32
	for i := 0; i < n; i++ {
		primarySum += a[i][i]
		secondarySum += a[i][n-i-1]
	}

	if primarySum > secondarySum {
		return primarySum - secondarySum
	}
	return secondarySum - primarySum
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 1024*1024)

	outputFile, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer outputFile.Close()

	writer := bufio.NewWriterSize(outputFile, 1024*1024)

	nTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
	checkError(err)
	n := int32(nTemp)

	var a [][]int32
	for aRowItr := 0; aRowItr < int(n); aRowItr++ {
		aRowTemp := strings.Split(readLine(reader), " ")

		var aRow []int32
		for _, aRowItem := range aRowTemp {
			aItemTemp, err := strconv.ParseInt(aRowItem, 10, 64)
			checkError(err)
			aItem := int32(aItemTemp)
			aRow = append(aRow, aItem)
		}

		if len(aRow) != int(int(n)) {
			panic("Bad input")
		}

		a = append(a, aRow)
	}

	result := diagonalDifference(a)

	fmt.Fprintf(writer, "%d\n", result)

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
