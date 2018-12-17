package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type PointOfLight struct {
	xPosition int
	yPosition int
	xVelocity int
	yVelocity int
}

func (p *PointOfLight) move() {
	p.xPosition = p.xPosition + p.xVelocity
	p.yPosition = p.yPosition + p.yVelocity
}

func getMinimumAndMaximumValues(pointsOfLight []*PointOfLight) (int, int, int, int) {
	minimumX := pointsOfLight[0].xPosition
	minimumY := pointsOfLight[0].yPosition
	maximumX := pointsOfLight[0].xPosition
	maximumY := pointsOfLight[0].yPosition

	for _, pointOfLight := range pointsOfLight {
		if pointOfLight.xPosition < minimumX {
			minimumX = pointOfLight.xPosition
		}

		if pointOfLight.yPosition < minimumY {
			minimumY = pointOfLight.yPosition
		}

		if pointOfLight.xPosition > maximumX {
			maximumX = pointOfLight.xPosition
		}

		if pointOfLight.yPosition > maximumY {
			maximumY = pointOfLight.yPosition
		}
	}

	return minimumX, minimumY, maximumX, maximumY
}

func moveAllPointsOfLight(pointsOfLight []*PointOfLight) {
	for _, pointOfLight := range pointsOfLight {
		pointOfLight.move()
	}
}

func printPointsOfLight(pointsOfLight []*PointOfLight) {
	minimumX, minimumY, maximumX, maximumY := getMinimumAndMaximumValues(pointsOfLight)

	for y := minimumY - 2; y <= maximumY+2; y++ {
		for x := minimumX - 2; x <= maximumX+2; x++ {
			pointFound := false

			for _, pointOfLight := range pointsOfLight {
				if pointOfLight.xPosition == x && pointOfLight.yPosition == y {
					pointFound = true
				}
			}

			if pointFound {
				fmt.Print("#")
			} else {
				fmt.Print(".")
			}
		}

		fmt.Print("\n")
	}
}

func checkAllPointsOfLightAreClose(pointsOfLight []*PointOfLight) bool {
	minimumX, minimumY, maximumX, maximumY := getMinimumAndMaximumValues(pointsOfLight)

	if maximumX-minimumX > 100 || maximumY-minimumY > 100 {
		return false
	}

	return true
}

func getInput() []string {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var inputLines []string

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		inputLines = append(inputLines, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return inputLines
}

func stringToInt(stringValue string) int {
	intValue, err := strconv.Atoi(stringValue)

	if err != nil {
		log.Fatal(err)
	}

	return intValue
}

func partOneAndTwo() {
	inputLines := getInput()

	pointsOfLight := []*PointOfLight{}

	for _, inputLine := range inputLines {
		firstSplitStrings := strings.Split(inputLine, "<")

		firstCoordinatesString := strings.Split(firstSplitStrings[1], ">")[0]
		secondCoordinatesString := strings.Split(firstSplitStrings[2], ">")[0]

		xPositionAsString := strings.TrimSpace(strings.Split(firstCoordinatesString, ",")[0])
		yPositionAsString := strings.TrimSpace(strings.Split(firstCoordinatesString, ",")[1])
		xVelocityAsString := strings.TrimSpace(strings.Split(secondCoordinatesString, ",")[0])
		yVelocityAsString := strings.TrimSpace(strings.Split(secondCoordinatesString, ",")[1])

		xPosition := stringToInt(xPositionAsString)
		yPosition := stringToInt(yPositionAsString)
		xVelocity := stringToInt(xVelocityAsString)
		yVelocity := stringToInt(yVelocityAsString)

		pointOfLight := PointOfLight{
			xPosition: xPosition,
			yPosition: yPosition,
			xVelocity: xVelocity,
			yVelocity: yVelocity,
		}

		pointsOfLight = append(pointsOfLight, &pointOfLight)
	}

	for i := 0; i < 1000000; i++ {
		if checkAllPointsOfLightAreClose(pointsOfLight) {
			fmt.Println("Second: ", i)
			printPointsOfLight(pointsOfLight)
		}

		moveAllPointsOfLight(pointsOfLight)
	}
}

func main() {
	partOneAndTwo()
}
