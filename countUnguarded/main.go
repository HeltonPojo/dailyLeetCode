package main

import "fmt"

func main() {
	fmt.Println(countUnguarded(4, 6, [][]int{{0, 0}, {1, 1}, {2, 3}}, [][]int{{0, 1}, {2, 2}, {1, 4}}), 7)
	fmt.Println(countUnguarded(3, 3, [][]int{{1, 1}}, [][]int{{0, 1}, {1, 0}, {2, 1}, {1, 2}}), 4)
}

func countUnguarded(m int, n int, guards [][]int, walls [][]int) int {
	const guardId = 2
	const wallId = 3

	freeSpaces := m * n
	// criar uma matriz e preencher ela com zeros
	matriz := make([][]int, m)
	for i := range matriz {
		matriz[i] = make([]int, n)
	}
	// preencher a matriz com os guardas e muros
	for _, guard := range guards {
		row := guard[0]
		col := guard[1]
		matriz[row][col] = guardId
		freeSpaces--
	}
	for _, wall := range walls {
		row := wall[0]
		col := wall[1]
		matriz[row][col] = wallId
		freeSpaces--
	}

	for _, guard := range guards {
		newProtecteds := protectSpace(guard, matriz, m, n)
		freeSpaces = freeSpaces - newProtecteds
	}

	return freeSpaces
}

func protectSpace(guardPos []int, matrix [][]int, m int, n int) int {
	pretecteds := 0
	originalRow := guardPos[0]
	originalCol := guardPos[1]

	currCol := originalCol + 1
	// x rigth
	for currCol < n && matrix[originalRow][currCol] < 2 {
		if matrix[originalRow][currCol] == 0 {
			pretecteds++
			matrix[originalRow][currCol] = 1
		}
		currCol++
	}

	currCol = originalCol - 1
	// x left
	for currCol >= 0 && matrix[originalRow][currCol] < 2 {
		if matrix[originalRow][currCol] == 0 {
			pretecteds++
			matrix[originalRow][currCol] = 1
		}
		currCol--
	}

	currRow := originalRow + 1
	// y down
	for currRow < m && matrix[currRow][originalCol] < 2 {
		if matrix[currRow][originalCol] == 0 {
			pretecteds++
			matrix[currRow][originalCol] = 1
		}
		currRow++
	}

	currRow = originalRow - 1
	// y up
	for currRow >= 0 && matrix[currRow][originalCol] < 2 {
		if matrix[currRow][originalCol] == 0 {
			pretecteds++
			matrix[currRow][originalCol] = 1
		}
		currRow--
	}

	return pretecteds
}
