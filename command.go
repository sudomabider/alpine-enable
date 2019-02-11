package main

import (
	"errors"
	"fmt"
	"sort"
	"strings"
)

type command struct {
	system      []string
	build       []string
	phpInstall  []string
	phpEnable   []string
	peclInstall []string
	npmInstall  []string
}

func (c *command) addDep(d dep) error {
	pecl := make([]string, len(d.pecl))
	copy(pecl, d.pecl)
	npm := d.npm

	if d.version != "" && (len(pecl) > 1 || len(npm) > 1) {
		return errors.New("Too many versionable candidates")
	}

	if d.version != "" && len(pecl) == 1 {
		pecl[0] = fmt.Sprintf("%s-%s", pecl[0], d.version)
	}

	if d.version != "" && len(npm) == 1 {
		npm[0] = fmt.Sprintf("%s@%s", npm[0], d.version)
	}

	c.system = append(c.system, d.system...)
	c.build = append(c.build, d.build...)
	c.phpInstall = append(c.phpInstall, d.php...)
	c.peclInstall = append(c.peclInstall, pecl...)
	c.phpEnable = append(c.phpEnable, d.pecl...)
	c.npmInstall = append(c.npmInstall, d.npm...)

	return nil
}

func (c *command) dedupAndSort() {
	c.system = dedupAndSort(c.system)
	c.build = dedupAndSort(c.build)
	c.phpInstall = dedupAndSort(c.phpInstall)
	c.phpEnable = dedupAndSort(c.phpEnable)
	c.peclInstall = dedupAndSort(c.peclInstall)
	c.npmInstall = dedupAndSort(c.npmInstall)
}

func dedupAndSort(vs []string) []string {
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

func (c command) expand() string {
	var result []string

	if len(c.system) > 0 {
		result = append(result, "apk add --no-cache "+strings.Join(c.system, " "))
	}

	if len(c.build) > 0 {
		result = append(result, "apk add --no-cache --virtual .build "+strings.Join(c.build, " "))
	}

	if len(c.peclInstall) > 0 {
		result = append(result, fmt.Sprintf("pecl install %s", strings.Join(c.peclInstall, " ")))
	}

	if len(c.phpEnable) > 0 {
		result = append(result, fmt.Sprintf("docker-php-ext-enable %s", strings.Join(c.phpEnable, " ")))
	}

	if len(c.phpInstall) > 0 {
		result = append(result, "docker-php-ext-install "+strings.Join(c.phpInstall, " "))
	}

	if len(c.npmInstall) > 0 {
		result = append(result, "npm i -g "+strings.Join(c.npmInstall, " "))
	}

	if len(c.build) > 0 {
		result = append(result, "apk del .build")
	}

	return strings.Join(result, " && ")
}
