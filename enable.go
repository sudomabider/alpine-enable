package main

import (
	"flag"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func main() {
	args := parseArgs()
	dryRun := flag.Bool("dry", false, "Print the command but not execute")
	flag.Parse()

	deps, err := ParseDeps(args)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	cmd := deps.Expand()

	if *dryRun {
		fmt.Println(cmd)
	} else {
		exec.Command("sh", "-c", cmd).Output()
	}
}

func parseArgs() []string {
	args := make([]string, 0)

	for _, arg := range os.Args[1:] {
		if !strings.HasPrefix(arg, "-") {
			args = append(args, arg)
		}
	}

	if len(args) == 0 {
		help()
		return make([]string, 0)
	}

	return args
}

func help() {
	fmt.Print("This is help")
}
