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
	"gopkg.in/yaml.v2"
	"os"
	"strings"
)

var (
	componentSubCommand      = kingpin.Command("components", "Show Information about the Components shipped with 'statusbar'")
	listComponentsSubCommand = componentSubCommand.Command("list", "List the shipped Components")

	printDefaultConfigSubCommand = componentSubCommand.Command("default-config", "Print the Default YAML config of a component")
	commandNameFlag              = printDefaultConfigSubCommand.Flag("name", "The Name of the component").Short('n').Required().String()
	wrappedFlag                  = printDefaultConfigSubCommand.Flag("wrap", "Print Complete Component Configuration").Short('w').Default("false").Bool()
)

func printDefaultConfig() {
	components := bar.GetComponents()
	name := strings.ToLower(*commandNameFlag)
	for _, v := range components {
		if strings.ToLower(v.GetDescriptor()) == name {
			config := v.GetDefaultConfig()
			if config == nil {
				fmt.Printf("The component %q does not have a Configuration", *commandNameFlag)
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
	fmt.Printf("Component %q not found!\n", *commandNameFlag)
	os.Exit(1)
}

func listComponents() {
	components := bar.GetComponents()
	fmt.Printf("Found %d Components:\n", len(components))
	for _, v := range components {
		fmt.Printf(" - %s\n", v.GetDescriptor())
	}
}
