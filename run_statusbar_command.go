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
	"github.com/c-mueller/statusbar/bar"
	"gopkg.in/alecthomas/kingpin.v2"
	"io/ioutil"
	"os"
)

const (
	configFileOpenErrorFormat   = "Opening configuration at %q has failed with an error"
	statusBarInitErrorFormat    = "Building the 'statusbar' with the configuration at %q has failed with an error"
	componentInitErrorFormat    = "Component initialisation has failed!"
	renderingFailureErrorFormat = "Rendering failed"

	renderingEngineNotFoundMessageFormat = "Renderer %q not found!"
)

var (
	// Default Config path
	defaultConfigPath = os.Getenv("HOME") + "/.statusbar-config.yml"

	// Statusbar Sub Command
	runCommand = kingpin.Command("run",
		"Run 'statusbar' in default (single process) mode").Alias("r").Alias("show")
	configPath = runCommand.Flag("config",
		"The path to the configuration file").Default(defaultConfigPath).Short('c').ExistingFile()
	modeArg = runCommand.Arg("renderer",
		"The name of the renderer (use 'statusbar renderer' to list all rendering engines)").Default("terminal").String()
)

func runStatusBar() {
	initializeLogger()

	renderer := findRenderer(*modeArg)
	if renderer == nil {
		log.Errorf(renderingEngineNotFoundMessageFormat, *modeArg)
		os.Exit(1)
	}

	log.Debugf("Reading config from %q", *configPath)
	configFile, err := os.Open(*configPath)
	exitOnErr(err, 1, configFileOpenErrorFormat, *configPath)

	readConfigBytes, err := ioutil.ReadAll(configFile)
	exitOnErr(err, 1, configFileOpenErrorFormat, *configPath)

	statusBar, err := bar.BuildFromConfig(readConfigBytes)
	exitOnErr(err, 1, statusBarInitErrorFormat, *configPath)

	log.Debug("Initializing...")
	err = statusBar.Init()
	exitOnErr(err, 1, componentInitErrorFormat)

	log.Debug("Rendering...")
	err = statusBar.Render(renderer)
	exitOnErr(err, 1, renderingFailureErrorFormat)
}
