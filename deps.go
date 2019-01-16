package main

import (
	"fmt"
	"strings"
)

// Dep holds all dependency information for a module
type Dep struct {
	sys   string
	build string
	php   string
	pecl  string
	npm   string
}

var reqs = map[string]Dep{
	"git": Dep{
		sys: "git openssh-client",
	},
	"postgres": Dep{
		sys: "postgresql-dev",
		php: "pdo_pgsql",
	},
}

// Deps holds a list of deps to be installed
type Deps map[string]Dep

func (ds Deps) expand() string {
	cmd := struct {
		sys []string
	}{}

	for _, d := range ds {
		if d.sys != "" {
			cmd.sys = append(cmd.sys, d.sys)
		}
	}

	var r string

	if len(cmd.sys) > 0 {
		r = "apk add --no-cache " + strings.Join(cmd.sys, " ")
	}

	return r
}

func parseDeps(vs []string) (Deps, error) {
	deps := make(Deps)

	for _, v := range vs {
		d, ok := reqs[v]
		if !ok {
			return deps, fmt.Errorf("%s is not recogized", v)
		}

		deps[v] = d
	}

	return deps, nil
}
