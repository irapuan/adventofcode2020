package secondpuzzle

import (
	"strconv"
	"strings"
)

func SecondPuzzleCalc(input string) int {
	list := strings.Split(input, "\n")

	for idx := 0; idx < len(list); idx++ {
		firstNumber, _ := strconv.Atoi(list[idx])
		for idx2 := 0; idx2 < len(list); idx2++ {
			secondNumber, _ := strconv.Atoi(list[idx2])
			for idx3 := 0; idx3 < len(list); idx3++ {
				thirdNumber, _ := strconv.Atoi(list[idx3])
				if (firstNumber + secondNumber + thirdNumber) == 2020 {
					return firstNumber * secondNumber * thirdNumber
				}
			}
		}
	}
	return 0
}
