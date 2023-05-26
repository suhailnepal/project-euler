package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

// Time complexity: O(t * N)
// Space complexity: O(1)

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	tTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	checkError(err)
	t := int32(tTemp)

	// O(t)
	for tItr := 0; tItr < int(t); tItr++ {
		n, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
		checkError(err)
		largestPrimeNumber := calculateLargestPrimeNumber(int(n))
		fmt.Println(largestPrimeNumber)
	}
}

func calculateLargestPrimeNumber(N int) int {
	// Declaring a number as prime first
	maxPrime := 0
	// O(N)
	for i := 2; i <= N; i++ {
		//  To be prime, the number must not be divisible and should be divisible
		//  by itself and 1
		if N%i == 0 {
			N /= i
			maxPrime = i
		}
	}

	if N > 1 {
		maxPrime = N
	}

	return maxPrime
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
