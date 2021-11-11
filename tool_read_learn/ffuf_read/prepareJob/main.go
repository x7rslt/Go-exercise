package main

import (
	"sync"
	"fmt"
	"os"

)

const (
	BANNER_HEADER = `
        /'___\  /'___\           /'___\       
       /\ \__/ /\ \__/  __  __  /\ \__/       
       \ \ ,__\\ \ ,__\/\ \/\ \ \ \ ,__\      
        \ \ \_/ \ \ \_/\ \ \_\ \ \ \ \_/      
         \ \_\   \ \_\  \ \____/  \ \_\       
          \/_/    \/_/   \/___/    \/_/       
`
	BANNER_SEP = "________________________________________________"
)

type OutputProvider interface {
	Banner()
	
}

func (s *Stdoutput) Banner() {
	fmt.Fprintf(os.Stderr, "%s\n       v\n%s\n\n", BANNER_HEADER, BANNER_SEP)
}


type Config struct {
	AutoCalibration        bool                      `json:"autocalibration"`
	AutoCalibrationStrings []string                  `json:"autocalibration_strings"`
	Colors                 bool                      `json:"colors"`
	CommandKeywords        []string                  `json:"-"`
	CommandLine            string                    `json:"cmdline"`
	ConfigFile             string                    `json:"configfile"`
	Data                   string                    `json:"postdata"`
	DirSearchCompat        bool                      `json:"dirsearch_compatibility"`
	Extensions             []string                  `json:"extensions"`
	FollowRedirects        bool                      `json:"follow_redirects"`
	Headers                map[string]string         `json:"headers"`
	IgnoreBody             bool                      `json:"ignorebody"`
	IgnoreWordlistComments bool                      `json:"ignore_wordlist_comments"`
	InputMode              string                    `json:"inputmode"`
	InputNum               int                       `json:"cmd_inputnum"`
	MaxTime                int                       `json:"maxtime"`
	MaxTimeJob             int                       `json:"maxtime_job"`
	Method                 string                    `json:"method"`
	OutputDirectory        string                    `json:"outputdirectory"`
	OutputFile             string                    `json:"outputfile"`
	OutputFormat           string                    `json:"outputformat"`
	OutputCreateEmptyFile  bool	                     `json:"OutputCreateEmptyFile"`
	ProgressFrequency      int                       `json:"-"`
	ProxyURL               string                    `json:"proxyurl"`
	Quiet                  bool                      `json:"quiet"`
	Rate                   int64                     `json:"rate"`
	Recursion              bool                      `json:"recursion"`
	RecursionDepth         int                       `json:"recursion_depth"`
	ReplayProxyURL         string                    `json:"replayproxyurl"`
	StopOn403              bool                      `json:"stop_403"`
	StopOnAll              bool                      `json:"stop_all"`
	StopOnErrors           bool                      `json:"stop_errors"`
	Threads                int                       `json:"threads"`
	Timeout                int                       `json:"timeout"`
	Url                    string                    `json:"url"`
	Verbose                bool                      `json:"verbose"`
}
func NewConfig() Config {
	var conf Config
	conf.AutoCalibrationStrings = make([]string, 0)
	conf.CommandKeywords = make([]string, 0)
	conf.Data = ""
	conf.DirSearchCompat = false
	conf.Extensions = make([]string, 0)
	conf.FollowRedirects = false
	conf.Headers = make(map[string]string)
	conf.IgnoreWordlistComments = false
	conf.InputMode = "clusterbomb"
	conf.InputNum = 0
	conf.MaxTime = 0
	conf.MaxTimeJob = 0
	conf.Method = "GET"
	conf.ProgressFrequency = 125
	conf.ProxyURL = ""
	conf.Quiet = false
	conf.Rate = 0
	conf.Recursion = false
	conf.RecursionDepth = 0
	conf.StopOn403 = false
	conf.StopOnAll = false
	conf.StopOnErrors = false
	conf.Timeout = 10
	conf.Url = ""
	conf.Verbose = false
	return conf
}

type Job struct {
	Config               *Config
	ErrorMutex           sync.Mutex
	Output               OutputProvider
	Counter              int
	ErrorCounter         int
	SpuriousErrorCounter int
	Total                int
	Running              bool
	RunningJob           bool
	Count403             int
	Count429             int
	Error                string
	queuepos             int
	currentDepth         int
}
type QueueJob struct {
	Url   string
	depth int
}

func NewJob(conf *Config)*Job{
	var j Job
	j.Config = conf
	j.Counter = 0
	j.ErrorCounter = 0
	j.SpuriousErrorCounter = 0
	j.Running = false
	j.RunningJob = false
	j.queuepos = 0
	j.currentDepth = 0
	return &j
}

type Multierror struct {
	errors []error
}
func (m *Multierror) ErrorOrNil() error {
	var errString string
	if len(m.errors) > 0 {
		errString += fmt.Sprintf("%d errors occured.\n", len(m.errors))
		for _, e := range m.errors {
			errString += fmt.Sprintf("\t* %s\n", e)
		}
		return fmt.Errorf("%s", errString)
	}
	return nil
}
//NewMultierror returns a new Multierror
func NewMultierror() Multierror {
	return Multierror{}
}


type Stdoutput struct {
	config  *Config
	Results []Result
}

type Result struct {
	Input            map[string][]byte `json:"input"`
	Position         int               `json:"position"`
	StatusCode       int64             `json:"status"`
	ContentLength    int64             `json:"length"`
	ContentWords     int64             `json:"words"`
	ContentLines     int64             `json:"lines"`
	RedirectLocation string            `json:"redirectlocation"`
	Url              string            `json:"url"`
	ResultFile       string            `json:"resultfile"`
	Host             string            `json:"host"`
	HTMLColor        string            `json:"-"`
}

func NewStdoutput(conf *Config) *Stdoutput {
	var outp Stdoutput
	outp.config = conf
	outp.Results = []Result{}
	return &outp
}
func NewOutputProviderByName(name string, conf *Config) OutputProvider {
	//We have only one outputprovider at the moment
	return NewStdoutput(conf)
}

func prepareJob(conf *Config) (*Job, error) {
	job := NewJob(conf)
	var errs Multierror
	// We only have stdout outputprovider right now
	job.Output = NewOutputProviderByName("stdout", conf)
	return job, errs.ErrorOrNil()
}

func main(){
	confl := NewConfig()
	fmt.Println(confl)

	jobs,err := prepareJob(&confl)
	fmt.Println(err)
	jobs.Output.Banner()
}