package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type coordinate struct {
	x int
	y int
}

func getCoordinateFromString(coordinateString string) coordinate {
	coordinateStringFields := strings.Fields(coordinateString)

	x, xError := strconv.Atoi(strings.Trim(coordinateStringFields[0], ","))
	y, yError := strconv.Atoi(coordinateStringFields[1])

	if xError != nil {
		log.Fatal(xError)
	}

	if yError != nil {
		log.Fatal(yError)
	}

	return coordinate{
		x: x,
		y: y,
	}
}

func getManhattanDistance(coordinate1 coordinate, coordinate2 coordinate) int {
	xDistance := coordinate2.x - coordinate1.x
	yDistance := coordinate2.y - coordinate1.y

	return xDistance + yDistance
}

func partOne() int {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var coordinates []coordinate

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		coordinates = append(coordinates, getCoordinateFromString(scanner.Text()))
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println(coordinates)

	xMin := coordinates[0].x
	yMin := coordinates[0].y
	xMax := coordinates[0].x
	yMax := coordinates[0].y

	for _, coordinate := range coordinates {
		if coordinate.x < xMin {
			xMin = coordinate.x
		}

		if coordinate.y < yMin {
			yMin = coordinate.y
		}

		if coordinate.x > xMax {
			xMax = coordinate.x
		}

		if coordinate.y > yMax {
			yMax = coordinate.y
		}
	}

	fmt.Println(xMin)
	fmt.Println(yMin)
	fmt.Println(xMax)
	fmt.Println(yMax)

	for x := xMin; x < xMax; x++ {
		for y := yMin; y < yMax; y++ {
			currentCoordinate := coordinate{x: x, y: y}

			closestCoordinate := coordinates[0]
			secondClosestCoordinate := coordinates[0]

			for _, coordinate := range coordinates {
				xDistance := coordinate.x - currentCoordinate.x
				yDistance := coordinate.y - currentCoordinate.y

				manhattanDistance := xDistance + yDistance

				closestCoordinateDistance := getManhattanDistance(currentCoordinate, closestCoordinate)

				if manhattanDistance < closestCoordinateDistance {
					secondClosestCoordinate = closestCoordinate
					closestCoordinate = coordinate
				}
			}

			fmt.Println(closestCoordinate)
			fmt.Println(secondClosestCoordinate)
		}
	}

	return 0
}

func main() {
	fmt.Println(partOne())
}
