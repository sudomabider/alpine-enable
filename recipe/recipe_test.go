package recipe

import (
    "errors"
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
                t.Errorf("Expected recipe: %v\nGot: %v", tt.expected, r)
            }
        })
    }
}

var getPHPRecipeTests = []struct{
    name string
    majorVersion string
    error error
}{
    {"should get php7 recipe", "7", nil},
    {"should get php5 recipe", "5", nil},
    {"should error getting php6 recipe", "6", errors.New("no recipes are found for PHP version [6]")},
}

func TestGetPHPRecipe(t *testing.T) {
    for _, tt := range getPHPRecipeTests {
        t.Run(tt.name, func(t *testing.T) {
            _, err := GetPHPRecipe(tt.majorVersion)

            if !reflect.DeepEqual(tt.error, err) {
                t.Errorf("Expected errpr: %v\nGot: %v", tt.error, err)
            }
        })
    }
}

func TestRecipe_Modules(t *testing.T) {
    r := Recipe{
        "a": Def{},
        "b": Def{},
    }

    expected := []string{"a", "b"}
    got := r.Modules()

    if !reflect.DeepEqual(expected, got) {
        t.Errorf("Expected modules: %v\nGot: %v", expected, got)
    }
}
