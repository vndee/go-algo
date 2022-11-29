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
 * Complete the 'maxRegion' function below.
 *
 * The function is expected to return an INTEGER.
 * The function accepts 2D_INTEGER_ARRAY grid as parameter.
 */

func maxRegion(grid [][]int32) int32 {
	// Write your code here
	// I technically adapt BFS to solve a typical DFS problem

	var visited [10][10]int32
	dx := []int32{-1, -1, -1, 0, +1, +1, +1, 0}
	dy := []int32{-1, 0, +1, +1, +1, 0, -1, -1}
	compCount := int32(0)

	n, m := int32(len(grid)), int32(len(grid[0]))
	for i := int32(0); i < n; i++ {
		for j := int32(0); j < m; j++ {
			if visited[i][j] == 0 && grid[i][j] > 0 {
				// do BFS
				compCount++
				l, r := int32(0), int32(1)
				queue_x := make([]int32, 0)
				queue_y := make([]int32, 0)

				queue_x = append(queue_x, i)
				queue_y = append(queue_y, j)
				visited[i][j] = compCount

				for l < r {
					u, v := queue_x[l], queue_y[l]
					visited[u][v] = compCount
					l++

					for _, x := range dx {
						for _, y := range dy {
							newX, newY := x+u, y+v

							if newX < 0 || newY < 0 || newX >= n || newY >= m {
								continue
							}

							if visited[newX][newY] > 0 || grid[newX][newY] == 0 {
								continue
							}

							r++
							queue_x = append(queue_x, newX)
							queue_y = append(queue_y, newY)
						}
					}
				}
			}
		}
	}

	var compSize [101]int32
	var ans int32 = 0
	for i := int32(0); i < n; i++ {
		for j := int32(0); j < m; j++ {
			kComp := visited[i][j]
			if kComp == 0 {
				continue
			}
			compSize[kComp]++
			if compSize[kComp] > ans {
				ans = compSize[kComp]
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

	nTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	checkError(err)
	n := int32(nTemp)

	mTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	checkError(err)
	m := int32(mTemp)

	var grid [][]int32
	for i := 0; i < int(n); i++ {
		gridRowTemp := strings.Split(strings.TrimRight(readLine(reader), " \t\r\n"), " ")

		var gridRow []int32
		for _, gridRowItem := range gridRowTemp {
			gridItemTemp, err := strconv.ParseInt(gridRowItem, 10, 64)
			checkError(err)
			gridItem := int32(gridItemTemp)
			gridRow = append(gridRow, gridItem)
		}

		if len(gridRow) != int(m) {
			panic("Bad input")
		}

		grid = append(grid, gridRow)
	}

	res := maxRegion(grid)

	fmt.Fprintf(writer, "%d\n", res)

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
