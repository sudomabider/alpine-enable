package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"

	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()
	app.Name = "Alpine enable"
	app.Usage = fmt.Sprintf("easily enable [%s]", strings.Join(SupportedModules(), ","))
	app.Version = "0.0.0"
	app.Flags = []cli.Flag{
		cli.BoolFlag{
			Name:   "dry, d",
			Hidden: false,
			Usage:  "print the full command but not execute",
		},
	}
	app.HideHelp = true
	app.Action = do

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}

func do(c *cli.Context) error {
	args := c.Args()
	if len(args) == 0 {
		cli.ShowAppHelp(c)
		return nil
	}

	deps, err := ParseDeps(args)
	if err != nil {
		return err
	}

	cmdStr := deps.Expand()

	fmt.Println("Command: " + cmdStr)
	if c.Bool("dry") {
		return nil
	}

	err = execCmd(cmdStr)
	if err != nil {
		return err
	}

	return nil
}

func execCmd(cmdStr string) error {
	cmd := exec.Command("sh", "-c", cmdStr)

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin

	err := cmd.Run()
	if err != nil {
		return err
	}

	return nil
}
