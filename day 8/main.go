package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	input := readInput("./in.txt")
	part1(input)
	part2(input)
}

const (
	// Instructions
	accum = 0
	jump  = 1
	noop  = -1
)

const (
	ok = iota
	infiniteLoop
	inconsistentMemory
	magic // this code should never be executed
)

type instruction struct {
	opCode   int
	argument int
}

func makeMemory(data []string) []instruction {
	memory := make([]instruction, len(data))
	convertOpcode := func(opCode string) int {
		switch opCode {
		case "acc":
			return accum
		case "jmp":
			return jump
		case "nop":
			return noop
		default:
			panic("Unknown opcode")
		}
	}
	for idx, line := range data {
		instructionSplit := strings.Split(line, " ")
		argument, _ := strconv.Atoi(instructionSplit[1])
		memory[idx] = instruction{
			opCode:   convertOpcode(instructionSplit[0]),
			argument: argument,
		}
	}

	return memory
}

func runMachine(memory []instruction) (int, int) {
	visited := make([]bool, len(memory))

	acc := 0
	pc := 0

	executeAcc := func(arg int) {
		acc += arg
		pc += 1
	}

	executeJmp := func(arg int) {
		pc += arg
	}

	executeNop := func(arg int) {
		pc += 1
	}

	execute := func(opCode, arg int) {
		switch opCode {
		case accum:
			executeAcc(arg)
		case jump:
			executeJmp(arg)
		case noop:
			executeNop(arg)
		}
	}

	for true {
		if pc == len(memory) {
			//	execute the last instruction
			return ok, acc
		} else if pc > len(memory) {
			fmt.Print("inconsistent state\n")
			return inconsistentMemory, acc
		} else {
			if visited[pc] == false {
				visited[pc] = true
			} else {
				return infiniteLoop, acc
			}
			currentInstruction := memory[pc]
			execute(currentInstruction.opCode, currentInstruction.argument)
		}
	}
	return magic, acc
}

func part1(data []string) {
	memory := makeMemory(data)
	statusCode, acc := runMachine(memory)
	print("Part 1: " )
	if statusCode == infiniteLoop {
		println("ACC: ",acc)
	} else {
		println("Unexpected statusCode")
	}
}

func part2(data []string) {
	inputMemory := makeMemory(data)
	swapOp := []int{}
	for idx, instr := range inputMemory {
		if (instr.opCode == noop && instr.argument > 0) || (instr.opCode == jump && instr.argument < 0) {
			swapOp = append(swapOp, idx)
		}
	}

	for _, memidx := range swapOp {
		memory := make([]instruction, len(inputMemory))
		copy(memory, inputMemory)
		memory[memidx].opCode *= -1
		if code, acc := runMachine(memory); code == 0 {
			println("Successful run: memidx", memidx, " acc:", acc)
			return
		}
	}
}
