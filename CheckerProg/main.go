package main

import (
	"bufio"
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
	seen := make(map[int]bool)
	for _, arg := range args {
		value, err := strconv.Atoi(arg)
		if err != nil {
			return nil, fmt.Errorf("error: %v", err)
		}
		if seen[value] {
			return nil, fmt.Errorf("error: duplicate integer %d", value)
		}
		seen[value] = true
		a.PUSH(value)
	}
	return a, nil
}

func executeInstruction(instruction string, a, b *Stack) error {
	switch instruction {
	case "sa":
		a.Swap()

	case "sb":
		b.Swap()

	case "ss":
		a.Swap()
		b.Swap()

	case "pa":
		if len(b.elements) == 0 {
			return fmt.Errorf("error: cannot pop from empty stack b")
		}
		a.Push(b.Pop())

	case "pb":
		if len(a.elements) == 0 {
			return fmt.Errorf("error: cannot pop from empty stack a")
		}
		b.Push(a.Pop())

	case "ra":
		a.Rotate()

	case "rb":
		b.Rotate()

	case "rr":
		a.Rotate()
		b.Rotate()

	case "rra":
		a.ReverseRotate()

	case "rrb":
		b.ReverseRotate()

	case "rrr":
		a.ReverseRotate()
		b.ReverseRotate()

	default:
		return fmt.Errorf("error: invalid instruction %s", instruction)
	}
	return nil
}

func main() {
	if len(os.Args) < 2 {
		return
	}
	if len(os.Args) != 2 {
		fmt.Fprintln(os.Stderr, "Error")
		os.Exit(1)
	}

	args := strings.Fields(os.Args[1])
	if len(args) == 0 {
		fmt.Fprintln(os.Stderr, "Error")
		os.Exit(1)
	}
	a, err := parseArguments(args)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error")
		os.Exit(1)
	}
	b := &Stack{}

	var instructions []string
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		// Keep the entire line so malformed instructions are rejected below.
		instructions = append(instructions, scanner.Text())
	}
	if scanner.Err() != nil {
		fmt.Fprintln(os.Stderr, "Error")
		os.Exit(1)
	}

	for _, instruction := range instructions {
		if err := executeInstruction(instruction, a, b); err != nil {
			fmt.Fprintln(os.Stderr, "Error")
			os.Exit(1)
		}
	}

	if a.IsSorted() && len(b.elements) == 0 {
		fmt.Println("OK")
	} else {
		fmt.Println("KO")
	}
}
