package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

// Complete the freqQuery function below.
func freqQuery(queries [][]int32) []int32 {
	ans := make([]int32, 0)
	freqMap := make(map[int32]int32)
	freqOccur := make(map[int32]int32)

	for _, query := range queries {
		switch query[0] {
		case 1:
			freqMap[query[1]]++
			freqOccur[freqMap[query[1]]]++

			if freqOccur[freqMap[query[1]]-1] > 0 {
				freqOccur[freqMap[query[1]]-1]--
			}

		case 2:
			if freqMap[query[1]] == 0 {
				continue
			}
			freqMap[query[1]]--
			freqOccur[freqMap[query[1]]]++

			if freqOccur[freqMap[query[1]]+1] > 0 {
				freqOccur[freqMap[query[1]]+1]--
			}

		case 3:
			if freqOccur[query[1]] > 0 {
				ans = append(ans, 1)
			} else {
				ans = append(ans, 0)
			}
		}
	}

	return ans
}

func main() {
	f, err := os.Open("in.txt")
	reader := bufio.NewReaderSize(f, 16*1024*1024)

	stdout, err := os.Create("out.txt")
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 16*1024*1024)

	qTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	checkError(err)
	q := int32(qTemp)

	var queries [][]int32
	for i := 0; i < int(q); i++ {
		queriesRowTemp := strings.Split(strings.TrimRight(readLine(reader), " \t\r\n"), " ")

		var queriesRow []int32
		for _, queriesRowItem := range queriesRowTemp {
			queriesItemTemp, err := strconv.ParseInt(queriesRowItem, 10, 64)
			checkError(err)
			queriesItem := int32(queriesItemTemp)
			queriesRow = append(queriesRow, queriesItem)
		}

		if len(queriesRow) != 2 {
			panic("Bad input")
		}

		queries = append(queries, queriesRow)
	}

	ans := freqQuery(queries)

	for i, ansItem := range ans {
		fmt.Fprintf(writer, "%d", ansItem)

		if i != len(ans)-1 {
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
