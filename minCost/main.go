package main

import "fmt"

func main() {
	fmt.Println(minCost("abaac", []int{1, 2, 3, 4, 5}), 3)
}

func minCost(colors string, neededTime []int) int {
	n := len(colors)
	if n == 0 {
		return 0
	}
	res := 0
	prevTime := neededTime[0]

	for i := 1; i < n; i++ {
		curTime := neededTime[i]
		if colors[i] == colors[i-1] {
			if prevTime < curTime {
				res += prevTime
				prevTime = curTime
			} else {
				res += curTime
			}
		} else {
			prevTime = curTime
		}
	}
	return res
}
