package main

import (
	"errors"
	"flag"
	"fmt"
	"os"
	"os/exec"
)

var dry bool
var version string

func main() {
	flag.BoolVar(&dry, "d", false, "Print the full command but not execute")
	flag.StringVar(&version, "version", "", "Specific package version to enable")

	flag.Usage = func() {
		fmt.Fprintf(os.Stdout, "Supported arguments: %v\n", SupportedModules())
		fmt.Fprintln(os.Stdout, "")
		fmt.Fprintln(os.Stdout, "Options:")
		flag.PrintDefaults()
	}
	flag.Parse()

	args := flag.Args()
	if len(args) == 0 {
		handleError(errors.New("No modules provided"))
	}

	deps, err := ParseDeps(args)
	handleError(err)

	cmdStr := deps.Expand()

	fmt.Printf("Command [%s]\n", cmdStr)
	if dry {
		os.Exit(0)
	}

	fmt.Println("")
	err = execCmd(cmdStr)
	handleError(err)
}

func handleError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
	}
}

func execCmd(cmdStr string) error {
	cmd := exec.Command("sh", "-c", cmdStr)

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin

	return cmd.Run()
}
