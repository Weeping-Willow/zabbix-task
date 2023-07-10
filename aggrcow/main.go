package main

import (
	"fmt"
	"sort"
)

func main() {
	var testCount int

	// solution tester no likey
	//fmt.Print("Enter the number of test cases: ")
	_, err := fmt.Scan(&testCount)
	if err != nil {
		panic(fmt.Sprintf("failed to read number of test cases: %v", err))
	}

	for i := 0; i < testCount; i++ {
		cows, stallCount := 0, 0
		// solution tester no likey
		//fmt.Print("Enter the number of stalls and cows: ")
		_, err := fmt.Scan(&stallCount, &cows)
		if err != nil {
			panic(fmt.Sprintf("failed to read count of cows and/or stalls: %v", err))
		}

		stalls := make([]int, stallCount)
		for j := 0; j < stallCount; j++ {
			// solution tester no likey
			//fmt.Printf("Enter the stall %d: ", j+1)
			_, err := fmt.Scan(&stalls[j])
			if err != nil {
				panic(fmt.Sprintf("failed to read stall placement at index %d: %v", j, err))
			}
		}
		fmt.Println(findMaxSpace(cows, stalls))
	}
}

func findMaxSpace(cows int, stalls []int) int {
	sort.SliceStable(stalls, func(i, j int) bool {
		return stalls[i] < stalls[j]
	})

	low := 0
	high := stalls[len(stalls)-1] - stalls[0]
	maxDistanceBetween := 0

	for low <= high {
		median := (low + high) / 2

		if cowsFit(cows, median, stalls) {
			maxDistanceBetween = median
			low = median + 1
		} else {
			high = median - 1
		}
	}

	return maxDistanceBetween
}

func cowsFit(cows, spaceBetween int, stalls []int) bool {
	cowsCounted := 1
	lastStall := stalls[0]

	for _, stall := range stalls[1:] {
		if stall-lastStall >= spaceBetween {
			cowsCounted++
			lastStall = stall
		}

		if cowsCounted == cows {
			return true
		}
	}

	return false
}
