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
}

func processNumbers(listOfNumbers *[]int, currentIndex int, nodes *[]Node) {
	numberOfChildNodes := (*listOfNumbers)[currentIndex]
	numberOfMetaDataEntries := (*listOfNumbers)[currentIndex+1]

	for i := 0; i < numberOfChildNodes; i++ {
		processNumbers(listOfNumbers, currentIndex+2, nodes)
	}

	node := Node{numberOfChildNodes: numberOfChildNodes, numberOfMetaDataEntries: numberOfMetaDataEntries}

	node.metaDataEntries = append(node.metaDataEntries, (*listOfNumbers)[currentIndex+2:currentIndex+2+node.numberOfMetaDataEntries]...)
	*listOfNumbers = append((*listOfNumbers)[:currentIndex], (*listOfNumbers)[currentIndex+2+node.numberOfMetaDataEntries:]...)

	*nodes = append(*nodes, node)
}

func partOne() {
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

	fmt.Println(sum)
}

func partTwo() {

}

func main() {
	partOne()
	partTwo()
}
