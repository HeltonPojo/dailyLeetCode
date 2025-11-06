package main

import (
	"fmt"
)

func main() {
	fmt.Println(processQueries(5, [][]int{{1, 2}, {2, 3}, {3, 4}, {4, 5}}, [][]int{{1, 3}, {2, 1}, {1, 1}, {2, 2}, {1, 2}}))
	fmt.Println(processQueries(5, [][]int{}, [][]int{{1, 1}, {2, 1}, {1, 1}}))
	fmt.Println(processQueries(2, [][]int{{1, 2}}, [][]int{{1, 1}, {1, 2}, {1, 2}, {2, 2}, {2, 2}, {1, 1}, {1, 2}, {1, 1}}))
}

func processQueries(c int, connections [][]int, queries [][]int) []int {
	if c == 0 {
		return []int{}
	}

	// Grafo de adjacência (não precisa estar ordenado)
	graph := make([][]int, c+1)
	online := make([]bool, c+1)

	// Inicializa todos como online
	for i := 1; i <= c; i++ {
		online[i] = true
	}

	// Constrói o grafo bidirecional
	for _, conn := range connections {
		graph[conn[0]] = append(graph[conn[0]], conn[1])
		graph[conn[1]] = append(graph[conn[1]], conn[0])
	}

	result := make([]int, 0, len(queries))

	for _, query := range queries {
		if query[0] == 2 {
			// Marca como offline
			online[query[1]] = false
		} else {
			// Busca o maior vizinho online usando BFS
			maxOnline := bfsMaxOnline(query[1], graph, online, c)
			result = append(result, maxOnline)
		}
	}

	return result
}

func bfsMaxOnline(start int, graph [][]int, online []bool, c int) int {
	// Se o próprio nó está online, retorna ele
	if online[start] {
		return start
	}

	// BFS para encontrar o maior vizinho online
	visited := make([]bool, c+1)
	queue := []int{start}
	visited[start] = true
	maxOnline := -1

	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]

		// Verifica os vizinhos
		for _, neighbor := range graph[node] {
			if visited[neighbor] {
				continue
			}
			visited[neighbor] = true

			if online[neighbor] {
				// Encontrou um vizinho online
				if neighbor > maxOnline {
					maxOnline = neighbor
				}
			} else {
				// Se está offline, adiciona na fila para continuar buscando
				queue = append(queue, neighbor)
			}
		}
	}

	return maxOnline
}
