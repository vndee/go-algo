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
 * Complete the 'poisonousPlants' function below.
 *
 * The function is expected to return an INTEGER.
 * The function accepts INTEGER_ARRAY p as parameter.
 */

func poisonousPlants(p []int32) int32 {
	// Write your code here

	stack := make([]int32, 0)
	days := make([]int32, len(p))
	dayToDie := int32(0)

	for i, elm := range p {
		dayToDie = 0
		x := int32(-1)
		for len(stack) > 0 && p[stack[len(stack)-1]] >= elm {
			x = stack[len(stack)-1]
			stack = stack[0 : len(stack)-1]
			if dayToDie < days[x] {
				dayToDie = days[x]
			}
		}

		if len(stack) == 0 {
			days[i] = 0
		} else if x == -1 {
			days[i] = 1
		} else {
			days[i] = dayToDie + 1
		}

		stack = append(stack, int32(i))
	}

	ans := int32(0)
	for _, d := range days {
		if d > ans {
			ans = d
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

	nTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	checkError(err)
	n := int32(nTemp)

	pTemp := strings.Split(strings.TrimSpace(readLine(reader)), " ")

	var p []int32

	for i := 0; i < int(n); i++ {
		pItemTemp, err := strconv.ParseInt(pTemp[i], 10, 64)
		checkError(err)
		pItem := int32(pItemTemp)
		p = append(p, pItem)
	}

	result := poisonousPlants(p)

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
