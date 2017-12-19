/*

Day 8. 

This works, but I'm really not happy with it. The error handling is just 
log.Fatalf, which is crappy - untestable. 
*/

package main

import (
	"io/ioutil"
	"strings"
	"strconv"
	"regexp"
	"log"
	"fmt"
)

type Statement struct {
	register string
	inc bool
	changeValue int

	testRegister string
	condition Condition
	testValue int
}

var registers = make(map[string]int)
var highestValue = 0

type Condition int

const (
	INVALID Condition = iota
	GT
	LT
	LTE
	GTE
	EQ
	NE
)

func getInput() []string {
	contentsBuf, _ := ioutil.ReadFile("input")
	contents := strings.TrimSpace(string(contentsBuf))

	ret := strings.Split(contents, "\n")

	return ret
}

func filterCondition(in string) Condition {
	switch (in) {
		case ">": return GT
		case "<": return LT
		case "==": return EQ
		case ">=": return GTE
		case "<=": return LTE
		case "!=": return NE
		default:
			log.Fatalf("Unrecognized condition: %v", in)
			return INVALID
	}
}

func filterInt(in string) int {
	v, _ := strconv.Atoi(in)

	return v
}

func filterString(in string) string {
	return in
}

func matchStatement(line string) Statement {
	var ret Statement

	re := regexp.MustCompile(`(\w+) (inc|dec) (-?\d+) if (\w+) ([!<>=]+) (-?\d+)`)

	if !re.MatchString(line) {
		fmt.Printf("Statement could not be matched: %v\n", line)
	}

	match := re.FindStringSubmatch(line)

	ret.register = filterString(match[1])
	ret.inc = match[2] == "inc"
	ret.changeValue = filterInt(match[3])

	ret.testRegister = filterString(match[4])
	ret.condition = filterCondition(match[5])
	ret.testValue = filterInt(match[6])

	return ret
}

func evaluateCondition(statement Statement) bool {
	switch (statement.condition) {
		case GT: return registers[statement.testRegister] > statement.testValue
		case LT: return registers[statement.testRegister] < statement.testValue
		case LTE: return registers[statement.testRegister] <= statement.testValue
		case GTE: return registers[statement.testRegister] >= statement.testValue
		case EQ: return registers[statement.testRegister] == statement.testValue
		case NE: return registers[statement.testRegister] != statement.testValue
		default:
			log.Fatalf("Unhandled condition: %v\n", statement.condition)
			return false
	}
}

func evaluateStatement(statement Statement) {
	if _, exists := registers[statement.register]; !exists {
		registers[statement.register] = 0
	}

	if (evaluateCondition(statement)) {
		if statement.inc {
			registers[statement.register] += statement.changeValue
		} else {
			registers[statement.register] -= statement.changeValue
		}
	}

	if registers[statement.register] > highestValue {
		highestValue = registers[statement.register]
	}
}

func findLargestRegister() string {
	var largest = "fake"

	for current, _ := range registers {
		if registers[current] > registers[largest] {
			largest = current
		}
	}

	return largest
}

func main() {
	var input = getInput()

	registers["fake"] = 0

	for _, line := range input {
		var statement = matchStatement(line)
	
		evaluateStatement(statement)
	}

	var largestRegister = findLargestRegister()

	log.Printf("The largest value in any register is %v with a value of %v\n", largestRegister, registers[largestRegister])
	log.Printf("The highest value during the process was %v", highestValue)
}
