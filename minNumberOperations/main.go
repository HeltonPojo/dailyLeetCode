package main

import (
	"fmt"
)

func main() {
	fmt.Println("Saida: ", minNumberOperations([]int{1, 2}), "Expected: 2")
	fmt.Println("Saida: ", minNumberOperations([]int{1, 2, 3, 2, 1}), "Expected: 3")
	fmt.Println("Saida: ", minNumberOperations([]int{3, 1, 1, 2}), "Expected: 4")
	fmt.Println("Saida: ", minNumberOperations([]int{3, 1, 5, 4, 2}), "Expected: 7")
	fmt.Println("Saida: ", minNumberOperations([]int{3, 1, 1, 3}), "Expected: 5")
	fmt.Println("Saida: ", minNumberOperations([]int{8, 1, 1, 1}), "Expected: 8")
}

func minNumberOperations(target []int) int {
	if len(target) == 0 {
		return 0
	}

	operations := target[0]

	for i := 1; i < len(target); i++ {
		if target[i] > target[i-1] {
			operations += target[i] - target[i-1]
		}
	}

	return operations
}
