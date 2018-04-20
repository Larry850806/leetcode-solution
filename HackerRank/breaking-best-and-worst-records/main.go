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
 * Complete the breakingRecords function below.
 */
func breakingRecords(score []int32) []int32 {
	var breakBest, breakWorst int32
	var (
		highest = score[0]
		lowest  = score[0]
	)
	for _, e := range score {
		if e > highest {
			highest = e
			breakBest++
			continue
		}
		if e < lowest {
			lowest = e
			breakWorst++
			continue
		}
	}
	return []int32{breakBest, breakWorst}
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

	scoreTemp := strings.Split(readLine(reader), " ")

	var score []int32

	for scoreItr := 0; scoreItr < int(n); scoreItr++ {
		scoreItemTemp, err := strconv.ParseInt(scoreTemp[scoreItr], 10, 64)
		checkError(err)
		scoreItem := int32(scoreItemTemp)
		score = append(score, scoreItem)
	}

	result := breakingRecords(score)

	for resultItr, resultItem := range result {
		fmt.Fprintf(writer, "%d", resultItem)

		if resultItr != len(result)-1 {
			fmt.Fprintf(writer, " ")
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
