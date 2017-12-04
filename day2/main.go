package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"strings"
	"strconv"
)

func getInput() string {
	contentsBuf, _ := ioutil.ReadFile("input")
	contents := strings.TrimSpace(string(contentsBuf))

	return contents
}

func getChecksum(input string) int {
	var runningTotal = 0;

	for _, line := range strings.Split(input, "\n") {
		highest := 0
		lowest := 99999999

		scanner := bufio.NewScanner(strings.NewReader(line))
		scanner.Split(bufio.ScanWords)

		for scanner.Scan() {
			val, _ := strconv.Atoi(scanner.Text())

			if val > highest {
				highest = val
			}

			if val < lowest {
				lowest = val
			}
		}

		runningTotal += highest - lowest
	}

	return runningTotal
}

func main() {
	var input = getInput()

	var checksum = getChecksum(input)

	fmt.Printf("Checksum %d\n", checksum);
}
