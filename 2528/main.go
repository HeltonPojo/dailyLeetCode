package main

func main() {
}

var (
	SUM  = make([]int, 100001)
	DIFF = make([]int, 100001)
)

func maxPower(stations []int, r, k int) int64 {
	minv, total := stations[0], 0
	for i, v := range stations {
		SUM[i] = 0
		minv = min(minv, v)
		total += v
	}
	n := len(stations)
	SUM[n] = 0
	for i := range n {
		left := max(i-r, 0)
		right := min(i+r+1, n)
		SUM[left] += stations[i]
		SUM[right] -= stations[i]
	}
	lo, hi := minv, total+k
	res := 0
	for lo <= hi {
		mid := lo + (hi-lo)/2
		if check(n, mid, r, k) {
			res = mid
			lo = mid + 1
		} else {
			hi = mid - 1
		}
	}
	return int64(res)
}

func check(n, val, r, k int) bool {
	copy(DIFF[:n], SUM[:n])
	sum, remaining := 0, k
	for i := range DIFF[:n] {
		sum += DIFF[i]
		if sum < val {
			add := val - sum
			if remaining < add {
				return false
			}
			remaining -= add
			end := min(2*r+i+1, n)
			DIFF[end] -= add
			sum += add
		}
	}
	return true
}
