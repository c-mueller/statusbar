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
	"fmt"
	"github.com/c-mueller/statusbar/util/braillechart"
	"github.com/shirou/gopsutil/cpu"
	"runtime"
	"time"
)

type LoadMonitor struct {
	StoreAvg        bool
	LoadAvgCount    int
	UpdateInterval  int
	cpuUpdateTicker *time.Ticker
	cpuLoads        []float64
	updateTimestamp time.Time
	recentAverages  []float64
	currentAverage  float64
}

func (c *LoadMonitor) init() error {
	// Initialize Load Bar
	c.cpuLoads = make([]float64, runtime.NumCPU())
	c.recentAverages = make([]float64, c.LoadAvgCount)
	c.updateTimestamp = time.Now()

	// Build Tickers
	c.cpuUpdateTicker = time.NewTicker(c.getRefreshDuration())

	// Start Goroutines
	go c.cpuUpdateHandler()
	return nil
}

func (c *LoadMonitor) stop() error {
	c.cpuUpdateTicker.Stop()
	return nil
}

func (c *LoadMonitor) getCPUBrailleLoadBars() string {
	cpuLoads := ""
	for i := 0; i < len(c.cpuLoads); i = i + 2 {
		rightLoad := c.cpuLoads[i]
		leftLoad := 0.0
		if (i + 1) < len(c.cpuLoads) {
			leftLoad = c.cpuLoads[i+1]
		}

		bc := braillechart.NewChartChar(leftLoad/100, rightLoad/100)

		chr, _ := bc.ToBrailleChar().MapToBrailleChar()

		cpuLoads += fmt.Sprintf("%c", chr)
	}
	return cpuLoads
}

func (c *LoadMonitor) getBrailleAverageLoadChart() string {
	dataCount := len(c.recentAverages)

	output := ""

	for i := dataCount - 1; i >= 0; i = i - 2 {
		left := c.recentAverages[i]
		right := 0.0
		if (i - 1) >= 0 {
			right = c.recentAverages[i-1]
		}

		bc := braillechart.NewChartChar(left/100, right/100)

		chr, _ := bc.ToBrailleChar().MapToBrailleChar()

		output += fmt.Sprintf("%c", chr)
	}

	return output
}

func (c *LoadMonitor) cpuUpdateHandler() {
	for range c.cpuUpdateTicker.C {
		data, _ := cpu.Percent(0, true)
		avg := 0.0
		for k, v := range data {
			c.cpuLoads[k] = v
			avg += v
		}

		if c.StoreAvg {
			avg = avg / float64(len(c.cpuLoads))
			c.recentAverages = append([]float64{avg}, c.recentAverages[:c.LoadAvgCount-1]...)

			currentAvg := 0.0
			for _, v := range c.recentAverages {
				currentAvg += v
			}
			currentAvg = currentAvg / float64(len(c.recentAverages))
			c.currentAverage = currentAvg
		}

		c.updateTimestamp = time.Now()
	}
}

func (c *LoadMonitor) getRefreshDuration() time.Duration {
	return time.Duration(c.UpdateInterval) * time.Millisecond
}
