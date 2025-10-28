package main

import "fmt"

// the fastest way of doing this is to split array
// divide o array no meio e olha o 0 da segunda parte
// se for maior ele continua a busca na segunda
// se for menor ele repete a busca na primeira
// se for igual ele devolve

type TestCase struct {
	Nums   []int
	Target int
	Output int
}

func main() {
	testCases := []TestCase{
		{[]int{1, 3, 5, 6}, 5, 2},
		{[]int{1, 3, 5, 6}, 2, 1},
		{[]int{1, 3, 5, 6}, 4, 2},
		{[]int{1, 3, 5, 6}, 7, 4},
		{[]int{2, 3, 5, 6}, 1, 0},
	}

	for _, testCase := range testCases {
		value := searchInsert(testCase.Nums, testCase.Target)

		fmt.Printf("Valor: %d, Saida: %d", testCase.Output, value)
		if value == testCase.Output {
			fmt.Print(" Passou")
		} else {
			fmt.Print(" Não Passou")
		}
		fmt.Println("")
	}
}

func searchInsert(nums []int, target int) int {
	for i, num := range nums {
		if num == target {
			return i
		}
		if target < num && i > 0 {
			return i
		} else if target < num {
			return 0
		}
	}
	return len(nums)
}
