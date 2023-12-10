package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

const (
	RED   = "red"
	GREEN = "green"
	BLUE  = "blue"
)

var finalResult int

func parseRound(rounds []string, game int) int {
	var extractionMap = map[string]int{
		"red":   0,
		"green": 0,
		"blue":  0,
	}

	for i, round := range rounds {
		if i > 0 {
			round = round[1:]
		}
		fmt.Printf("[%d]: '%s'\n", i, round)

		extractions := strings.Split(round, ",")

		for j, extraction := range extractions {
			if j > 0 {
				extraction = extraction[1:]
			}
			fmt.Printf("\t- [%d]: '%s'\n", j, extraction)

			splitExtraction := strings.Split(extraction, " ")

			colour := splitExtraction[1]

			amount, err := strconv.Atoi(splitExtraction[0])
			if err != nil {
				log.Fatal(err)
			}

            if amount > extractionMap[colour] {
                extractionMap[colour] = amount
            }
		}
	}

    power := 1
    fmt.Println(extractionMap)
    for _, value := range extractionMap {
        power *= value
    }
    return power
}

func main() {
	start := time.Now()

	file, err := os.Open("./input.txt")
	defer file.Close()

	if err != nil {
		log.Fatal(err)
	}

	fileScanner := bufio.NewScanner(file)
	fileScanner.Split(bufio.ScanLines)

	for fileScanner.Scan() {
		inputString := fileScanner.Text()

		// Remove the "Game" from the string
		inputString = strings.ReplaceAll(inputString, "Game ", "")

		splitString := strings.Split(inputString, ":")
		gameString := splitString[0]

		// Get extraction results removing first trailing whitespace
		gameResultsString := splitString[1][1:]

		game, err := strconv.Atoi(gameString)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("##########\nGame number: %d\n##########\n", game)
		if game == 100 {
			fmt.Println()
		}

		rounds := strings.Split(gameResultsString, ";")

        power := parseRound(rounds, game)
        finalResult += power 
	}

	fmt.Printf("FINAL RESULT: %d\n", finalResult)
	elapsed := time.Since(start)
	fmt.Printf("Execution time: %s", elapsed)
}
