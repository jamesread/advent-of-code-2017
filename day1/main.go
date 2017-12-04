package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

// SolvedCaptcha is used to store the result and the original captcha.
type SolvedCaptcha struct {
	Result int;
	Captcha string;
}

func numberAtStringIndex(contents string, i int) int {
	v, _ := strconv.Atoi((string)([]rune(contents)[i]))

	return v
}

func getInput() string {
	contentsBuf, _ := ioutil.ReadFile("input")
	contents := strings.TrimSpace(string(contentsBuf))

	return contents
}

// SolveCaptchaHalfway solves a Captcha as defined in the spec of the AOC2017 day1,
// problem, using lookups half way around the list (part 2)
//
// http://adventofcode.com/2017/day/1
func SolveCaptchaHalfway(contents string) *SolvedCaptcha {
	total := 0

	for i := range contents {
		curr := numberAtStringIndex(contents, i)

		indexToLookup := i + (len(contents) / 2);

		if indexToLookup >= len(contents) {
			indexToLookup -= len(contents);
		}

		next := numberAtStringIndex(contents, indexToLookup)

		if curr == next {
			total += curr
		}
	}

	return &SolvedCaptcha {
		Captcha: contents,
		Result: total,
	};
}


// SolveCaptchaNext solves a Captcha as defined in the spec of the AOC2017 day1,
// problem, using i+1 lookups (part 1)
//
// http://adventofcode.com/2017/day/1
func SolveCaptchaNext(contents string) *SolvedCaptcha {
	total := 0;

	for i := range contents {
		curr := numberAtStringIndex(contents, i)

		nextValidIndex := -1;

		if i+1 < len(contents) {
			nextValidIndex = i+1;
		} else if i+1 == len(contents) {
			nextValidIndex = 0
		}

		if nextValidIndex != -1 {
			next := numberAtStringIndex(contents, nextValidIndex)

			if curr == next {
				total += curr;
			}
		}
	}

	return &SolvedCaptcha {
		Captcha: contents,
		Result: total,
	};
}

func main() {
	contents := getInput();
	solution := SolveCaptchaHalfway(contents);

	fmt.Println("Checksum:", solution.Result);
}
