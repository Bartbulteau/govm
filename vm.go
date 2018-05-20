package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"
)

var version string = "0.0.1 alpha"

func main() {
	// fmt.Println("GoVM 0.0.1")

	if len(os.Args) < 2 {
		println("\nMissing arguments, type \"govm -h\" to get help\n")
	} else {
		if os.Args[1] == "-h" {
			println("\ngovm command help : \n")
			println("[1] govm [filename.asm]")
			println("[2] govm -h")
			println("[3] govm -v\n")
		} else if os.Args[1] == "-v" {
			println("\nGoVM version", version, "\n")
		} else {
			data, err := ioutil.ReadFile(os.Args[1])
			check(err)
			content := string(data)
			splited := SpeSplit(content)

			// for i := 0; i < len(splited); i++ {
			// 	println("[", i, "] element :", splited[i])
			// }

			run(splited)

		}
	}
}

func check(e error) {
	if e != nil {
		log.Fatal(e)
	}
}

func SpeSplit(s string) []string {
	w := strings.FieldsFunc(s, func(r rune) bool {
		switch r {
		case ' ', '\n', ';', '\t':
			return true
		}
		return false
	})
	return w
}

func run(p []string) {

	running := true
	regs := map[string]int{
		"ax": 0,
		"bx": 0,
		"cx": 0,
		"dx": 0,
	}

	pc := 0

	for running {
		switch p[pc] {
		case "halt":
			running = false

		// memory
		case "mov":
			pc += 2
			if value, err := strconv.Atoi(p[pc]); err == nil {
				pc--
				regs[p[pc]] = value
				pc++
			} else {
				pc--
				regs[p[pc]] = regs[p[pc+1]]
				pc++
			}

		// I/O
		case "printr":
			if regs["ax"] == 0 {
				pc++
				fmt.Printf("%d", regs[p[pc]])
			} else if regs["ax"] == 1 {
				pc++
				fmt.Printf("%c", regs[p[pc]])
			}

		// Math
		case "add":
			pc += 2
			if value, err := strconv.Atoi(p[pc]); err == nil {
				pc--
				regs[p[pc]] += value
				pc++
			} else {
				pc--
				regs[p[pc]] += regs[p[pc+1]]
				pc++
			}
		case "sub":
			pc += 2
			if value, err := strconv.Atoi(p[pc]); err == nil {
				pc--
				regs[p[pc]] -= value
				pc++
			} else {
				pc--
				regs[p[pc]] -= regs[p[pc+1]]
				pc++
			}
		case "mul":
			pc += 2
			if value, err := strconv.Atoi(p[pc]); err == nil {
				pc--
				regs[p[pc]] *= value
				pc++
			} else {
				pc--
				regs[p[pc]] *= regs[p[pc+1]]
				pc++
			}
		case "div":
			pc += 2
			if value, err := strconv.Atoi(p[pc]); err == nil {
				pc--
				regs[p[pc]] /= value
				pc++
			} else {
				pc--
				regs[p[pc]] /= regs[p[pc+1]]
				pc++
			}

		default:
			println("Error invalid instruction : \"", p[pc], "\" [halt]")
			running = false
		}
		pc += 1
	}
}
