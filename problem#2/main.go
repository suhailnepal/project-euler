package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	tTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	checkError(err)
	t := int32(tTemp)

	for tItr := 0; tItr < int(t); tItr++ {
		n, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
		checkError(err)
		sum := febonacciSum(int(n))
		start := time.Now()
		fmt.Println(sum)
		fmt.Printf("Time taken: %v\n", time.Since(start))

	}
}

func febonacciSum(N int) (sum int) {
	// We use two pointers , current and previous value.
	//  We already know that current = 1 as its the second value
	// and prev to be 0 as its the first value
	prev := 0
	curr := 1
	// initial sum to be 0
	sum = 0

	for curr <= N {
		// Add the current term to the sum if it is even
		if curr%2 == 0 {
			sum += curr
		}

		// Generate the next term by adding the previous two term
		// Note we need to add before re-assigning
		next := prev + curr
		prev = curr
		curr = next
	}

	return
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
