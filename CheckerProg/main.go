package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Stack struct {
	elements []int
}

func (s *Stack) Push(value int) {
	s.elements = append([]int{value}, s.elements...)
}
func (s *Stack) PUSH(value int) {
	s.elements = append(s.elements, value)
}

func (s *Stack) Pop() int {
	if len(s.elements) == 0 {
		return -1 // Handle underflow error
	}
	value := s.elements[0]
	s.elements = s.elements[1:]
	return value
}

func (s *Stack) Swap() {
	if len(s.elements) < 2 {
		return
	}
	s.elements[0], s.elements[1] = s.elements[1], s.elements[0]
}

func (s *Stack) Rotate() {
	if len(s.elements) == 0 {
		return
	}
	value := s.elements[0]
	copy(s.elements, s.elements[1:])
	s.elements[len(s.elements)-1] = value
}

func (s *Stack) ReverseRotate() {
	if len(s.elements) == 0 {
		return
	}
	value := s.elements[len(s.elements)-1]
	s.elements = append([]int{value}, s.elements[:len(s.elements)-1]...)
}

func (s *Stack) IsSorted() bool {
	for i := 0; i < len(s.elements)-1; i++ {
		if s.elements[i] > s.elements[i+1] {
			return false
		}
	}
	return true
}

func parseArguments(args []string) (*Stack, error) {
	a := &Stack{}
	for _, arg := range args {
		value, err := strconv.Atoi(arg)
		if err != nil {
			return nil, fmt.Errorf("error: %v", err)
		}
		//FIXME
		a.PUSH(value)
	}
	return a, nil
}

func executeInstruction(instruction string, a, b *Stack) error {
	switch instruction {
	case "sa":
		//FIXME 
		a.Swap()

		// fmt.Println("sa")
		// fmt.Println("Stake A: ", a.elements)
		// fmt.Println("Stake B: ", b.elements)
		// fmt.Println()


	case "sb":
		b.Swap()

		// fmt.Println("sb")
		// fmt.Println("Stake A: ", a.elements)
		// fmt.Println("Stake B: ", b.elements)
		// fmt.Println()


	case "ss":
		a.Swap()
		b.Swap()

		// fmt.Println("ss")
		// fmt.Println("Stake A: ", a.elements)
		// fmt.Println("Stake B: ", b.elements)
		// fmt.Println()


	case "pa":
		if len(b.elements) == 0 {
			return fmt.Errorf("error: cannot pop from empty stack b")
		}
		a.Push(b.Pop())

		// fmt.Println("pa")
		// fmt.Println("Stake A: ", a.elements)
		// fmt.Println("Stake B: ", b.elements)
		// fmt.Println()


	case "pb":
		if len(a.elements) == 0 {
			return fmt.Errorf("error: cannot pop from empty stack a")
		}
		b.Push(a.Pop())

		// fmt.Println("pb")
		// fmt.Println("Stake A: ", a.elements)
		// fmt.Println("Stake B: ", b.elements)
		// fmt.Println()


	case "ra":
		// sheft up
		a.Rotate()

		// fmt.Println("ra")
		// fmt.Println("Stake A: ", a.elements)
		// fmt.Println("Stake B: ", b.elements)
		// fmt.Println()


	case "rb":
		b.Rotate()

		// fmt.Println("rb")
		// fmt.Println("Stake A: ", a.elements)
		// fmt.Println("Stake B: ", b.elements)
		// fmt.Println()


	case "rr":
		a.Rotate()
		b.Rotate()

		// fmt.Println("rr")
		// fmt.Println("Stake A: ", a.elements)
		// fmt.Println("Stake B: ", b.elements)
		// fmt.Println()


	case "rra":
		a.ReverseRotate()

		// fmt.Println("rra")
		// fmt.Println("Stake A: ", a.elements)
		// fmt.Println("Stake B: ", b.elements)
		// fmt.Println()


	case "rrb":
		b.ReverseRotate()

		// fmt.Println("rrb")
		// fmt.Println("Stake A: ", a.elements)
		// fmt.Println("Stake B: ", b.elements)
		// fmt.Println()

	case "rrr":
		a.ReverseRotate()
		b.ReverseRotate()

		// fmt.Println("rrr")
		// fmt.Println("Stake A: ", a.elements)
		// fmt.Println("Stake B: ", b.elements)
		// fmt.Println()

	default:
		return fmt.Errorf("error: invalid instruction %s", instruction)
	}
	return nil
}

func main() {
	if len(os.Args) < 2 {
		return
	}

	a, err := parseArguments(strings.Split(os.Args[1], " "))
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error")
		return
	}
	b := &Stack{}

	var instructions []string
	for {
		var instruction string
		if _, err := fmt.Scanln(&instruction); err != nil {
			break
		}
		instructions = append(instructions, instruction)
	}
	
	// fmt.Println("---------------------------")
	// fmt.Println("Stake A: ", a.elements)
	// fmt.Println("Stake B: ", b.elements)
	// fmt.Println("---------------------------")

	for _, instruction := range instructions {
		if err := executeInstruction(instruction, a, b); err != nil {
			fmt.Fprintln(os.Stderr, "Error")
			return
		}
	}

	if a.IsSorted() && len(b.elements) == 0 {
		// fmt.Println("Stake A: ", a.elements)
		// fmt.Println("Stake B: ", b.elements)
		fmt.Println("OK")
	} else {
		// fmt.Println("Stake A: ", a.elements)
		// fmt.Println("Stake B: ", b.elements)
		fmt.Println("KO")
	}
}




