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

// Deps holds a list of deps to be installed
type Deps map[string]Dep

// Expand the Deps into executable shell command
func (ds Deps) Expand() string {
	var sys, php, build, pecl, npm []string

	for _, d := range ds {
		if d.sys != "" {
			sys = append(sys, d.sys)
		}
		if d.php != "" {
			php = append(php, d.php)
		}
		if d.build != "" {
			build = append(build, d.build)
		}
		if d.pecl != "" {
			pecl = append(pecl, d.pecl)
		}
		if d.npm != "" {
			npm = append(npm, d.npm)
		}
	}

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
