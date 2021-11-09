package main

import (
	"errors"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
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

func generate_test_sequence() string {

	moves := []string{"U", "D", "R", "L", "F", "B", "U2", "D2", "R2", "L2", "F2", "B2", "U'", "D'", "R'", "L'", "F'", "B'"}
	length := rand.Intn(100) + 1

	move := moves[rand.Intn(len(moves))]

	for i := 0; i < length; i++ {
		move = move + " " + moves[rand.Intn(len(moves))]
	}

	return move
}

func test() {
	test_sequences := make([]string, 0)

	// test_sequences = append(test_sequences, "D B2 D2 B2 B B' F2 R")
	// test_sequences = append(test_sequences, "B L' F2 F' F' R F2 R' F' L' B2 F F' B2 U' U2 U2 L' L' B U2 R2 B2 D D L'")
	// test_sequences = append(test_sequences, "F U2 B F D' L2 U L' D2 R2")
	//random test sequences
	// test_sequences = append(test_sequences, "F")
	// test_sequences = append(test_sequences, "B")
	// test_sequences = append(test_sequences, "R")
	// test_sequences = append(test_sequences, "L")
	// test_sequences = append(test_sequences, "U")
	// test_sequences = append(test_sequences, "D")
	// test_sequences = append(test_sequences, "F2")
	// test_sequences = append(test_sequences, "B2")
	// test_sequences = append(test_sequences, "R2")
	// test_sequences = append(test_sequences, "L2")
	// test_sequences = append(test_sequences, "U2")
	// test_sequences = append(test_sequences, "D2")
	// test_sequences = append(test_sequences, "R2 U D' L F2")
	// test_sequences = append(test_sequences, "R' F2 U2 L2 D2")
	// test_sequences = append(test_sequences, "U' L R2 F' U'")
	// test_sequences = append(test_sequences, "L2 B2 R2 L D'")
	// test_sequences = append(test_sequences, "D F2 B R' U2 R'")
	// test_sequences = append(test_sequences, "B' D U' F2 B L B2 F' D' F")
	// test_sequences = append(test_sequences, "F2 U' D' B' L' B' R2 U2 F L'")
	// test_sequences = append(test_sequences, "F R' U R F' R L F D B")
	// test_sequences = append(test_sequences, "D2 B2 L' F D' F' B' U L B")
	// test_sequences = append(test_sequences, "D2 L B' R' F2 U B' L D B'")
	// test_sequences = append(test_sequences, "U L' F L D L' R U B2 U R2 B U2 D2 L' U' R B' R' F' B D L R B2")
	// test_sequences = append(test_sequences, "U' F L2 U B D' U L' D U2 L' R' B2 D U' R D' U2 R' F B R2 F' B D")
	// test_sequences = append(test_sequences, "D2 F2 L' R' D B R2 B L B2 U' F2 R' F D2 B U2 B2 L B2 R' B R L' F2")
	// test_sequences = append(test_sequences, "B2 L2 D' F2 R2 L' F L R' D2 B2 D2 F L2 D' L U B' L U' L D L' B2 U2 F2 L2 U2 L' D L2 B2 D L B")
	// test_sequences = append(test_sequences, "R2 D' F' L' R U2 B2 L' R' U' D2 B' F U' L2 R2 B2 D2 U F2 R2 F' R U2 R U' F L' R' U' R' U F R2 F'")
	// test_sequences = append(test_sequences, "F B2 R F' D2 R' L U' B' U' F2 D2 B F L D F' R' D' F' L' U' F2 L' D2 B2 D' F U' L U' D2 F2 U2 L2")
	// test_sequences = append(test_sequences, "F B2 R F' F2 D2 B F L D F' R' D' F' L' U' F2 L' D2 B2 D' F U' L U' D2 F2 U2 L2")
	// test_sequences = append(test_sequences, "D2 R2 F' U B D2 R2 L' U D' B L D2 R2 B U F2 D U B' F L' F2 B' R2 U' L U' B' R U2 F2 R F' D2")
	// test_sequences = append(test_sequences, "R2 L2 B R2 F2 R L2 D2 F2 R F R U' R' D2 F2 U' R D U' F2 B2 L2 F' B2 L' U2 F' D2 L D L' B' L2 F")
	// test_sequences = append(test_sequences, "F2 L2 R D2 R' F' B2 L R F2 D2 L F2 B' L2 R' F' R U' F2 L2 U2 B L2 B' L2 B F2 R2 B2 D' R2 F2 D B' F L2 U2 R' U' R2 L' F' D2 B' F U' R L2 B2")
	// test_sequences = append(test_sequences, "F2 L2 R D2 R' F' B2 L R F2 B2 D' R2 F2 D B' F L2 U2 R' U' R2 L' F' D2 B' F U' R L2 B2")
	// test_sequences = append(test_sequences, "F B' U2 L D2 U B R L U F R2 F' R D' B U2 L R' U' R2 B2 U L B D U L2 R2 B2 D' L2 D' L2 F' D B' U2 D2 R' D' U2 R D' U2 F' D R' L B")
	// test_sequences = append(test_sequences, "B' L' D2 R U2 D' R2 B' L R' F2 R L B2 F R L2 D' U' B' D2 B D2 R2 L F' D F R' L F D U L' B2 U' F2 B D2 R' F B' L F R2 F2 D' U2 B R2")
	// test_sequences = append(test_sequences, "F' U R2 F' R U B2 R D2 L2 F' D' L' U B' L2 D' B' L D2 R2 D2 R F2 B' D2 F2 D2 F2 R L2 D U' F2 R2 B R' L B U2 F' B R2 U2 B F2 D2 B' R2 L2")
	// test_sequences = append(test_sequences, "L R' U' R' L' U F2 R' F' U' L2 F R F B D2 F L' R B' U' F' B2 D F' B' R' L F2 B2 U F' L2 B2 D' U B L' D' L F D2 L R' B' F2 L' B2 F2 R")
	// test_sequences = append(test_sequences, "F' B' R B2 U2 D F D U' F L U B2 U2 R' D' B' L F D B2 D2 L' U L'")
	// test_sequences = append(test_sequences, "U F U2 D' R B2 R D B2 L2 B L2 D2 F B' D' L2 R2 D R2 U' F2 L2 D2 L'")
	// test_sequences = append(test_sequences, "U2 F' L2 F2 L2")
	// test_sequences = append(test_sequences, "B D2 R2 L' U D' B L D2 R2 B U F2 D U B' F U2 F' L2 F2 L2 D2 L F2 B' L2 R'")
	// test_sequences = append(test_sequences, "F' U' L2 F2 L2")
	// test_sequences = append(test_sequences, "B D2 R2 L' U D' B L D2' L2 R'")
	// test_sequences = append(test_sequences, "F' U' L2 F2 L2 F2 R L2 D2 F2 R F R U' R' D2 ")
	// test_sequences = append(test_sequences, "U R2 F B R B2 R U2 L B2 R U' D' R2 F R' L B2 U2 F2")
	// test_sequences = append(test_sequences, "R' U2 B L' F U' B D F U D' L D2 F' R B' D F' U' B' U D'")
	wrong := false
	for i := 0; i < 1000; i++ {
		test_sequences = append(test_sequences, generate_test_sequence())
	}

	for i := range test_sequences {
		start := time.Now()
		cube := init_cube()
		cube = sequence(cube, test_sequences[i])
		sequence_length := len(cube.solution)
		cube.solution = nil
		println(i, sequence_length)
		println(test_sequences[i])
		// print_cube_map(cube)

		// fmt.Println("DOING:", test_sequences[i])
		// print_cube_map(cube)
		// starttime := time.Since(start).Seconds()
		cube = astar(cube, bottom_cross_check, bottom_cross_heuristic, bottom_cross_get_moves)
		// fmt.Println("BOTTOM CROSS", len(cube.solution), time.Since(start).Seconds()-starttime)
		// old_solution_lenght := len(cube.solution)
		// print(len(cube.solution), " SOLUTION: ")
		// print_cube_map(cube)
		// print_solution(cube)
		// starttime = time.Since(start).Seconds()
		cube = astar(cube, first_layer_check, first_layer_heuristic, first_layer_get_moves)
		// fmt.Println("FIRST LAYER", len(cube.solution)-old_solution_lenght, time.Since(start).Seconds()-starttime)
		// old_solution_lenght = len(cube.solution)

		// print(len(cube.solution), " SOLUTION: ")
		// print_cube_map(cube)
		// print_solution(cube)
		// starttime = time.Since(start).Seconds()
		cube = astar(cube, second_layer_check, second_layer_heuristic, second_layer_get_moves)
		// fmt.Println("SECOND LAYER", len(cube.solution)-old_solution_lenght, time.Since(start).Seconds()-starttime)
		// old_solution_lenght = len(cube.solution)

		// print(len(cube.solution), " SOLUTION: ")
		// print_cube_map(cube)
		// print_solution(cube)
		// starttime = time.Since(start).Seconds()
		cube = astar(cube, top_cross_check, top_cross_heuristic, top_cross_get_moves)
		// fmt.Println("TOP CROSS", len(cube.solution)-old_solution_lenght, time.Since(start).Seconds()-starttime)
		// old_solution_lenght = len(cube.solution)

		// print(len(cube.solution), " SOLUTION: ")
		// print_cube_map(cube)
		// print_solution(cube)
		// starttime = time.Since(start).Seconds()
		cube = astar(cube, top_layer_check, top_layer_heuristic, top_layer_get_moves)
		// fmt.Println("TOP LAYER", len(cube.solution)-old_solution_lenght, time.Since(start).Seconds()-starttime)
		cube = optimize(cube)
		// print_cube_map(cube)

		duration := time.Since(start)

		color := "\033[1;32m"
		if len(cube.solution) > 150 {
			wrong = true
			color = "\033[1;31m"
		}
		fmt.Printf("%s%d%s\n", color, len(cube.solution), "\033[0m")
		color = "\033[1;32m"
		if duration.Seconds() > 3 {
			wrong = true
			color = "\033[1;31m"
		}
		fmt.Printf("%s%f%d%s\n", color, time.Since(start).Seconds(), len(cube.solution), "\033[0m")
		println()
		// print_solution(cube)
	}
	println(wrong)
}

func rubik(seq string) {
	validate(seq)
	cube := init_cube()
	cube = sequence(cube, seq)
	cube.solution = nil

	fmt.Println("BOTTOM CROSS")
	print_cube_map(cube)
	cube = astar(cube, bottom_cross_check, bottom_cross_heuristic, bottom_cross_get_moves)
	fmt.Println("FIRST LAYER")
	print_cube_map(cube)
	cube = astar(cube, first_layer_check, first_layer_heuristic, first_layer_get_moves)
	fmt.Println("SECOND LAYER")
	print_cube_map(cube)
	cube = astar(cube, second_layer_check, second_layer_heuristic, second_layer_get_moves)
	fmt.Println("TOP CROSS")
	print_cube_map(cube)
	cube = astar(cube, top_cross_check, top_cross_heuristic, top_cross_get_moves)
	fmt.Println("TOP LAYER")
	print_cube_map(cube)
	cube = astar(cube, top_layer_check, top_layer_heuristic, top_layer_get_moves)
	cube = optimize(cube)
	print_solution(cube)
	fmt.Println(len(cube.solution))
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
