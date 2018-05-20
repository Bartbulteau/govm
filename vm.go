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

func run(program []string) {

	running := true
	regs := map[string]int{
		"ax": 0,
		"bx": 0,
		"cx": 0,
		"dx": 0,
		"jx": 0,
	}

	pc := 0

	p, labels := Labelprocess(program)
	// for key, val := range labels {
	// 	println("key:", key, "\tval:", val)
	// }

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
		// case "print":
		// 	if regs["ax"] == 0 {
		// 		pc++
		// 		fmt.Printf("%d", p[pc][1])
		// 	} else if regs["ax"] == 1 {
		// 		pc++
		// 		val, err := strconv.Atoi(p[pc])
		// 		println(val)
		// 		if err != nil {
		// 			fmt.Printf("%c", val)
		// 		}

		// 	}

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

		// Logic
		case "jmp":
			pc++
			adress, err := strconv.Atoi(p[pc])
			if err != nil {
				println("Error trying to jump to invalid adress :", p[pc])
				running = false
			} else {
				pc = adress
			}
		case "jz":
			pc++
			adress, err := strconv.Atoi(p[pc])
			if err != nil {
				println("Error trying to jump to invalid adress :", p[pc])
			} else {
				if regs["jx"] == 0 {
					pc = adress
				}
			}
		case "jnz":
			pc++
			adress, err := strconv.Atoi(p[pc])
			if err != nil {
				println("Error trying to jump to invalid adress :", p[pc])
			} else {
				if regs["jx"] != 0 {
					pc = adress
				}
			}

		default:
			if val, ok := labels[p[pc][:len(p[pc])-1]]; ok {
				val++
				val--
			} else {
				println("Error invalid instruction : \"", p[pc], "\" [halt]")
				running = false
			}

		}
		pc += 1
	}
}

func Labelprocess(p []string) ([]string, map[string]int) {
	labels := make(map[string]int)
	for i := 0; i < len(p); i++ {
		if p[i][len(p[i])-1] == ':' {
			labels[p[i][:len(p[i])-1]] = i
		}
	}
	for j := 0; j < len(p); j++ {
		if val, ok := labels[p[j]]; ok {
			p[j] = strconv.Itoa(val)
		}
	}
	return p, labels
}
