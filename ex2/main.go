package main

import (
	"fmt"
	"os"
	"sorter/utils"
)

func main() {
	args := os.Args[1:]

	fmt.Println("output: ", utils.Sorter(args))
}
