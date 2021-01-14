package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/BurntSushi/toml"
)

type ConfigOptions struct {
	HTTP HTTPOptions
}
type HTTPOptions struct {
	Method string
}

func NewConfigOptions() *ConfigOptions {
	c := &ConfigOptions{}
	c.HTTP.Method = "Get"
	return c
}

func ReadDefaultConfig() (*ConfigOptions, error) {
	userhome, err := os.UserHomeDir()
	if err != nil {
		return NewConfigOptions(), err
	}
	defaultconf := filepath.Join(userhome, ".ffufrc")
	return ReadConfig(defaultconf)

}
func ReadConfig(configFile string) (*ConfigOptions, error) {
	conf := NewConfigOptions()
	configData, err := ioutil.ReadFile(configFile)
	if err == nil {
		err = toml.Unmarshal(configData, conf)
	}
	return conf, err
}

func ParseFlags(opts *ConfigOptions) *ConfigOptions {
	flag.StringVar(&opts.HTTP.Method, "X", opts.HTTP.Method, "HTTP method to use")
	flag.Parse()
	return opts
}
func main() {
	//var err, optserr error

	var opts *ConfigOptions

	opts, optserr := ReadDefaultConfig()
	opts = ParseFlags(opts)
	fmt.Println(opts, optserr)
}
