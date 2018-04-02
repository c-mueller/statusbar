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

package main

import (
	"fmt"
	"github.com/c-mueller/statusbar/bar"
	"gopkg.in/alecthomas/kingpin.v2"
)

const rendererListFormat = "%-20s %-100s\n"

var (
	rendererCmd = kingpin.Command("renderer",
		"List all available rendering engines")
)

func listRenderer() {
	fmt.Printf(rendererListFormat, "NAME", "DESCRIPTION")
	for _, v := range bar.GetRenderer() {
		fmt.Printf(rendererListFormat, v.GetName(), v.GetDescription())
	}
}
