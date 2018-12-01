package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func frequencyStringToInt(frequency string) int {
	integer, err := strconv.Atoi(frequency)

	check(err)

	return integer
}

func partOne() int {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var frequencies []string
	frequencyTotal := 0

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		frequencies = append(frequencies, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	for _, element := range frequencies {
		frequencyTotal += frequencyStringToInt(element)
	}

	return frequencyTotal
}

func partTwo() int {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var frequencies []string
	frequencyTotals := make(map[int]bool)
	frequencyTotal := 0
	firstRepeatFrequency := 0
	repeatFrequencyFound := false

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		frequencies = append(frequencies, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	for !repeatFrequencyFound {
		for _, element := range frequencies {
			frequencyTotal += frequencyStringToInt(element)

			_, frequencyExists := frequencyTotals[frequencyTotal]

			if frequencyExists {
				firstRepeatFrequency = frequencyTotal
				repeatFrequencyFound = true
				break
			} else {
				frequencyTotals[frequencyTotal] = true
			}
		}
	}

	return firstRepeatFrequency
}

func main() {
	fmt.Println(partOne())
	fmt.Println(partTwo())
}
