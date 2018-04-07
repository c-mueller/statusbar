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

package terminal

import (
	"fmt"
	"github.com/c-mueller/statusbar/bar/statusbarlib"
	"github.com/op/go-logging"
	"time"
)

var log = logging.MustGetLogger("terminal_renderer")

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

func (t *TerminalRenderer) Render(bar statusbarlib.StatusBar) error {
	for _, v := range bar.GetComponents() {
		defer v.Component.Stop()
	}
	oldLength := 0
	for {
		writeBlanksOnLine(oldLength)

		components := bar.GetComponents()

		// Print new Output
		l, s, err := components.RenderComponentsAsString()
		if err != nil {
			return err
		}

		r := l
		if t.ShortMode {
			r = s
		}

		fmt.Printf("\r%s", r)
		oldLength = len(r)
		time.Sleep(bar.GetRefreshInterval())
	}
	return nil
}

func (t *TerminalRenderer) Init(sb statusbarlib.StatusBar) error {
	log.Debug("Initializing terminal renderer")
	return nil
}

func (t *TerminalRenderer) GetName() string {
	return t.Name
}

func (t *TerminalRenderer) GetDescription() string {
	return t.Description
}

func writeBlanksOnLine(count int) {
	fmt.Printf("\r")
	for i := 0; i < count; i++ {
		fmt.Printf(" ")
	}
}
