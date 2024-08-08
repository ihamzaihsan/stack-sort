# Push-Swap

Push-Swap is a project written in Go that involves sorting a list of integers using two stacks (a and b) and a set of predefined instructions. The project comprises two programs: `push-swap` and `checker`.

## Project Structure
```
.
|-- Checker
|   `-- main.go
|-- PKG
|   `-- PS_oprations.go
|-- go.mod
|-- main.go
`-- swap
    `-- main.go
```

## Instructions

- **pa**: Push the top element of stack B to stack A.
- **pb**: Push the top element of stack A to stack B.
- **sa**: Swap the first two elements of stack A.
- **sb**: Swap the first two elements of stack B.
- **ss**: Execute sa and sb.
- **ra**: Rotate stack A (shift up all elements by 1, the first element becomes the last).
- **rb**: Rotate stack B.
- **rr**: Execute ra and rb.
- **rra**: Reverse rotate A (shift down all elements by 1, the last element becomes the first).
- **rrb**: Reverse rotate B.
- **rrr**: Execute rra and rrb.

Return `n` size of instructions for sorting `x` number of values: 
- if `x` = 3 then `n` <= 3
- if `x` = 5 then `n` <= 12
- if `x` = 100 then `n` <= 1500
- if `x` = 500 then `n` <= 11500

<rb>

## Programs

- push-swap

This program calculates and displays the smallest set of instructions to sort the stack `a` in ascending order.

**Usage:**

```sh
$ ./push-swap "2 1 3 6 5 8"
pb
pb
ra
sa
rrr
pa
pa
```

<br>

- checker

This program reads instructions from standard input and executes them on the given stack a. It then checks if the stack is sorted and stack b is empty.

**Usage:**
```shell
$ ./checker "3 2 1 0"
sa
rra
pb

KO

$ echo -e "rra\npb\nsa\nrra\npa" | ./checker "3 2 1 0"
OK
```


1. Clone the repository:
```shell
git clone https://github.com/yourusername/push-swap.git

```

2. Navigate to the project directory:
```shell
cd push-swap
```

3. Build the project:
```shell
go build -o push-swap ./SwapProg
go build -o checker ./CheckerProg
```

## Example 

```shell
$ ARG="2 1 3 6 5 8"; ./push-swap "$ARG" | wc -l
8

$ ARG="2 1 3 6 5 8"; ./push-swap "$ARG" | ./checker "$ARG"
OK
```