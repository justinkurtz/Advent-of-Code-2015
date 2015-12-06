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

	floor := 0
	found := false
	for i, c := range line {
		if c == '(' {
			floor++
		} else if c == ')' {
			floor--
		} else {
			fmt.Printf("Didn't recognize %c\n", c)
		}

		if floor == -1 && !found {
			fmt.Printf("The answer to part 2 is %d\n", i+1)
			found = true
		}
	}

	fmt.Printf("Santa is on floor %d\n", floor)
}
