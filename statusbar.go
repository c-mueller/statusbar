package main

import (
	"time"
	"github.com/c-mueller/statusbar/cpucomponent"
)

func main() {
	sb := cpucomponent.StatusBar{
		UpdateInterval: time.Second * 1,
	}
	sb.Run()
}
