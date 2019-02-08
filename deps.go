package main

import (
	"fmt"
	"sort"
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

// Command holds the build parts of the final command
type Command struct {
	sys   []string
	build []string
	php   []string
	pecl  []string
	npm   []string
}

func (ds Deps) buildCmd() Command {
	c := Command{}

	for _, d := range ds {
		if len(d.sys) != 0 {
			c.sys = append(c.sys, d.sys...)
		}
		if len(d.php) != 0 {
			c.php = append(c.php, d.php...)
		}
		if len(d.build) != 0 {
			c.build = append(c.build, d.build...)
		}
		if len(d.pecl) != 0 {
			c.pecl = append(c.pecl, d.pecl...)
		}
		if len(d.npm) != 0 {
			c.npm = append(c.npm, d.npm...)
		}
	}

	c.sys = refine(c.sys)
	c.php = refine(c.php)
	c.build = refine(c.build)
	c.pecl = refine(c.pecl)
	c.npm = refine(c.npm)

	return c
}

// Expand the Deps into executable shell command
func (ds Deps) Expand() string {
	c := ds.buildCmd()

	var cmd []string
	var sysCmd, phpCmd, buildCmd, peclCmd, packageVersion, npmCmd string

	if len(c.sys) > 0 {
		sysCmd = "apk add --no-cache " + strings.Join(c.sys, " ")
		cmd = append(cmd, sysCmd)
	}

	if len(c.build) > 0 {
		buildCmd = "apk add --no-cache --virtual .build " + strings.Join(c.build, " ")
		cmd = append(cmd, buildCmd)
	}

	if len(c.pecl) > 0 {
		pecl := strings.Join(c.pecl, " ")

        packageVersion = pecl;

		if (version != "") {
            packageVersion = fmt.Sprintf("%s-%s", pecl, version);
        }

		peclCmd = fmt.Sprintf("pecl install %s && docker-php-ext-enable %s", packageVersion, pecl)
		cmd = append(cmd, peclCmd)
	}

	if len(c.php) > 0 {
		phpCmd = "docker-php-ext-install " + strings.Join(c.php, " ")
		cmd = append(cmd, phpCmd)
	}

	if len(c.npm) > 0 {
		npmCmd = "npm i -g " + strings.Join(c.npm, " ")
		cmd = append(cmd, npmCmd)
	}

	if len(c.build) > 0 {
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

func refine(vs []string) []string {
	keys := make(map[string]bool)
	list := []string{}
	for _, entry := range vs {
		if _, value := keys[entry]; !value {
			keys[entry] = true
			list = append(list, entry)
		}
	}

	sort.Strings(list)

	return list
}
