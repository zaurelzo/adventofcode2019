package main

import (
	"fmt"
	"log"
	"strconv"
)

func diagnosticCode(programCode []int, input int) ([]int, error) {
	printedValue := make([]int, 0)
	output := make([]int, len(programCode))
	copy(output, programCode)
	instructionPointer := 0
	halt := false
	for !halt && instructionPointer < len(output) {
		var instAsByte []byte = formatInstruction(output[instructionPointer])
		instCode, _ := strconv.Atoi(string(instAsByte[3:]))
		switch instCode {
		// add
		case 1:
			output[output[instructionPointer+3]] = getValue(output, instructionPointer+1, instAsByte[2]) + getValue(output, instructionPointer+2, instAsByte[1])
			instructionPointer += 4
		//mul
		case 2:
			output[output[instructionPointer+3]] = getValue(output, instructionPointer+1, instAsByte[2]) * getValue(output, instructionPointer+2, instAsByte[1])
			instructionPointer += 4
		//set input
		case 3:
			output[output[instructionPointer+1]] = input
			instructionPointer += 2
		//print output
		case 4:
			printedValue = append(printedValue, getValue(output, instructionPointer+1, instAsByte[2]))
			fmt.Printf("%d\n", getValue(output, instructionPointer+1, instAsByte[2]))
			instructionPointer += 2
		//jump if true
		case 5:
			if getValue(output, instructionPointer+1, instAsByte[2]) != 0 {
				instructionPointer = getValue(output, instructionPointer+2, instAsByte[1])
			} else {
				instructionPointer += 3
			}
		//jump if false
		case 6:
			if getValue(output, instructionPointer+1, instAsByte[2]) == 0 {
				instructionPointer = getValue(output, instructionPointer+2, instAsByte[1])
			} else {
				instructionPointer += 3
			}
		//less than
		case 7:
			val1 := getValue(output, instructionPointer+1, instAsByte[2])
			val2 := getValue(output, instructionPointer+2, instAsByte[1])
			if val1 < val2 {
				output[output[instructionPointer+3]] = 1
			} else {
				output[output[instructionPointer+3]] = 0
			}
			instructionPointer += 4
		//equal
		case 8:
			if getValue(output, instructionPointer+1, instAsByte[2]) == getValue(output, instructionPointer+2, instAsByte[1]) {
				output[output[instructionPointer+3]] = 1
			} else {
				output[output[instructionPointer+3]] = 0
			}
			instructionPointer += 4
			//halt
		case 99:
			halt = true
			//instruction not supported
		default:
			return []int{}, fmt.Errorf("WARN, UNKNOWN INSTRUCTION %+v, extracted from %+v at index %+v ", instCode, output[instructionPointer], instructionPointer)
			//index++
			//panic("unknow instruction")
		}
	}
	return printedValue, nil
}

//transform an instruction to follow the pattern ABCDE, with A being optional
func formatInstruction(instruction int) []byte {
	var abcdeInstruction = make([]byte, 5)
	codeAsByte := strconv.Itoa(instruction)
	for i := 0; i < 5-len(codeAsByte); i++ {
		abcdeInstruction[i] = byte('0')
	}

	for i := 0; i < len(codeAsByte); i++ {
		abcdeInstruction[i+5-len(codeAsByte)] = codeAsByte[i]
	}

	return abcdeInstruction
}

func getValue(array []int, index int, mode byte) int {
	switch mode {
	case byte('0'):
		return array[array[index]]
	case byte('1'):
		return array[index]
	default:
		panic("unknow mode ")
	}

}

func main() {
	input := []int{3, 225, 1, 225, 6, 6, 1100, 1, 238, 225, 104, 0, 1102, 7, 85, 225, 1102, 67, 12, 225, 102, 36, 65, 224, 1001, 224, -3096, 224, 4, 224, 1002, 223, 8, 223, 101, 4, 224, 224, 1, 224, 223, 223, 1001, 17, 31, 224, 1001, 224, -98, 224, 4, 224, 1002, 223, 8, 223, 101, 5, 224, 224, 1, 223, 224, 223, 1101, 86, 19, 225, 1101, 5, 27, 225, 1102, 18, 37, 225, 2, 125, 74, 224, 1001, 224, -1406, 224, 4, 224, 102, 8, 223, 223, 101, 2, 224, 224, 1, 224, 223, 223, 1102, 13, 47, 225, 1, 99, 14, 224, 1001, 224, -98, 224, 4, 224, 102, 8, 223, 223, 1001, 224, 2, 224, 1, 224, 223, 223, 1101, 38, 88, 225, 1102, 91, 36, 224, 101, -3276, 224, 224, 4, 224, 1002, 223, 8, 223, 101, 3, 224, 224, 1, 224, 223, 223, 1101, 59, 76, 224, 1001, 224, -135, 224, 4, 224, 102, 8, 223, 223, 1001, 224, 6, 224, 1, 223, 224, 223, 101, 90, 195, 224, 1001, 224, -112, 224, 4, 224, 102, 8, 223, 223, 1001, 224, 7, 224, 1, 224, 223, 223, 1102, 22, 28, 225, 1002, 69, 47, 224, 1001, 224, -235, 224, 4, 224, 1002, 223, 8, 223, 101, 5, 224, 224, 1, 223, 224, 223, 4, 223, 99, 0, 0, 0, 677, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1105, 0, 99999, 1105, 227, 247, 1105, 1, 99999, 1005, 227, 99999, 1005, 0, 256, 1105, 1, 99999, 1106, 227, 99999, 1106, 0, 265, 1105, 1, 99999, 1006, 0, 99999, 1006, 227, 274, 1105, 1, 99999, 1105, 1, 280, 1105, 1, 99999, 1, 225, 225, 225, 1101, 294, 0, 0, 105, 1, 0, 1105, 1, 99999, 1106, 0, 300, 1105, 1, 99999, 1, 225, 225, 225, 1101, 314, 0, 0, 106, 0, 0, 1105, 1, 99999, 107, 226, 226, 224, 102, 2, 223, 223, 1006, 224, 329, 1001, 223, 1, 223, 1107, 677, 226, 224, 1002, 223, 2, 223, 1005, 224, 344, 101, 1, 223, 223, 108, 677, 226, 224, 102, 2, 223, 223, 1006, 224, 359, 101, 1, 223, 223, 7, 677, 226, 224, 102, 2, 223, 223, 1005, 224, 374, 101, 1, 223, 223, 1008, 677, 226, 224, 1002, 223, 2, 223, 1006, 224, 389, 1001, 223, 1, 223, 7, 226, 677, 224, 102, 2, 223, 223, 1005, 224, 404, 101, 1, 223, 223, 1007, 226, 226, 224, 102, 2, 223, 223, 1006, 224, 419, 101, 1, 223, 223, 7, 226, 226, 224, 102, 2, 223, 223, 1005, 224, 434, 1001, 223, 1, 223, 8, 226, 226, 224, 1002, 223, 2, 223, 1006, 224, 449, 101, 1, 223, 223, 1007, 677, 677, 224, 102, 2, 223, 223, 1006, 224, 464, 101, 1, 223, 223, 1007, 226, 677, 224, 1002, 223, 2, 223, 1006, 224, 479, 101, 1, 223, 223, 108, 226, 226, 224, 102, 2, 223, 223, 1005, 224, 494, 1001, 223, 1, 223, 1108, 677, 677, 224, 102, 2, 223, 223, 1005, 224, 509, 1001, 223, 1, 223, 107, 226, 677, 224, 1002, 223, 2, 223, 1005, 224, 524, 101, 1, 223, 223, 1108, 677, 226, 224, 1002, 223, 2, 223, 1005, 224, 539, 1001, 223, 1, 223, 1008, 677, 677, 224, 1002, 223, 2, 223, 1006, 224, 554, 101, 1, 223, 223, 1008, 226, 226, 224, 102, 2, 223, 223, 1005, 224, 569, 1001, 223, 1, 223, 8, 677, 226, 224, 102, 2, 223, 223, 1006, 224, 584, 101, 1, 223, 223, 107, 677, 677, 224, 102, 2, 223, 223, 1006, 224, 599, 101, 1, 223, 223, 8, 226, 677, 224, 102, 2, 223, 223, 1006, 224, 614, 101, 1, 223, 223, 1107, 226, 677, 224, 102, 2, 223, 223, 1006, 224, 629, 101, 1, 223, 223, 108, 677, 677, 224, 1002, 223, 2, 223, 1005, 224, 644, 1001, 223, 1, 223, 1107, 226, 226, 224, 102, 2, 223, 223, 1005, 224, 659, 101, 1, 223, 223, 1108, 226, 677, 224, 102, 2, 223, 223, 1005, 224, 674, 101, 1, 223, 223, 4, 223, 99, 226}
	//input := []int{3, 0, 4, 0, 99}
	//input := []int{1002, 4, 3, 4, 33}
	partOnePrintedValues, err := diagnosticCode(input, 1)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Part one, response, diagnostic code %d \n", partOnePrintedValues[len(partOnePrintedValues)-1])

	partTwoPrintedValues, err2 := diagnosticCode(input, 5)
	if err2 != nil {
		log.Fatal(err2)
	}
	fmt.Printf("Part two, response, diagnostic code %d \n", partTwoPrintedValues[len(partTwoPrintedValues)-1])
}
