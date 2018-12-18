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
			distance := 3

			for zx := 0; zx < distance; zx++ {
				for zy := 0; zy < distance; zy++ {
					value += powerLevels[x+zx][y+zy]
				}
			}

			if value > highestValue {
				highestValue = value
				highestXCoordinate = x
				highestYCoordinate = y
			}
		}
	}

	fmt.Println("Coordinates:", highestXCoordinate+1, ",", highestYCoordinate+1)
}

func dayTwo(serialNumber int) {
	powerLevels := make([][]int, 300)

	for x := 1; x <= 300; x++ {
		powerLevels[x-1] = make([]int, 300)

		for y := 1; y <= 300; y++ {
			powerLevels[x-1][y-1] = getPowerLevelForCoordinates(x, y, serialNumber)
		}
	}

	highestXCoordinate := -1
	highestYCoordinate := -1
	highestSize := -1
	highestValue := -100

	for distance := 0; distance < len(powerLevels)-2; distance++ {
		for x := 0; x < len(powerLevels)-distance; x++ {
			for y := 0; y < len(powerLevels[x])-distance; y++ {
				value := 0

				if len(powerLevels)-distance < len(powerLevels) && len(powerLevels[x])-distance < len(powerLevels[x]) {
					for zx := 0; zx <= distance; zx++ {
						for zy := 0; zy <= distance; zy++ {
							value += powerLevels[x+zx][y+zy]
						}
					}

					if value > highestValue {
						highestValue = value
						highestXCoordinate = x
						highestYCoordinate = y
						highestSize = distance
					}
				}
			}
		}
	}

	fmt.Println("Coordinates:", highestXCoordinate+1, ",", highestYCoordinate+1, ",", highestSize+1)
}

func main() {
	dayOne(6303)
	dayTwo(6303)
}
