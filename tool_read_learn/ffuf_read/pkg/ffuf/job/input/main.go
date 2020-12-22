package main

type InputProvider interface{
	AddProvider(InputProviderConfig) error
	Next() bool
	Position() int
	Reset()
	Value() map[string][]byte
	Total() int
}

type InternalInputProvider interface{
	Keyword() string
	Next() bool
	Position() int
	ResetPosition()
	Value() []byte
	Total() int
}
type InputProviderConfig struct {
	Name string `json:"name"`
	Keyword string `json:"keyword"`
	Value string `json:"value"`
}
type Config struct{
	InputProviders []InputProviderConfig

}

//作为InputProvider的实例
type MainInputProvider struct {
	Providers   [].InternalInputProvider
	Config      *Config
	position    int
	msbIterator int
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

func main(){
	job.Input, errs = input.NewInputProvider(conf)

}