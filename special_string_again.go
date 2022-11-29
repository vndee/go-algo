package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

var x string
var m = make(map[string]int64)

func dp(l, r int) int64 {
	k := fmt.Sprintf("%d_%d", l, r)
	if m[k] > 0 {
		return m[k]
	}

	if r-l+1 == 3 {
		if x[l] == x[r] && x[l+1] == x[r] {
			// ooo
			m[k] = 6
			return m[k]
		}

		if x[l] == x[r] && x[l+1] != x[r] {
			// non
			m[k] = 4
			return m[k]
		}

		if x[l] == x[l+1] || x[l+1] == x[r] {
			// nno || noo
			m[k] = 4
			return m[k]
		}

		// mon
		m[k] = 3
		return m[k]
	}

	if l > r {
		return 0
	}
	if r-l == 0 {
		return 1
	}

	m[k] = dp(l+1, r) + dp(l, r-1) - dp(l+1, r-1)
	return m[k]
}

// Complete the substrCount function below.
func substrCount_bruteforce(n int32, s string) int64 {
	x = s
	ans := dp(0, len(x)-1)
	return ans
}

func min(x, y int64) int64 {
	if x > y {
		return y
	} else {
		return x
	}
}

func substrCount(n int32, s string) int64 {
	ans := int64(0)
	for i := int32(0); i < n; i++ {
		j := int32(0)

		for i+j < n && s[i] == s[i+j] {
			j++
			ans++
		}

		if i+2*j > n {
			continue
		}

		flag := false
		for k := int32(1); k < j+1; k++ {
			if i+j+k >= n || s[i] != s[i+j+k] {
				flag = true
				break
			}
		}

		if !flag {
			ans++
		}
	}
	return ans
}

func main() {
	f, err := os.Open("in.txt")
	checkError(err)
	reader := bufio.NewReaderSize(f, 1024*1024)

	stdout, err := os.Create("out.txt")
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 1024*1024)

	nTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
	checkError(err)
	n := int32(nTemp)

	s := readLine(reader)

	result := substrCount(n, s)

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
