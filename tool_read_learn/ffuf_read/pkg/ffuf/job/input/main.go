package main

import (
	"bytes"
	"context"
	"fmt"
	"os"
	"os/exec"
	"strconv"
)

const (
	SHELL_CMD = "cmd.exe"
	SHELL_ARG = "/C"
)

type Request struct {
	Method   string
	Host     string
	Url      string
	Headers  map[string]string
	Data     []byte
	Input    map[string][]byte
	Position int
	Raw      string
}

// Response struct holds the meaningful data returned from request and is meant for passing to filters
type Response struct {
	StatusCode    int64
	Headers       map[string][]string
	Data          []byte
	ContentLength int64
	ContentWords  int64
	ContentLines  int64
	Cancelled     bool
	Request       *Request
	Raw           string
	ResultFile    string
}

//RunnerProvider是一个请求执行者的接口
type RunnerProvider interface {
	Prepare(input map[string][]byte) (Request, error)
	Execute(req *Request) (Response, error)
}

//InputProvider接口处理RunnerProvider的输入数据
type InputProvider interface {
	AddProvider(InputProviderConfig) error
	Next() bool
	Position() int
	Reset()
	Value() map[string][]byte
	Total() int
}

//InternalInputProvider接口负责向InputProvider提供输入数据
type InternalInputProvider interface {
	Keyword() string
	Next() bool
	Position() int
	ResetPosition()
	Value() []byte
	Total() int
}

type InputProviderConfig struct {
	Value string
}

//作为InputProvider的实例
type MainInputProvider struct {
	Providers   []InternalInputProvider
	Config      *Config
	position    int
	msbIterator int
}
type Multierror struct {
	errors []error
}

//NewMultierror returns a new Multierror
func NewMultierror() Multierror {
	return Multierror{}
}

func NewInputProvider(conf *Config) InputProvider {
	mainip := MainInputProvider{Config: conf, msbIterator: 0}
	return &mainip
}

func (i *MainInputProvider) AddProvider(provider InputProviderConfig) error {
	return nil
}
func (i *MainInputProvider) Position() int {
	return i.position
}

//Reset resets all the inputproviders and counters
func (i *MainInputProvider) Reset() {
	for _, p := range i.Providers {
		p.ResetPosition()
	}
	i.position = 0
	i.msbIterator = 0
}

//Next will increment the cursor position, and return a boolean telling if there's inputs left
func (i *MainInputProvider) Next() bool {
	if i.position >= i.Total() {
		return false
	}
	i.position++
	return true
}

//Value returns a map of inputs for keywords
func (i *MainInputProvider) Value() map[string][]byte {
	retval := make(map[string][]byte)
	retval["Value"] = []byte("No input value")
	return retval
}
func (i *MainInputProvider) Total() int {
	count := 0
	if i.Config.InputMode == "pitchfork" {
		for _, p := range i.Providers {
			if p.Total() > count {
				count = p.Total()
			}
		}
	}
	if i.Config.InputMode == "clusterbomb" {
		count = 1
		for _, p := range i.Providers {
			count = count * p.Total()
		}
	}
	return count
}

type Job struct {
	Input InputProvider
}

func NewConfig(ctx context.Context, cancel context.CancelFunc) Config {
	var conf Config
	conf.InputProviders = make([]InputProviderConfig, 0)
	return conf

}

type CommandInput struct {
	config  *Config
	count   int
	command string
}

func NewCommandInput(conf *Config) (*CommandInput, error) {
	var cmd CommandInput
	cmd.config = conf
	cmd.count = 0
	return &cmd, nil
}

//Keyword returns the keyword assigned to this InternalInputProvider
func (c *CommandInput) Keyword() string {
	return "keyword for null"
}

//Position will return the current position in the input list
func (c *CommandInput) Position() int {
	return c.count
}

//ResetPosition will reset the current position of the InternalInputProvider
func (c *CommandInput) ResetPosition() {
	c.count = 0
}

//IncrementPosition increments the current position in the inputprovider
func (c *CommandInput) IncrementPosition() {
	c.count += 1
}

//Next will increment the cursor position, and return a boolean telling if there's iterations left
func (c *CommandInput) Next() bool {
	return c.count < c.config.InputNum
}

//Value returns the input from command stdoutput
func (c *CommandInput) Value() []byte {
	var stdout bytes.Buffer
	os.Setenv("FFUF_NUM", strconv.Itoa(c.count))
	cmd := exec.Command(SHELL_CMD, SHELL_ARG, c.command)
	cmd.Stdout = &stdout
	err := cmd.Run()
	if err != nil {
		return []byte("")
	}
	return stdout.Bytes()
}

//Total returns the size of wordlist
func (c *CommandInput) Total() int {
	return c.config.InputNum
}

type Config struct {
	InputNum       int
	InputProviders []InputProviderConfig
	InputMode      string
}

func main() {
	ipc := InputProviderConfig{"ipc for null"}
	conf := &Config{123, []InputProviderConfig{ipc}, "InputMode null"}
	newin := NewInputProvider(conf)
	j := Job{newin}
	j.Input.Reset()
	fmt.Println(j)

}
