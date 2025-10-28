package main

import "fmt"

func main() {
	fmt.Print(countValidSelections([]int{1, 0, 2, 0, 3}))
}

func countValidSelections(nums []int) int {
	count := 0
	for i, num := range nums {
		if num == 0 {
			copy1 := make([]int, len(nums))
			copy(copy1, nums)
			if validateSelctions(copy1, i, 1) {
				count++
			}

			copy2 := make([]int, len(nums))
			copy(copy2, nums)

			if validateSelctions(copy2, i, -1) {
				count++
			}
		}
	}
	return count
}

func validateSelctions(nums []int, init int, dir int) bool {
	curr := init
	fmt.Println(nums)
	for curr >= 0 && curr < len(nums) {
		if nums[curr] > 0 {
			dir = dir * (-1)
			nums[curr]--
		}
		curr += dir
	}

	countZeros := 0
	for _, num := range nums {
		if num == 0 {
			countZeros++
		}
	}
	return countZeros == len(nums)
}
