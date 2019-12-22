package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"
)

func readModulesMass() []int {
	file, err := os.Open("day1-input")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	content, err := ioutil.ReadAll(file)
	elts := strings.Split(string(content), "\n")
	result := make([]int, len(elts))
	for index, value := range elts {
		i, _ := strconv.Atoi(value)
		result[index] = i
	}
	//fmt.Println(result)
	return result
}

func totalFuel(masses []int) int {
	sumOfFuel := 0
	for _, value := range readModulesMass() {
		sumOfFuel += (value/3 - 2)
	}
	return sumOfFuel
}

func totalFuelConsideringFuelRequirement(masses []int) int {
	sumOfFuel := 0
	for _, value := range masses {
		fuelOfFuel := value/3 - 2
		for ; fuelOfFuel > 0; fuelOfFuel = fuelOfFuel/3 - 2 {
			sumOfFuel += fuelOfFuel
		}
	}
	return sumOfFuel
}

func main() {
	fmt.Printf("part one :  %d \n", totalFuel(readModulesMass()))
	fmt.Printf("part two : fuelOfFuel  %d \n", totalFuelConsideringFuelRequirement(readModulesMass()))
}
