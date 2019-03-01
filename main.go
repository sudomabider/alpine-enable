package main

import (
	"errors"
	"flag"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

const (
	usage = `Usage: enable OPTIONS [args...]

Options:
  --version|-v      Print version
  --help|-h         Print usage
  --dry|-d          Print the command but not execute

args: [%s]`
)

var version string
var buildTime string
var commitHash string

var dryFlag bool
var versionFlag bool
var helpFlag bool

func main() {
	Run(recipe)
}

//Run the program
func Run(recipe Recipe) {
	flag.BoolVar(&dryFlag, "dry", false, "")
	flag.BoolVar(&dryFlag, "d", false, "")
	flag.BoolVar(&versionFlag, "version", false, "")
	flag.BoolVar(&versionFlag, "v", false, "")
	flag.BoolVar(&helpFlag, "help", false, "")
	flag.BoolVar(&helpFlag, "h", false, "")

	flag.Usage = func() {
		var modules []string
		for k := range recipe {
			modules = append(modules, k)
		}

		fmt.Fprintf(os.Stderr, usage, strings.Join(modules, " "))
	}
	flag.Parse()

	if versionFlag {
		fmt.Fprintf(os.Stderr, "Version: %s\nBuilt time: %s\nCommit hash: %s\n", version, buildTime, commitHash)
		os.Exit(0)
	}

	if helpFlag {
		flag.Usage()
		os.Exit(0)
	}

	args := flag.Args()
	c, err := parseArgs(args, recipe)
	if err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(1)
	}

	cmdStr := c.expand()

	if dryFlag {
		fmt.Fprintf(os.Stderr, "Command [%s]", cmdStr)
		os.Exit(0)
	}

	fmt.Printf("Command [%s]\n", cmdStr)
	fmt.Println("")

	execCmd(cmdStr)
}

func parseArgs(args []string, recipe Recipe) (command, error) {
	if len(args) == 0 {
		return command{}, errors.New("No modules provided")
	}

	c := command{}

	for _, a := range args {
		var name, version string
		if p := strings.Split(a, "@"); len(p) > 1 {
			name = p[0]
			version = p[1]
		} else {
			name = a
		}

		dep, ok := recipe[name]
		if !ok {
			return command{}, fmt.Errorf("[%s] is not recogized", name)
		}

		dep.version = version
		err := c.addDep(dep)
		if err != nil {
			return command{}, err
		}
	}

	c.dedupAndSort()

	return c, nil
}

func execCmd(cmdStr string) {
	cmd := exec.Command("sh", "-c", cmdStr)

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin

	err := cmd.Run()

	if err != nil {
		fmt.Fprint(os.Stderr, err.Error())
		os.Exit(1)
	}
}
