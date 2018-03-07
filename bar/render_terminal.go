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
	log.Debug("Initializing terminal renderer")
	return nil
}
