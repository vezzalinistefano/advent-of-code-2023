package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
	"unicode"
)

const numberFirstLetters = "otfsen"

var numbers = map[string]int{
	"one":   1,
	"two":   2,
	"three": 3,
	"four":  4,
	"five":  5,
	"six":   6,
	"seven": 7,
	"eight": 8,
	"nine":  9,
}

func extractNumber(s string) int {
	i := 0
	var ok bool
	value, ok := numbers[s]

	for ok == false {
		i = i + 1
		subs := s[:len(s)-i]

		// 3 is the least amount of characters
		if len(subs) < len("one") {
			break
		}

		value, ok = numbers[subs]
	}
	return value
}

func main() {
	start := time.Now()

	file, err := os.Open("./input.txt")
	defer file.Close()

	if err != nil {
		log.Println(err)
	}

	fileScanner := bufio.NewScanner(file)
	fileScanner.Split(bufio.ScanLines)

	var sum int

	for fileScanner.Scan() {
		line := fileScanner.Text()

		var first *int
		var last *int
		last = nil
		first = nil

		for index, r := range line {
			if unicode.IsDigit(r) {
				if first == nil {
					i, err := strconv.Atoi(string(r))
					if err != nil {
						log.Println(err)
					}

					first = &i
				} else {
					i, err := strconv.Atoi(string(r))
					if err != nil {
						log.Println(err)
					}

					last = &i
				}
			} else if strings.Contains(numberFirstLetters, string(r)) {
				var i int
				if (index + 5) > (len(line) - 1) {
					i = extractNumber(line[index:])
				} else {
					i = extractNumber(line[index : index+5])
				}

				if i > 0 {
					if first == nil {
						first = &i
					} else {
						last = &i
					}
				}
			}
		}
		if last == nil {
			last = first
		}
		sum += (*first * 10) + *last
	}

	fmt.Printf("RESULT IS: %d\n", sum)
	elapsed := time.Since(start)
	fmt.Printf("Execution time: %s", elapsed)
}
