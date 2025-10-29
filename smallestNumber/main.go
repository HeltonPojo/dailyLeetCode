package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	smallestNumber(5)
	smallestNumber(10)
	smallestNumber(3)
	smallestNumber(1000)
	smallestNumber(999)
}

func smallestNumber(n int) int {
	binaryString := strconv.FormatInt(int64(n), 2)
	allOnes := strings.ReplaceAll(binaryString, "0", "1")
	out, err := strconv.ParseInt(allOnes, 2, 64)
	if err != nil {
		return 0
	}
	fmt.Println(int(out))
	return int(out)
}
