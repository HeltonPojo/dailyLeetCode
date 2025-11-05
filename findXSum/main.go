package main

import "sort"

func main() {
}

func findXSum(nums []int, k int, x int) []int {
	n := len(nums)
	res := []int{}

	for i := 0; i <= n-k; i++ {
		freq := map[int]int{}
		for j := i; j < i+k; j++ {
			freq[nums[j]]++
		}

		type pair struct {
			num, count int
		}
		arr := []pair{}
		for num, count := range freq {
			arr = append(arr, pair{num, count})
		}

		// Ordena por frequência decrescente, e depois por valor decrescente
		sort.Slice(arr, func(i, j int) bool {
			if arr[i].count == arr[j].count {
				return arr[i].num > arr[j].num
			}
			return arr[i].count > arr[j].count
		})

		sum := 0
		for j := 0; j < len(arr) && j < x; j++ {
			sum += arr[j].num * arr[j].count
		}
		res = append(res, sum)
	}

	return res
}
