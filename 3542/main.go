package main

func main() {
	minOperations([]int{3, 1, 2, 1})
}

func minOperations(nums []int) int {
	stack := make([]int, len(nums)+1)
	top, ans := 0, 0
	for _, num := range nums {
		for stack[top] > num {
			top--
			ans++
		}
		if stack[top] != num {
			top++
			stack[top] = num
		}
	}
	return ans + top
}
