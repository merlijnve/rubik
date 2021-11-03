package main

// import (
// 	"fmt"
// )

func top_cross_get_moves(cube Cube) []Cube {
	moves := make([]Cube, 0)
	sequences := make([]string, 0)

	sequences = append(sequences, "U")
	sequences = append(sequences, "U'")
	sequences = append(sequences, "F R U R' U' F'")
	sequences = append(sequences, "U R U R' U R U' U' R' U")
	sequences = append(sequences, "U B U B' U B U' U' B' U")
	sequences = append(sequences, "U L U L' U L U' U' L' U")
	sequences = append(sequences, "U F U F' U F U' U' F' U")

	// fmt.Println(cube.pattern[1],cube.pattern[3],cube.pattern[5],cube.pattern[7])
	if cube.pattern[3] == "U" && cube.pattern[7] != "U" {
		sequences = nil
		sequences = append(sequences, "F R U R' U' F'")
	}
	if cube.pattern[7] == "U" && cube.pattern[5] != "U" {
		sequences = nil
		sequences = append(sequences, "R B U B' U' R'")
	}
	if cube.pattern[5] == "U" && cube.pattern[1] != "U" {
		sequences = nil
		sequences = append(sequences, "B L U L' U' B'")
	}
	if cube.pattern[1] == "U" && cube.pattern[3] != "U" {
		sequences = nil
		sequences = append(sequences, "L F U F' U' L'")
	}

	for i := range sequences {
		temp := copy(cube)
		temp = sequence(temp, sequences[i])
		moves = append(moves, temp)
	}
	return moves
}

func top_cross_check(cube Cube) bool {
	if get_edge_value(cube.pattern, "UR") != "UR" {
		return false
	}
	if get_edge_value(cube.pattern, "UF") != "UF" {
		return false
	}
	if get_edge_value(cube.pattern, "UL") != "UL" {
		return false
	}
	if get_edge_value(cube.pattern, "UB") != "UB" {
		return false
	}
	return true
}

func top_cross_heuristic(cube Cube) int {
	return 0
}
