package main

func bottom_cross_heuristic(cube Cube) int {
	value := 0

	_, f := edges()
	edges := f()
	for i := range edges {
		edge_key := get_edge_value(cube.pattern, i)
		if edge_key[0] != i[0] && edge_key[1] != i[0] {
			value += 1
		}
		if edge_key[0] != i[1] && edge_key[1] != i[1] {
			value += 1
		}
	}
	return value * 2
}

func bottom_cross_get_moves(cube Cube) []Cube {
	moves := make([]Cube, 0)

	sequences := make([]string, 0)
	sequences = append(sequences, "U")
	sequences = append(sequences, "U'")
	sequences = append(sequences, "D")
	sequences = append(sequences, "D'")
	sequences = append(sequences, "R")
	sequences = append(sequences, "R'")
	sequences = append(sequences, "L")
	sequences = append(sequences, "L'")
	sequences = append(sequences, "F")
	sequences = append(sequences, "F'")
	sequences = append(sequences, "B")
	sequences = append(sequences, "B'")

	if get_edge_value(cube.pattern, "FD") == "DF" {
		sequences = nil
		sequences = append(sequences, "F D' L D")
	}
	if get_edge_value(cube.pattern, "DL") == "LD" {
		sequences = nil
		sequences = append(sequences, "L D' B D")
	}
	if get_edge_value(cube.pattern, "DB") == "DB" {
		sequences = nil
		sequences = append(sequences, "B D' R D")
	}
	if get_edge_value(cube.pattern, "RD") == "DR" {
		sequences = nil
		sequences = append(sequences, "R D' F D")
	}

	if get_edge_value(cube.pattern, "UF") == "FD" {
		sequences = nil
		sequences = append(sequences, "F' D' L D")
	}
	if get_edge_value(cube.pattern, "UL") == "LD" {
		sequences = nil
		sequences = append(sequences, "L' D' B D")
	}
	if get_edge_value(cube.pattern, "UB") == "BD" {
		sequences = nil
		sequences = append(sequences, "B' D' R D")
	}
	if get_edge_value(cube.pattern, "UR") == "RD" {
		sequences = nil
		sequences = append(sequences, "R' D' F D")
	}

	if get_edge_value(cube.pattern, "FL") == "DF" {
		sequences = nil
		sequences = append(sequences, "D' L D")
	}
	if get_edge_value(cube.pattern, "RF") == "DR" {
		sequences = nil
		sequences = append(sequences, "D' F D")
	}
	if get_edge_value(cube.pattern, "RB") == "BD" {
		sequences = nil
		sequences = append(sequences, "D' R D")
	}
	if get_edge_value(cube.pattern, "LB") == "DL" {
		sequences = nil
		sequences = append(sequences, "D' B D")
	}

	if get_edge_value(cube.pattern, "RF") == "FD" {
		sequences = nil
		sequences = append(sequences, "D R' D'")
	}
	if get_edge_value(cube.pattern, "RB") == "DR" {
		sequences = nil
		sequences = append(sequences, "D B' D'")
	}
	if get_edge_value(cube.pattern, "LB") == "BD" {
		sequences = nil
		sequences = append(sequences, "D L' D'")
	}
	if get_edge_value(cube.pattern, "FL") == "LD" {
		sequences = nil
		sequences = append(sequences, "D F' D'")
	}

	if len(cube.solution) < 35 {
		if get_edge_value(cube.pattern, "UF") == "DF" {
			sequences = nil
			sequences = append(sequences, "F2")
		}
		if get_edge_value(cube.pattern, "UB") == "DB" {
			sequences = nil
			sequences = append(sequences, "B2")
		}
		if get_edge_value(cube.pattern, "UL") == "DL" {
			sequences = nil
			sequences = append(sequences, "L2")
		}
		if get_edge_value(cube.pattern, "UR") == "DR" {
			sequences = nil
			sequences = append(sequences, "R2")
		}

		if get_edge_value(cube.pattern, "RF") == "RD" {
			sequences = nil
			sequences = append(sequences, "R'")
		}
		if get_edge_value(cube.pattern, "FL") == "DL" {
			sequences = nil
			sequences = append(sequences, "L")
		}

		if get_edge_value(cube.pattern, "RB") == "RD" {
			sequences = nil
			sequences = append(sequences, "R")
		}
		if get_edge_value(cube.pattern, "LB") == "LD" {
			sequences = nil
			sequences = append(sequences, "L'")
		}

		if get_edge_value(cube.pattern, "RF") == "DF" {
			sequences = nil
			sequences = append(sequences, "F")
		}
		if get_edge_value(cube.pattern, "FL") == "FD" {
			sequences = nil
			sequences = append(sequences, "F'")
		}

		if get_edge_value(cube.pattern, "RB") == "DB" {
			sequences = nil
			sequences = append(sequences, "B'")
		}
		if get_edge_value(cube.pattern, "LB") == "DB" {
			sequences = nil
			sequences = append(sequences, "B")
		}
	}

	for i := range sequences {
		temp := copy(cube)
		temp = sequence(temp, sequences[i])
		moves = append(moves, temp)
	}
	return moves
}

func bottom_cross_check(cube Cube) bool {
	if get_edge_value(cube.pattern, "RD") != "RD" {
		return false
	}
	if get_edge_value(cube.pattern, "FD") != "FD" {
		return false
	}
	if get_edge_value(cube.pattern, "DL") != "DL" {
		return false
	}
	if get_edge_value(cube.pattern, "DB") != "DB" {
		return false
	}
	return true
}
