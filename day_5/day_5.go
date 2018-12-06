package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func reactPolymer(polymer string) string {
	newString := ""

	for _, character := range polymer {
		if len(newString) == 0 {
			newString = string(character)
		} else {
			currentCharacter := string(newString[len(newString)-1])

			characterAsString := string(character)

			if (characterAsString == strings.ToUpper(characterAsString) && currentCharacter == strings.ToLower(characterAsString)) || (characterAsString == strings.ToLower(characterAsString) && currentCharacter == strings.ToUpper(characterAsString)) {
				newString = newString[:len(newString)-1]
			} else {
				newString += string(character)
			}
		}
	}

	return newString
}

func collapsePolymer(polymer string, charToRemove string) string {
	collapsedPolymer := strings.Replace(polymer, charToRemove, "", -1)
	collapsedPolymer = strings.Replace(collapsedPolymer, strings.ToUpper(charToRemove), "", -1)

	return collapsedPolymer
}

func partOne() int {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var polymer string

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		polymer = scanner.Text()
	}

	answer := reactPolymer(polymer)

	return len(answer)
}

func partTwo() int {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var polymer string

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		polymer = scanner.Text()
	}

	alphabet := "abcdefghijklmnopqrstuvwxyz"

	var polymers []string

	for _, character := range alphabet {
		polymers = append(polymers, collapsePolymer(polymer, string(character)))
	}

	currentLowestCount := len(polymer)

	for _, polymer := range polymers {
		reactedPolymer := reactPolymer(polymer)

		if len(reactedPolymer) < currentLowestCount {
			currentLowestCount = len(reactedPolymer)
		}
	}

	return currentLowestCount
}

func main() {
	fmt.Println(partOne())
	fmt.Println(partTwo())
}
