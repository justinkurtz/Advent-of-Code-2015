package main

import (
	"bufio"
	"crypto/md5"
	"fmt"
	"os"
	"strconv"
)

const numZeros int = 6

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	line := scanner.Bytes()
	if err := scanner.Err(); err != nil {
		fmt.Println(err)
		return
	}

	for i := 0; ; i++ {
		iStr := []byte(strconv.Itoa(i))
		newLine := append(line, iStr...)
		sum := fmt.Sprintf("%x", md5.Sum(newLine))
		if isValid(&sum) {
			fmt.Printf("Answer: %d\n", i)
			break
		}
	}
}

func isValid(sum *string) bool {
	for i := 0; i < numZeros; i++ {
		if (*sum)[i] != '0' {
			return false
		}
	}

	return true
}
