package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func reactPolymer(polymer string) string {
	newString := polymer

	possibilities := []string{
		"Aa",
		"Bb",
		"Cc",
		"Dd",
		"Ee",
		"Ff",
		"Gg",
		"Hh",
		"Ii",
		"Jj",
		"Kk",
		"Ll",
		"Mm",
		"Nn",
		"Oo",
		"Pp",
		"Qq",
		"Rr",
		"Ss",
		"Tt",
		"Uu",
		"Vv",
		"Ww",
		"Xx",
		"Yy",
		"Zz",
		"aA",
		"bB",
		"cC",
		"dD",
		"eE",
		"fF",
		"gG",
		"hH",
		"iI",
		"jJ",
		"kK",
		"lL",
		"mM",
		"nN",
		"oO",
		"pP",
		"qQ",
		"rR",
		"sS",
		"tT",
		"uU",
		"vV",
		"wW",
		"xX",
		"yY",
		"zZ",
	}

	for _, possibility := range possibilities {
		newString = strings.Replace(newString, possibility, "", -1)
	}

	if len(polymer) != len(newString) {
		return reactPolymer(newString)
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
