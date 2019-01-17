package main

import (
	"bufio"
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

	cmdStr := deps.Expand()

	fmt.Println("Running: " + cmdStr)
	if *dryRun {
		return
	}

	err = execCmd(cmdStr)
	if err != nil {
		fmt.Println(err.Error())
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

func execCmd(cmdStr string) error {
	cmd := exec.Command("sh", "-c", cmdStr)

	cmdReaderStd, err := cmd.StdoutPipe()
	if err != nil {
		return err
	}

	scannerStd := bufio.NewScanner(cmdReaderStd)
	go func() {
		for scannerStd.Scan() {
			fmt.Printf("output | %s\n", scannerStd.Text())
		}
	}()

	cmdReaderErr, err := cmd.StderrPipe()
	if err != nil {
		return err
	}

	scannerErr := bufio.NewScanner(cmdReaderErr)
	go func() {
		for scannerErr.Scan() {
			fmt.Printf("error | %s\n", scannerErr.Text())
		}
	}()

	err = cmd.Start()
	if err != nil {
		return err
	}

	err = cmd.Wait()
	if err != nil {
		return err
	}

	return nil
}
