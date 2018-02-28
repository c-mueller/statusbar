package main

import (
	"github.com/c-mueller/statusbar/bar"
	"github.com/c-mueller/statusbar/components/cpu"
	"github.com/c-mueller/statusbar/components/text"
)

func main() {
	sb := bar.NewStatusBar()

	cpuComponent, err := cpu.Builder.BuildComponent(nil)
	textComponent, _ := text.Builder.BuildComponent(nil)
	if err != nil {
		panic(err)
	}

	sb.AddComponent(textComponent)
	sb.AddComponent(cpuComponent)

	sb.Init()
	sb.RenderTerminal()
}
