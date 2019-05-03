package command

import (
    "enable/recipe"
    "errors"
    "fmt"
    "sort"
    "strings"
)

type Builder struct {
	pre         []string
	post        []string
	system      []string
	build       []string
	phpInstall  []string
	phpEnable   []string
	peclInstall []string
	npmInstall  []string
}

func NewBuilder(r recipe.Recipe, args []string) (Builder, error) {
    if len(args) == 0 {
        return Builder{}, errors.New("no modules provided")
    }

    var b Builder

    for _, a := range args {
        var name, ver string
        if m := strings.Split(a, "@"); len(m) > 1 {
            name = m[0]
            ver = m[1]
        } else {
            name = a
        }

        m, ok := r[name]
        if !ok {
            return b, fmt.Errorf("[%s] is invalid or not supported", name)
        }

        m.Version = ver
        err := b.Add(m)
        if err != nil {
            return b, err
        }
    }

    return b, nil
}

func (b *Builder) Add(d recipe.Def) error {
	pecl := make([]string, len(d.Pecl))
	copy(pecl, d.Pecl)

	if d.Version != "" && len(pecl) > 1 {
		return errors.New("multiple versionable candidates")
	}

	if d.Version != "" && len(pecl) == 1 {
		pecl[0] = fmt.Sprintf("%s-%s", pecl[0], d.Version)
	}

	if d.Pre != "" {
		b.pre = append(b.pre, d.Pre)
	}

	if d.Post != "" {
		b.post = append(b.post, d.Post)
	}

	b.system = append(b.system, d.System...)
	b.build = append(b.build, d.Build...)
	b.phpInstall = append(b.phpInstall, d.Php...)
	b.peclInstall = append(b.peclInstall, pecl...)
	b.phpEnable = append(b.phpEnable, d.Pecl...)
	b.npmInstall = append(b.npmInstall, d.Npm...)

	return nil
}

func (b *Builder) dedupAndSort() {
	b.system = dedupAndSort(b.system)
	b.build = dedupAndSort(b.build)
	b.phpInstall = dedupAndSort(b.phpInstall)
	b.phpEnable = dedupAndSort(b.phpEnable)
	b.peclInstall = dedupAndSort(b.peclInstall)
	b.npmInstall = dedupAndSort(b.npmInstall)
}

func dedupAndSort(vs []string) []string {
	keys := make(map[string]bool)
    var list []string
	for _, entry := range vs {
		if _, value := keys[entry]; !value {
			keys[entry] = true
			list = append(list, entry)
		}
	}

	sort.Strings(list)

	return list
}

func (b Builder) Expand() string {
    b.dedupAndSort()

	var result []string

	if len(b.pre) > 0 {
		result = append(result, strings.Join(b.pre, " && "))
	}

	if len(b.system) > 0 {
		result = append(result, "apk add --no-cache "+strings.Join(b.system, " "))
	}

	if len(b.build) > 0 {
		result = append(result, "apk add --no-cache --virtual .build "+strings.Join(b.build, " "))
	}

	if len(b.peclInstall) > 0 {
		result = append(result, fmt.Sprintf("pecl install %s", strings.Join(b.peclInstall, " ")))
	}

	if len(b.phpEnable) > 0 {
		result = append(result, fmt.Sprintf("docker-php-ext-enable %s", strings.Join(b.phpEnable, " ")))
	}

	if len(b.phpInstall) > 0 {
		result = append(result, "docker-php-ext-install "+strings.Join(b.phpInstall, " "))
	}

	if len(b.npmInstall) > 0 {
		result = append(result, "npm i -g "+strings.Join(b.npmInstall, " "))
	}

	if len(b.post) > 0 {
		result = append(result, strings.Join(b.post, " && "))
	}

	if len(b.build) > 0 {
		result = append(result, "apk del .build")
	}

	return strings.Join(result, " && \\\n")
}
