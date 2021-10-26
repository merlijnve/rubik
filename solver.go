package main

import (
	"container/heap"
	"fmt"
	"strconv"
)

type Node struct {
	heurValue int
	toStart   int
	total     int
	cube      Cube
	parent    *Node
	Index     int
}

type PriorityQueue []Node

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].total < pq[j].total
}

// PriorityQueue function to retrieve the MIN element from the heap
// gets called on heap.Pop
func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	item.Index = -1
	*pq = old[0 : n-1]
	return item
}

// PriorityQueue function to push an element to the heap
// gets called on heap.Push, x is the element to push
func (pq *PriorityQueue) Push(x interface{}) {
	n := len(*pq)
	item := x.(Node)
	item.Index = n
	*pq = append(*pq, item)
}

// PriorityQueue function to swap an element in the heap
func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].Index = i
	pq[j].Index = j
}

// Function to retrieve the lowest value Node from the queue and move it to [closed]
func move_lowest_to_closed(open map[string]bool, closed map[string]Node, root *Node, priorityQueue PriorityQueue) (Node, PriorityQueue) {
	var key_of_lowest string

	lowest := heap.Pop(&priorityQueue).(Node)
	key_of_lowest = get_pattern(lowest.cube)
	closed[key_of_lowest] = lowest
	delete(open, key_of_lowest)
	return lowest, priorityQueue
}

// Calculate steps taken by iterating through the parent pointer until nil
func calc_to_start(current Node) int {
	i := 1
	for current.parent != nil {
		i++
		current = *current.parent
	}
	return i
}

// Executes all possible moves and returns an array of all successor puzzles
func get_successors(current Node, get_moves func(cube Cube) []Cube, heur_value int) []Node {
	moves := get_moves(current.cube)
	successors := make([]Node, 0)

	for i := range moves {
		successors = append(successors, Node{heur_value, len(moves[i].solution), heur_value + len(moves[i].solution), moves[i], &current, 0})
	}
	return successors
}

// Creates key out of the puzzle state (for use as key in closed/open map)
func stateToString(numbers []int, n int) string {
	var s string

	for i := 0; i < n*n; i++ {
		s = s + strconv.Itoa(numbers[i])
	}
	return s
}

// Checks if state exists in list and if the current state was reached in less steps
// (never happens because cost is always 1 so states that have been found before cannot be reached by a shortcut)
func find_and_compare_states(list map[string]Node, current Node) bool {
	if len(list) == 0 {
		return false
	}
	key := get_pattern(current.cube)
	_, ok := list[key]
	if ok == true && list[key].total < current.total {
		return true
	}
	return false
}

// Implementation of A* using:
// 1. PriorityQueue for quick access to the MIN node
// 2. Maps with the puzzle state as key for quick checking if a state was found before
func astar(cube Cube, checker func(cube Cube) bool, heur func(cube Cube) int, get_moves func(cube Cube) []Cube) Cube {
	priorityQueue := make(PriorityQueue, 1)
	open := make(map[string]bool)
	closed := make(map[string]Node)
	node_current := Node{}

	heur_value := heur(cube)
	node_start := Node{heur_value, 0, heur_value + 0, cube, nil, 0}
	open[get_pattern(cube)] = true
	priorityQueue[0] = node_start
	heap.Init(&priorityQueue)
	for true {
		node_current, priorityQueue = move_lowest_to_closed(open, closed, &node_start, priorityQueue)
		if checker(node_current.cube) {
			return node_current.cube
		} else {
			successors := get_successors(node_current, get_moves, heur(node_current.cube))
			for s := range successors {
				if find_and_compare_states(closed, successors[s]) == false &&
					!open[get_pattern(successors[s].cube)] {
					open[get_pattern(successors[s].cube)] = true
					heap.Push(&priorityQueue, successors[s])
				}
			}
		}
	}
	fmt.Println("Could not solve")
	return Cube{}
}
