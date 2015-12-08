package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var grid [1000][1000]int

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		line := scanner.Text()
		if err := scanner.Err(); err != nil {
			fmt.Println(err)
			return
		} else if line == "" {
			break
		}

		processLine(&line)
	}

	fmt.Printf("The brightness is at %d.\n", count())
}

func processLine(str *string) {
	words := strings.Split(*str, " ")
	if len(words) < 4 {
		return
	}

	if words[0] == "turn" {
		x1, y1 := parseIndices(&words[2])
		x2, y2 := parseIndices(&words[4])
		iter(x1, y1, x2, y2, func(b int) int {
			if words[1] == "on" {
				return b + 1
			} else if b == 0 {
				return 0
			}

			return b - 1
		})
	} else if words[0] == "toggle" {
		x1, y1 := parseIndices(&words[1])
		x2, y2 := parseIndices(&words[3])
		iter(x1, y1, x2, y2, func(b int) int {
			return b + 2
		})
	}
}

func parseIndices(str *string) (int, int) {
	indices := strings.Split(*str, ",")
	x1, _ := strconv.Atoi(indices[0])
	y1, _ := strconv.Atoi(indices[1])
	return x1, y1
}

func iter(x1 int, y1 int, x2 int, y2 int, f func(int) int) {
	for x := x1; x <= x2; x++ {
		for y := y1; y <= y2; y++ {
			grid[x][y] = f(grid[x][y])
		}
	}
}

func count() int {
	count := 0
	iter(0, 0, 999, 999, func(b int) int {
		count += b
		return b
	})

	return count
}
