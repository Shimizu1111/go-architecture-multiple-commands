package main

import (
	"fmt"
	"os"

	"github.com/shimizu/calculator/internal/validation"
)

func main() {
	a, err := validation.ValidateInput(os.Args[1])
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	b, err := validation.ValidateInput(os.Args[2])
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println(a + b)
}
