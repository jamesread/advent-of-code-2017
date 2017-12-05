package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"strings"
	"strconv"
)

func getHigherLower(a int, b int) (int, int) {
	var higher = 0
	var lower = 0

	if a > b {
		higher = a
		lower = b
	} else {
		higher = b
		lower = a
	}

	return higher, lower
}

func getInput() string {
	contentsBuf, _ := ioutil.ReadFile("input")
	contents := strings.TrimSpace(string(contentsBuf))

	return contents
}

type Gridline map[int]int;
type Grid map[int]Gridline;

func getDivideChecksum(grid Grid) int {
	var total = 0;

	lineLoop: for lineNo, row := range grid {
		for _, cellCurrent := range(row) {
			for _, cellMatch := range(row) {
				if cellCurrent == cellMatch {
					continue
				}

				var higher, lower = getHigherLower(cellCurrent, cellMatch)

				if higher % lower == 0 {
					total += higher / lower

					fmt.Printf("Line %d. Higher %d / Lower %d = %d \n", lineNo, higher, lower , higher / lower)
					continue lineLoop
				}
			}
		}

	}

	return total
}

func buildGrid(input string) Grid {
	grid := make(Grid)

	for lineNo, line := range strings.Split(input, "\n") {
		scanner := bufio.NewScanner(strings.NewReader(line))
		scanner.Split(bufio.ScanWords)

		grid[lineNo] = make(Gridline)

		var i = 0
		for scanner.Scan() {
			val, _ := strconv.Atoi(scanner.Text())

			grid[lineNo][i] = val
			i++
		}
	}

	return grid
}

func main() {
	var input = getInput()

	var checksum = getChecksum(buildGrid(input))

	fmt.Printf("Checksum %d\n", checksum);
}
