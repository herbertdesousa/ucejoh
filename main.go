package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	file, err := os.ReadFile("code.uc")

	if err != nil {
		log.Fatal(err)
	}

	instructions := strings.Split(string(file), "\n")

	for i := 0; i < len(instructions); i++ {
		fmt.Println(instructions[i])
	}

	/* stack := Stack{}

	stack.Push(NewStackValue(1))

	stack.Push(NewStackValue(2))

	stack.Push(NewStackValue(3))
	stack.Print()
	fmt.Println() */
}
