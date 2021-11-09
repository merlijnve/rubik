package main

func second_layer_get_moves(cube Cube) []Cube {
	moves := make([]Cube, 0)
	sequences := make([]string, 0)

	// 0 1 2
	// 3 4 5
	// 6 7 8

	if cube.pattern[19] == "F" && cube.pattern[7] == "R" {
		// sequences = nil
		sequences = append(sequences, "U R U' R' U' F' U F")
	} else if cube.pattern[19] == "F" && cube.pattern[7] == "L" {
		// sequences = nil
		sequences = append(sequences, "U' L' U L U F U' F'")
	}
	if cube.pattern[10] == "R" && cube.pattern[5] == "B" {
		// sequences = nil
		sequences = append(sequences, "U B U' B' U' R' U R")
	} else if cube.pattern[10] == "R" && cube.pattern[5] == "F" {
		// sequences = nil
		sequences = append(sequences, "U' F' U F U R U' R'")
	}
	if cube.pattern[46] == "B" && cube.pattern[1] == "L" {
		// sequences = nil
		sequences = append(sequences, "U L U' L' U' B' U B")
	} else if cube.pattern[46] == "B" && cube.pattern[1] == "R" {
		// sequences = nil
		sequences = append(sequences, "U' R' U R U B U' B'")
	}
	if cube.pattern[37] == "L" && cube.pattern[3] == "F" {
		// sequences = nil
		sequences = append(sequences, "U F U' F' U' L' U L")
	} else if cube.pattern[37] == "L" && cube.pattern[3] == "B" {
		// sequences = nil
		sequences = append(sequences, "U' B' U B U L U' L'")
	}

	if len(sequences) == 0 {
		if get_edge_value(cube.pattern, "RF") == "FR" {
			sequences = append(sequences, "U R U' R' U' F' U F U U U R U' R' U' F' U F")
		} else if get_edge_value(cube.pattern, "RB") == "BR" {
			sequences = append(sequences, "U B U' B' U' R' U R U U U B U' B' U' R' U R")
		} else if get_edge_value(cube.pattern, "LB") == "BL" {
			sequences = append(sequences, "U L U' L' U' B' U B U U U L U' L' U' B' U B")
		} else if get_edge_value(cube.pattern, "FL") == "LF" {
			sequences = append(sequences, "U F U' F' U' L' U L U U U F U' F' U' L' U L")
		} else if get_edge_value(cube.pattern, "RF") != "RF" {
			sequences = append(sequences, "U R U' R' U' F' U F")
			sequences = append(sequences, "U' F' U F U R U' R'")
		} else if get_edge_value(cube.pattern, "RB") != "RB" {
			sequences = append(sequences, "U B U' B' U' R' U R")
			sequences = append(sequences, "U' R' U R U B U' B'")
		} else if get_edge_value(cube.pattern, "LB") != "LB" {
			sequences = append(sequences, "U L U' L' U' B' U B")
			sequences = append(sequences, "U' B' U B U L U' L'")
		} else if get_edge_value(cube.pattern, "FL") != "FL" {
			sequences = append(sequences, "U F U' F' U' L' U L")
			sequences = append(sequences, "U' L' U L U F U' F'")
		} else {
			sequences = append(sequences, "U R U' R' U' F' U F")
			sequences = append(sequences, "U' L' U L U F U' F'")
		}
		sequences = append(sequences, "U")
		sequences = append(sequences, "U'")
		// sequences = append(sequences, "U R U' R' U' F' U F U2 U R U' R' U' F' U F")
		// sequences = append(sequences, "U B U' B' U' R' U R U2 U B U' B' U' R' U R")
		// sequences = append(sequences, "U L U' L' U' B' U B U2 U L U' L' U' B' U B")
		// sequences = append(sequences, "U F U' F' U' L' U L U2 U F U' F' U' L' U L")

		// sequences = append(sequences, "U R U' R' U' F' U F")
		// sequences = append(sequences, "U' L' U L U F U' F'")
		// sequences = append(sequences, "U B U' B' U' R' U R")
		// sequences = append(sequences, "U' F' U F U R U' R'")
		// sequences = append(sequences, "U L U' L' U' B' U B")
		// sequences = append(sequences, "U' R' U R U B U' B'")
		// sequences = append(sequences, "U F U' F' U' L' U L")
		// sequences = append(sequences, "U' B' U B U L U' L'")

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
