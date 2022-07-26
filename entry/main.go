package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"plugin"
)

func main() {
	// Load the plugin
	pluginPath := "plugin.so"
	p, err := plugin.Open(pluginPath)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error loading plugin %q\n", pluginPath)
		panic(err)
	}

	// Verify that we can find the expected symbol
	expectedSymbol := "MustExist"
	sym, err := p.Lookup(expectedSymbol)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error connecting symbol %q\n", expectedSymbol)
		panic(err)
	}
	_ = sym

	// Verify value of the symbol is "bunny"
	expectedValue := "bunny"
	value := sym.(*string)
	if *value != expectedValue {
		fmt.Fprintf(os.Stderr, "error: expected symbol %q to have value %q, got %q\n", expectedSymbol, expectedValue, *value)
		panic(err)
	}

	// Echo input
	in, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		log.Fatal(err)
	}
	if string(in) == "fail" {
		fmt.Fprintf(os.Stderr, "%s", in)
		os.Exit(1)
	}
	fmt.Fprintf(os.Stdout, "%s", string(in))
}
