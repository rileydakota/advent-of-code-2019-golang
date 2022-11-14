package day2

import (
	"log"
	"strconv"
	"strings"
)

func Map(vs []string, f func(string) (int, error)) []int {
	vsm := make([]int, len(vs))
	for i, v := range vs {
		vsm[i], _ = f(v)
	}
	return vsm
}

func Solvepart1(input []string) ([]int, error) {
	inputArray := Map(strings.Split(input[0], ","), strconv.Atoi)
	inputArray[1] = 12
	inputArray[2] = 2

	place := 0
	for place < len(inputArray) {
		opcode := inputArray[place]
		switch opcode {
		case 1:

			value1 := inputArray[place+1]
			value2 := inputArray[place+2]
			target := inputArray[place+3]
			inputArray[target] = inputArray[value1] + inputArray[value2]
			place += 4
		case 2:
			value1 := inputArray[place+1]
			value2 := inputArray[place+2]
			target := inputArray[place+3]
			inputArray[target] = inputArray[value1] * inputArray[value2]
			place += 4
		case 99:
			log.Printf("day2 part1 - opcode 99 encountered, exiting - place: %s", strconv.Itoa(place))
			return inputArray, nil
		default:
			panic("lol")
		}
	}
	log.Printf("%v", inputArray)
	return inputArray, nil

}
