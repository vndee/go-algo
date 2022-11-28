package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"sort"
	"strconv"
	"strings"
)

/*
 * Complete the 'activityNotifications' function below.
 *
 * The function is expected to return an INTEGER.
 * The function accepts following parameters:
 *  1. INTEGER_ARRAY expenditure
 *  2. INTEGER d
 */

func activityNotifications_bruteforce(expenditure []int32, d int32) int32 {
	// Write your code here
	var ans int32 = 0
	var n int32 = int32(len(expenditure))

	for i := d; i < n; i++ {
		slice := expenditure[i-d : i]
		sort.Slice(slice, func(i, j int) bool { return slice[i] < slice[j] })

		median := slice[d>>1]
		// fmt.Println(slice, "->", median, expenditure[i])
		if expenditure[i] >= 2*median {
			ans = ans + 1
		}
	}

	return ans
}

func activityNotifications(expenditure []int32, d int32) int32 {
	var ans int32 = 0
	var count [201]int32
	var runningMax int32 = 0

	for i := int32(0); i < d; i++ {
		count[expenditure[i]]++
		if runningMax < expenditure[i] {
			runningMax = expenditure[i]
		}
	}

	var pivot int32 = (d + 1) >> 1
	for i := d; i < int32(len(expenditure)); i++ {
		// find median
		var s int32 = 0
		var median float32

		for j := int32(0); j <= runningMax; j++ {
			if count[j] != 0 && s+count[j] >= pivot {
				median = float32(j)
				if d%2 == 0 {
					// find the second median
					if s+count[j] >= pivot+1 {
						median = (median + float32(j)) / 2
					} else {
						for j := int32(median) + 1; j <= runningMax; j++ {
							if count[j] != 0 {
								median = (median + float32(j)) / 2
								break
							}
						}
					}
				}
				break
			}
			s = s + count[j]
		}

		// fmt.Println(expenditure[i-d:i], median, expenditure[i])
		count[expenditure[i-d]]--
		count[expenditure[i]]++
		if expenditure[i] > runningMax {
			runningMax = expenditure[i]
		}

		if float32(expenditure[i]) >= 2*median {
			ans++
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

	firstMultipleInput := strings.Split(strings.TrimSpace(readLine(reader)), " ")

	nTemp, err := strconv.ParseInt(firstMultipleInput[0], 10, 64)
	checkError(err)
	n := int32(nTemp)

	dTemp, err := strconv.ParseInt(firstMultipleInput[1], 10, 64)
	checkError(err)
	d := int32(dTemp)

	expenditureTemp := strings.Split(strings.TrimSpace(readLine(reader)), " ")

	var expenditure []int32

	for i := 0; i < int(n); i++ {
		expenditureItemTemp, err := strconv.ParseInt(expenditureTemp[i], 10, 64)
		checkError(err)
		expenditureItem := int32(expenditureItemTemp)
		expenditure = append(expenditure, expenditureItem)
	}

	result := activityNotifications(expenditure, d)

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
