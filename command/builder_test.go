package command

import (
    . "enable/recipe"
    "errors"
    "reflect"
    "testing"
)

var newBuilderTests = []struct {
    name    string
    recipe  Recipe
    args    []string
    builder Builder
    error   error
}{
    {
        "should error when no arg is provided",
        Recipe{},
        []string{},
        Builder{},
        errors.New("no modules provided"),
    },
    {
        "should error when arg is invalid",
        Recipe{"dep1": Def{System: []string{"lib1"}}},
        []string{"module"},
        Builder{},
        errors.New("[module] is invalid or not supported"),
    },
    {
        "should return new builder",
        Recipe{"m1": Def{
            Pre:    "pre1 command",
            Post:   "post1 command",
            Build:  []string{"build1"},
            Php:    []string{"php1"},
            Pecl:   []string{"pecl1"},
            Npm:    []string{"npm1"},
            System: []string{"sys1"},
        }, "m2": Def{
            Pre:    "pre2 command",
            Post:   "post2 command",
            Build:  []string{"build2"},
            Php:    []string{"php2"},
            Pecl:   []string{"pecl2"},
            Npm:    []string{"npm2"},
            System: []string{"sys2"},
        }},
        []string{"m1", "m2"},
        Builder{
            pre:         []string{"pre1 command", "pre2 command"},
            post:        []string{"post1 command", "post2 command"},
            system:      []string{"sys1", "sys2"},
            build:       []string{"build1", "build2"},
            phpInstall:  []string{"php1", "php2"},
            phpEnable:   []string{"pecl1", "pecl2"},
            peclInstall: []string{"pecl1", "pecl2"},
            npmInstall:  []string{"npm1", "npm2"},
        },
        nil,
    },
    {
        "should error when version has multiple candidates",
        Recipe{"module": Def{Pecl: []string{"pecl1", "pecl2"}}},
        []string{"module@1.0"},
        Builder{},
        errors.New("multiple versionable candidates"),
    },
    {
        "should return new builder with versioned pecl install",
        Recipe{"module": Def{
            Php:  []string{"php1"},
            Pecl: []string{"pecl1"},
        }},
        []string{"module@1.0"},
        Builder{
            phpInstall:  []string{"php1"},
            phpEnable:   []string{"pecl1"},
            peclInstall: []string{"pecl1-1.0"},
        },
        nil,
    },
}

func TestNewBuilder(t *testing.T) {
    for _, tt := range newBuilderTests {
        t.Run(tt.name, func(t *testing.T) {
            b, err := NewBuilder(tt.recipe, tt.args)

            if !reflect.DeepEqual(tt.builder, b) {
                t.Errorf("Expected builder: %v\nGot builder: %v", tt.builder, b)
            }

            if !reflect.DeepEqual(tt.error, err) {
                t.Errorf("Expected error: %v\nGot error: %v", tt.error, err)
            }
        })
    }
}

func TestBuilder_Add(t *testing.T) {
    b := Builder{
        system: []string{"sys1"},
        phpEnable: []string{"pecl1"},
        peclInstall: []string{"pecl1"},
    }

    err := b.Add(Def{System: []string{"sys2"}, Pecl: []string{"pecl2"}})

    if err != nil {
        t.Errorf("Got error [%v]; expecting none", err)
    }

    expected := Builder{
        system: []string{"sys1", "sys2"},
        phpEnable: []string{"pecl1", "pecl2"},
        peclInstall: []string{"pecl1", "pecl2"},
    }

    if !reflect.DeepEqual(expected, b) {
        t.Errorf("Expected: %v\n Got: %v", expected, b)
    }
}

func TestDedupAndSort(t *testing.T) {
    ss := []string{"b", "c", "a", "b"}
    expected := []string{"a", "b", "c"}
    result := dedupAndSort(ss)

    if !reflect.DeepEqual(expected, result) {
        t.Errorf("Expected: %v\nGot: %v", expected, result)
    }
}

func TestBuilder_Expand(t *testing.T) {
    b := Builder{
        pre: []string{"pre command", "another pre command"},
        post: []string{"post command", "another post command"},
        system: []string{"sys2", "sys1"},
        build: []string{"build2", "build1"},
        phpInstall: []string{"php2", "php1"},
        phpEnable: []string{"pecl2", "pecl1"},
        peclInstall: []string{"pecl2", "pecl1"},
        npmInstall: []string{"npm2", "npm1"},
    }

    expected := `pre command && another pre command && \
apk add --no-cache sys1 sys2 && \
apk add --no-cache --virtual .build build1 build2 && \
pecl install pecl1 pecl2 && \
docker-php-ext-enable pecl1 pecl2 && \
docker-php-ext-install php1 php2 && \
npm i -g npm1 npm2 && \
post command && another post command && \
apk del .build`

    expanded := b.Expand()
    if expanded != expected {
        t.Errorf("Expected: %s\nGot: %s", expected, expanded)
    }
}
