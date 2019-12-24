package main

import (
	"fmt"
	"strconv"
)

func diagnosticCode(programCode []int, inputChan chan int, outputChan chan int) ([]int, error) {
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
			//use golang chan to read input. Since reading from chan is a blocking operation
			//that feed perfectly the definition of blocking input describe by the challenge to read the signal input
			inputValue := <-inputChan
			output[output[instructionPointer+1]] = inputValue
			instructionPointer += 2
		//print output
		case 4:
			printedValue = append(printedValue, getValue(output, instructionPointer+1, instAsByte[2]))
			//fmt.Printf("%d\n", getValue(output, instructionPointer+1, instAsByte[2]))
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
			//panic("unknow instruction")
		}
	}
	//send  output signal
	outputChan <- printedValue[len(printedValue)-1]
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

// code taken from https://stackoverflow.com/questions/30226438/generate-all-permutations-in-go
func permutations(arr []int) [][]int {
	var helper func([]int, int)
	res := [][]int{}

	helper = func(arr []int, n int) {
		if n == 1 {
			tmp := make([]int, len(arr))
			copy(tmp, arr)
			res = append(res, tmp)
		} else {
			for i := 0; i < n; i++ {
				helper(arr, n-1)
				if n%2 == 1 {
					tmp := arr[i]
					arr[i] = arr[n-1]
					arr[n-1] = tmp
				} else {
					tmp := arr[0]
					arr[0] = arr[n-1]
					arr[n-1] = tmp
				}
			}
		}
	}
	helper(arr, len(arr))
	return res
}

func computeAmplifiersValue(programCode []int, phaseSettingValues []int) int {
	amplifierValues := make([]int, 0)
	amplifierValues = append(amplifierValues, 0)
	//use chan to simulate signal input and output of amplifier
	//each amplifier is represented by a goroutine and the signal mechanism is simulate by chan
	//a output chan of amplifier is the input chan of the next amplifier
	//Because reading from a chan and writing to a chan block when the other side of the chan is not
	//ready, this mechanism naturaly synchronise amplifiers (the first amplifier execute completely,
	// then the second amplifer and son on)
	var arrayInputOutputChan []pairOfChan = initializePairOfChan(len(phaseSettingValues))
	for ampNumber, phaseSetting := range phaseSettingValues {
		go diagnosticCode(programCode, arrayInputOutputChan[ampNumber].inputChan, arrayInputOutputChan[ampNumber].outputChan)
		//provide input signal of the current amplifier
		arrayInputOutputChan[ampNumber].inputChan <- phaseSetting
		arrayInputOutputChan[ampNumber].inputChan <- amplifierValues[len(amplifierValues)-1]
		//get the output signal of the current amplifier and save this value the use it as the input signal of the next
		//amplifier
		amplifierValues = append(amplifierValues, <-arrayInputOutputChan[ampNumber].outputChan)
	}
	return amplifierValues[len(amplifierValues)-1]
}

type pairOfChan struct {
	inputChan  chan int
	outputChan chan int
}

func initializePairOfChan(nbOfElement int) []pairOfChan {
	array := make([]pairOfChan, nbOfElement)
	for i := 0; i < nbOfElement; i++ {
		if i == 0 {
			array[i] = pairOfChan{inputChan: make(chan int), outputChan: make(chan int)}
		} else {
			array[i] = pairOfChan{inputChan: array[i-1].outputChan, outputChan: make(chan int)}
		}
	}
	return array
}

//solve part one
func maxAmplifiersValue(programCode []int) int {
	var max *int
	for _, phaseSettingValues := range permutations([]int{0, 1, 2, 3, 4}) {
		var amplifierValue int = computeAmplifiersValue(programCode, phaseSettingValues)
		if max == nil || amplifierValue > *max {
			max = &amplifierValue
		}
	}
	if max == nil {
		panic("No max found")
	}
	return *max
}

func main() {
	input := []int{3, 8, 1001, 8, 10, 8, 105, 1, 0, 0, 21, 30, 47, 64, 81, 98, 179, 260, 341, 422, 99999, 3, 9, 1001, 9, 5, 9, 4, 9, 99, 3, 9, 1002, 9, 5, 9, 101, 4, 9, 9, 102, 2, 9, 9, 4, 9, 99, 3, 9, 102, 3, 9, 9, 101, 2, 9, 9, 1002, 9, 3, 9, 4, 9, 99, 3, 9, 1001, 9, 5, 9, 1002, 9, 3, 9, 1001, 9, 3, 9, 4, 9, 99, 3, 9, 1002, 9, 3, 9, 101, 2, 9, 9, 102, 5, 9, 9, 4, 9, 99, 3, 9, 1001, 9, 2, 9, 4, 9, 3, 9, 1001, 9, 2, 9, 4, 9, 3, 9, 1001, 9, 2, 9, 4, 9, 3, 9, 1001, 9, 2, 9, 4, 9, 3, 9, 101, 2, 9, 9, 4, 9, 3, 9, 1001, 9, 2, 9, 4, 9, 3, 9, 1001, 9, 1, 9, 4, 9, 3, 9, 1001, 9, 2, 9, 4, 9, 3, 9, 101, 2, 9, 9, 4, 9, 3, 9, 1001, 9, 2, 9, 4, 9, 99, 3, 9, 101, 2, 9, 9, 4, 9, 3, 9, 102, 2, 9, 9, 4, 9, 3, 9, 1001, 9, 1, 9, 4, 9, 3, 9, 1001, 9, 2, 9, 4, 9, 3, 9, 102, 2, 9, 9, 4, 9, 3, 9, 1002, 9, 2, 9, 4, 9, 3, 9, 102, 2, 9, 9, 4, 9, 3, 9, 102, 2, 9, 9, 4, 9, 3, 9, 102, 2, 9, 9, 4, 9, 3, 9, 1002, 9, 2, 9, 4, 9, 99, 3, 9, 101, 1, 9, 9, 4, 9, 3, 9, 101, 2, 9, 9, 4, 9, 3, 9, 101, 2, 9, 9, 4, 9, 3, 9, 101, 1, 9, 9, 4, 9, 3, 9, 101, 2, 9, 9, 4, 9, 3, 9, 1002, 9, 2, 9, 4, 9, 3, 9, 101, 2, 9, 9, 4, 9, 3, 9, 1001, 9, 1, 9, 4, 9, 3, 9, 102, 2, 9, 9, 4, 9, 3, 9, 1002, 9, 2, 9, 4, 9, 99, 3, 9, 1001, 9, 2, 9, 4, 9, 3, 9, 1001, 9, 1, 9, 4, 9, 3, 9, 102, 2, 9, 9, 4, 9, 3, 9, 101, 1, 9, 9, 4, 9, 3, 9, 1002, 9, 2, 9, 4, 9, 3, 9, 101, 1, 9, 9, 4, 9, 3, 9, 102, 2, 9, 9, 4, 9, 3, 9, 1001, 9, 2, 9, 4, 9, 3, 9, 1002, 9, 2, 9, 4, 9, 3, 9, 101, 2, 9, 9, 4, 9, 99, 3, 9, 1001, 9, 2, 9, 4, 9, 3, 9, 101, 2, 9, 9, 4, 9, 3, 9, 1001, 9, 1, 9, 4, 9, 3, 9, 1002, 9, 2, 9, 4, 9, 3, 9, 101, 1, 9, 9, 4, 9, 3, 9, 1001, 9, 1, 9, 4, 9, 3, 9, 1002, 9, 2, 9, 4, 9, 3, 9, 101, 1, 9, 9, 4, 9, 3, 9, 1002, 9, 2, 9, 4, 9, 3, 9, 102, 2, 9, 9, 4, 9, 99}
	fmt.Printf("Part one, highest signal %d", maxAmplifiersValue(input))
}
