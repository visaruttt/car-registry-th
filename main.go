package main

import (
	"fmt"
	"strconv"
)

func main() {
	CarRegistry()
}

func CarRegistry() {
	targetNumbers := []int{9, 1, 6}
	perfectedSumResult := buildPerfectResult()
	frontNumber := 9

	plateNumbers := getPlateNumbers(targetNumbers, perfectedSumResult, "2 ขห ", frontNumber, make(map[string]bool))

	for plateNumber, _ := range plateNumbers {
		fmt.Println(plateNumber)
	}

	frontNumber = 10
	plateNumbers = getPlateNumbers(targetNumbers, perfectedSumResult, "2 ขอ ", frontNumber, make(map[string]bool))
	for plateNumber, _ := range plateNumbers {
		fmt.Println(plateNumber)
	}
}

func getPlateNumbers(targetNumbers []int, perfectedSumResult map[int]bool, plateNumber string, sum int, plateNumbers map[string]bool) map[string]bool {
	for _, targetNumber := range targetNumbers {
		newPlateNumber := plateNumber
		sumNewPlate := sum
		newPlateNumber = plateNumber + strconv.Itoa(targetNumber)
		sumNewPlate += targetNumber
		if len(newPlateNumber) >= 14 {
			return plateNumbers
		}
		if perfectedSumResult[sumNewPlate] && !plateNumbers[newPlateNumber] {
			fmt.Println("received new plateNumber", newPlateNumber, "with sum", sumNewPlate)
			plateNumbers[newPlateNumber] = true
		}
		getPlateNumbers(targetNumbers, perfectedSumResult, newPlateNumber, sumNewPlate, plateNumbers)
	}
	return plateNumbers
}

func buildPerfectResult() map[int]bool {
	perfectedSumResult := make(map[int]bool)
	perfectedSumResultList := []int{9, 14, 15, 19, 24, 36, 40, 41, 42, 44, 45, 46, 50, 51, 54, 55, 56, 59, 63, 64, 65}

	for _, result := range perfectedSumResultList {
		perfectedSumResult[result] = true
	}
	return perfectedSumResult
}
