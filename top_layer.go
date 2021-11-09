package main

import "strings"

func top_layer_get_moves(cube Cube) []Cube {
	moves := make([]Cube, 0)
	sequences := make([]string, 0)

	sequences = append(sequences, "U R U' L' U R' U' L")
	sequences = append(sequences, "U B U' F' U B' U' F")
	sequences = append(sequences, "U L U' R' U L' U' R")
	sequences = append(sequences, "U F U' B' U F' U' B")

	for i := range sequences {
		temp := copy(cube)
		temp = sequence(temp, sequences[i])
		moves = append(moves, temp)
	}
	return moves
}

func top_layer_check(cube Cube) bool {
	if !strings.Contains(get_corner_value(cube.pattern, "URF"), "R") ||
		!strings.Contains(get_corner_value(cube.pattern, "URF"), "F") {
		return false
	}
	if !strings.Contains(get_corner_value(cube.pattern, "URB"), "R") ||
		!strings.Contains(get_corner_value(cube.pattern, "URB"), "B") {
		return false
	}
	if !strings.Contains(get_corner_value(cube.pattern, "UFL"), "L") ||
		!strings.Contains(get_corner_value(cube.pattern, "UFL"), "F") {
		return false
	}
	if !strings.Contains(get_corner_value(cube.pattern, "ULB"), "L") ||
		!strings.Contains(get_corner_value(cube.pattern, "ULB"), "B") {
		return false
	}

	return true
}

func top_layer_heuristic(cube Cube) int {
	return 0
}
