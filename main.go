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
	"io/ioutil"
	"os"

	"fmt"
	"github.com/c-mueller/statusbar/bar"
	"gopkg.in/alecthomas/kingpin.v2"
)

var (
	configPath   = kingpin.Flag("config", "The Path to the Configuration file").Default("config.yml").Short('c').ExistingFile()
	terminalMode = kingpin.Flag("terminal", "Render the Statusbar in Terminal Mode").Short('t').Bool()
	i3wmMode     = kingpin.Flag("i3", "Render the Statusbar in i3wm Mode").Short('i').Bool()
)

func main() {
	kingpin.Parse()

	if *terminalMode == *i3wmMode {
		fmt.Println("The Application can either Run in i3wm or Terminal Mode!")
		os.Exit(1)
	}

	f, err := os.Open(*configPath)
	if err != nil {
		panic(err)
	}
	cfgbytes, err := ioutil.ReadAll(f)
	if err != nil {
		panic(err)
	}
	sb, err := bar.BuildFromConfig(cfgbytes)
	if err != nil {
		panic(err)
	}

	sb.Init()

	if *terminalMode {
		sb.RenderTerminal()
	} else if *i3wmMode {
		sb.RenderI3()
	}
}
