package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func stringToInt(stringToConvert string) int {
	integer, err := strconv.Atoi(stringToConvert)

	if err != nil {
		log.Fatal(err)
	}

	return integer
}

func dayOne() int {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	claimMap := make(map[string]int)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		claim := scanner.Text()

		splitClaim := strings.Fields(claim)

		coordinateString := splitClaim[2]
		widthHeightString := splitClaim[3]

		splitXAndYCoordinateStrings := strings.Split(coordinateString, ",")

		xCoordinateString := splitXAndYCoordinateStrings[0]
		yCoordinateString := strings.Trim(splitXAndYCoordinateStrings[1], ":")

		xCoordinate := stringToInt(xCoordinateString)
		yCoordinate := stringToInt(yCoordinateString)

		widthHeightSplitStrings := strings.Split(widthHeightString, "x")
		width := stringToInt(widthHeightSplitStrings[0])
		height := stringToInt(widthHeightSplitStrings[1])

		for y := yCoordinate; y < yCoordinate+height; y++ {
			for x := xCoordinate; x < xCoordinate+width; x++ {
				coordinateAsString := fmt.Sprintf("%d,%d", x, y)

				_, coordinateExists := claimMap[coordinateAsString]

				if coordinateExists {
					claimMap[coordinateAsString]++

				} else {
					claimMap[coordinateAsString] = 1
				}
			}
		}
	}

	overlapCounter := 0

	for y := 0; y < 1000; y++ {
		for x := 0; x < 1000; x++ {
			coordinateAsString := fmt.Sprintf("%d,%d", x, y)

			coordinateValue, coordinateExists := claimMap[coordinateAsString]

			if coordinateExists {
				if coordinateValue >= 2 {
					overlapCounter++
				}
			}
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return overlapCounter
}

func dayTwo() int {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var claims []string
	claimMap := make(map[string]int)
	idsWithoutOverlapMap := make(map[int]bool)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		claims = append(claims, scanner.Text())
	}

	for _, claim := range claims {
		splitClaim := strings.Fields(claim)

		claimID := stringToInt(strings.Trim(splitClaim[0], "#"))
		coordinateString := splitClaim[2]
		widthHeightString := splitClaim[3]

		splitXAndYCoordinateStrings := strings.Split(coordinateString, ",")

		xCoordinateString := splitXAndYCoordinateStrings[0]
		yCoordinateString := strings.Trim(splitXAndYCoordinateStrings[1], ":")

		xCoordinate := stringToInt(xCoordinateString)
		yCoordinate := stringToInt(yCoordinateString)

		widthHeightSplitStrings := strings.Split(widthHeightString, "x")
		width := stringToInt(widthHeightSplitStrings[0])
		height := stringToInt(widthHeightSplitStrings[1])

		idsWithoutOverlapMap[claimID] = true

		for y := yCoordinate; y < yCoordinate+height; y++ {
			for x := xCoordinate; x < xCoordinate+width; x++ {
				coordinateAsString := fmt.Sprintf("%d,%d", x, y)

				_, coordinateExists := claimMap[coordinateAsString]

				if coordinateExists {
					claimMap[coordinateAsString]++

					if claimMap[coordinateAsString] >= 2 {
						idsWithoutOverlapMap[claimID] = false
					}
				} else {
					claimMap[coordinateAsString] = 1
				}
			}
		}
	}

	for _, claim := range claims {
		splitClaim := strings.Fields(claim)

		claimID := stringToInt(strings.Trim(splitClaim[0], "#"))
		coordinateString := splitClaim[2]
		widthHeightString := splitClaim[3]

		splitXAndYCoordinateStrings := strings.Split(coordinateString, ",")

		xCoordinateString := splitXAndYCoordinateStrings[0]
		yCoordinateString := strings.Trim(splitXAndYCoordinateStrings[1], ":")

		xCoordinate := stringToInt(xCoordinateString)
		yCoordinate := stringToInt(yCoordinateString)

		widthHeightSplitStrings := strings.Split(widthHeightString, "x")
		width := stringToInt(widthHeightSplitStrings[0])
		height := stringToInt(widthHeightSplitStrings[1])

		idsWithoutOverlapMap[claimID] = true

		for y := yCoordinate; y < yCoordinate+height; y++ {
			for x := xCoordinate; x < xCoordinate+width; x++ {
				coordinateAsString := fmt.Sprintf("%d,%d", x, y)

				_, coordinateExists := claimMap[coordinateAsString]

				if coordinateExists {
					if claimMap[coordinateAsString] >= 2 {
						idsWithoutOverlapMap[claimID] = false
					}
				}
			}
		}
	}

	nonOverlappingID := -1

	for key, value := range idsWithoutOverlapMap {
		if value {
			nonOverlappingID = key
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return nonOverlappingID
}

func main() {
	fmt.Println(dayOne())
	fmt.Println(dayTwo())
}
