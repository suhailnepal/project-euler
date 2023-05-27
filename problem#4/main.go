package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

// Time Complexity: O(t) * O(n log i) = O(t * n log i)
// Space complexity: O(1)

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	tTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	checkError(err)
	t := int32(tTemp)

	// O(t)
	for tItr := 0; tItr < int(t); tItr++ {
		nTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
		checkError(err)
		n := int32(nTemp)
		number, isPalindrome := CalculateLargestPalindrome(int(n))
		if isPalindrome {
			fmt.Println(number)
		}
	}
}

// O(log i)
func CheckPalindrome(n int) bool {
	str := strconv.Itoa(n)
	length := len(str)
	for i := 0; i < length/2; i++ {
		if str[i] != str[length-1-i] {
			return false
		}
	}
	return true
}

// O(n ) * O (log i ) = O(n log i)
func CalculateLargestPalindrome(n int) (int, bool) {
	// O(n)
	for i := n - 1; i >= 100000; i-- {
		// O(log i)
		if CheckPalindrome(i) {
			for j := 999; j >= 100; j-- {
				if i%j == 0 && i/j >= 100 && i/j <= 999 {
					return i, true
				}
			}
		}
	}
	return 0, false
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
