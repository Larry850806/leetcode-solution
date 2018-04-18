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
 * Complete the miniMaxSum function below.
 */
func miniMaxSum(arr []int32) {
	var (
		min = arr[0]
		max = arr[0]
		sum int64
	)

	for _, e := range arr {
		sum += int64(e)
		if e > max {
			max = e
		}
		if e < min {
			min = e
		}
	}
	fmt.Printf("%d %d", sum-int64(max), sum-int64(min))
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 1024*1024)

	arrTemp := strings.Split(readLine(reader), " ")

	var arr []int32

	for arrItr := 0; arrItr < 5; arrItr++ {
		arrItemTemp, err := strconv.ParseInt(arrTemp[arrItr], 10, 64)
		checkError(err)
		arrItem := int32(arrItemTemp)
		arr = append(arr, arrItem)
	}

	miniMaxSum(arr)
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
