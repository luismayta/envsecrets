package main

import (
	"github.com/luismayta/envsecrets/v1/cmd/envsecrets"
	"github.com/luismayta/envsecrets/v1/internal/errors"
)

func main() {
	err := envsecrets.Execute()
	if err != nil {
		errors.Must(err, errors.ErrorUnknown, "Error in Execute")
	}
}
