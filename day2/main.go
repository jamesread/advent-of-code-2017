package main

import (
	"fmt"
	"io/ioutil"
	"strings"
	"bufio"
)

func getInput() string {
	contentsBuf, _ := ioutil.ReadFile("input")
	contents := strings.TrimSpace(string(contentsBuf))

	return contents
}

func getChecksum(input string) int {
	var scanner = bufio.NewScanner(input)

	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}

	return 0;
}

func main() {
	var input = getInput()

	var checksum = getChecksum(input)

	fmt.Sprintf("Checksum %d", checksum);
}
