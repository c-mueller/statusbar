package command

import (
	"github.com/c-mueller/statusbar/bar/statusbarlib"
	"github.com/mitchellh/mapstructure"
	"time"
)

var Builder = CommandBuilder{}
var DefaultConfig = Config{
	Command:           "echo \"Hello World!\"",
	ExecutionInterval: 6000,
}

type CommandBuilder struct {
}

type Component struct {
	id            string
	commandRunner *CommandRunner
	Config        Config
}

type Config struct {
	Command           string `yaml:"command" mapstructure:"command"`
	ExecutionInterval int    `yaml:"execution_interval" mapstructure:"execution_interval"`
}

func (c *CommandBuilder) BuildComponent(identifier string, i interface{}) (statusbarlib.BarComponent, error) {
	cfg := &DefaultConfig
	if i != nil {
		var ic *Config
		err := mapstructure.Decode(i, &ic)
		if err != nil {
			return nil, err
		}
		cfg = ic
	}
	component := &Component{
		Config: *cfg,
		id:     identifier,
	}

	return statusbarlib.BarComponent(component), nil
}

func (c *CommandBuilder) GetDescriptor() string {
	return "Command"
}

func (c *CommandBuilder) GetDefaultConfig() interface{} {
	return DefaultConfig
}

func (c *Component) Init() error {
	c.commandRunner = NewCommandRunner(c.Config.Command, time.Duration(c.Config.ExecutionInterval)*time.Millisecond)
	c.commandRunner.Start()
	return nil
}

func (c *Component) Render() (*statusbarlib.RenderingOutput, error) {
	value := c.commandRunner.lastValue
	return &statusbarlib.RenderingOutput{LongText: value, ShortText: value}, nil
}

func (c *Component) Stop() error {
	c.commandRunner.Stop()
	return nil
}

func (c *Component) GetIdentifier() string {
	return c.id
}
