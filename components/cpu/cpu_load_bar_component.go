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

package cpu

import (
	"github.com/c-mueller/statusbar/bar/statusbarlib"
	"github.com/mitchellh/mapstructure"
)

var LoadBarBuilder = LoadBarComponentBuilder{}

func (c *LoadBarComponentBuilder) BuildComponent(identifier string, i interface{}) (statusbarlib.BarComponent, error) {
	cfg := &LoadBarConfiguration{}
	if i == nil {
		cfg = &DefaultLoadBarConfig
	} else {
		var ic *LoadBarConfiguration
		err := mapstructure.Decode(i, &ic)
		if err != nil {
			return nil, err
		}
		cfg = ic
	}
	component := &LoadBarComponent{
		Config: cfg,
		id:     identifier,
	}

	return statusbarlib.BarComponent(component), nil
}

func (c *LoadBarComponentBuilder) GetDescriptor() string {
	return "CPULoadBar"
}

func (c *LoadBarComponentBuilder) GetDefaultConfig() interface{} {
	return DefaultLoadBarConfig
}

func (c *LoadBarComponent) Init() error {
	c.updateHandler = &LoadMonitor{
		UpdateInterval: c.Config.UpdateInterval,
		LoadAvgCount:   0,
		StoreAvg:       false,
	}
	return c.updateHandler.init()
}

func (c *LoadBarComponent) Render() (*statusbarlib.RenderingOutput, error) {
	value := c.updateHandler.getCPUBrailleLoadBars()
	return &statusbarlib.RenderingOutput{LongText: value, ShortText: value}, nil
}

func (c *LoadBarComponent) Stop() error {
	c.updateHandler.stop()
	return nil
}

func (c *LoadBarComponent) GetIdentifier() string {
	return c.id
}
