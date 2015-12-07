package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	niceOne := 0
	niceTwo := 0
	for scanner.Scan() {
		line := scanner.Text()
		if err := scanner.Err(); err != nil {
			fmt.Println(err)
			return
		} else if line == "" {
			break
		}

		if isNicePartOne(&line) {
			niceOne++
		}
		if isNicePartTwo(&line) {
			niceTwo++
		}
	}

	fmt.Printf("Part 1: There are %d nice strings.\n", niceOne)
	fmt.Printf("Part 2: There are %d nice strings.\n", niceTwo)
}

func isNicePartOne(str *string) bool {
	// fmt.Printf("hasThreeVowels=%t\n", hasThreeVowels(str))
	// fmt.Printf("hasDoubleLetter=%t\n", hasDoubleLetter(str))
	// fmt.Printf("hasNaughtyStr=%t\n", hasNaughtyStr(str))
	return hasThreeVowels(str) && hasDoubleLetter(str) && !hasNaughtyStr(str)
}

func isNicePartTwo(str *string) bool {
	//fmt.Printf("hasTwoPair=%t\n", hasTwoPair(str))
	//fmt.Printf("hasDoubleJumpLetter=%t\n", hasDoubleLetter(str))
	return hasTwoPair(str) && hasDoubleJumpLetter(str)
}

// It contains at least three vowels
func hasThreeVowels(str *string) bool {
	count := 0
	for _, c := range *str {
		if c == 'a' || c == 'e' || c == 'i' || c == 'o' || c == 'u' {
			count++
			if count >= 3 {
				return true
			}
		}
	}

	return false
}

// It contains at least one letter that appears twice in a row
func hasDoubleLetter(str *string) bool {
	prev := rune((*str)[0])
	for _, c := range (*str)[1:] {
		if c == prev {
			return true
		}

		prev = c
	}

	return false
}

// It does not contain the strings ab, cd, pq, or xy
func hasNaughtyStr(str *string) bool {
	naughtyStrs := []string{"ab", "cd", "pq", "xy"}

	for _, naughty := range naughtyStrs {
		if strings.Contains(*str, naughty) {
			return true
		}
	}

	return false
}

// It contains a pair of any two letters that appears at least twice in the string without overlapping
func hasTwoPair(str *string) bool {
	strLen := len(*str)
	// O(N^2) horribleness
	for i := 0; i < strLen-1; i++ {
		if i > 1 {
			if strings.Contains((*str)[:i-1], (*str)[i:i+2]) {
				//fmt.Printf("%s is found on the left (%s)\n", (*str)[i:i+2], (*str)[:i-1])
				return true
			}
		}

		if i < strLen-3 {
			if strings.Contains((*str)[i+2:], (*str)[i:i+2]) {
				//fmt.Printf("%s is found on the right (%s)\n", (*str)[i:i+2], (*str)[i+2:])
				return true
			}
		}
	}

	return false
}

// It contains at least one letter which repeats with exactly one letter between them
func hasDoubleJumpLetter(str *string) bool {
	if len(*str) < 3 {
		return false
	}

	prev := 0
	for i := 2; i < len(*str); i++ {
		if (*str)[prev] == (*str)[i] {
			return true
		}

		prev = i - 1
	}

	return false
}
