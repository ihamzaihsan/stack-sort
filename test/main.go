package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

const (
	red   = "\033[31m"
	reset = "\033[0m"
)

// Helper functions for push-swap operations
func pa(stackA, stackB *[]int) {
	if len(*stackB) > 0 {
		*stackA = append([]int{(*stackB)[0]}, *stackA...)
		*stackB = (*stackB)[1:]
	}
}

func pb(stackA, stackB *[]int) {
	if len(*stackA) > 0 {
		*stackB = append([]int{(*stackA)[0]}, *stackB...)
		*stackA = (*stackA)[1:]
	}
}

func sa(stackA *[]int) {
	if len(*stackA) > 1 {
		(*stackA)[0], (*stackA)[1] = (*stackA)[1], (*stackA)[0]
	}
}

func sb(stackB *[]int) {
	if len(*stackB) > 1 {
		(*stackB)[0], (*stackB)[1] = (*stackB)[1], (*stackB)[0]
	}
}

func ss(stackA, stackB *[]int) {
	sa(stackA)
	sb(stackB)
}

func ra(stackA *[]int) {
	if len(*stackA) > 1 {
		*stackA = append((*stackA)[1:], (*stackA)[0])
	}
}

func rb(stackB *[]int) {
	if len(*stackB) > 1 {
		*stackB = append((*stackB)[1:], (*stackB)[0])
	}
}

func rr(stackA, stackB *[]int) {
	ra(stackA)
	rb(stackB)
}

func rra(stackA *[]int) {
	if len(*stackA) > 1 {
		*stackA = append([]int{(*stackA)[len(*stackA)-1]}, (*stackA)[:len(*stackA)-1]...)
	}
}

func rrb(stackB *[]int) {
	if len(*stackB) > 1 {
		*stackB = append([]int{(*stackB)[len(*stackB)-1]}, (*stackB)[:len(*stackB)-1]...)
	}
}

func rrr(stackA, stackB *[]int) {
	rra(stackA)
	rrb(stackB)
}

// Function to sort the stack using the defined operations
func sortStack(stackA, stackB *[]int) []string {
	var instructions []string

	// Example sorting algorithm using more operations
	for !isSorted(*stackA) {
		minIndex := findMinIndex(*stackA)
		if minIndex == 0 {
			pb(stackA, stackB)
			instructions = append(instructions, "pb")
		} else if minIndex == 1 {
			sa(stackA)
			instructions = append(instructions, "sa")
		} else if minIndex <= len(*stackA)/2 {
			ra(stackA)
			instructions = append(instructions, "ra")
		} else {
			rra(stackA)
			instructions = append(instructions, "rra")
		}
	}
	for len(*stackB) > 0 {
		pa(stackA, stackB)
		instructions = append(instructions, "pa")
	}

	return instructions
}

func isRepeated(stack []int) bool {
	seen := make(map[int]bool)
	for _, num := range stack {
		if seen[num] {
			return true
		}
		seen[num] = true
	}
	return false
}


// Helper function to check if the stack is sorted
func isSorted(stack []int) bool {
	for i := 0; i < len(stack)-1; i++ {
		if stack[i] > stack[i+1] {
			return false
		}
	}
	return true
}

// Helper function to find the index of the minimum element in the stack
func findMinIndex(stack []int) int {
	minIndex := 0
	for i := 1; i < len(stack); i++ {
		if stack[i] < stack[minIndex] {
			minIndex = i
		}
	}
	return minIndex
}

func findMaxIndex(stack []int) int {
	maxIndex := 0
	for i := 1; i < len(stack); i++ {
		if stack[i] > stack[maxIndex] {
			maxIndex = i
		}
	}
	return maxIndex
}


func main() {
	if len(os.Args) < 2 {
		return
	}

	input := os.Args[1]

	elements := strings.Split(input, " ")

	var stackA []int

	for _, elem := range elements {
		num, err := strconv.Atoi(elem)
		if err != nil {
			fmt.Println(red, "Error", reset)
			return
		}
		stackA = append(stackA, num)
	}

	// Now stackA is initialized with the input integers
	var stackB []int



	if isSorted(stackA) {
		return

	} else if isRepeated(stackA) {
		return

	}

	// Sort the stack
	instructions := sortStack(&stackA, &stackB)

	in := ""

	// Print the instructions
	for _, instr := range instructions {
		in = in + instr + "\\n"
		fmt.Println(instr)
	}
	// fmt.Println()
	// fmt.Println(in)
}