package main

import (
	"fmt"
	"io/ioutil"
	"strings"
	"strconv"
)

func getInput() []string {
	contentsBuf, _ := ioutil.ReadFile("input")
	contents := strings.TrimSpace(string(contentsBuf))

	ret := strings.Split(contents, "\n")

	return ret
}

func getNumberList(contents []string) []int {
	var list = []int {}

	for _, item := range contents {
		var v, _ = strconv.Atoi(item)
		list = append(list, v)
	}

	return list;
}

func getNumberOfStepsTilExit(numberList []int, advancedRule bool) int {
	var stepCount = 0
	var currentIndex = 0
	var newIndex = 0

	for {
		if currentIndex >= len(numberList) || currentIndex < 0 { 
			fmt.Println("Freeee!")
			break
		}

		var adjustmentAmount = 1

		if advancedRule {
			if numberList[currentIndex] >= 3 {
				adjustmentAmount = -1
			}
		}

		newIndex = currentIndex + numberList[currentIndex]
		numberList[currentIndex] += adjustmentAmount
		currentIndex = newIndex
	
		stepCount++
	}

	return stepCount
}

func main() {
	//contents := strings.Split("0 3 0 1 -3", " ")
	contents := getInput()

	numberList := getNumberList(contents)

	fmt.Println("Number of steps to exit: ", getNumberOfStepsTilExit(numberList, true))
}
