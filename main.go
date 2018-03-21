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
	"github.com/op/go-logging"
	"gopkg.in/alecthomas/kingpin.v2"
	"strings"
	"gopkg.in/yaml.v2"
)

var (
	// Global Flags
	verbose = kingpin.Flag("verbose", "Print Verbose Information to Stderr").Short('v').Default("false").Bool()
	debug   = kingpin.Flag("debug", "Print debug Information to Stderr (Includes verbose mode)").Short('d').Default("false").Bool()

	// Statusbar Sub Command
	barSubCommand = kingpin.Command("run", "Run statusbar in default mode")
	configPath    = barSubCommand.Flag("config", "The Path to the Configuration file").Default("config.yml").Short('c').ExistingFile()
	terminalMode  = barSubCommand.Flag("terminal", "Render the Statusbar in Terminal Mode").Short('t').Bool()
	short         = barSubCommand.Flag("short", "Render Short version (Only works in Terminal mode)").Short('s').Default("false").Bool()
	i3wmMode      = barSubCommand.Flag("i3", "Render the Statusbar in i3wm Mode").Short('i').Bool()

	componentSubCommand          = kingpin.Command("components", "Show Information about the Components shipped with 'statusbar'")
	listComponentsSubCommand     = componentSubCommand.Command("list", "List the shipped Components")
	printDefaultConfigSubCommand = componentSubCommand.Command("default-config", "Print the Default YAML config of a component")
	commandNameArgument          = printDefaultConfigSubCommand.Flag("name", "The Name of the component").Short('n').Required().String()
	wrappedFlag                  = printDefaultConfigSubCommand.Flag("wrap", "Print Complete Component Configuration").Short('w').Default("false").Bool()
)

var format = logging.MustStringFormatter(
	`%{color}[%{time:15:04:05} - %{level}] - %{module}:%{color:reset} %{message}`,
)

var log = logging.MustGetLogger("sb_main")

func main() {
	switch kingpin.Parse() {
	case "run":
		runStatusBar()
	case "components list":
		listComponents()
	case "components default-config":
		printDefaultConfig()
	}

}

func printDefaultConfig() {
	components := bar.GetComponents()
	name := strings.ToLower(*commandNameArgument)
	for _, v := range components {
		if strings.ToLower(v.GetDescriptor()) == name {
			config := v.GetDefaultConfig()
			if config == nil {
				fmt.Printf("The component %q does not have a Configuration", *commandNameArgument)
				return
			}
			var yamlConfig []byte
			var err error
			if !*wrappedFlag {
				yamlConfig, err = yaml.Marshal(config)

			} else {
				wrappedConfig := bar.StatusBarComponentConfig{
					Identifier:           "my-identifier",
					Type:                 v.GetDescriptor(),
					CustomSeparator:      false,
					HideInShortMode:      false,
					Spec:                 config,
					CustomSeparatorValue: "|",
				}
				yamlConfig, err = yaml.Marshal(wrappedConfig)
			}
			if err != nil {
				panic(err)
			}
			fmt.Println(string(yamlConfig))
			return
		}
	}
	fmt.Printf("Component %q not found!\n", *commandNameArgument)
	os.Exit(1)
}

func listComponents() {
	components := bar.GetComponents()
	fmt.Printf("Found %d Components:\n", len(components))
	for _, v := range components {
		fmt.Printf(" - %s\n", v.GetDescriptor())
	}
}

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

func exitOnErr(err error, code int, format string, values ...interface{}) {
	if err != nil {
		log.Errorf(format, values...)
		log.Errorf("Error Message: %q", err.Error())
		os.Exit(code)
	}
}

func initializeLogger() {
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
}
