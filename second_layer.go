package main

func second_layer_get_moves(cube Cube) []Cube {
	moves := make([]Cube, 0)
	sequences := make([]string, 0)

	sequences = append(sequences, "U R U' R' U' F' U F")
	sequences = append(sequences, "U")
	sequences = append(sequences, "U'")

	if cube.pattern[19] == "F" {
		sequences = nil
		sequences = append(sequences, "U R U' R' U' F' U F")
		sequences = append(sequences, "U' L' U L U F U' F'")
	}
	if cube.pattern[19] == "R" {
		sequences = nil
		sequences = append(sequences, "U B U' B' U' R' U R")
		sequences = append(sequences, "U' F' U F U R U' R'")
	}
	if cube.pattern[19] == "B" {
		sequences = nil
		sequences = append(sequences, "U L U' L' U' B' U B")
		sequences = append(sequences, "U' R' U R U B U' B'")
	}
	if cube.pattern[19] == "L" {
		sequences = nil
		sequences = append(sequences, "U F U' F' U' L' U L")
		sequences = append(sequences, "U' B' U B U L U' L'")
	}
	
	for i := range sequences {
		temp := copy(cube)
		temp = sequence(temp, sequences[i])
		moves = append(moves, temp)
	}
	return moves
}

func second_layer_check(cube Cube) bool {
	if get_edge_value(cube.pattern, "FL") != "FL" {
		return false
	}
	if get_edge_value(cube.pattern, "RF") != "RF" {
		return false
	}
	if get_edge_value(cube.pattern, "RB") != "RB" {
		return false
	}
	if get_edge_value(cube.pattern, "LB") != "LB" {
		return false
	}
	return true
}

func second_layer_heuristic(cube Cube) int {
	return 0
}
