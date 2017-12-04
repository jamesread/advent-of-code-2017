package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"strings"
	"sort"
)

func getInput() string {
	contentsBuf, _ := ioutil.ReadFile("input")
	contents := strings.TrimSpace(string(contentsBuf))

	return contents
}

// Thanks: https://stackoverflow.com/questions/22688651/golang-how-to-sort-string-or-byte
type unsortedRunes []rune;

func (s unsortedRunes) Less(i, j int) bool {
    return s[i] < s[j]
}

func (s unsortedRunes) Swap(i, j int) {
    s[i], s[j] = s[j], s[i]
}

func (s unsortedRunes) Len() int {
	return len(s)
}

func sortTok(tok string) string {
	var r = []rune(tok)

	sort.Sort(unsortedRunes(r))

	return string(r)
}

func countValidPassphrases(contents string) int {
	var totalValid = 0;

	for _, line := range strings.Split(contents, "\n") {
		var valid = true

		scanner := bufio.NewScanner(strings.NewReader(line))
		scanner.Split(bufio.ScanWords)

		var toks = make(map[string]bool)

		for scanner.Scan() {
			var tok = scanner.Text()
			var sortedTok = sortTok(tok)

			_, exists := toks[sortedTok]

			if exists {
				fmt.Println("Line invalid. Repeated token ", tok, " (", sortedTok, ") in line ", line)
				valid = false;
			}

			toks[sortedTok] = true
		}

		if valid {
			totalValid++
		}
	}

	return totalValid;
}

func main() {
	var input = getInput()

	fmt.Printf("Valid phrases: %d\n", countValidPassphrases(input))
}
