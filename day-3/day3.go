package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	line := scanner.Text()
	if err := scanner.Err(); err != nil {
		fmt.Println(err)
		return
	}

	result := 1
	currPos := coord{0, 0}
	m := make(map[coord]bool)
	m[currPos] = true
	for _, dir := range line {
		switch dir {
		case '>':
			currPos.x++
		case '^':
			currPos.y++
		case '<':
			currPos.x--
		case 'v':
			currPos.y--
		}

		if !m[currPos] {
			result++
		}

		m[currPos] = true
	}

	fmt.Printf("Santa delivered to %d houses.\n", result)
}

type coord struct {
	x int
	y int
}
