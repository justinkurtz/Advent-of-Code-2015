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

type pair struct {
	x byte
	y byte
}

// It contains a pair of any two letters that appears at least twice in the string without overlapping
func hasTwoPair(str *string) bool {
	m := make(map[pair]int)
	for i := 0; i < len(*str)-1; i++ {
		p := pair{(*str)[i], (*str)[i+1]}
		if m[p] != 0 && m[p] != i {
			return true
		}

		m[p] = i + 1
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
