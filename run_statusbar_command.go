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
	"io/ioutil"
	"os"
)

var (
	// Statusbar Sub Command
	barSubCommand = kingpin.Command("run", "Run statusbar in default mode")
	configPath    = barSubCommand.Flag("config", "The Path to the Configuration file").Default("config.yml").Short('c').ExistingFile()
	terminalMode  = barSubCommand.Flag("terminal", "Render the Statusbar in Terminal Mode").Short('t').Bool()
	short         = barSubCommand.Flag("short", "Render Short version (Only works in Terminal mode)").Short('s').Default("false").Bool()
	i3wmMode      = barSubCommand.Flag("i3", "Render the Statusbar in i3wm Mode").Short('i').Bool()
)

func runStatusBar() {
	initializeLogger()
	if *terminalMode == *i3wmMode {
		fmt.Println("The Application can either Run in i3wm (-i) or Terminal Mode (-t)!")
		os.Exit(1)
	}
	log.Debugf("Reading Config from %q", *configPath)
	f, err := os.Open(*configPath)
	exitOnErr(err, 1, "Opening Configuration at %q has failed with an error", *configPath)

	cfgbytes, err := ioutil.ReadAll(f)
	exitOnErr(err, 1, "Opening Configuration at %q has failed with an error", *configPath)

	sb, err := bar.BuildFromConfig(cfgbytes)
	exitOnErr(err, 1, "Building the Statusbar with the configuration at %q has failed with an error", *configPath)

	log.Debug("Initializing...")
	err = sb.Init()
	exitOnErr(err, 1, "Component Initialisation has failed!")

	log.Debug("Rendering...")
	if *terminalMode {
		sb.RenderTerminal(*short)
	} else if *i3wmMode {
		sb.RenderI3()
	}
}
