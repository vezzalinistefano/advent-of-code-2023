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

func main() {
	start := time.Now()

	file, err := os.Open("./test.txt")
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
			fmt.Println(index)
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
                fmt.Println(index)
				if (index + 4) > (len(line) - 1) {
					fmt.Println(line[index:])
				} else {
					fmt.Println(line[index : index+4])
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
