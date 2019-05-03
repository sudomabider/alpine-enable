package main

import (
    "enable/command"
    "enable/php"
    "enable/recipe"
    "flag"
    "fmt"
    "os"
    "os/exec"
    "strings"
)

const usage = `Usage: enable OPTIONS [args...]

Options:
  --all|-a          Enable all available modules
  --version|-v      Print version
  --help|-h         Print usage
  --dry|-d          Print the command but not execute

args: [%s]`

var (
    version    string
    buildTime  string
    commitHash string

    dryFlag     bool
    versionFlag bool
    helpFlag    bool
    allFlag     bool
)

func main() {
    v, err := php.GetMajorVersion()
    handleError(err)

    r, err := recipe.GetPHPRecipe(v)
    handleError(err)

    parseVars(r)
    run(r)
}

func parseVars(r recipe.Recipe) {
    flag.BoolVar(&dryFlag, "dry", false, "")
    flag.BoolVar(&dryFlag, "d", false, "")
    flag.BoolVar(&versionFlag, "version", false, "")
    flag.BoolVar(&versionFlag, "v", false, "")
    flag.BoolVar(&helpFlag, "help", false, "")
    flag.BoolVar(&helpFlag, "h", false, "")
    flag.BoolVar(&allFlag, "all", false, "")
    flag.BoolVar(&allFlag, "a", false, "")

    flag.Usage = func() {
        mods := r.Modules()
        fmt.Printf(usage, strings.Join(mods, " "))
    }
    flag.Parse()
}

func run(r recipe.Recipe) {
    if versionFlag {
        fmt.Printf("Version: %s\nBuilt time: %s\nCommit hash: %s\n", version, buildTime, commitHash)
        os.Exit(0)
    }

    if helpFlag {
        flag.Usage()
        os.Exit(0)
    }

    var args []string
    if allFlag {
        args = r.Modules()
    } else {
        args = flag.Args()
    }
    b, err := command.NewBuilder(r, args)
    handleError(err)

    fmt.Printf("[Command]\n%s\n", b.Expand())
    if dryFlag {
        os.Exit(0)
    }

    fmt.Println("")

    c := exec.Command("sh", "-c", b.Expand())
    c.Stdout = os.Stdout
    c.Stderr = os.Stderr
    c.Stdin = os.Stdin

    err = c.Run()
    handleError(err)
}

func handleError(err error) {
    if err != nil {
        _, _ = fmt.Fprintln(os.Stderr, err.Error())
        os.Exit(1)
    }
}
