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
	santaPos := coord{0, 0}
	roboPos := coord{0, 0}
	var currPos *coord
	m := make(map[coord]bool)
	m[santaPos] = true
	for i, dir := range line {
		if i%2 == 0 {
			currPos = &santaPos
		} else {
			currPos = &roboPos
		}

		move(dir, currPos)
		if !m[*currPos] {
			result++
		}

		m[*currPos] = true
	}

	fmt.Printf("Santa delivered to %d houses.\n", result)
}

func move(dir rune, pos *coord) {
	switch dir {
	case '>':
		pos.x++
	case '^':
		pos.y++
	case '<':
		pos.x--
	case 'v':
		pos.y--
	}
}

type coord struct {
	x int
	y int
}
