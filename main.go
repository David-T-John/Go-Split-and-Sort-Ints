package main

//Concurrency in Go - Module 3 Activity
import (
	"fmt"
	"math"
	"sort"
	"strconv"
	"strings"
	"sync"
)

func main() {
	// var outputSlice []string
	var input string

	var inputSlice []string
	var inputSliceInts []int

	var subSlice1 []int
	var subSlice2 []int
	var subSlice3 []int
	var subSlice4 []int

	var outputSortedSlice []int

	var wg sync.WaitGroup
	setSubSlices := func(inputSlice []int) {
		var length = int(len(inputSlice))
		var threequarter = int(math.Floor(float64(length) * 0.75))
		var half = int(math.Floor(float64(length) * 0.5))
		var quarter = int(math.Floor(float64(length) * 0.25))
		subSlice1 = inputSlice[0:quarter]
		subSlice2 = inputSlice[quarter:half]
		subSlice3 = inputSlice[half:threequarter]
		subSlice4 = inputSlice[threequarter:length]

		fmt.Println("Slice 1 will sort: ", subSlice1)
		fmt.Println("Slice 2 will sort: ", subSlice2)
		fmt.Println("Slice 3 will sort: ", subSlice3)
		fmt.Println("Slice 4 will sort: ", subSlice4)
	}

	var sortSubSlice = func(subSlice []int) []int {
		sort.Slice(subSlice, func(i, j int) bool { return subSlice[i] < subSlice[j] })
		wg.Done()
		fmt.Println(subSlice)
		return subSlice
	}

	wg.Add(4)
	fmt.Println("Enter a series of integers")
	fmt.Scanln(&input)
	fmt.Println("Your input: ", input)

	inputSlice = strings.Split(input, "")
	fmt.Println("Your input slice: ", inputSlice)

	for i, char := range inputSlice {
		num, err := strconv.Atoi(char)
		inputSliceInts = append(inputSliceInts, num)
		fmt.Println("Error: ", err)

		fmt.Println(i)
		fmt.Println(num)

	}
	fmt.Println("Your input slice integers: ", inputSliceInts)
	setSubSlices(inputSliceInts)
	go sortSubSlice(subSlice1)
	go sortSubSlice(subSlice2)
	go sortSubSlice(subSlice3)
	go sortSubSlice(subSlice4)
	// outputSlice = append(subSlice1, subSlice2, subSlice3, subSlice4)

	wg.Wait()
	sort.Ints(inputSliceInts)
	outputSortedSlice = inputSliceInts
	fmt.Println("Output Sorted Integers Slice: ", outputSortedSlice)

}
