package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func partOne() int {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	idsWithTwo := 0
	idsWithThree := 0

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		id := scanner.Text()

		characterMap := make(map[string]int)

		for _, char := range id {
			character := string(char)

			_, characterKeyExists := characterMap[character]

			if characterKeyExists {
				characterMap[character]++
			} else {
				characterMap[character] = 1
			}
		}

		twosFound := false
		threesFound := false

		for _, value := range characterMap {
			if value == 2 && !twosFound {
				idsWithTwo++
				twosFound = true
			}

			if value == 3 && !threesFound {
				idsWithThree++
				threesFound = true
			}
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return idsWithTwo * idsWithThree
}

func partTwo() string {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var lines []string

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	matchFound := false
	var answerID string

	for firstLineIterator := 0; firstLineIterator < len(lines)-1; firstLineIterator++ {
		firstLine := lines[firstLineIterator]

		for secondLineIterator := firstLineIterator + 1; secondLineIterator < len(lines); secondLineIterator++ {
			secondLine := lines[secondLineIterator]

			differences := 0

			for charIterator := 0; charIterator < len(firstLine); charIterator++ {
				if string(firstLine[charIterator]) != string(secondLine[charIterator]) {
					differences++
				}
			}

			if differences <= 1 {
				for charIterator := 0; charIterator < len(firstLine); charIterator++ {
					if string(firstLine[charIterator]) == string(secondLine[charIterator]) {
						answerID += string(firstLine[charIterator])
					}
				}

				matchFound = true
				break
			}
		}

		if matchFound {
			break
		}
	}

	return answerID
}

func main() {
	fmt.Println(partOne())
	fmt.Println(partTwo())
}
