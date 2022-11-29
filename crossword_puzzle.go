package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

/*
 * Complete the 'crosswordPuzzle' function below.
 *
 * The function is expected to return a STRING_ARRAY.
 * The function accepts following parameters:
 *  1. STRING_ARRAY crossword
 *  2. STRING words
 */

var ans []string

type slot struct {
	positionX     int
	positionY     int
	positionStart int
	positionEnd   int
}

func isValidGrid() bool {
	return false
}

func backtrack(i, j int) {

}

func crosswordPuzzle(crossword []string, words string) []string {
	// Write your code here

	ans = make([]string, 10)
	blankChar := byte(45) // "_"

	// find empty rows and cols
	col, row := make([]slot, 0), make([]slot, 0)
	for i := 0; i < 10; i++ {
		for j := 0; j < 10; j++ {
			if crossword[i][j] == blankChar {
				// find empty row start at i, j
				var k int
				for k = 0; j+k < 10 && crossword[i][j+k] == blankChar; k++ {
				}
				row = append(row, slot{
					positionX:     i,
					positionY:     j,
					positionStart: j,
					positionEnd:   j + k,
				})

				// find empty col start at i, j
				for k = 0; i+k < 10 && crossword[i+k][j] == blankChar; k++ {
				}
				col = append(row, slot{
					positionX:     i,
					positionY:     j,
					positionStart: i,
					positionEnd:   i + k,
				})
			}
		}
	}

	fmt.Println("col", col)
	fmt.Println("row", row)
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

	var crossword []string

	for i := 0; i < 10; i++ {
		crosswordItem := readLine(reader)
		crossword = append(crossword, crosswordItem)
	}

	words := readLine(reader)

	result := crosswordPuzzle(crossword, words)

	for i, resultItem := range result {
		fmt.Fprintf(writer, "%s", resultItem)

		if i != len(result)-1 {
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
