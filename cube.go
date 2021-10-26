package main

import "strings"

func edges() (func(key string) []int, func() map[string][]int) {
	edges := map[string][]int{
		"UR": {5, 10},
		"UF": {7, 19},
		"UL": {3, 37},
		"UB": {1, 46},
		"RF": {12, 23},
		"RD": {16, 32},
		"RB": {14, 48},
		"FD": {25, 28},
		"FL": {21, 41},
		"DL": {30, 43},
		"DB": {34, 52},
		"LB": {39, 50}}

	return func(key string) []int {
			return edges[key]
		},
		func() map[string][]int {
			return edges
		}
}

func corners() (func(key string) []int, func() map[string][]int) {
	corners := map[string][]int{
		"URF": {8, 9, 20},
		"URB": {2, 11, 45},
		"UFL": {6, 18, 38},
		"ULB": {0, 36, 47},
		"RFD": {15, 26, 29},
		"RDB": {17, 35, 51},
		"FDL": {24, 27, 44},
		"DLB": {33, 42, 53}}

	return func(key string) []int {
			return corners[key]
		},
		func() map[string][]int {
			return corners
		}
}

func move_is_legal(move string) bool {
	if move == "U" ||
		move == "D" ||
		move == "R" ||
		move == "L" ||
		move == "F" ||
		move == "B" ||
		move == "U2" ||
		move == "D2" ||
		move == "R2" ||
		move == "L2" ||
		move == "F2" ||
		move == "B2" ||
		move == "U'" ||
		move == "D'" ||
		move == "R'" ||
		move == "L'" ||
		move == "F'" ||
		move == "B'" {
		return true
	}
	return false
}

func copy(cube Cube) Cube {
	new_cube := Cube{make([]string, 0, 0), make([]string, 0, 0), make([]string, 0, 0)}

	for i := range cube.pattern {
		new_cube.pattern = append(new_cube.pattern, cube.pattern[i])
	}
	for i := range cube.solution {
		new_cube.solution = append(new_cube.solution, cube.solution[i])
	}
	for i := range cube.optimal_solution {
		new_cube.optimal_solution = append(new_cube.optimal_solution, cube.optimal_solution[i])
	}

	return new_cube
}

func get_corner_value(pattern []string, corner_key string) string {
	if len(corner_key) != 3 {
		return ""
	}
	f, _ := corners()
	return pattern[f(corner_key)[0]] + pattern[f(corner_key)[1]] + pattern[f(corner_key)[2]]
}

func get_edge_value(pattern []string, edge_key string) string {
	if len(edge_key) != 2 {
		return ""
	}
	f, _ := edges()
	return pattern[f(edge_key)[0]] + pattern[f(edge_key)[1]]
}

func get_pattern(cube Cube) string {
	return strings.Join(cube.pattern, "")
}

func rotate_face(cube Cube, face int) Cube {
	offset := face * 9

	temp := cube.pattern[6+offset]
	cube.pattern[6+offset] = cube.pattern[8+offset]
	cube.pattern[8+offset] = cube.pattern[2+offset]
	cube.pattern[2+offset] = cube.pattern[0+offset]
	cube.pattern[0+offset] = temp

	temp = cube.pattern[3+offset]
	cube.pattern[3+offset] = cube.pattern[7+offset]
	cube.pattern[7+offset] = cube.pattern[5+offset]
	cube.pattern[5+offset] = cube.pattern[1+offset]
	cube.pattern[1+offset] = temp

	for i := 0; i < 3; i++ {
		if face == 0 {
			temp = cube.pattern[36+i]
			cube.pattern[36+i] = cube.pattern[18+i]
			cube.pattern[18+i] = cube.pattern[9+i]
			cube.pattern[9+i] = cube.pattern[45+i]
			cube.pattern[45+i] = temp
		}
		if face == 1 {
			temp = cube.pattern[20+3*i]
			cube.pattern[20+3*i] = cube.pattern[29+3*i]
			cube.pattern[29+3*i] = cube.pattern[51-3*i]
			cube.pattern[51-3*i] = cube.pattern[2+3*i]
			cube.pattern[2+3*i] = temp
		}
		if face == 2 {
			temp = cube.pattern[38+3*i]
			cube.pattern[38+3*i] = cube.pattern[27+i]
			cube.pattern[27+i] = cube.pattern[15-3*i]
			cube.pattern[15-3*i] = cube.pattern[8-i]
			cube.pattern[8-i] = temp
		}
		if face == 3 {
			temp = cube.pattern[45+i+3*2]
			cube.pattern[45+i+3*2] = cube.pattern[9+i+3*2]
			cube.pattern[9+i+3*2] = cube.pattern[18+i+3*2]
			cube.pattern[18+i+3*2] = cube.pattern[36+i+3*2]
			cube.pattern[36+i+3*2] = temp
		}
		if face == 4 {
			temp = cube.pattern[3*i]
			cube.pattern[3*i] = cube.pattern[53-3*i]
			cube.pattern[53-3*i] = cube.pattern[27+3*i]
			cube.pattern[27+3*i] = cube.pattern[18+3*i]
			cube.pattern[18+3*i] = temp
		}
		if face == 5 {
			temp = cube.pattern[35-i]
			cube.pattern[35-i] = cube.pattern[42-3*i]
			cube.pattern[42-3*i] = cube.pattern[i]
			cube.pattern[i] = cube.pattern[11+3*i]
			cube.pattern[11+3*i] = temp
		}
	}
	return cube
}

func rotate_face_inverse(cube Cube, face int) Cube {
	offset := face * 9

	temp := cube.pattern[0+offset]
	cube.pattern[0+offset] = cube.pattern[2+offset]
	cube.pattern[2+offset] = cube.pattern[8+offset]
	cube.pattern[8+offset] = cube.pattern[6+offset]
	cube.pattern[6+offset] = temp

	temp = cube.pattern[1+offset]
	cube.pattern[1+offset] = cube.pattern[5+offset]
	cube.pattern[5+offset] = cube.pattern[7+offset]
	cube.pattern[7+offset] = cube.pattern[3+offset]
	cube.pattern[3+offset] = temp

	for i := 0; i < 3; i++ {
		if face == 0 {
			temp = cube.pattern[45+i]
			cube.pattern[45+i] = cube.pattern[9+i]
			cube.pattern[9+i] = cube.pattern[18+i]
			cube.pattern[18+i] = cube.pattern[36+i]
			cube.pattern[36+i] = temp
		}
		if face == 1 {
			temp = cube.pattern[2+3*i]
			cube.pattern[2+3*i] = cube.pattern[51-3*i]
			cube.pattern[51-3*i] = cube.pattern[29+3*i]
			cube.pattern[29+3*i] = cube.pattern[20+3*i]
			cube.pattern[20+3*i] = temp
		}
		if face == 2 {
			temp = cube.pattern[8-i]
			cube.pattern[8-i] = cube.pattern[15-3*i]
			cube.pattern[15-3*i] = cube.pattern[27+i]
			cube.pattern[27+i] = cube.pattern[38+3*i]
			cube.pattern[38+3*i] = temp
		}
		if face == 3 {
			temp = cube.pattern[36+i+3*2]
			cube.pattern[36+i+3*2] = cube.pattern[18+i+3*2]
			cube.pattern[18+i+3*2] = cube.pattern[9+i+3*2]
			cube.pattern[9+i+3*2] = cube.pattern[45+i+3*2]
			cube.pattern[45+i+3*2] = temp
		}
		if face == 4 {
			temp = cube.pattern[18+3*i]
			cube.pattern[18+3*i] = cube.pattern[27+3*i]
			cube.pattern[27+3*i] = cube.pattern[53-3*i]
			cube.pattern[53-3*i] = cube.pattern[3*i]
			cube.pattern[3*i] = temp
		}
		if face == 5 {
			temp = cube.pattern[11+3*i]
			cube.pattern[11+3*i] = cube.pattern[i]
			cube.pattern[i] = cube.pattern[42-3*i]
			cube.pattern[42-3*i] = cube.pattern[35-i]
			cube.pattern[35-i] = temp
		}
	}
	return cube
}

func U(cube Cube) Cube {
	return rotate_face(cube, 0)
}

func Ui(cube Cube) Cube {
	return rotate_face_inverse(cube, 0)
}

func U2(cube Cube) Cube {
	cube = Ui(cube)
	return Ui(cube)
}

func D(cube Cube) Cube {
	return rotate_face(cube, 3)
}

func Di(cube Cube) Cube {
	return rotate_face_inverse(cube, 3)
}

func D2(cube Cube) Cube {
	cube = Di(cube)
	return Di(cube)
}

func R(cube Cube) Cube {
	return rotate_face(cube, 1)
}

func Ri(cube Cube) Cube {
	return rotate_face_inverse(cube, 1)
}

func R2(cube Cube) Cube {
	cube = Ri(cube)
	return Ri(cube)
}

func L(cube Cube) Cube {
	return rotate_face(cube, 4)
}

func Li(cube Cube) Cube {
	return rotate_face_inverse(cube, 4)
}

func L2(cube Cube) Cube {
	cube = Li(cube)
	return Li(cube)
}

func F(cube Cube) Cube {
	return rotate_face(cube, 2)
}

func Fi(cube Cube) Cube {
	return rotate_face_inverse(cube, 2)
}

func F2(cube Cube) Cube {
	cube = Fi(cube)
	return Fi(cube)
}

func B(cube Cube) Cube {
	return rotate_face(cube, 5)
}

func Bi(cube Cube) Cube {
	return rotate_face_inverse(cube, 5)
}

func B2(cube Cube) Cube {
	cube = Bi(cube)
	return Bi(cube)
}

func sequence(cube Cube, seq string) Cube {
	sequence := strings.Fields(seq)
	for i := range sequence {
		switch sequence[i] {
		case "U":
			cube = U(cube)
		case "D":
			cube = D(cube)
		case "R":
			cube = R(cube)
		case "L":
			cube = L(cube)
		case "F":
			cube = F(cube)
		case "B":
			cube = B(cube)
		case "U2":
			cube = U2(cube)
		case "D2":
			cube = D2(cube)
		case "R2":
			cube = R2(cube)
		case "L2":
			cube = L2(cube)
		case "F2":
			cube = F2(cube)
		case "B2":
			cube = B2(cube)
		case "U'":
			cube = Ui(cube)
		case "D'":
			cube = Di(cube)
		case "R'":
			cube = Ri(cube)
		case "L'":
			cube = Li(cube)
		case "F'":
			cube = Fi(cube)
		case "B'":
			cube = Bi(cube)
		}
		cube.solution = append(cube.solution, sequence[i])
	}
	return cube
}
