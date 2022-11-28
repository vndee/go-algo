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
 * Complete the 'countInversions' function below.
 *
 * The function is expected to return a LONG_INTEGER.
 * The function accepts INTEGER_ARRAY arr as parameter.
 */

func swap(x, y *int32) {
	*x, *y = *y, *x
}

func merge(arr []int32) ([]int32, int64) {
	if len(arr) == 1 {
		return arr, 0
	}

	if len(arr) == 2 {
		if arr[0] > arr[1] {
			swap(&arr[0], &arr[1])
			return arr, 1
		} else {
			return arr, 0
		}
	}

	m := len(arr) >> 1

	left, cLeft := merge(arr[0 : m+1])
	right, cRight := merge(arr[m+1 : len(arr)])

	var cMerge int64 = 0
	var newArr = make([]int32, len(left)+len(right))

	var c, l, r int = 0, 0, 0

	for l < len(left) && r < len(right) {
		if left[l] <= right[r] {
			newArr[c] = left[l]
			l++
		} else {
			newArr[c] = right[r]
			r++

			cMerge += int64(len(left) - l)
		}

		c++
	}

	for l < len(left) {
		newArr[c] = left[l]
		c++
		l++
	}

	for r < len(right) {
		newArr[c] = right[r]
		c++
		r++
	}

	ans := cMerge + cLeft + cRight
	return newArr, ans
}

func countInversions(arr []int32) int64 {
	// Write your code here
	_, ans := merge(arr)
	return ans
}

func main() {
	f, err := os.Open("in.txt")
	checkError(err)
	reader := bufio.NewReaderSize(f, 16*1024*1024)

	stdout, err := os.Create("out.txt")
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 16*1024*1024)

	tTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	checkError(err)
	t := int32(tTemp)

	for tItr := 0; tItr < int(t); tItr++ {
		nTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
		checkError(err)
		n := int32(nTemp)

		arrTemp := strings.Split(strings.TrimSpace(readLine(reader)), " ")

		var arr []int32

		for i := 0; i < int(n); i++ {
			arrItemTemp, err := strconv.ParseInt(arrTemp[i], 10, 64)
			checkError(err)
			arrItem := int32(arrItemTemp)
			arr = append(arr, arrItem)
		}

		result := countInversions(arr)

		fmt.Fprintf(writer, "%d\n", result)
	}

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
