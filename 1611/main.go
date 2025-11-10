package main

import "fmt"

func main() {
	// minimumOneBitOperations(3)
	minimumOneBitOperations(13)
	// minimumOneBitOperations(53)
	// minimumOneBitOperations(63)
	// minimumOneBitOperations(51)
}

func minimumOneBitOperations(n int) int {
	var ans int

	for ; n > 0; n >>= 1 {
		beforeAns := ans
		ans ^= n
		fmt.Println(n, ans, beforeAns)
	}
	return ans
}
