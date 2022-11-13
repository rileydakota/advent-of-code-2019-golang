package day1

import (
	"log"
	"math"
	"strconv"
)

func calcFuel(mass int) int {
	fuel := math.Floor(float64(mass)/3) - 2
	if math.Signbit(fuel) {
		return 0
	}
	return int(fuel)
}

func Solvepart1(input []string) (int, error) {

	total := 0

	for _, element := range input {
		integer, err := strconv.Atoi(element)
		if err != nil {
			log.Fatal(err)
			return 0, err
		}
		total = total + calcFuel(integer)
	}

	return total, nil
}

func Solvepart2(input []string) (int, error) {
	total := 0

	for _, element := range input {
		integer, err := strconv.Atoi(element)
		if err != nil {
			log.Fatal(err)
			return 0, err
		}
		fuelArray := []int{}
		fuel := calcFuel(integer)
		for fuel > 0 {
			newFuel := calcFuel(fuel)
			fuelArray = append(fuelArray, fuel)
			fuel = newFuel
		}

		totalFuel := 0

		for _, element := range fuelArray {
			totalFuel = totalFuel + element
		}

		total = total + totalFuel

	}

	return total, nil
}
