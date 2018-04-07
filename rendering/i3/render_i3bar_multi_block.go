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

package i3

import (
	"encoding/json"
	"fmt"
	"github.com/c-mueller/statusbar/bar/statusbarlib"
	"time"
)

const (
	i3barMultiBlockRenderName        = "i3mb"
	i3barMultiBlockRenderDescription = "Render the statusbar in i3bar mode (multi block)"
)

type I3MultiBlockRenderer struct {
}

func (r *I3MultiBlockRenderer) Render(bar statusbarlib.StatusBar) error {
	for _, v := range bar.GetComponents() {
		defer v.Component.Stop()
	}

	//Send Array Opening Bracket
	fmt.Print("[[]")
	for {
		//Begin new Block
		fmt.Print(",[")

		components := bar.GetComponents()

		for idx, v := range components {

			d, err := v.Component.Render()
			if err != nil {
				return err
			}

			block := i3BarBlock{
				Name:      v.Identifier,
				Instance:  v.Identifier,
				FullText:  d.LongText,
				ShortText: d.ShortText,
			}
			obj, _ := json.Marshal(block)
			fmt.Print(string(obj))
			if idx != len(components)-1 {
				fmt.Print(", ")
			}
		}

		//"Flush" output
		fmt.Println("]")

		//Wait for next refresh
		time.Sleep(bar.GetRefreshInterval())
	}
	return nil
}

func (r *I3MultiBlockRenderer) Init(sb statusbarlib.StatusBar) error {
	log.Debug("Initializing i3wm renderer")
	writeI3BarHeader()
	return nil
}

func (r *I3MultiBlockRenderer) GetName() string {
	return i3barMultiBlockRenderName
}

func (r *I3MultiBlockRenderer) GetDescription() string {
	return i3barMultiBlockRenderDescription
}
