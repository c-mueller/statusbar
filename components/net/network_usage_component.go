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

package net

import (
	"errors"
	"fmt"
	"github.com/c-mueller/statusbar/bar/statusbarlib"
	"github.com/c-mueller/statusbar/components/net/netlib"
	"github.com/mitchellh/mapstructure"
	"github.com/shirou/gopsutil/net"
	"time"
)

var Builder = ComponentBuilder{}

func (c *ComponentBuilder) BuildComponent(identifier string, i interface{}) (statusbarlib.BarComponent, error) {
	cfg := &Configuration{}
	if i == nil {
		cfg = &DefaultConfig
	} else {
		var ic *Configuration
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

func (c *ComponentBuilder) GetDescriptor() string {
	return "Network"
}

func (b *ComponentBuilder) GetDefaultConfig() interface{} {
	return DefaultConfig
}

func (c *Component) Init() error {
	if !c.Config.Global {
		interfaces, err := net.Interfaces()
		if err != nil {
			return err
		}
		interfaceFound := false
		for _, v := range interfaces {
			if v.Name == c.Config.InterfaceName {
				interfaceFound = true
				break
			}
		}
		if !interfaceFound {
			return errors.New(fmt.Sprintf("network_component: Interface %q not found", c.Config.InterfaceName))
		}
	}

	c.collector = &netlib.ThroughputLogger{
		UpdateInterval:   time.Duration(c.Config.UpdateInterval) * time.Millisecond,
		Global:           c.Config.Global,
		InterfaceName:    c.Config.InterfaceName,
		StoredValueCount: c.Config.RecentCount,
	}

	c.collector.Start()

	return nil
}

func (c *Component) Render() (*statusbarlib.RenderingOutput, error) {
	avg := c.collector.RecentMeasurements.ComputeAverage().ToSpeedPerSecond(c.Config.UpdateInterval)

	outputString := avg.FormatToString()

	if c.Config.ShowTotal {
		outputString += fmt.Sprintf(" (%s)", c.collector.LastMeasurement.FormatToString())
	}

	return &statusbarlib.RenderingOutput{LongText: outputString, ShortText: outputString}, nil
}

func (c *Component) Stop() error {
	c.collector.Stop()
	return nil
}

func (c *Component) GetIdentifier() string {
	return c.id
}
