package cpu

import (
	"fmt"
	"github.com/c-mueller/statusbar/braillegraph"
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
	c.cpuUpdateTicker = time.NewTicker(c.Config.UpdateInterval)

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

func (c *CPULoadComponent) composeString() string {
	cpuLoads := "CPU: "
	for i := 0; i < len(c.cpuLoads); i = i + 2 {
		rightLoad := c.cpuLoads[i]
		leftLoad := 0.0
		if (i + 1) < len(c.cpuLoads) {
			leftLoad = c.cpuLoads[i+1]
		}
		chr, _ := braillegraph.NewGraphChar(leftLoad/100, rightLoad/100).ToBrailleChar().MapToBrailleChar()
		cpuLoads += fmt.Sprintf("%c", chr)
	}
	if c.Config.ShowAverageLoad {
		cpuLoads += fmt.Sprintf(" | AVG: %03d%%", int(c.currentAverage))
	}
	return cpuLoads
}

func (c *CPULoadComponent) cpuUpdateHandler() {
	for range c.cpuUpdateTicker.C {
		data, _ := cpu.Percent(c.Config.UpdateInterval, true)
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
