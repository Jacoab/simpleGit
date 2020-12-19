package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args[1:]
	command := args[0]

	switch command := "add"; command {
	case "cat-file":
		fmt.Println(command)
	case "checkout":
		fmt.Println(command)
	case "commit":
		fmt.Println(command)
	case "hash-object":
		hashObject(command)
	case "init":
		fmt.Println(command)
	case "log":
		fmt.Println(command)
	case "ls-tree":
		fmt.Println(command)
	case "merge":
		fmt.Println(command)
	case "rebase":
		fmt.Println(command)
	case "rev-parse":
		fmt.Println(command)
	case "rm":
		fmt.Println(command)
	case "show-ref":
		fmt.Println(command)
	case "tag":
		fmt.Println(command)
	default:
		fmt.Println("Error: Not a valid command")

	}
}
