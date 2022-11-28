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
 * Complete the 'arrayManipulation' function below.
 *
 * The function is expected to return a LONG_INTEGER.
 * The function accepts following parameters:
 *  1. INTEGER n
 *  2. 2D_INTEGER_ARRAY queries
 */

type Node struct {
	// segment tree with lazy update implementation
	Key        int64
	Lazy       int64
	Left       int32
	Right      int32
	LeftChild  *Node
	RightChild *Node
}

func newNode() *Node {
	return &Node{
		Key:        0,
		Lazy:       0,
		Left:       0,
		Right:      0,
		LeftChild:  nil,
		RightChild: nil,
	}
}

func (node *Node) lazyUpdate() *Node {
	if node.Lazy <= 0 {
		return node
	}

	if node.LeftChild != nil {
		node.LeftChild.Lazy += node.Lazy
		node.LeftChild.Key += node.Lazy
	}

	if node.RightChild != nil {
		node.RightChild.Lazy += node.Lazy
		node.RightChild.Key += node.Lazy
	}

	node.Lazy = 0
	return node
}

func (node *Node) update(l, r int32, value int64) *Node {
	if l > node.Right || r < node.Left {
		return node
	}

	if node.Left >= l && node.Right <= r {
		node.Key += value
		node.Lazy += value
		node = node.lazyUpdate()
		return node
	}

	node = node.lazyUpdate()

	if node.LeftChild != nil {
		node.LeftChild = node.LeftChild.update(l, r, value)
		if node.Key < node.LeftChild.Key {
			node.Key = node.LeftChild.Key
		}
	}

	if node.RightChild != nil {
		node.RightChild = node.RightChild.update(l, r, value)
		if node.Key < node.RightChild.Key {
			node.Key = node.RightChild.Key
		}
	}

	return node
}

func (node *Node) query(l, r int32) int64 {
	if l > node.Right || r < node.Left {
		return -int64(1e9 + 1)
	}

	if node.Left >= l && node.Right <= r {
		return node.Key
	}

	node.lazyUpdate()

	mid := (l + r) >> 1
	leftMax := node.LeftChild.query(l, mid)
	rightMax := node.RightChild.query(mid+1, r)

	if leftMax > rightMax {
		return leftMax
	}

	return rightMax
}

func (node *Node) buildTree(l, r int32) {
	if l == r {
		node.Key = 0
		node.Left = l
		node.Right = r
		node.LeftChild = nil
		node.RightChild = nil
	} else {
		mid := (l + r) >> 1

		// create left child
		node.LeftChild = newNode()
		node.LeftChild.buildTree(l, mid)

		// create right child
		node.RightChild = newNode()
		node.RightChild.buildTree(mid+1, r)

		if node.LeftChild.Key > node.RightChild.Key {
			node.Key = node.LeftChild.Key
		} else {
			node.Key = node.RightChild.Key
		}

		node.Left = l
		node.Right = r
	}
}

func (node *Node) lnr() {
	if node.LeftChild != nil {
		node.LeftChild.lnr()
	}

	if node.LeftChild == nil && node.RightChild == nil && node.Left == node.Right {
		fmt.Print(node.Key, " ")
	}

	if node.RightChild != nil {
		node.RightChild.lnr()
	}
}

func arrayManipulation(n int32, queries [][]int32) int64 {
	// Write your code here
	root := newNode()
	root.buildTree(0, n-1)

	for _, query := range queries {
		root.update(int32(query[0]-1), int32(query[1]-1), int64(query[2]))
	}

	ans := root.query(0, n-1)
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

	firstMultipleInput := strings.Split(strings.TrimSpace(readLine(reader)), " ")

	nTemp, err := strconv.ParseInt(firstMultipleInput[0], 10, 64)
	checkError(err)
	n := int32(nTemp)

	mTemp, err := strconv.ParseInt(firstMultipleInput[1], 10, 64)
	checkError(err)
	m := int32(mTemp)

	var queries [][]int32
	for i := 0; i < int(m); i++ {
		queriesRowTemp := strings.Split(strings.TrimRight(readLine(reader), " \t\r\n"), " ")

		var queriesRow []int32
		for _, queriesRowItem := range queriesRowTemp {
			queriesItemTemp, err := strconv.ParseInt(queriesRowItem, 10, 64)
			checkError(err)
			queriesItem := int32(queriesItemTemp)
			queriesRow = append(queriesRow, queriesItem)
		}

		if len(queriesRow) != 3 {
			panic("Bad input")
		}

		queries = append(queries, queriesRow)
	}

	result := arrayManipulation(n, queries)

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
