package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func max(x, y int64) int64 {
	if x > y {
		return x
	} else {
		return y
	}
}

// Complete the riddle function below.
func riddle(arr []int64) []int64 {
	// complete this function
	n := len(arr)
	ans := make([]int64, n)

	currentMax, posCurrentMax := 0, 0
	for i, elm := range arr {
		ans[0] = max(ans[0], elm)
		if i == 0 {
			continue
		}

		if elm > currentMax {

		}
	}
}

func main() {
	f, err := os.Open("in.txt")
	checkError(f)
	reader := bufio.NewReaderSize(f, 1024*1024)

	stdout, err := os.Create("out.txt")
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 1024*1024)

	nTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
	checkError(err)
	n := int32(nTemp)

	arrTemp := strings.Split(readLine(reader), " ")

	var arr []int64

	for i := 0; i < int(n); i++ {
		arrItem, err := strconv.ParseInt(arrTemp[i], 10, 64)
		checkError(err)
		arr = append(arr, arrItem)
	}

	res := riddle(arr)

	for i, resItem := range res {
		fmt.Fprintf(writer, "%d", resItem)

		if i != len(res)-1 {
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
