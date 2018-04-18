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
 * Complete the plusMinus function below.
 */
func plusMinus(arr []int32) {
	n := len(arr)
	var positive, negative, zero int
	for _, v := range arr {
		if v > 0 {
			positive++
		} else if v < 0 {
			negative++
		} else if v == 0 {
			zero++
		}
	}
	fmt.Printf("%v\n%v\n%v", float32(positive)/float32(n), float32(negative)/float32(n), float32(zero)/float32(n))
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 1024*1024)

	nTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
	checkError(err)
	n := int32(nTemp)

	arrTemp := strings.Split(readLine(reader), " ")

	var arr []int32

	for arrItr := 0; arrItr < int(n); arrItr++ {
		arrItemTemp, err := strconv.ParseInt(arrTemp[arrItr], 10, 64)
		checkError(err)
		arrItem := int32(arrItemTemp)
		arr = append(arr, arrItem)
	}

	plusMinus(arr)
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
