package main_test

import (
	// "os"
	"testing"

	enable "github.com/sudomabider/docker-enable"
)

func TestInvalidArg(t *testing.T) {
	// os.Setenv("PHPIZE_DEPS", "autoconf dpkg-dev dpkg file g++ gcc libc-dev make pkgconf re2c")

	_, err := enable.ParseDeps([]string{"a"})

	if err == nil {
		t.Error("Expected error of unrecoginized argument; got none")
	}
}

func TestExpand(t *testing.T) {
	// os.Setenv("PHPIZE_DEPS", "autoconf dpkg-dev dpkg file g++ gcc libc-dev make pkgconf re2c")

	deps, err := enable.ParseDeps([]string{"git", "postgres", "mysql", "zip", "soap", "crux", "swoole"})

	if err != nil {
		t.Errorf("Expected no error; got %s", err.Error())
	}

	// TODO: instead of testing strings directly, test the intermediate slice values (needs refactoring)
	expected := "apk add --no-cache git openssh-client postgresql-dev mysql-client zlib-dev libxml2-dev && apk add --no-cache --virtual .build $PHPIZE_DEPS && pecl install swoole && docker-php-ext-enable swoole && docker-php-ext-install bcmath sockets pdo_pgsql pdo_mysql zip soap && apk del .build"
	got := deps.Expand()
	if got != expected {
		t.Errorf("Expected: %s\nGot: %s", expected, got)
	}
}
