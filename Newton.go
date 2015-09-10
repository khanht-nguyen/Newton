package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/robertkrimen/otto"
	"github.com/robertkrimen/otto/underscore"
)

var g_vm *otto.Otto

func main() {
	flag.Parse()

	if !*flag_underscore {
		underscore.Disable()
	}

	err := func() error {
		src, err := readSource(flag.Arg(0))
		if err != nil {
			return err
		}

		vm := otto.New()
    g_vm = vm

    vm.Set("setRoute", setRoute)
    vm.Set("startServer", startServer)

		_, err = vm.Run(src)
		return err
	}()
	if err != nil {
		switch err := err.(type) {
		case *otto.Error:
			fmt.Print(err.String())
		default:
			fmt.Println(err)
		}
		os.Exit(64)
	}
}
