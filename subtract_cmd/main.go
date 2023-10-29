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
		return
	}
	b, err := validation.ValidateInput(os.Args[2])
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(a - b)
}
