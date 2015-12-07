package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	nice := 0
	for scanner.Scan() {
		line := scanner.Text()
		if err := scanner.Err(); err != nil {
			fmt.Println(err)
			return
		} else if line == "" {
			break
		}

		if isNice(&line) {
			nice++
		}
	}

	fmt.Printf("There are %d nice strings.\n", nice)
}

func isNice(str *string) bool {
	// fmt.Printf("hasThreeVowels=%t\n", hasThreeVowels(str))
	// fmt.Printf("hasDoubleLetter=%t\n", hasDoubleLetter(str))
	// fmt.Printf("hasNaughtyStr=%t\n", hasNaughtyStr(str))

	return hasThreeVowels(str) && hasDoubleLetter(str) && !hasNaughtyStr(str)
}

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

func hasNaughtyStr(str *string) bool {
	naughtyStrs := []string{"ab", "cd", "pq", "xy"}

	for _, naughty := range naughtyStrs {
		if strings.Contains(*str, naughty) {
			return true
		}
	}

	return false
}
