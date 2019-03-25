package main

import (
	"testing"
)

func TestInvalidDep(t *testing.T) {
	_, err := parseArgs([]string{"wrong_dep"}, Recipe{})

	if err == nil {
		t.Error("Expected error of unrecoginized argument; got none")
	}
}

func TestCannotAddVersionedDepWithMultipleCandidates(t *testing.T) {
	testCases := []struct {
		name   string
		recipe Recipe
	}{
		{"pecl", Recipe{"dep": dep{pecl: []string{"a", "b"}}}},
		{"npm", Recipe{"dep": dep{npm: []string{"a", "b"}}}},
		{"pecl and npm", Recipe{"dep": dep{pecl: []string{"a", "b"}, npm: []string{"c", "d"}}}},
	}

	for _, tc := range testCases {
		_, err := parseArgs([]string{"dep@1.0"}, tc.recipe)

		if err == nil {
			t.Error("Expected error of too many versionable candidates; got none")
		}
	}
}

func TestVersionIsIgnoredForUnversionableDeps(t *testing.T) {
	testCases := []struct {
		name   string
		recipe Recipe
	}{
		{"system", Recipe{"dep": dep{system: []string{"a"}}}},
		{"php", Recipe{"dep": dep{php: []string{"a"}}}},
	}

	for _, tc := range testCases {
		c, err := parseArgs([]string{"dep@1.0"}, tc.recipe)

		if err != nil {
			t.Errorf("Expected no error; got %s", err.Error())
		}

		switch tc.name {
		case "system":
			v := c.system[0]
			if v != "a" {
				t.Errorf("Expected version to be dropped; got %s", v)
			}
		case "php":
			v := c.phpInstall[0]
			if v != "a" {
				t.Errorf("Expected version to be dropped; got %s", v)
			}
		}
	}
}

func TestExpand(t *testing.T) {
	c, err := parseArgs([]string{"git", "postgres", "mysql", "zip", "soap", "crux", "swoole", "pm2", "mcrypt"}, recipe)

	if err != nil {
		t.Errorf("Expected no error; got %s", err.Error())
	}

	expected := command{
		system:      []string{"git", "libltdl", "libmcrypt-dev", "libxml2-dev", "libzip-dev", "mysql-client", "nodejs", "npm", "openssh-client", "postgresql-dev", "zlib-dev"},
		build:       []string{"$PHPIZE_DEPS"},
		phpInstall:  []string{"bcmath", "mcrypt", "mysql", "pcntl", "pdo_mysql", "pdo_pgsql", "pgsql", "soap", "sockets", "zip"},
		peclInstall: []string{"swoole"},
		phpEnable:   []string{"swoole"},
		npmInstall:  []string{"pm2"},
	}

	if !commandEquals(expected, c) {
		t.Errorf("Expected: %s\nGot: %s", expected, c)
	}
}

func TestExpandWithVersion(t *testing.T) {
	c, err := parseArgs([]string{"git", "mysql", "zip", "swoole@1.1", "pm2@2.2"}, recipe)

	if err != nil {
		t.Errorf("Expected no error; got %s", err.Error())
	}

	expected := command{
		system:      []string{"git", "libzip-dev", "mysql-client", "nodejs", "npm", "openssh-client", "zlib-dev"},
		build:       []string{"$PHPIZE_DEPS"},
		phpInstall:  []string{"mysql", "pcntl", "pdo_mysql", "zip"},
		peclInstall: []string{"swoole-1.1"},
		phpEnable:   []string{"swoole"},
		npmInstall:  []string{"pm2@2.2"},
	}

	if !commandEquals(expected, c) {
		t.Errorf("Expected: %s\nGot: %s", expected, c)
	}
}

func commandEquals(a, b command) bool {
	return sliceEquals(a.system, b.system) &&
		sliceEquals(a.build, b.build) &&
		sliceEquals(a.phpInstall, b.phpInstall) &&
		sliceEquals(a.peclInstall, b.peclInstall) &&
		sliceEquals(a.phpEnable, b.phpEnable) &&
		sliceEquals(a.npmInstall, b.npmInstall)
}

func sliceEquals(a, b []string) bool {
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
