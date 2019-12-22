package main

import (
	"fmt"
	"strconv"
	"strings"
)

func numberOfpassword(passwordRange string, meetCriteria func(int) bool) int {
	counter := 0
	splitResult := strings.Split(passwordRange, "-")
	lowerBound, _ := strconv.Atoi(splitResult[0])
	upperBound, _ := strconv.Atoi(splitResult[1])
	for currentPassword := lowerBound; currentPassword < upperBound; currentPassword++ {
		if meetCriteria(currentPassword) {
			counter++
		}
	}
	return counter
}

func twoAdjacentDigitsAreSame(password int) bool {
	digit := strconv.Itoa(password)
	for index := 0; index < len(digit)-1; index++ {
		if digit[index] == digit[index+1] {
			return true
		}
	}
	return false
}

func digitNeverDecrease(password int) bool {
	digit := strconv.Itoa(password)
	for index := 0; index < len(digit)-1; index++ {
		if digit[index] > digit[index+1] {
			return false
		}
	}
	return true
}

func containsExactlyTwoConsecutiveDigits(password int) bool {
	digit := strconv.Itoa(password)
	mapDigitToPosition := make(map[byte][]int)
	for index := 0; index < len(digit); index++ {
		mapDigitToPosition[digit[index]] = append(mapDigitToPosition[digit[index]], index)
	}
	for _, listPosOccurences := range mapDigitToPosition {
		if len(listPosOccurences) == 2 && listPosOccurences[0]+1 == listPosOccurences[1] {
			return true
		}
	}
	return false
}

func main() {
	partOneCriteriaFunction := func(password int) bool {
		return twoAdjacentDigitsAreSame(password) && digitNeverDecrease(password)
	}
	fmt.Printf(" Part one, number of passwords : %d \n", numberOfpassword("130254-678275", partOneCriteriaFunction))

	partTwoCriteriaFunction := func(password int) bool {
		return digitNeverDecrease(password) &&
			containsExactlyTwoConsecutiveDigits(password)
	}
	fmt.Printf("Part two :  number of passwords   %+v", numberOfpassword("130254-678275", partTwoCriteriaFunction))
}
