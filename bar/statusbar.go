package bar

import (
	"fmt"
	"time"
)

func NewStatusBar() *StatusBar {
	return &StatusBar{
		Components:      make([]BarComponent, 0),
		RefreshInterval: 500 * time.Millisecond,
	}
}

func (bar *StatusBar) AddComponent(component BarComponent) {
	bar.Components = append(bar.Components, component)
}

func (bar *StatusBar) Init() error {
	for _, v := range bar.Components {
		err := v.Init()
		if err != nil {
			return err
		}
	}
	return nil
}

func (bar *StatusBar) RenderTerminal() error {
	for _, v := range bar.Components {
		defer v.Stop()
	}
	for {
		resultString := ""
		for i, v := range bar.Components {
			r, err := v.Render()
			if err != nil {
				return err
			}
			resultString += r
			if i < len(bar.Components)-1 {
				resultString += " | "
			}
		}
		fmt.Println(resultString)
		time.Sleep(bar.RefreshInterval)
	}
}
