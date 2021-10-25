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
		print(cube.solution[i] + "\n")
	}
}

func test() {
	print("test mode triggered\n")
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
	fmt.Println(cube.pattern)
	cube = sequence(cube, seq)
	fmt.Println(cube.pattern)

	cube = astar(cube, bottom_cross_check, bottom_cross_heuristic, bottom_cross_get_moves)
	cube = astar(cube, first_layer_check, first_layer_heuristic, first_layer_get_moves)
}

func main() {
	argsWithoutProg := os.Args[1:]

	if len(argsWithoutProg) == 1 {
		if argsWithoutProg[0] == "test" {
			test()
		} else {
			rubik(argsWithoutProg[0])
		}
	} else {
		check(errors.New("Usage: ./rubik [sequence]"))
	}
}
