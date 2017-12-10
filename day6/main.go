package main

import (
	"bufio"
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

	scanner := bufio.NewScanner(strings.NewReader(contents[0]))
	scanner.Split(bufio.ScanWords)

	for scanner.Scan() {
		var v, _ = strconv.Atoi(scanner.Text())
		list = append(list, v)
	}

	return list;
}

func getIndexWithMostBlocks(bl []int) int {
	var biggestBank = 0;
	var largestNumberOfBlocks = 0;

	for i := 0; i < len(bl); i++ {
		if bl[i] > largestNumberOfBlocks {
			largestNumberOfBlocks = bl[i]
			biggestBank = i
		}
	}

	return biggestBank
}

func serialize(bl []int) string {
	return strings.Trim(strings.Join(strings.Fields(fmt.Sprint(bl)), "."), "[]")
}

var seenBanks = make(map[string]int)

func isBalancePossible(bl []int) bool {
	if _, ok := seenBanks[serialize(bl)]; ok {
		fmt.Println("Further balancing not possible, as I've seen the current pattern before. Here's all the patterns I've seen", seenBanks)
		return false;
	} else {
		return true
	}
}

func rebalance(banks []int) int {
	var indexToBalance = getIndexWithMostBlocks(banks)
	var blocksLeftToBalance = banks[indexToBalance]

	fmt.Println("Selected bank", indexToBalance, "to rebalance. It has the most blocks, with ", blocksLeftToBalance, "currently.")

	banks[indexToBalance] = 0

	var currentBank = indexToBalance + 1

	for blocksLeftToBalance > 0 {
		if currentBank == len(banks) {
			fmt.Println("wrapping banks")
			currentBank = 0
		}

		fmt.Println(blocksLeftToBalance, "blocks from bank", indexToBalance, "left to balance. Allocating a block to bank", currentBank)
		banks[currentBank]++

		blocksLeftToBalance--
		
		currentBank++
	}

	return indexToBalance
}

func main() {
	var contents = getInput()
	var bl = getNumberList(contents)

	var patternPosition = 0
	var balanceAttempts = 0

	fmt.Println("Starting bank definition:", bl);

	for isBalancePossible(bl) {
		seenBanks[serialize(bl)] = patternPosition

		var rebalancedIndex = rebalance(bl)
		fmt.Println("Finished rebalancing bank", rebalancedIndex, "\t Bank List is now:", bl);

		balanceAttempts++
		patternPosition++
	}

	fmt.Println("Balance attempts:", balanceAttempts)
	fmt.Println("Final Result:", bl)
	fmt.Println("Cycles since the pattern was first seen:", len(seenBanks) - seenBanks[serialize(bl)]);
}
