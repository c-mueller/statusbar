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

const (
	terminalRendererName             = "terminal"
	terminalRendererDescription      = "Render the statusbar within the terminal, in long mode (all components)."
	terminalRendererNameShort        = "terminal-short"
	terminalRendererDescriptionShort = "Render the statusbar within the terminal, in short mode (only enabled components)."
)

type TerminalRenderer struct {
	ShortMode   bool
	Name        string
	Description string
}

func NewTerminalRenderer(short bool) *TerminalRenderer {
	name := terminalRendererName
	description := terminalRendererDescription
	if short {
		name = terminalRendererNameShort
		description = terminalRendererDescriptionShort
	}
	return &TerminalRenderer{
		ShortMode:   short,
		Name:        name,
		Description: description,
	}
}

func (t *TerminalRenderer) Render(bar *StatusBar) error {
	for _, v := range bar.components {
		defer v.component.Stop()
	}
	oldLength := 0
	for {
		writeBlanksOnLine(oldLength)

		// Print new Output
		longString := ""
		shortString := ""
		for i, v := range bar.components {
			l, s, err := renderComponent(i, bar, v)
			if err != nil {
				return err
			}
			longString += l
			shortString += s
		}

		r := longString
		if t.ShortMode {
			r = shortString
		}

		fmt.Printf("\r%s", r)
		oldLength = len(r)
		time.Sleep(bar.RefreshInterval)
	}
	return nil
}

func (t *TerminalRenderer) Init(sb *StatusBar) error {
	log.Debug("Initializing terminal renderer")
	return nil
}

func (t *TerminalRenderer) GetName() string {
	return t.Name
}

func (t *TerminalRenderer) GetDescription() string {
	return t.Description
}
