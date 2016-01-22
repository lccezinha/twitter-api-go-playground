package main

import (
  "fmt"
  "errors"
)

func checkError(err error) {
    if err != nil {
      errors.New(fmt.Sprintf("Some shit happen: %s", err))
    }
}