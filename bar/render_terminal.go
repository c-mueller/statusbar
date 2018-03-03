package bar

import (
	"fmt"
	"time"
)

type TerminalRenderer struct {
}

func (t *TerminalRenderer) Render(bar *StatusBar) error {
	for _, v := range bar.components {
		defer v.component.Stop()
	}
	oldlen := 0
	for {
		// Remove old output
		fmt.Printf("\r")
		for i := 0; i < oldlen; i++ {
			fmt.Printf(" ")
		}
		// Print new Output
		resultString := ""
		for i, v := range bar.components {
			r, err := v.component.Render()
			if err != nil {
				return err
			}
			resultString += r
			if i < len(bar.components)-1 {
				if v.config.CustomSeparator {
					resultString += v.config.CustomSeparatorValue
				} else {
					resultString += " | "
				}
			}
		}
		fmt.Printf("\r%s", resultString)
		oldlen = len(resultString)
		time.Sleep(bar.RefreshInterval)
	}
	return nil
}

func (t *TerminalRenderer) Init(sb *StatusBar) error {

	return nil
}
