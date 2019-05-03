package recipe

import "sort"

// Def is the definition of how a module should be enabled
type Def struct {
    Pre     string
    Post    string
    System  []string
    Build   []string
    Php     []string
    Pecl    []string
    Npm     []string
    Version string
}

// Recipe is a list of module definitions
type Recipe map[string]Def

func (r Recipe) Modules() []string {
    var mods []string
    for k := range r {
        mods = append(mods, k)
    }

    sort.Strings(mods)
    return mods
}
