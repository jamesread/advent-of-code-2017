package main

import (
	"fmt"
	"io/ioutil"
	"regexp"
	"strings"
	"strconv"
)

type Node struct {
	name string
	weight int
	childrenNames []string

	children map[string]*Node
	parent *Node
}

func getInput() []string {
	contentsBuf, _ := ioutil.ReadFile("input")
	contents := strings.TrimSpace(string(contentsBuf))

	ret := strings.Split(contents, "\n")

	return ret
}

func matchChildren(line string) ([]string, error) {
	children := []string {}

	re, _ := regexp.Compile(`(\w+),?`)

	if re.MatchString(line) != true {
		return children, fmt.Errorf("Children string could not be matched: %s", line)
	}

	groups := re.FindAllStringSubmatch(line, -1)

	for _, child := range groups {
		children = append(children, child[1])
	}

	return children, nil
}

func matchLine(line string) *Node {
	re, _ := regexp.Compile(`(\w+) \((\d+)\)( -> (.+))?`)

	if re.MatchString(line) != true {
		fmt.Printf("no match: %s \n", line)
	}

	var node = new(Node)
	node.children = make(map[string]*Node)

	groups := re.FindAllStringSubmatch(line, -1)[0]

	node.name = groups[1]
	node.weight, _ = strconv.Atoi(groups[2])

	if len(groups) > 2 {
		node.childrenNames, _ = matchChildren(groups[4])
	}

	fmt.Printf("Matched lined: %+v \n", node)

	return node
}

/**
This function builds a tree from a list/map of nodes, but iterating over all
nodes and assigning the children/parent links (pointers). 

This is a crude way of building a tree. The traditional is to insert and
perform a rebalance, but that only words for ordered trees. 

That isn't possible here because our tree node data appears to be un-ordered.

We return the name of the root node.
*/
func buildTree(nodeList map[string]*Node) {
	fmt.Println("Building tree...")

	for _, node := range(nodeList) {
		for _, childName := range node.childrenNames {
			var childNode = nodeList[childName]
			node.children[childName] = childNode

			nodeList[childName].parent = node
		}
	}
}

func findRoot(nodeList map[string]*Node) *Node {
	var ret = new(Node)
	ret.name = "?"

	// Don't break on finding the root, as this checks we don't have multiple roots.
	for _, node := range(nodeList) {
		if node.parent == nil {
			if ret.name == "?" {
				ret = node
			} else {
				fmt.Println("Already found a root!")
			}
		}
	}

	return ret
}

func checkChildrenBalance(node *Node) {
	var check = -1

	for _, child := range node.children {
		if check == -1 {
			check = child.weight
		} else {
			if child.weight != check {
				fmt.Printf("\tChild weight check failed, should be %v, but it is %v. Parent node is %v \n", check, child.weight, node.name)
			}
		}

		checkChildrenBalance(child)
	}
}

func main() {
	var input = getInput()

	nodeList := map[string]*Node {}
	
	for _, line := range input {
		node := matchLine(strings.TrimSpace(line))

		nodeList[node.name] = node
	}

	buildTree(nodeList)

	var root = findRoot(nodeList)

	fmt.Printf("Root node: %v\n", root.name)

	//checkChildrenBalance(root)
}
