package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"
	"unicode"
)

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

		for _, r := range line {
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
