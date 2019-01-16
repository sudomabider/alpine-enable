package main

import (
	"fmt"
	"os"
)

func main() {
	args := parseArgs()

	deps, err := parseDeps(args)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	cmd := deps.expand()
	fmt.Println(cmd)
}

func parseArgs() []string {
	args := os.Args[1:]

	if len(args) == 0 {
		help()
		return make([]string, 0)
	}

	return args
}

func help() {
	fmt.Print("This is help")
}
