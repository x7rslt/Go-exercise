package main

import (
	_ "github.com/projectdiscovery/fdmax/autofdmax"
	"github.com/projectdiscovery/gologger"
	"github.com/projectdiscovery/naabu/v2/pkg/runner"
)
/*
pkg/runner/options:125行，直接赋值Host值扫描
showBanner()
options.Host = "127.0.0.1"
*/
func main() {
	// Parse the command line flags and read config files
	options := runner.ParseOptions()


	naabuRunner, err := runner.NewRunner(options)

	if err != nil {
		gologger.Fatal().Msgf("Could not create runner: %s\n", err)
	}

	err = naabuRunner.RunEnumeration()
	if err != nil {
		gologger.Fatal().Msgf("Could not run enumeration: %s\n", err)
	}
}
