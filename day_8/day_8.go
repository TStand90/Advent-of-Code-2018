package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Node struct {
	numberOfChildNodes      int
	numberOfMetaDataEntries int
	metaDataEntries         []int
	childNodes              []Node
}

func processNumbers(listOfNumbers *[]int, currentIndex int, nodes *[]Node) Node {
	numberOfChildNodes := (*listOfNumbers)[currentIndex]
	numberOfMetaDataEntries := (*listOfNumbers)[currentIndex+1]

	var childNodes []Node

	for i := 0; i < numberOfChildNodes; i++ {
		childNodes = append(childNodes, processNumbers(listOfNumbers, currentIndex+2, nodes))
	}

	node := Node{numberOfChildNodes: numberOfChildNodes, numberOfMetaDataEntries: numberOfMetaDataEntries}
	node.childNodes = append(node.childNodes, childNodes...)

	node.metaDataEntries = append(node.metaDataEntries, (*listOfNumbers)[currentIndex+2:currentIndex+2+node.numberOfMetaDataEntries]...)
	*listOfNumbers = append((*listOfNumbers)[:currentIndex], (*listOfNumbers)[currentIndex+2+node.numberOfMetaDataEntries:]...)

	*nodes = append(*nodes, node)

	return node
}

func sumNodes(rootNode Node) int {
	sum := 0

	if rootNode.numberOfChildNodes == 0 {
		for _, metaDataValue := range rootNode.metaDataEntries {
			sum += metaDataValue
		}
	} else {
		for _, metaDataValue := range rootNode.metaDataEntries {
			if len(rootNode.childNodes) >= metaDataValue {
				childNode := rootNode.childNodes[metaDataValue-1]

				sum += sumNodes(childNode)
			}

		}
	}

	return sum
}

func partOne() int {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var listOfNumbersAsString string

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		listOfNumbersAsString = scanner.Text()
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	listOfNumbersAsStrings := strings.Fields(listOfNumbersAsString)

	var listOfNumbers []int

	for _, numberAsString := range listOfNumbersAsStrings {
		integer, err := strconv.Atoi(numberAsString)

		if err != nil {
			log.Fatal(err)
		}

		listOfNumbers = append(listOfNumbers, integer)
	}

	var nodes []Node

	processNumbers(&listOfNumbers, 0, &nodes)

	sum := 0

	for _, node := range nodes {
		for _, value := range node.metaDataEntries {
			sum += value
		}
	}

	return sum
}

func partTwo() int {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var listOfNumbersAsString string

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		listOfNumbersAsString = scanner.Text()
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	listOfNumbersAsStrings := strings.Fields(listOfNumbersAsString)

	var listOfNumbers []int

	for _, numberAsString := range listOfNumbersAsStrings {
		integer, err := strconv.Atoi(numberAsString)

		if err != nil {
			log.Fatal(err)
		}

		listOfNumbers = append(listOfNumbers, integer)
	}

	var nodes []Node

	rootNode := processNumbers(&listOfNumbers, 0, &nodes)

	sum := sumNodes(rootNode)

	return sum
}

func main() {
	fmt.Println(partOne())
	fmt.Println(partTwo())
}
