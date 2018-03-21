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
	"github.com/mitchellh/mapstructure"
	"github.com/shirou/gopsutil/net"
	"time"
)

var Builder = ComponentBuilder{}

func (c *ComponentBuilder) BuildComponent(identifier string, i interface{}) (statusbarlib.BarComponent, error) {
	cfg := &Configuration{}
	if i == nil {
		cfg = &DefaultConfiguration
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

	c.recentThroughputs = make(recentNetworkThroughputs, 0)
	c.totalThroughput = fromNetworkStats(c.getInterfaceStats())

	c.updateTicker = time.NewTicker(time.Duration(c.Config.UpdateInterval) * time.Millisecond)

	go c.collect()

	return nil
}

func (c *Component) Render() (string, error) {
	avg := c.recentThroughputs.computeAverage().toSpeedPerSecond(c.Config.UpdateInterval)

	outputString := avg.formatToString()

	if c.Config.ShowTotal {
		outputString += fmt.Sprintf(" (%s)", c.totalThroughput.formatToString())
	}

	return outputString, nil
}

func (c *Component) Stop() error {
	c.updateTicker.Stop()
	return nil
}

func (c *Component) GetIdentifier() string {
	return c.id
}

func (c *Component) collect() {
	for range c.updateTicker.C {
		// Collect the Current Network Stats
		current := fromNetworkStats(c.getInterfaceStats())

		// Calculate Difference
		diff := current.subtract(c.totalThroughput)

		// Append to Recent List
		c.appendThroughputStats(diff)

		// Store Latest value
		c.totalThroughput = current
	}
}

func (c *Component) appendThroughputStats(t *networkThroughput) {
	if len(c.recentThroughputs) < c.Config.RecentCount {
		c.recentThroughputs = append(recentNetworkThroughputs{*t}, c.recentThroughputs...)
	} else {
		c.recentThroughputs = append(recentNetworkThroughputs{*t}, c.recentThroughputs[:c.Config.RecentCount]...)
	}
}

func (c *Component) getInterfaceStats() net.IOCountersStat {
	stats, err := net.IOCounters(!c.Config.Global)

	if err != nil {
		log.Error("Collecting Network Data has failed:", err)
		panic(err)
	}

	if c.Config.Global {
		return stats[0]
	} else {
		for _, v := range stats {
			if v.Name == c.Config.InterfaceName {
				return v
			}
		}
		log.Errorf("Invalid interface %q!", c.Config.InterfaceName)
		panic("Network Interface is not existing anymore!")
	}
}
