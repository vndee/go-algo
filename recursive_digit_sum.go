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
 * Complete the 'superDigit' function below.
 *
 * The function is expected to return an INTEGER.
 * The function accepts following parameters:
 *  1. STRING n
 *  2. INTEGER k
 */

func calcSum(x int64) int64 {
	newSum := int64(0)
	for x/10 > 0 {
		newSum += x % 10
		x /= 10
	}
	newSum += x
	return newSum
}

func superDigit(n string, k int32) int32 {
	// calc sum k by k times
	sum := int64(0)
	for _, ch := range n {
		sum += int64(ch - 48)
	}

	sum *= int64(k)

	for sum/10 > 0 {
		runningSum := calcSum(sum)
		sum = runningSum
	}

	return int32(sum)
}

func main() {
	f, err := os.Open("in.txt")
	reader := bufio.NewReaderSize(f, 16*1024*1024)

	stdout, err := os.Create("out.txt")
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 16*1024*1024)

	firstMultipleInput := strings.Split(strings.TrimSpace(readLine(reader)), " ")

	n := firstMultipleInput[0]

	kTemp, err := strconv.ParseInt(firstMultipleInput[1], 10, 64)
	checkError(err)
	k := int32(kTemp)

	result := superDigit(n, k)

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
