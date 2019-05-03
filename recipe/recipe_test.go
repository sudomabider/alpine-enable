package recipe

import (
    "reflect"
    "testing"
)

var combineTests = []struct {
    name     string
    input    []Recipe
    expected Recipe
}{
    {"should combine 1 recipe", []Recipe{
        {"dep1": Def{System: []string{"lib1"}}},
    }, Recipe{
        "dep1": Def{System: []string{"lib1"}},
    }},
    {"should combine 2 recipes", []Recipe{
        {"dep1": Def{System: []string{"lib1"}}},
        {"dep2": Def{System: []string{"lib2"}}},
    }, Recipe{
        "dep1": Def{System: []string{"lib1"}},
        "dep2": Def{System: []string{"lib2"}},
    }},
    {"should combine 3 recipes", []Recipe{
        {"dep1": Def{System: []string{"lib1"}}},
        {"dep2": Def{System: []string{"lib2"}}},
        {"dep3": Def{System: []string{"lib3"}}},
    }, Recipe{
        "dep1": Def{System: []string{"lib1"}},
        "dep2": Def{System: []string{"lib2"}},
        "dep3": Def{System: []string{"lib3"}},
    }},
}

func TestCombiningRecipes(t *testing.T) {
    for _, tt := range combineTests {
        t.Run(tt.name, func(t *testing.T) {
            r := combine(tt.input...)

            if !reflect.DeepEqual(tt.expected, r) {
                t.Error("wrong")
            }
        })
    }
}
