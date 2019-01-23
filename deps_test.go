package main

import (
	"testing"
)

func TestInvalidDep(t *testing.T) {
	_, err := ParseDeps([]string{"wrong_dep"})

	if err == nil {
		t.Error("Expected error of unrecoginized argument; got none")
	}
}

func TestExpand(t *testing.T) {
	deps, err := ParseDeps([]string{"git", "postgres", "mysql", "zip", "soap", "crux", "swoole", "pm2"})

	if err != nil {
		t.Errorf("Expected no error; got %s", err.Error())
	}

	expected := Command{
		sys:   []string{"git", "libxml2-dev", "libzip-dev", "mysql-client", "nodejs", "npm", "openssh-client", "postgresql-dev", "zlib-dev"},
		build: []string{"$PHPIZE_DEPS"},
		php:   []string{"bcmath", "pdo_mysql", "pdo_pgsql", "soap", "sockets", "zip"},
		pecl:  []string{"swoole"},
		npm:   []string{"pm2"},
	}
	parts := deps.buildCmd()
	if !commandsEqual(expected, parts) {
		t.Errorf("Expected: %s\nGot: %s", expected, parts)
	}
}

func commandsEqual(a, b Command) bool {
	return slicesEqual(a.sys, b.sys) &&
		slicesEqual(a.build, b.build) &&
		slicesEqual(a.php, b.php) &&
		slicesEqual(a.pecl, b.pecl) &&
		slicesEqual(a.npm, b.npm)
}

func slicesEqual(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}

	for k, v := range a {
		if v != b[k] {
			return false
		}
	}

	return true
}
