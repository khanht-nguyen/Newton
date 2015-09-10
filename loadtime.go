package main

import (
  "flag"
  "os"
  "io/ioutil"
)

var flag_underscore *bool = flag.Bool("underscore", true, "Load underscore into the runtime environment")

func readSource(filename string) ([]byte, error) {
	if filename == "" || filename == "-" {
		return ioutil.ReadAll(os.Stdin)
	}
	return ioutil.ReadFile(filename)
}
