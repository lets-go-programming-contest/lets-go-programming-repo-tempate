package main

import (
	"fmt"
	"os"

	"github.com/Tuc0Sa1amanka/task-1/pkg/calculator"
)

func main() {
	var a int
	var b int
	var c rune

	if _, err := fmt.Scanln(&a); err != nil {
		fmt.Println("Invalid first operand")
		os.Exit(1)
	}

	if _, err := fmt.Scanln(&b); err != nil {
		fmt.Println("Invalid second operand")
		os.Exit(1)
	}

	_, err := fmt.Scanf("%c", &c)
	if err != nil || !calculator.IsOperator(c) {
		fmt.Println("Invalid operation")
		os.Exit(1)
	}

	switch c {
	case '+':
		fmt.Println(calculator.Sum(a, b))
	case '-':
		fmt.Println(calculator.Dif(a, b))
	case '*':
		fmt.Println(calculator.Mul(a, b))
	case '/':
		res, err := calculator.Div(a, b)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println(res)
	}
}
