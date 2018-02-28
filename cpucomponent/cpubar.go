package cpucomponent

import (
	"fmt"
	"github.com/c-mueller/statusbar/braillegraph"
	"github.com/shirou/gopsutil/cpu"
	"runtime"
	"time"
)



type StatusBar struct {
	UpdateInterval  time.Duration
	cpuUpdateTicker *time.Ticker
	renderTicker    *time.Ticker
	cpuLoads        []float64
}

func (sb *StatusBar) Run() {
	sb.cpuLoads = make([]float64, runtime.NumCPU())
	sb.renderTicker = time.NewTicker(sb.UpdateInterval)
	sb.cpuUpdateTicker = time.NewTicker(sb.UpdateInterval)

	go sb.cpuUpdateHandler()

	sb.renderBar()
}

func (sb *StatusBar) buildCPULoadString() string {
	cpuLoads := "CPU: "
	for i := 0; i < len(sb.cpuLoads); i = i + 2 {
		rightLoad := sb.cpuLoads[i]
		leftLoad := 0.0
		if (i + 1) < len(sb.cpuLoads) {
			leftLoad = sb.cpuLoads[i+1]
		}
		chr, _ := braillegraph.NewGraphChar(leftLoad/100, rightLoad/100).ToBrailleChar().MapToBrailleChar()
		cpuLoads += fmt.Sprintf("%c", chr)
	}
	return cpuLoads
}

func (sb *StatusBar) renderBar() {
	for range sb.renderTicker.C {
		fmt.Printf("\r%s", sb.buildCPULoadString())
	}
}

func (sb *StatusBar) cpuUpdateHandler() {
	for range sb.cpuUpdateTicker.C {
		data, _ := cpu.Percent(sb.UpdateInterval, true)
		for k, v := range data {
			sb.cpuLoads[k] = v
		}
	}
}
