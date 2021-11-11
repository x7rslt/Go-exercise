package main

import (
	"fmt"
	"os"
)

const (
	TERMINAL_CLEAR_LINE = "\r\r"
	ANSI_CLEAR          = ""
	ANSI_RED            = ""
	ANSI_GREEN          = ""
	ANSI_BLUE           = ""
	ANSI_YELLOW         = ""
)

type Config struct {
	Colors bool
	Quiet  bool
}

type OutputProvider interface {
	Info(infostring string)
	//Result(resp string)
}

type Stdoutput struct {
	config *Config
	//Results []Result
}

func (s *Stdoutput) Info(infostring string) {
	if s.config.Quiet {
		fmt.Fprintf(os.Stderr, "%s", infostring)
	} else {
		if !s.config.Colors {
			fmt.Fprintf(os.Stderr, "%s[INFO] %s\n", TERMINAL_CLEAR_LINE, infostring)
		} else {
			fmt.Fprintf(os.Stderr, "%s[INFO] %s\n", TERMINAL_CLEAR_LINE, ANSI_BLUE, ANSI_CLEAR, infostring)
		}
	}
}

func NewStdoutput(conf *Config) *Stdoutput {
	var outp Stdoutput
	outp.config = conf
	//outp.Results = []Result{}
	return &outp
}

func NewOutputProviderByName(name string, conf *Config) OutputProvider {
	//We have only one outputprovider at the moment
	return NewStdoutput(conf)
}

type Job struct {
	Config *Config
	Output OutputProvider
}

func NewJob(conf *Config) *Job {
	var j Job
	j.Config = conf
	return &j
}
func main() {
	conf := &Config{true, true}
	//job := NewJob(conf)
	var job Job
	job.Output = NewOutputProviderByName("stdout", conf)
	job.Output.Info(fmt.Sprintf("Scanning :%s", "lehehe"))
}
