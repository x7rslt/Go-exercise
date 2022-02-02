package main

import (
	"github.com/projectdiscovery/gologger"
	"github.com/projectdiscovery/naabu/tree/master/v2"
)

func main() {
	// Parse the command line flags and read config files
	options := (naabu).ParseOptions()

	naabuRunner, err := runner.NewRunner(options)
	if err != nil {
		gologger.Fatal().Msgf("Could not create runner: %s\n", err)
	}

	err = naabuRunner.RunEnumeration()
	if err != nil {
		gologger.Fatal().Msgf("Could not run enumeration: %s\n", err)
	}
}