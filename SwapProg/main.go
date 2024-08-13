package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
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

// Function to sort the stack using the defined operations
func sortStack(stackA, stackB *[]int) []string {
	var instructions []string

	// Use short solutions for small stacks and grouped pushes for larger ones.
	if len(*stackA) == 2 {
		for !isSorted(*stackA) {
			sa(stackA)
			instructions = append(instructions, "sa")
		}

	} else if len(*stackA) == 3 {
		minIndex := findMinIndex(*stackA)
		maxIndex := findMaxIndex(*stackA)

		if (minIndex == 0 && maxIndex == 1) || (minIndex == 2 && maxIndex == 0) {
			if minIndex == 0 && maxIndex == 1 {
				sa(stackA)
				instructions = append(instructions, "sa")

				ra(stackA)
				instructions = append(instructions, "ra")

			} else if minIndex == 2 && maxIndex == 0 {
				ra(stackA)
				instructions = append(instructions, "ra")

				sa(stackA)
				instructions = append(instructions, "sa")
			}

		} else {
			if minIndex == 1 && maxIndex == 0 {
				ra(stackA)
				instructions = append(instructions, "ra")

			} else if minIndex == 1 && maxIndex == 2 {
				sa(stackA)
				instructions = append(instructions, "sa")

			} else if minIndex == 2 && maxIndex == 1 {
				rra(stackA)
				instructions = append(instructions, "rra")

			}
		}

	} else if len(*stackA) > 6 {
		return sortLargeStack(stackA, stackB)
	} else {
		for !isSorted(*stackA) && len(*stackA) > 3 {
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
		// Reuse the existing three-value solution before restoring the minima.
		if len(*stackA) == 3 {
			var emptyStack []int
			instructions = append(instructions, sortStack(stackA, &emptyStack)...)
		}

	}

	for len(*stackB) > 0 {
		pa(stackA, stackB)
		instructions = append(instructions, "pa")
	}

	return instructions
}

// sortLargeStack keeps the same push/rotate approach, but moves groups of low
// values together instead of rotating to every individual minimum.
func sortLargeStack(stackA, stackB *[]int) []string {
	if isSorted(*stackA) {
		return nil
	}
	// A rank lets the same group sizes work with negative or widely spaced values.
	ranks := make(map[int]int)
	for _, value := range *stackA {
		for _, other := range *stackA {
			if other < value {
				ranks[value]++
			}
		}
	}
	// Try nearby group sizes and retain the shortest valid sequence.
	groupSize := 1
	for groupSize*groupSize < len(*stackA) {
		groupSize++
	}
	var best []string
	var bestA []int
	for _, window := range []int{groupSize, groupSize * 3 / 2, groupSize * 2} {
		a := append([]int(nil), (*stackA)...)
		var b []int
		instructions := sortGroups(&a, &b, ranks, window)
		if best == nil || len(instructions) < len(best) {
			best = instructions
			bestA = a
		}
	}
	*stackA = bestA
	*stackB = nil
	return best
}

func sortGroups(stackA, stackB *[]int, ranks map[int]int, window int) []string {
	var instructions []string
	for len(*stackA) > 0 {
		rank := ranks[(*stackA)[0]]
		if rank <= len(*stackB) {
			pb(stackA, stackB)
			instructions = append(instructions, "pb")
			rb(stackB)
			instructions = append(instructions, "rb")
		} else if rank <= len(*stackB)+window {
			pb(stackA, stackB)
			instructions = append(instructions, "pb")
		} else {
			ra(stackA)
			instructions = append(instructions, "ra")
		}
	}
	for len(*stackB) > 0 {
		maxIndex := findMaxIndex(*stackB)
		if maxIndex == 0 {
			pa(stackA, stackB)
			instructions = append(instructions, "pa")
		} else if maxIndex <= len(*stackB)/2 {
			rb(stackB)
			instructions = append(instructions, "rb")
		} else {
			rrb(stackB)
			instructions = append(instructions, "rrb")
		}
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
	if len(os.Args) != 2 {
		fmt.Fprintln(os.Stderr, "Error")
		os.Exit(1)
	}

	input := os.Args[1]

	elements := strings.Fields(input)
	if len(elements) == 0 {
		fmt.Fprintln(os.Stderr, "Error")
		os.Exit(1)
	}

	var stackA []int

	for _, elem := range elements {
		num, err := strconv.Atoi(elem)
		if err != nil {
			fmt.Fprintln(os.Stderr, "Error")
			os.Exit(1)
		}
		stackA = append(stackA, num)
	}

	// Now stackA is initialized with the input integers
	var stackB []int

	if isRepeated(stackA) {
		fmt.Fprintln(os.Stderr, "Error")
		os.Exit(1)

	} else if isSorted(stackA) {
		return

	}

	// Sort the stack
	instructions := sortStack(&stackA, &stackB)

	// Print the instructions
	for _, instr := range instructions {
		fmt.Println(instr)
	}
}
