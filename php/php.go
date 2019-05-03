package php

import (
    "errors"
    "os"
    "regexp"
)

var regex = regexp.MustCompile(`(\d+)\.\d+\.\d+`)

func GetMajorVersion() (string, error) {
    v := os.Getenv("PHP_VERSION")

    if v == "" {
        return "", errors.New("could not find PHP version")
    }

    match := regex.FindStringSubmatch(v)

    if match == nil {
        return "", errors.New("could not find PHP major version")
    }

    return match[1], nil
}
