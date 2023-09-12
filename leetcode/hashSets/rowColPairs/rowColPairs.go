package main

import "fmt"

func colMaps(n [][]int) []map[int]int {
	rows, cols := len(n), len(n[0])

	m := make([]map[int]int, cols)

	for i := range m {
		m[i] = make(map[int]int)
	}

	for i := 0; i < cols; i++ {
		for j := 0; j < rows; j++ {
			// map at index i, key at j;  value=  list at value j, item at index i
			m[i][j] = n[j][i]
		}
	}
	return m
}
//  *** 2352. Equal Row and Column Pairs ***
// Given a 0-indexed n x n integer matrix grid, return the number of pairs (ri, cj) such that row ri and column cj are equal.
// A row and column pair is considered equal if they contain the same elements in the same order (i.e., an equal array).
func equalPairs(n [][]int) int {
	rows, cols := len(n), len(n[0])
	found := 0
	// num list loop
	for i := 0; i < rows; i++ {
		// only move cols when row not match
		for j := 0; j < cols; j++ {
			match := true
			for r := 0; r < rows; r++ {
				if n[i][r] != n[r][j] {
					match = false
					break
				}
			}
			if match {
				found++
			}
		}

	}
	return found
}

func main() {

	n := [][]int{{3, 1, 2, 2}, {1, 4, 4, 5}, {2, 4, 2, 2}, {2, 4, 2, 2}}

	fmt.Println(colMaps(n))
	fmt.Println(equalPairs(n))
}
