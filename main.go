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
	"github.com/op/go-logging"
)

var (
	configPath   = kingpin.Flag("config", "The Path to the Configuration file").Default("config.yml").Short('c').ExistingFile()
	terminalMode = kingpin.Flag("terminal", "Render the Statusbar in Terminal Mode").Short('t').Bool()
	i3wmMode     = kingpin.Flag("i3", "Render the Statusbar in i3wm Mode").Short('i').Bool()

	verbose      = kingpin.Flag("verbose", "Print Verbose Information to Stderr").Short('v').Default("false").Bool()
	debug        = kingpin.Flag("debug", "Print debug Information to Stderr (Includes verbose mode)").Short('d').Default("false").Bool()
)

var format = logging.MustStringFormatter(
	`%{color}[%{time:15:04:05} - %{level}] - %{module}:%{color:reset} %{message}`,
)

var log = logging.MustGetLogger("sb_main")

func main() {
	kingpin.Parse()

	backend := logging.NewLogBackend(os.Stderr, "", 0)
	backendFormatter := logging.NewBackendFormatter(backend, format)
	leveledBackend := logging.AddModuleLevel(backendFormatter)
	if *debug {
		leveledBackend.SetLevel(logging.DEBUG, "")
	} else if *verbose {
		leveledBackend.SetLevel(logging.INFO, "")
	} else {
		leveledBackend.SetLevel(logging.ERROR, "")
	}

	logging.SetBackend(leveledBackend)

	log.Debug("Parsed Command Line arguments")

	if *terminalMode == *i3wmMode {
		fmt.Println("The Application can either Run in i3wm or Terminal Mode!")
		os.Exit(1)
	}

	log.Debugf("Reading Config from %q", *configPath)
	f, err := os.Open(*configPath)
	if err != nil {
		log.Error("Opening the config file", *configPath, "has failed. Error Message:", err.Error())
		os.Exit(1)
	}
	cfgbytes, err := ioutil.ReadAll(f)
	if err != nil {
		log.Error("Reading the config file", *configPath, "has failed. Error Message:", err.Error())
		os.Exit(1)
	}
	sb, err := bar.BuildFromConfig(cfgbytes)
	if err != nil {
		log.Error("Building the Statusbar from the config file", *configPath, "has failed. Error Message:", err.Error())
		os.Exit(1)
	}

	log.Debug("Initializing...")
	sb.Init()

	log.Debug("Rendering...")
	if *terminalMode {
		sb.RenderTerminal()
	} else if *i3wmMode {
		sb.RenderI3()
	}
}
