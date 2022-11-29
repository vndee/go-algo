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
 * Complete the 'minimumMoves' function below.
 *
 * The function is expected to return an INTEGER.
 * The function accepts following parameters:
 *  1. STRING_ARRAY grid
 *  2. INTEGER startX
 *  3. INTEGER startY
 *  4. INTEGER goalX
 *  5. INTEGER goalY
 */

func minimumMoves(grid []string, startX int32, startY int32, goalX int32, goalY int32) int32 {
	// Write your code here
	n := int32(len(grid))
	head, tail := 0, 1
	blockedMark := []byte("X")[0]
	queueX, queueY := make([]int32, 0), make([]int32, 0)

	var shortestPath [100][100]int32

	shortestPath[startX][startY] = 1
	queueX, queueY = append(queueX, startX), append(queueY, startY)

	for head < tail {
		x, y := queueX[head], queueY[head]
		head++

		if x == goalX && y == goalY {
			break
		}

		isBlockedCol, isBlockedRow := false, false
		for k := int32(1); x+k < n || y+k < n; k++ {
			if x+k < n && !isBlockedCol {
				// move forward along the col
				u, v := x+k, y

				if grid[u][v] == blockedMark {
					isBlockedCol = true
				} else {
					if shortestPath[u][v] == 0 {
						shortestPath[u][v] = shortestPath[x][y] + 1

						tail++
						queueX, queueY = append(queueX, u), append(queueY, v)
					}
				}
			}

			if y+k < n && !isBlockedRow {
				// move forward along the row
				u, v := x, y+k

				if grid[u][v] == blockedMark {
					isBlockedRow = true
				} else {
					if shortestPath[u][v] == 0 {
						shortestPath[u][v] = shortestPath[x][y] + 1

						tail++
						queueX, queueY = append(queueX, u), append(queueY, v)
					}
				}
			}
		}

		isBlockedCol, isBlockedRow = false, false
		for k := int32(1); x-k >= 0 || y-k >= 0; k++ {
			if x-k >= 0 && !isBlockedCol {
				// move backward along the col
				u, v := x-k, y

				if grid[u][v] == blockedMark {
					isBlockedCol = true
				} else {
					if shortestPath[u][v] == 0 {
						shortestPath[u][v] = shortestPath[x][y] + 1

						tail++
						queueX, queueY = append(queueX, u), append(queueY, v)
					}
				}
			}

			if y-k >= 0 && !isBlockedRow {
				// move backward along the row
				u, v := x, y-k

				if grid[u][v] == blockedMark {
					isBlockedRow = true
				} else {
					if shortestPath[u][v] == 0 {
						shortestPath[u][v] = shortestPath[x][y] + 1

						tail++
						queueX, queueY = append(queueX, u), append(queueY, v)
					}
				}
			}
		}
	}

	for i := int32(0); i < n; i++ {
		for j := int32(0); j < n; j++ {
			fmt.Print(shortestPath[i][j], " ")
		}
		fmt.Println()
	}
	fmt.Println(shortestPath[0][2])

	return shortestPath[goalX][goalY] - 1
}

func main() {
	f, err := os.Open("in.txt")
	checkError(err)
	reader := bufio.NewReaderSize(f, 16*1024*1024)

	stdout, err := os.Create("out.txt")
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 16*1024*1024)

	nTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	checkError(err)
	n := int32(nTemp)

	var grid []string

	for i := 0; i < int(n); i++ {
		gridItem := readLine(reader)
		grid = append(grid, gridItem)
	}

	firstMultipleInput := strings.Split(strings.TrimSpace(readLine(reader)), " ")

	startXTemp, err := strconv.ParseInt(firstMultipleInput[0], 10, 64)
	checkError(err)
	startX := int32(startXTemp)

	startYTemp, err := strconv.ParseInt(firstMultipleInput[1], 10, 64)
	checkError(err)
	startY := int32(startYTemp)

	goalXTemp, err := strconv.ParseInt(firstMultipleInput[2], 10, 64)
	checkError(err)
	goalX := int32(goalXTemp)

	goalYTemp, err := strconv.ParseInt(firstMultipleInput[3], 10, 64)
	checkError(err)
	goalY := int32(goalYTemp)

	result := minimumMoves(grid, startX, startY, goalX, goalY)

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
