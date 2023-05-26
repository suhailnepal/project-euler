//  Problem #1
// If we list all the natural numbers below  that are multiples of  or , we get  and . The sum of these multiples is .

// Find the sum of all the multiples of  or  below .

// Input Format

// First line contains  that denotes the number of test cases. This is followed by  lines, each containing an integer, .

// Constraints

// Output Format

// For each test case, print an integer that denotes the sum of all the multiples of  or  below .

// Sample Input 0

// 2
// 10
// 100
// Sample Output 0

// 23
// 2318
// Explanation 0

// For , if we list all the natural numbers below  that are multiples of  or , we get  and . The sum of these multiples is .

// Similarly for , we get .

package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func main() {
	// O(t)
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	tTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	checkError(err)
	t := int32(tTemp)

	for tItr := 0; tItr < int(t); tItr++ {
		nTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
		checkError(err)
		n := int(nTemp)
		fmt.Println(FindTotalSumOfMultiples(n))
	}
}

func FindTotalSumOfMultiples(N int) int {
	// O(N)
	// sum := 0
	// for i := 0; i < N; i++ {
	// 	if i%3 == 0 || i%5 == 0 {
	// 		sum += i
	// 	}
	// }
	// return sum

	// O(1)
	numMultiples3 := (N - 1) / 3
	sumMultiples3 := 3 * (numMultiples3 * (numMultiples3 + 1)) / 2

	// Calculate the sum of multiples of 5
	numMultiples5 := (N - 1) / 5
	sumMultiples5 := 5 * (numMultiples5 * (numMultiples5 + 1)) / 2

	// Calculate the sum of multiples of 15 (to avoid double-counting)
	numMultiples15 := (N - 1) / 15
	sumMultiples15 := 15 * (numMultiples15 * (numMultiples15 + 1)) / 2

	// Calculate the final sum
	sum := sumMultiples3 + sumMultiples5 - sumMultiples15

	return sum
}

func readLine(reader *bufio.Reader) string {
	// O(L)
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
