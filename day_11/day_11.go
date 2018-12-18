package main

import (
	"fmt"
	"log"
	"strconv"
)

func getPowerLevelForCoordinates(x int, y int, serialNumber int) int {
	rackID := x + 10

	powerLevel := rackID * y
	powerLevel = powerLevel + serialNumber
	powerLevel = powerLevel * rackID

	powerLevelAsString := strconv.Itoa(powerLevel)

	hundredsDigit, err := strconv.Atoi(string(powerLevelAsString[len(powerLevelAsString)-3]))

	if err != nil {
		log.Fatal(err)
	}

	return hundredsDigit - 5
}

func dayOne(serialNumber int) {
	powerLevels := make([][]int, 300)

	for x := 1; x <= 300; x++ {
		powerLevels[x-1] = make([]int, 300)

		for y := 1; y <= 300; y++ {
			powerLevels[x-1][y-1] = getPowerLevelForCoordinates(x, y, serialNumber)
		}
	}

	highestXCoordinate := -1
	highestYCoordinate := -1
	highestValue := -100

	for x := 0; x < len(powerLevels)-2; x++ {
		for y := 0; y < len(powerLevels[x])-2; y++ {
			value := 0

			value += powerLevels[x][y]
			value += powerLevels[x][y+1]
			value += powerLevels[x][y+2]
			value += powerLevels[x+1][y]
			value += powerLevels[x+1][y+1]
			value += powerLevels[x+1][y+2]
			value += powerLevels[x+2][y]
			value += powerLevels[x+2][y+1]
			value += powerLevels[x+2][y+2]

			if value > highestValue {
				highestValue = value
				highestXCoordinate = x
				highestYCoordinate = y
			}
		}
	}

	fmt.Println("Coordinates:", highestXCoordinate+1, ",", highestYCoordinate+1)
}

func main() {
	dayOne(6303)
}
