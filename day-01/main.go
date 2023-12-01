package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"unicode"
)

func main() {
	file, err := os.Open("./test.txt")
	defer file.Close()

	if err != nil {
		log.Println(err)
	}

	fileScanner := bufio.NewScanner(file)
	fileScanner.Split(bufio.ScanLines)

	for fileScanner.Scan() {
		line := fileScanner.Text()
		fmt.Println(line)

		var first *int
		first = nil
		var last *int
		last = nil

		for _, r := range line {
			if unicode.IsDigit(r) {
				if first == nil {
					i, err := strconv.Atoi(string(r))
					if err != nil {
						log.Println(err)
					}
					i = i * 10
					first = &i
					fmt.Println(*first)
				} else {
					i, err := strconv.Atoi(string(r))
					if err != nil {
						log.Println(err)
					}
					last = &i
					fmt.Printf("last %d\n", *last)
				}
			}
		}
	}
}
