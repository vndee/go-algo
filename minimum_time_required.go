package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

// Complete the minTime function below.
func minTime(machines []int64, goal int64) int64 {
	var l, r, m int64 = 0, int64(1e13), int64(0)

	for l < r {
		m = (l + r) >> 1

		// check if m days is valid
		var producedItems int64 = int64(0)
		for _, machine := range machines {
			producedItems += int64(m / machine)
		}

		if producedItems >= goal {
			r = m
		} else {
			l = m + 1
		}
	}

	return r
}

func main() {
	f, err := os.Open("in.txt")
	checkError(err)
	reader := bufio.NewReaderSize(f, 1024*1024)

	stdout, err := os.Create("out.txt")
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 1024*1024)

	nGoal := strings.Split(readLine(reader), " ")

	nTemp, err := strconv.ParseInt(nGoal[0], 10, 64)
	checkError(err)
	n := int32(nTemp)

	goal, err := strconv.ParseInt(nGoal[1], 10, 64)
	checkError(err)

	machinesTemp := strings.Split(readLine(reader), " ")

	var machines []int64

	for i := 0; i < int(n); i++ {
		machinesItem, err := strconv.ParseInt(machinesTemp[i], 10, 64)
		checkError(err)
		machines = append(machines, machinesItem)
	}

	ans := minTime(machines, goal)

	fmt.Fprintf(writer, "%d\n", ans)

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
