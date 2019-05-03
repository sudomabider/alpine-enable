package php

import (
    "errors"
    "os"
    "reflect"
    "testing"
)

var getMajorVersionTests = []struct {
    name string
    phpVersionEnv string
    version string
    error error
}{
    {
        name: "should error when no version is set",
        phpVersionEnv: "",
        version: "",
        error: errors.New("could not find PHP version"),
    },
    {
        name: "should error when version is invalid",
        phpVersionEnv: "a.b.c",
        version: "",
        error: errors.New("could not find PHP major version"),
    },
    {
        name: "should return major version",
        phpVersionEnv: "2.3.4",
        version: "2",
        error: nil,
    },
}

func TestGetMajorVersion(t *testing.T) {
    for _, tt := range getMajorVersionTests {
        t.Run(tt.name, func(t *testing.T) {
            _ = os.Setenv("PHP_VERSION", tt.phpVersionEnv)
            defer func() {_ = os.Unsetenv("PHP_VERSION")}()

            mv, err := GetMajorVersion()

            if !reflect.DeepEqual(tt.error, err) {
                t.Errorf("Expected error: %v\nGot error: %v", tt.error, err)
            }

            if mv != tt.version {
                t.Errorf("Expected version: %s\nGot version: %s", tt.version, mv)
            }
        })
    }
}
