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
	"github.com/logrusorgru/aurora"
	"github.com/shirou/gopsutil/cpu"
	"runtime"
	"time"
)

func (c *CPULoadComponent) Init() error {
	//Initialize Load Bar
	c.cpuLoads = make([]float64, runtime.NumCPU())
	c.recentAverages = make([]float64, c.Config.LoadAverageCount)
	c.updateTimestamp = time.Now()
	c.currentValue = c.composeString()

	//Build Tickers
	c.cpuUpdateTicker = time.NewTicker(time.Duration(c.Config.UpdateInterval) * time.Second)

	//Start Goroutines
	go c.cpuUpdateHandler()

	return nil
}

func (c *CPULoadComponent) Render() (string, error) {
	return c.currentValue, nil
}

func (c *CPULoadComponent) IsLatest(t time.Time) bool {
	return t == c.updateTimestamp
}

func (c *CPULoadComponent) Stop() error {
	c.cpuUpdateTicker.Stop()
	return nil
}

func (c *CPULoadComponent) GetIdentifier() string {
	return c.id
}

func (c *CPULoadComponent) composeString() string {
	cpuLoads := "CPU: "
	for i := 0; i < len(c.cpuLoads); i = i + 2 {
		rightLoad := c.cpuLoads[i]
		leftLoad := 0.0
		if (i + 1) < len(c.cpuLoads) {
			leftLoad = c.cpuLoads[i+1]
		}

		bc := braillechart.NewChartChar(leftLoad/100, rightLoad/100)

		chr, _ := bc.ToBrailleChar().MapToBrailleChar()

		format := aurora.Bold("%c")
		if bc.SumValues() >= 7 {
			format = aurora.Red(format)
		} else {
			format = aurora.Green(format)
		}

		cpuLoads += aurora.Sprintf(format, chr)
	}
	if c.Config.ShowAverageLoad {
		formatString := "%03d%%"
		currentColor := aurora.Bold(formatString)
		if c.currentAverage <= 75 {
			currentColor = aurora.Green(currentColor)
		} else if c.currentAverage > 75 {
			currentColor = aurora.Red(currentColor)
		}

		cpuLoads += fmt.Sprintf(" | AVG: %s", aurora.Sprintf(currentColor, int(c.currentAverage)))
	}
	return cpuLoads
}

func (c *CPULoadComponent) cpuUpdateHandler() {
	for range c.cpuUpdateTicker.C {
		data, _ := cpu.Percent(time.Duration(c.Config.UpdateInterval)*time.Second, true)
		avg := 0.0
		for k, v := range data {
			c.cpuLoads[k] = v
			avg += v
		}

		if c.Config.ShowAverageLoad {
			avg = avg / float64(len(c.cpuLoads))
			c.recentAverages = append([]float64{avg}, c.recentAverages[:c.Config.LoadAverageCount-1]...)

			currentAvg := 0.0
			for _, v := range c.recentAverages {
				currentAvg += v
			}
			currentAvg = currentAvg / float64(len(c.recentAverages))
			c.currentAverage = currentAvg
		}

		c.updateTimestamp = time.Now()
		c.currentValue = c.composeString()
	}
}
