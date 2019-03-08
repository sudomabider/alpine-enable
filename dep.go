package main

type dep struct {
	pre     string
	post    string
	system  []string
	build   []string
	php     []string
	pecl    []string
	npm     []string
	version string
}
