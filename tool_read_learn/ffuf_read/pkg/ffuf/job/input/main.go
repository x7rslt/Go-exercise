package main

//RunnerProvider是一个请求执行者的接口
type RunnerProvider interface {
	Prepare(input map[string][]byte) (Request, error)
	Execute(req *Request) (Response, error)
}

//InputProvider接口处理RunnerProvider的输入数据
type InputProvider interface{
	AddProvider(InputProviderConfig) error
	Next() bool
	Position() int
	Reset()
	Value() map[string][]byte
	Total() int
}

//InternalInputProvider接口负责向InputProvider提供输入数据
type InternalInputProvider interface{
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
	Providers   [].InternalInputProvider
	Config      *Config
	position    int
	msbIterator int
}
func NewInputProvider(conf *ffuf.Config) (ffuf.InputProvider, ffuf.Multierror) {
	validmode := false
	errs := ffuf.NewMultierror()
	for _, mode := range []string{"clusterbomb", "pitchfork"} {
		if conf.InputMode == mode {
			validmode = true
		}
	}
	if !validmode {
		errs.Add(fmt.Errorf("Input mode (-mode) %s not recognized", conf.InputMode))
		return &MainInputProvider{}, errs
	}
	mainip := MainInputProvider{Config: conf, msbIterator: 0}
	// Initialize the correct inputprovider
	for _, v := range conf.InputProviders {
		err := mainip.AddProvider(v)
		if err != nil {
			errs.Add(err)
		}
	}
	return &mainip, errs
}

//Reset resets all the inputproviders and counters
func (i *MainInputProvider) Reset() {
	for _, p := range i.Providers {
		p.ResetPosition()
	}
	i.position = 0
	i.msbIterator = 0
}

type Job struct{
	Input InputProvider
}




func NewConfig(ctx context.Context,cancel context.CancelFunc)Config{
	var conf Config
	conf.InputProviders = make([]InputProviderConfig,0)
	return conf

}


type CommandInput struct {
	config  *Config
	count   int
}

func NewCommandInput(conf *ffuf.Config) (*CommandInput, error) {
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
	bb := []byte("Value for null")
	return stdout.Bytes()
}

//Total returns the size of wordlist
func (c *CommandInput) Total() int {
	return c.config.InputNum
}
type Config struct {
	InputNum int
	InputProviders []InputProviderConfig
}


func main(){
	ipc := InputProviderConfig{"ipc for null"}
	conf := &Config{123, []InputProviderConfig{ipc}}
	j := Job{

	}
	j.Input.Reset()
	fmt.Println(j)

}