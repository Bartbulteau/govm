package stack

import (
	"fmt"
	"os"
)

type Stack struct {
	stack [2048]int
	sp    int
}

func Create(s *Stack) {
	for i := 0; i < len(s.stack); i++ {
		s.stack[i] = 0
	}
	s.sp = -1
}

func Push(s *Stack, value int) {
	s.sp++
	if s.sp <= 2048 {
		s.stack[s.sp] = value
	} else {
		println("GoVM : Fatal error, stack overflow")
		os.Exit(1)
	}
}

func Pop(s *Stack) int {
	var value int
	if s.sp == -1 {
		value = -1
		println("GoVM : Fatal error, trying to Pop from empty stack")
		os.Exit(1)
	} else if s.sp >= 0 {
		value = s.stack[s.sp]
		s.sp--
	}
	return value
}

func Add(s *Stack) {
	var a int
	var b int

	if s.sp >= 1 {
		a = Pop(s)
		b = Pop(s)
		Push(s, a+b)
	} else {
		println("GoVM : Fatal error, not enough values to add on stack")
		os.Exit(1)
	}
}

func Sub(s *Stack) {
	var a int
	var b int

	if s.sp >= 1 {
		a = Pop(s)
		b = Pop(s)
		Push(s, a-b)
	} else {
		println("GoVM : Fatal error, not enough values to subtract on stack")
		os.Exit(1)
	}
}

func Mul(s *Stack) {
	var a int
	var b int

	if s.sp >= 1 {
		a = Pop(s)
		b = Pop(s)
		Push(s, a*b)
	} else {
		println("GoVM : Fatal error, not enough values to multiply on stack")
		os.Exit(1)
	}
}

func Div(s *Stack) {
	var a int
	var b int

	if s.sp >= 1 {
		a = Pop(s)
		b = Pop(s)
		Push(s, a/b)
	} else {
		println("GoVM : Fatal error, not enough values to divide on stack")
		os.Exit(1)
	}
}

func Print(s *Stack) { // prints a string stored in the stack
	if s.sp >= 0 {
		ok := true
		var value int
		for ok {
			value = Pop(s)
			if value == 0 {
				ok = false
			}
			fmt.Printf("%c", value)
		}
	} else {
		println("GoVM : Fatal error, trying to print from empty stack")
		os.Exit(1)
	}
}
