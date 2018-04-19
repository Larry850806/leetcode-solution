package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func gcd(arr []int32) int32 {
	n := len(arr)
	if n == 1 {
		return arr[0]
	}
	if n == 2 {
		a := arr[0]
		b := arr[1]
		if b == 0 {
			return a
		}
		return gcd([]int32{b, a % b})
	}
	return gcd([]int32{gcd(arr[0 : n/2]), gcd(arr[n/2 : n])})
}

func lcm(arr []int32) int32 {
	n := len(arr)
	if n == 1 {
		return arr[0]
	}
	if n == 2 {
		a := arr[0]
		b := arr[1]
		return a * b / gcd([]int32{a, b})
	}
	return lcm([]int32{lcm(arr[0 : n/2]), lcm(arr[n/2 : n])})
}

/*
 * Complete the getTotalX function below.
 */
func getTotalX(a []int32, b []int32) (count int32) {
	lcmA := lcm(a)
	gcdB := gcd(b)
	for i := lcmA; i <= gcdB; i++ {
		if i%lcmA == 0 && gcdB%i == 0 {
			count++
		}
	}
	return
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 1024*1024)

	outputFile, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer outputFile.Close()

	writer := bufio.NewWriterSize(outputFile, 1024*1024)

	nm := strings.Split(readLine(reader), " ")

	nTemp, err := strconv.ParseInt(nm[0], 10, 64)
	checkError(err)
	n := int32(nTemp)

	mTemp, err := strconv.ParseInt(nm[1], 10, 64)
	checkError(err)
	m := int32(mTemp)

	aTemp := strings.Split(readLine(reader), " ")

	var a []int32

	for aItr := 0; aItr < int(n); aItr++ {
		aItemTemp, err := strconv.ParseInt(aTemp[aItr], 10, 64)
		checkError(err)
		aItem := int32(aItemTemp)
		a = append(a, aItem)
	}

	bTemp := strings.Split(readLine(reader), " ")

	var b []int32

	for bItr := 0; bItr < int(m); bItr++ {
		bItemTemp, err := strconv.ParseInt(bTemp[bItr], 10, 64)
		checkError(err)
		bItem := int32(bItemTemp)
		b = append(b, bItem)
	}

	total := getTotalX(a, b)

	fmt.Fprintf(writer, "%d\n", total)

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
