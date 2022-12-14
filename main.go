package main

import (
	day1 "adventofcode2019/solutions/day1"
	"adventofcode2019/solutions/day2"
	"bufio"
	"log"
	"os"
	"strconv"
)

func loadFile(filename string) ([]string, error) {
	f, err := os.Open(filename)

	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	defer f.Close()

	var lines []string
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	return lines, scanner.Err()

}

func main() {
	day1input, err := loadFile("input/day1.txt")
	if err != nil {
		log.Fatal(err)
	}
	day1part1answer, err := day1.Solvepart1(day1input)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("day 1 part 1 answer: %s", strconv.Itoa(day1part1answer))
	day1part2answer, err := day1.Solvepart2(day1input)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("day 1 part 2 answer: %s", strconv.Itoa(day1part2answer))

	day2input, err := loadFile("input/day2.txt")
	if err != nil {
		log.Fatal(err)
	}
	day2part1answer, err := day2.Solvepart1(day2input)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("day 1 part 1 answer: %s", strconv.Itoa(day2part1answer[0]))
}
