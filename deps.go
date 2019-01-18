package main

import (
	"fmt"
	"strings"
)

// Dep holds all dependency information for a module
type Dep struct {
	sys   []string
	build []string
	php   []string
	pecl  []string
	npm   []string
}

// Deps holds a list of deps to be installed
type Deps map[string]Dep

// Expand the Deps into executable shell command
func (ds Deps) Expand() string {
	var sys, php, build, pecl, npm []string

	for _, d := range ds {
		if len(d.sys) != 0 {
			sys = append(sys, d.sys...)
		}
		if len(d.php) != 0 {
			php = append(php, d.php...)
		}
		if len(d.build) != 0 {
			build = append(build, d.build...)
		}
		if len(d.pecl) != 0 {
			pecl = append(pecl, d.pecl...)
		}
		if len(d.npm) != 0 {
			npm = append(npm, d.npm...)
		}
	}

	sys = dedup(sys)
	php = dedup(php)
	build = dedup(build)
	pecl = dedup(pecl)
	npm = dedup(npm)

	var cmd []string
	var sysCmd, phpCmd, buildCmd, peclCmd, npmCmd string

	if len(sys) > 0 {
		sysCmd = "apk add --no-cache " + strings.Join(sys, " ")
		cmd = append(cmd, sysCmd)
	}

	if len(build) > 0 {
		buildCmd = "apk add --no-cache --virtual .build " + strings.Join(build, " ")
		cmd = append(cmd, buildCmd)
	}

	if len(pecl) > 0 {
		peclCmd = fmt.Sprintf("pecl install %s && docker-php-ext-enable %s", strings.Join(pecl, " "), strings.Join(pecl, " "))
		cmd = append(cmd, peclCmd)
	}

	if len(php) > 0 {
		phpCmd = "docker-php-ext-install " + strings.Join(php, " ")
		cmd = append(cmd, phpCmd)
	}

	if len(npm) > 0 {
		npmCmd = "npm i -g " + strings.Join(npm, " ")
		cmd = append(cmd, npmCmd)
	}

	if len(build) > 0 {
		cmd = append(cmd, "apk del .build")
	}

	return strings.Join(cmd, " && ")
}

// ParseDeps parses a slice of strings into Deps
func ParseDeps(vs []string) (Deps, error) {
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

func dedup(vs []string) []string {
	keys := make(map[string]bool)
	list := []string{}
	for _, entry := range vs {
		if _, value := keys[entry]; !value {
			keys[entry] = true
			list = append(list, entry)
		}
	}
	return list
}
