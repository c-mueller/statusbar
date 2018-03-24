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

package netlib

import (
	"github.com/op/go-logging"
	"github.com/shirou/gopsutil/net"
	"time"
)

var log = logging.MustGetLogger("netlib_throughput_logger")

func (c *ThroughputLogger) Start() {
	c.RecentMeasurements = make(ThroughputList, 0)
	c.LastMeasurement = *FromNetworkStats(c.getInterfaceStats())

	c.updateTicker = time.NewTicker(c.UpdateInterval)

	go c.collectionLoop()
}

func (c *ThroughputLogger) Stop() {
	c.updateTicker.Stop()
}

func (c *ThroughputLogger) collectionLoop() {
	for range c.updateTicker.C {
		c.collect()
	}
}

func (c *ThroughputLogger) collect() {
	// Collect the Current Network Stats
	current := FromNetworkStats(c.getInterfaceStats())

	// Calculate Difference
	diff := current.Subtract(&c.LastMeasurement)

	// Append to Recent List
	c.appendMeasurement(diff)

	// Store Latest value
	c.LastMeasurement = *current
}

func (c *ThroughputLogger) appendMeasurement(t *NetworkThroughput) {
	if len(c.RecentMeasurements) < c.StoredValueCount {
		c.RecentMeasurements = append(ThroughputList{*t}, c.RecentMeasurements...)
	} else {
		c.RecentMeasurements = append(ThroughputList{*t}, c.RecentMeasurements[:c.StoredValueCount]...)
	}
}

func (c *ThroughputLogger) getInterfaceStats() net.IOCountersStat {
	stats, err := net.IOCounters(!c.Global)

	if err != nil {
		log.Error("Collecting Network Data has failed:", err)
		panic(err)
	}

	if c.Global {
		return stats[0]
	} else {
		for _, v := range stats {
			if v.Name == c.InterfaceName {
				return v
			}
		}
		log.Errorf("Invalid interface %q!", c.InterfaceName)
		panic("Network Interface is not existing anymore!")
	}
}
