package main

import (
	"fmt"
	"os"

	mymodule "github.com/johnny880319/project-starters/go/internal/my-module"
)

func main() {
	if err := run(); err != nil {
		fmt.Fprintln(os.Stderr, "Error:", err)
		os.Exit(1)
	}
	fmt.Println("All tests passed successfully!")
}

func run() error {
	if mymodule.MyFunction1() != "Hello, World!" {
		return fmt.Errorf("unexpected result from MyFunction1")
	}
	if mymodule.MyFunction2() != 42 {
		return fmt.Errorf("unexpected result from MyFunction2")
	}
	return nil
}
