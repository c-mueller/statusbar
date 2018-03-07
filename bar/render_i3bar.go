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
	"github.com/gin-gonic/gin/json"
	"time"
)

type I3BarRenderer struct {
}

type i3BarHeader struct {
	Version     int  `json:"version"`
	ClickEvents bool `json:"click_events"`
}

type i3BarBlock struct {
	Name     string `json:"name"`
	Instance string `json:"instance"`
	FullText string `json:"full_text"`
}

func (r *I3BarRenderer) writeHeader() {
	header := i3BarHeader{
		Version:     1,
		ClickEvents: false,
	}
	data, _ := json.Marshal(header)
	fmt.Println(string(data))
}

func (r *I3BarRenderer) Render(bar *StatusBar) error {
	for _, v := range bar.components {
		defer v.component.Stop()
	}

	//Send Array Opening Bracket
	fmt.Print("[[]")
	for {
		//Begin new Block
		fmt.Print(",[")

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

		block := i3BarBlock{
			Name:     "block",
			Instance: "block",
			FullText: resultString,
		}
		obj, _ := json.Marshal(block)
		fmt.Print(string(obj))

		//"Flush" output
		fmt.Println("]")

		//Wait for next refresh
		time.Sleep(bar.RefreshInterval)
	}
	return nil
}

func (r *I3BarRenderer) Init(sb *StatusBar) error {
	log.Debug("Initializing i3wm renderer")
	r.writeHeader()
	return nil
}
