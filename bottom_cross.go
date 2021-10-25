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
	return value
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
		sequences = append(sequences, "F D' L D")
	}
	if get_edge_value(cube.pattern, "DL") == "LD" {
		sequences = append(sequences, "L D' B D")
	}
	if get_edge_value(cube.pattern, "DB") == "DB" {
		sequences = append(sequences, "B D' R D")
	}
	if get_edge_value(cube.pattern, "RD") == "DR" {
		sequences = append(sequences, "R D' F D")
	}

	if get_edge_value(cube.pattern, "UF") == "FD" {
		sequences = append(sequences, "F' D' L D")
	}
	if get_edge_value(cube.pattern, "UL") == "LD" {
		sequences = append(sequences, "L' D' B D")
	}
	if get_edge_value(cube.pattern, "UB") == "BD" {
		sequences = append(sequences, "B' D' R D")
	}
	if get_edge_value(cube.pattern, "UR") == "RD" {
		sequences = append(sequences, "R' D' F D")
	}

	if get_edge_value(cube.pattern, "FL") == "DF" {
		sequences = append(sequences, "D' L D")
	}
	if get_edge_value(cube.pattern, "RF") == "DR" {
		sequences = append(sequences, "D' F D")
	}
	if get_edge_value(cube.pattern, "RB") == "BD" {
		sequences = append(sequences, "D' R D")
	}
	if get_edge_value(cube.pattern, "LB") == "DL" {
		sequences = append(sequences, "D' B D")
	}

	if get_edge_value(cube.pattern, "RF") == "FD" {
		sequences = append(sequences, "D R' D'")
	}
	if get_edge_value(cube.pattern, "RB") == "DR" {
		sequences = append(sequences, "D B' D'")
	}
	if get_edge_value(cube.pattern, "LB") == "BD" {
		sequences = append(sequences, "D L' D'")
	}
	if get_edge_value(cube.pattern, "FL") == "LD" {
		sequences = append(sequences, "D F' D'")
	}

	for i := range sequences {
		temp := copy(cube)
		sequence(temp, sequences[i])
		moves = append(moves, temp)
	}
	return moves
}

func fix_edges_bottom_cross(cube Cube) Cube {
	state := cube.pattern
	if state[1] == "D" {
		sequence(cube, "B' B'")
	}
	if state[3] == "D" {
		sequence(cube, "L' L'")
	}
	if state[5] == "D" {
		sequence(cube, "R' R'")
	}
	if state[7] == "D" {
		sequence(cube, "F' R'")
	}

	state = cube.pattern
	if state[28] != "D" {
		sequence(cube, "F D' L D")
	}
	if state[30] != "D" {
		sequence(cube, "L D' B D")
	}
	if state[32] != "D" {
		sequence(cube, "R D' F D")
	}
	if state[34] != "D" {
		sequence(cube, "B D' R D")
	}
	return cube
}

func bottom_cross_check(cube Cube) bool {
	if get_edge_value(cube.pattern, "RD") != "RD" && get_edge_value(cube.pattern, "RD") != "DR" {
		return false
	}
	if get_edge_value(cube.pattern, "FD") != "FD" && get_edge_value(cube.pattern, "FD") != "DF" {
		return false
	}
	if get_edge_value(cube.pattern, "DL") != "DL" && get_edge_value(cube.pattern, "DL") != "LD" {
		return false
	}
	if get_edge_value(cube.pattern, "DB") != "DB" && get_edge_value(cube.pattern, "DB") != "BD" {
		return false
	}
	return true
}
