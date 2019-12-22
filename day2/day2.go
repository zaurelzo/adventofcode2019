package main

import "fmt"

func programAlarm(input []int, noun int, verb int) []int {
	output := make([]int, len(input))
	copy(output, input)
	output[1] = noun
	output[2] = verb
	index := 0
	halt := false
	for ; !halt && index < len(input); index += 4 {
		switch output[index] {
		case 1:
			output[output[index+3]] = output[output[index+1]] + output[output[index+2]]
		case 2:
			output[output[index+3]] = output[output[index+1]] * output[output[index+2]]
		case 99:
			halt = true
		default:
			halt = true
		}
	}
	return output
}

func findNounAndVerb(input []int) (int, int) {
	for noun := 0; noun <= 99; noun++ {
		for verb := 0; verb <= 99; verb++ {
			output := programAlarm(input, noun, verb)
			if output[0] == 19690720 {
				return noun, verb
			}
		}

	}
	return 0, 0
}

func main() {
	var input = []int{1, 0, 0, 3, 1, 1, 2, 3, 1, 3, 4, 3, 1, 5, 0, 3, 2, 6, 1, 19, 1, 19, 5, 23, 2, 9, 23, 27, 1, 5, 27, 31, 1, 5, 31, 35, 1, 35, 13, 39, 1, 39, 9, 43, 1, 5, 43, 47, 1, 47, 6, 51, 1, 51, 13, 55, 1, 55, 9, 59, 1, 59, 13, 63, 2, 63, 13, 67, 1, 67, 10, 71, 1, 71, 6, 75, 2, 10, 75, 79, 2, 10, 79, 83, 1, 5, 83, 87, 2, 6, 87, 91, 1, 91, 6, 95, 1, 95, 13, 99, 2, 99, 13, 103, 1, 103, 9, 107, 1, 10, 107, 111, 2, 111, 13, 115, 1, 10, 115, 119, 1, 10, 119, 123, 2, 13, 123, 127, 2, 6, 127, 131, 1, 13, 131, 135, 1, 135, 2, 139, 1, 139, 6, 0, 99, 2, 0, 14, 0}
	fmt.Printf(" part one : %d \n", programAlarm(input, 12, 2)[0])
	noun, verb := findNounAndVerb(input)
	fmt.Printf("part two : 100 * noun + verb = %d \n", 100*noun+verb)
}
