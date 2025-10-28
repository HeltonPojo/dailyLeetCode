package main

import (
	"fmt"
	"strings"
)

type TestCase struct {
	Banks  []string
	Target int
}

func main() {
	testCases := []TestCase{
		{[]string{"011001", "000000", "010100", "001000"}, 8},
	}
	for i, testCase := range testCases {
		output := numberOfBeams(testCase.Banks)
		if output == testCase.Target {
			fmt.Printf("%d Passou", i)
		} else {
			fmt.Printf("%d Erro com sainda %d", i, output)
		}
	}
}

func numberOfBeams(bank []string) int {
	lasers := 0
	lastLineSec := 0
	for i := range len(bank) {
		splitedStr := strings.Split(bank[i], "")
		lineSec := 0
		for _, char := range splitedStr {
			if char == "1" {
				lineSec++
			}
		}

		if lastLineSec > 0 && lineSec > 0 {
			lasers += lastLineSec * lineSec
			lastLineSec = lineSec
		} else if lineSec > 0 {
			lastLineSec = lineSec
		}
	}
	return lasers
}
