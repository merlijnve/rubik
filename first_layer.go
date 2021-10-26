package main

import "strings"

func corner_is_correct(corner_value string, key string) bool {
	return corner_value == key
}

func corner_is_above(corner_value string, key string) bool {
	corner_value = strings.Replace(corner_value, "D", "U", 1)
	return corner_is_twisted(corner_value, key)
}

func corner_is_twisted(corner_value, key string) bool {
	return corner_value == string(key[0]+key[2]+key[1]) ||
		corner_value == string(key[1]+key[0]+key[2]) ||
		corner_value == string(key[1]+key[2]+key[0]) ||
		corner_value == string(key[2]+key[0]+key[1]) ||
		corner_value == string(key[2]+key[1]+key[0])
}

func first_layer_heuristic(cube Cube) int {
	value := 0

	_, f := corners()
	corners := f()

	for i := range corners {
		if corner_is_correct(get_corner_value(cube.pattern, i), i) {
			value += 0
		} else {
			if corner_is_twisted(get_corner_value(cube.pattern, i), i) {
				value += 4
			} else if corner_is_above(get_corner_value(cube.pattern, i), i) == false {
				value += 1
			}
		}
	}
	return value
}

func first_layer_get_moves(cube Cube) []Cube {
	moves := make([]Cube, 0)
	sequences := make([]string, 0)

	sequences = append(sequences, "L' U' L U")
	sequences = append(sequences, "R' U' R U")
	sequences = append(sequences, "B' U' B U")
	sequences = append(sequences, "F' U' F U")

	if sorted(get_corner_value(cube.pattern, "UFL")) == "DFL" {
		sequences = append(sequences, "L' U' L")
		sequences = append(sequences, "F U F'")
		sequences = append(sequences, "F R U' U' R' F'")
	}
	if sorted(get_corner_value(cube.pattern, "ULB")) == "BDL" {
		sequences = append(sequences, "B' U' B")
		sequences = append(sequences, "L U L'")
		sequences = append(sequences, "L F U' U' F' L'")
	}
	if sorted(get_corner_value(cube.pattern, "URB")) == "BDR" {
		sequences = append(sequences, "R' U' R")
		sequences = append(sequences, "B U B'")
		sequences = append(sequences, "B L U' U' L' B'")
	}
	if sorted(get_corner_value(cube.pattern, "URF")) == "DFR" {
		sequences = append(sequences, "F' U' F")
		sequences = append(sequences, "R U R'")
		sequences = append(sequences, "R B U' U' B' R'")
	}

	for i := range sequences {
		temp := copy(cube)
		temp = sequence(temp, sequences[i])
		moves = append(moves, temp)
	}
	return moves
}

func first_layer_check(cube Cube) bool {
	if get_corner_value(cube.pattern, "RFD") != "RFD" {
		return false
	}
	if get_corner_value(cube.pattern, "RDB") != "RDB" {
		return false
	}
	if get_corner_value(cube.pattern, "FDL") != "FDL" {
		return false
	}
	if get_corner_value(cube.pattern, "DLB") != "DLB" {
		return false
	}
	return true
}

// func first_layer(cube Cube) Cube {
// 	get_moves(cube)
// 	first_layer_heuristic(cube)
// 	return cube
// }
