package main

import "fmt"

func main() {
	fmt.Print(getSneakyNumbers([]int{1, 1, 1, 1, 1}))
	fmt.Print(getSneakyNumbers([]int{0, 1, 1, 0}))
	fmt.Print(getSneakyNumbers([]int{0, 3, 2, 1, 3, 2}))
	fmt.Print(getSneakyNumbers([]int{7, 1, 5, 4, 3, 4, 6, 0, 9, 5, 8, 2}))
}

func getSneakyNumbers(nums []int) []int {
	setMap := make(map[int]int)
	setOutput := make(map[int]int)
	var output []int

	for _, num := range nums {
		if setMap[num] == 0 {
			setMap[num] = 1
		} else if setOutput[num] == 0 {
			setOutput[num] = 1
			output = append(output, num)
		}
	}

	return output
}
