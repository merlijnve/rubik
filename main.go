package main

import (
	"errors"
	"fmt"
	"os"
	"strings"
)

type Cube struct {
	pattern          []string
	solution         []string
	optimal_solution []string
}

func check(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(0)
	}
}

func print_solution(cube Cube) {
	for i := range cube.solution {
		fmt.Println(cube.solution[i])
	}
}

func init_cube() Cube {
	cube := Cube{make([]string, 0, 0), make([]string, 0, 0), make([]string, 0, 0)}

	faces := [6]string{"U", "R", "F", "D", "L", "B"}
	for f := range faces {
		for i := 1; i <= 9; i++ {
			cube.pattern = append(cube.pattern, faces[f])
		}
	}

	return cube
}

func validate(seq string) {
	arr := strings.Fields(seq)
	for i := range arr {
		if !move_is_legal(arr[i]) {
			check(errors.New("Illegal move in input string"))
		}
	}
}

func rubik(seq string) {
	validate(seq)
	cube := init_cube()
	cube = sequence(cube, seq)
	cube.solution = nil
	cube = astar(cube, bottom_cross_check, bottom_cross_heuristic, bottom_cross_get_moves)
	cube = astar(cube, first_layer_check, first_layer_heuristic, first_layer_get_moves)
	cube = astar(cube, second_layer_check, second_layer_heuristic, second_layer_get_moves)
	cube = astar(cube, top_cross_check, top_cross_heuristic, top_cross_get_moves)
	cube = astar(cube, top_layer_check, top_layer_heuristic, top_layer_get_moves)
	for i := 0; i < 4; i++ {
		for cube.pattern[8] != "U" {
			cube = sequence(cube, "R' D' R D")
		}
		cube = sequence(cube, "U")
	}
	cube = optimize(cube)
	print_solution(cube)
}

func main() {
	argsWithoutProg := os.Args[1:]

	if len(argsWithoutProg) == 1 {
		rubik(argsWithoutProg[0])

	} else {
		check(errors.New("Usage: ./rubik [sequence]"))
	}
}
