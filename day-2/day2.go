package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	result := 0
	for scanner.Scan() {
		line := scanner.Text()
		if err := scanner.Err(); err != nil {
			fmt.Println(err)
			return
		}

		dims := getDims(&line)
		result += getSurfaceArea(dims) + getExtraArea(dims)
	}

	fmt.Printf("The elves need %d square feet of wrapping paper.\n", result)
}

func getDims(str *string) *[]int {
	dimsstr := strings.Split(*str, "x")

	l, _ := strconv.Atoi(dimsstr[0])
	w, _ := strconv.Atoi(dimsstr[1])
	h, _ := strconv.Atoi(dimsstr[2])

	dims := []int{l, w, h}
	sort.Ints(dims)
	return &dims
}

func getSurfaceArea(dims *[]int) int {
	return 2*(*dims)[0]*(*dims)[1] +
		2*(*dims)[1]*(*dims)[2] +
		2*(*dims)[2]*(*dims)[0]
}

// dims should be a sorted array of ints in ascending order
func getExtraArea(dims *[]int) int {
	return (*dims)[0] * (*dims)[1]
}
