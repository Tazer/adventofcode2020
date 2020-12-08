package main

import (
	"bufio"
	"flag"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {

	var version = flag.Int("version", 1, "first or second part of the assignment")

	flag.Parse()

	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	lines := []string{}

	for scanner.Scan() {
		l := scanner.Text()
		if err != nil {
			log.Fatal(err)
		}
		lines = append(lines, l)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	log.Print(lines[0])

	instructions := parseIntstructions(lines)

	res, _ := runProgram(instructions)

	res2, _ := executeInstructions(instructions, instructions, 0)
	log.Printf("Version: %d ,Result: %d , Result v2: %d 🎄", *version, res, res2)

}

func parseIntstructions(lines []string) []instruction {
	instructions := []instruction{}

	for _, l := range lines {
		arr := strings.Split(l, " ")

		v, _ := strconv.Atoi(strings.ReplaceAll(arr[1], "+", ""))

		i := instruction{
			op:    arr[0],
			value: v,
		}
		instructions = append(instructions, i)
	}
	return instructions
}

func runProgram(instructions []instruction) (int, bool) {
	acc := 0
	index := 0
	completed := true

	for {
		i := instructions[index]

		if i.executed {
			log.Printf("whats executed %+v", i)
			completed = false
			break
		}

		instructions[index].executed = true
		switch i.op {
		case "nop":
			index++
		case "acc":
			acc += i.value
			index++
		case "jmp":
			index = index + i.value
		}
		if index > len(instructions)-1 {
			completed = true
			break
		}

	}
	return acc, completed
}

func executeInstructions(instructions []instruction, original []instruction, fixIndex int) (int, bool) {

	acc, completed := runProgram(instructions)

	if !completed {

		instructions := []instruction{}
		for _, o := range original {
			n := o
			n.executed = false
			instructions = append(instructions, n)
		}

		fix := instructions[fixIndex]

		log.Printf("Trying to fix inst: %d with %v", fixIndex, fix)
		switch fix.op {
		case "nop":
			instructions[fixIndex].op = "jmp"
		case "jmp":
			instructions[fixIndex].op = "nop"
		}
		fixIndex++
		log.Printf("Running following instructions \n %+v", instructions)
		return executeInstructions(instructions, original, fixIndex)
	}

	return acc, completed
}

type instruction struct {
	op       string
	value    int
	executed bool
}
