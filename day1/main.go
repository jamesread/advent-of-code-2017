package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func numberAt(contents string, i int) int {
	v, _ := strconv.Atoi((string)([]rune(contents)[i]))

	return v
}

func getInput() string {
	contentsBuf, _ := ioutil.ReadFile("input")
	contents := strings.TrimSpace(string(contentsBuf))

	return contents
}

// SolveCaptcha solves a Captcha as defined in the spec of the AOC2017 day1 problem.
//
// http://adventofcode.com/2017/day/1
func SolveCaptcha(contents string) int {
	total := 0;

	for i := range contents {
		curr := numberAt(contents, i)

		nextValidIndex := -1;

		if i+1 < len(contents) {
			nextValidIndex = i+1;
		} else if i+1 == len(contents) {
			nextValidIndex = 0
		}

		if nextValidIndex != -1 {
			next := numberAt(contents, nextValidIndex)

			if curr == next {
				total += curr;
			}
		}
	}

	return total;
}

func main() {
	contents := getInput();
	checksum := SolveCaptcha(contents);

	fmt.Println("Checksum:", checksum);
}
