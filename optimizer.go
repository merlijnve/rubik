package main

import (
	"strings"
)

func remove_move_and_inversion(cube Cube) Cube {
	solution_len := len(cube.solution)
	i := 0
	for i < solution_len {
		if solution_len-i >= 2 && strings.Contains(cube.solution[i+1], "'") && !strings.Contains(cube.solution[i], "'") && cube.solution[i][0] == cube.solution[i+1][0] {
			cube.solution = append(cube.solution[0:i], cube.solution[i+2:solution_len]...)
			i = 0
		} else if solution_len-i >= 2 && strings.Contains(cube.solution[i], "'") && !strings.Contains(cube.solution[i+1], "'") && cube.solution[i][0] == cube.solution[i+1][0] {
			cube.solution = append(cube.solution[0:i], cube.solution[i+2:solution_len]...)
			i = 0
		} else {
			i += 1
		}
		solution_len = len(cube.solution)
	}
	return cube
}

func remove_quarter_rotations(cube Cube) Cube {
	solution_len := len(cube.solution)
	i := 0

	for i < solution_len {
		if solution_len-i >= 3 && cube.solution[i] == cube.solution[i+1] && cube.solution[i] == cube.solution[i+2] {
			if i+3 == solution_len {
				str := make([]byte, 2)
				str[0] = cube.solution[i][0]
				str[1] = '\''

				temp := cube.solution[0:i]
				temp = append(temp, string(str))
				cube.solution = temp
			} else {
				str := make([]byte, 2)
				str[0] = cube.solution[i][0]
				str[1] = '\''

				temp := cube.solution[0:i]
				temp = append(temp, string(str))
				temp = append(temp, cube.solution[i+2:len(cube.solution)]...)
				cube.solution = temp
			}
			i = 0
			solution_len = len(cube.solution)
		} else {
			i += 1
		}
	}
	return cube
}

func remove_full_rotations(cube Cube) Cube {
	solution_len := len(cube.solution)
	i := 0

	for i < solution_len {
		if solution_len-i >= 4 && cube.solution[i] == cube.solution[i+1] && cube.solution[i] == cube.solution[i+2] && cube.solution[i] == cube.solution[i+3] {
			if i+4 == solution_len {
				cube.solution = cube.solution[0:i]
			} else {
				cube.solution = append(cube.solution[0:i], cube.solution[i+4:solution_len]...)
			}
			i = 0
			solution_len = len(cube.solution)
		} else {
			i += 1
		}
	}
	return cube
}

func smallest_notation(cube Cube) Cube {
	solution_len := len(cube.solution)
	i := 0

	for i < solution_len {
		if solution_len-i >= 2 && cube.solution[i] == cube.solution[i+1] {
			if i+2 == solution_len {
				str := make([]byte, 2)
				str[0] = cube.solution[i][0]
				str[1] = '2'
				cube.solution = cube.solution[0:i]
				cube.solution = append(cube.solution, string(str))
			} else {
				str := make([]byte, 2)
				str[0] = cube.solution[i][0]
				str[1] = '2'
				temp := cube.solution[0:i]
				temp = append(temp, string(str))
				cube.solution = append(temp, cube.solution[i+2:solution_len]...)
			}
			i = 0
		} else {
			i += 1
		}
		solution_len = len(cube.solution)
	}
	return cube
}

func optimize(cube Cube) Cube {
	new_cube := copy(cube)

	new_cube = remove_full_rotations(new_cube)
	new_cube = remove_quarter_rotations(new_cube)
	new_cube = remove_move_and_inversion(new_cube)
	new_cube = smallest_notation(new_cube)
	return new_cube
}
