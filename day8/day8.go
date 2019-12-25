package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"
)

func createLayers(fileName string, wide int, tall int) [][]int {
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	arrayOfLayer := make([][]int, 0)
	content, err := ioutil.ReadAll(file)
	layer := 1
	for index := 0; index < len(content); index, layer = wide*tall*layer, layer+1 {
		subLayer := make([]int, wide*tall)
		for j := index; j < wide*tall*layer; j++ {
			value, _ := strconv.Atoi(string(content[j]))
			subLayer[j%(wide*tall)] = value
		}
		arrayOfLayer = append(arrayOfLayer, subLayer)
	}
	/*for index := 0; index < len(content); index++ {
		array = append(array, content[index])
	}*/
	return arrayOfLayer
}

func countZeroOneAndTwoDigits(array []int) (*int, int, int) {
	var zero *int
	one := 0
	two := 0
	for index := 0; index < len(array); index++ {
		switch array[index] {
		case 0:
			if zero == nil {
				z := 1
				zero = &z
			} else {
				(*zero)++
			}
		case 1:
			one++
		case 2:
			two++
		}

	}
	return zero, one, two
}

//solve part one
func countByLayer(arrayOfLayers [][]int) int {
	var minZero *int
	var uno int
	var dos int
	for _, layer := range arrayOfLayers {
		zero, one, two := countZeroOneAndTwoDigits(layer)
		if (minZero == nil && zero != nil) || (zero != nil && *zero < *minZero) {
			minZero, uno, dos = zero, one, two
		}
	}
	if minZero == nil {
		panic("no zero found")
	}
	return uno * dos
}
func combineLayer(arrayOfLayers [][]int) []int {
	finalImage := make([]int, len(arrayOfLayers[0]))
	for indexDigit := 0; indexDigit < len(arrayOfLayers[0]); indexDigit++ {
		currentPixel := make([]int, 0)
		for _, layer := range arrayOfLayers {
			currentPixel = append(currentPixel, layer[indexDigit])
		}
		finalImage[indexDigit] = choosePixelColor(currentPixel)
	}
	return finalImage
}

func choosePixelColor(pixels []int) int {
	if pixels[0] == 0 || pixels[0] == 1 {
		return pixels[0]
	}
	indexNoTransparentPixel := 1
	for indexNoTransparentPixel < len(pixels) && pixels[indexNoTransparentPixel] == 2 {
		indexNoTransparentPixel++
	}
	if indexNoTransparentPixel == len(pixels) {
		return 2
	}
	return pixels[indexNoTransparentPixel]
}

func printMessage(pixels []int, wide int) {
	for i, pixel := range pixels {
		if i > 0 && i%(wide) == 0 {
			fmt.Print("\n")
		}
		if pixel != 2 {
			if pixel == 0 {
				fmt.Printf(" ")
			} else {
				fmt.Printf("#")
			}
		}
	}
	//fmt.Println("end")
}
func main() {

	fmt.Printf("Part one, number of 1 digits multiplied by the number of 2 = %+v \n",
		countByLayer(createLayers("input", 25, 6)))

	message := combineLayer(createLayers("input", 25, 6))
	fmt.Printf("Part two, message :  \n")
	printMessage(message, 25)

	/*message := combineLayer(createLayers("input2", 2, 2))
	fmt.Printf("Part two, message :  \n")
	printMessage(message, 2)
	*/
}
