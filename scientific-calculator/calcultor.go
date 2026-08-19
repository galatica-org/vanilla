package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
)

func Calculate(op string, a, b float64) float64 {
	switch op {
	case "add":
		return a + b
	case "subtract":
		return a - b
	case "multiply":
		return a * b
	case "divide":
		return a / b
	case "power":
		return math.Pow(a, b)
	case "sqrt":
		return math.Sqrt(a)
	case "sin":
		return math.Sin(a)
	case "cos":
		return math.Cos(a)
	case "tan":
		return math.Tan(a)
	case "log":
		return math.Log10(a)
	case "ln":
		return math.Log(a)
	default:
		panic("Unknown operation: " + op)
	}
}

func main() {
	if len(os.Args) < 3 {
		fmt.Println("Usage: go run calculator.go <op> <a> [b]")
		fmt.Println("Operations: add, subtract, multiply, divide, power, sqrt, sin, cos, tan, log, ln")
		fmt.Println("Examples:")
		fmt.Println("  go run calculator.go add 2 3")
		fmt.Println("  go run calculator.go sqrt 16")
		fmt.Println("  go run calculator.go sin 1.5708")
		os.Exit(1)
	}

	op := os.Args[1]
	a, err := strconv.ParseFloat(os.Args[2], 64)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error: invalid number for a")
		os.Exit(1)
	}

	var b float64
	if len(os.Args) > 3 {
		b, err = strconv.ParseFloat(os.Args[3], 64)
		if err != nil {
			fmt.Fprintln(os.Stderr, "Error: invalid number for b")
			os.Exit(1)
		}
	}

	defer func() {
		if r := recover(); r != nil {
			fmt.Fprintln(os.Stderr, "Error:", r)
			os.Exit(1)
		}
	}()

	result := Calculate(op, a, b)
	fmt.Println(result)
}