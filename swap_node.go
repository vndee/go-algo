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
 * Complete the 'swapNodes' function below.
 *
 * The function is expected to return a 2D_INTEGER_ARRAY.
 * The function accepts following parameters:
 *  1. 2D_INTEGER_ARRAY indexes
 *  2. INTEGER_ARRAY queries
 */

type Node struct {
	value   int32
	left    int32
	right   int32
	isLeaft bool
}

func inOderTraversal(tree []Node, u int32, ans *[]int32) {
	if u < 0 {
		return
	}

	inOderTraversal(tree, tree[u].left, ans)
	*ans = append(*ans, tree[u].value+1)
	inOderTraversal(tree, tree[u].right, ans)
}

func calcHeight(tree []Node, u int32, prev int32, H *[]int32) {
	if u < 0 {
		return
	}

	(*H)[u] = (*H)[prev] + 1
	calcHeight(tree, tree[u].left, u, H)
	calcHeight(tree, tree[u].right, u, H)
}

func swap(x, y *int32) {
	*x, *y = *y, *x
}

func swapNodes(indexes [][]int32, queries []int32) [][]int32 {
	// Write your code here
	tree := make([]Node, len(indexes))

	for i, child := range indexes {
		tree[i].value = int32(i)
		tree[i].left, tree[i].right = child[0]-1, child[1]-1
		if child[0] < 0 && child[1] < 0 {
			tree[i].isLeaft = true
		} else {
			tree[i].isLeaft = false
		}
	}

	ans := make([][]int32, 0)
	heights := make([]int32, len(indexes))
	calcHeight(tree, 0, 0, &heights)

	for _, q := range queries {
		for i, h := range heights {
			if h%q == 0 && !tree[i].isLeaft {
				// swap child
				swap(&tree[i].left, &tree[i].right)
			}
		}

		sol := make([]int32, 0)
		inOderTraversal(tree, 0, &sol)
		ans = append(ans, sol)
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

	var indexes [][]int32
	for i := 0; i < int(n); i++ {
		indexesRowTemp := strings.Split(strings.TrimRight(readLine(reader), " \t\r\n"), " ")

		var indexesRow []int32
		for _, indexesRowItem := range indexesRowTemp {
			indexesItemTemp, err := strconv.ParseInt(indexesRowItem, 10, 64)
			checkError(err)
			indexesItem := int32(indexesItemTemp)
			indexesRow = append(indexesRow, indexesItem)
		}

		if len(indexesRow) != 2 {
			panic("Bad input")
		}

		indexes = append(indexes, indexesRow)
	}

	queriesCount, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	checkError(err)

	var queries []int32

	for i := 0; i < int(queriesCount); i++ {
		queriesItemTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
		checkError(err)
		queriesItem := int32(queriesItemTemp)
		queries = append(queries, queriesItem)
	}

	result := swapNodes(indexes, queries)

	for i, rowItem := range result {
		for j, colItem := range rowItem {
			fmt.Fprintf(writer, "%d", colItem)

			if j != len(rowItem)-1 {
				fmt.Fprintf(writer, " ")
			}
		}

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
