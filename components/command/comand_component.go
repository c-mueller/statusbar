// statusbar - (https://github.com/c-mueller/statusbar)
// Copyright (c) 2018 Christian Müller <cmueller.dev@gmail.com>.
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with this program.  If not, see <http://www.gnu.org/licenses/>.

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

func (c *CommandBuilder) BuildComponent(identifier string, i interface{}, builders []statusbarlib.ComponentBuilder) (statusbarlib.BarComponent, error) {
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
