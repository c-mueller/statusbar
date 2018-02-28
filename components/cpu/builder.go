package cpu

import (
	bar "github.com/c-mueller/statusbar/bar"
)

var Builder = CPUComponentBuilder{}

func (c *CPUComponentBuilder) BuildComponent(i *interface{}) (bar.BarComponent, error) {
	component := &CPULoadComponent{
		Config: &DefaultConfiguration,
	}

	return bar.BarComponent(component), nil
}

func (c *CPUComponentBuilder) GetDescriptor() string {
	return "CPULoadDisplay"
}
