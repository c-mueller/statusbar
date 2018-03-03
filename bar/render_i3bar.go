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
)

type I3BarRenderer struct {
}

type i3BarHeader struct {
	Version     int  `json:"version"`
	ClickEvents bool `json:"click_events"`
}

func (r *I3BarRenderer) writeHeader() {
	header := i3BarHeader{
		Version:     1,
		ClickEvents: false,
	}
	data, _ := json.Marshal(header)
	fmt.Println(string(data))
}

func (r *I3BarRenderer) Render(sb *StatusBar) error {
	return nil
}

func (r *I3BarRenderer) Init(sb *StatusBar) error {
	r.writeHeader()
	return nil
}
