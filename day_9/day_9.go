package main

import (
	"bufio"
	"container/ring"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func getInput() string {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var inputLine string

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		inputLine = scanner.Text()
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return inputLine
}

func convertStringToInt(value string) int {
	integerValue, err := strconv.Atoi(value)

	if err != nil {
		log.Fatal(err)
	}

	return integerValue
}

func getWinningScore(mapWithValues map[int]int) int {
	maxValue := 0

	for _, value := range mapWithValues {
		if value > maxValue {
			maxValue = value
		}
	}

	return maxValue
}

func playGame(numberOfPlayers int, lastMarbleValue int) map[int]int {
	marbles := ring.New(1)
	marbles.Value = 0

	scores := make(map[int]int)

	for i := 1; i <= lastMarbleValue; i++ {
		if i%23 == 0 {
			currentPlayer := i % numberOfPlayers

			marbles = marbles.Move(-8)
			marbleToRemove := marbles.Unlink(1)

			scores[currentPlayer] += i + marbleToRemove.Value.(int)

			marbles = marbles.Next()
		} else {
			marbles = marbles.Next()

			newMarble := ring.New(1)
			newMarble.Value = i

			marbles.Link(newMarble)
			marbles = marbles.Next()
		}
	}

	return scores
}

func partOne() int {
	inputLine := getInput()

	inputLineFields := strings.Fields(inputLine)

	numberOfPlayers := convertStringToInt(inputLineFields[0])
	lastMarbleValue := convertStringToInt(inputLineFields[6])

	scores := playGame(numberOfPlayers, lastMarbleValue)

	return getWinningScore(scores)
}

func partTwo() int {
	inputLine := getInput()

	inputLineFields := strings.Fields(inputLine)

	numberOfPlayers := convertStringToInt(inputLineFields[0])
	lastMarbleValue := convertStringToInt(inputLineFields[6])

	scores := playGame(numberOfPlayers, lastMarbleValue*100)

	return getWinningScore(scores)
}

func main() {
	fmt.Println(partOne())
	fmt.Println(partTwo())
}
