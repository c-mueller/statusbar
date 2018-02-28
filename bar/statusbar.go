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
	"errors"
	"fmt"
	"github.com/c-mueller/statusbar/bar/bi"
	"gopkg.in/yaml.v2"
	"time"
)

func BuildFromConfig(config []byte) (*StatusBar, error) {
	var cfg *StatusBarConfig
	yaml.Unmarshal(config, &cfg)

	sb := newStatusBar()

	for _, v := range cfg.Components {
		componentFound := false
		for _, builder := range builders {
			if v.Type == builder.GetDescriptor() {
				componentFound = true
				component, err := builder.BuildComponent(v.Identifier, v.Spec)
				if err != nil {
					return nil, err
				}
				sb.addComponent(component, v)
			}
		}
		if !componentFound {
			return nil, errors.New(fmt.Sprintf("No Component of type %q found", v.Type))
		}
	}

	return sb, nil
}

func newStatusBar() *StatusBar {
	return &StatusBar{
		components:      make([]*componentInstance, 0),
		RefreshInterval: 500 * time.Millisecond,
	}
}

func (bar *StatusBar) addComponent(component bi.BarComponent, config StatusBarComponentConfig) error {
	for _, v := range bar.components {
		if v.GetIdentifier() == config.Identifier {
			return errors.New(fmt.Sprintf("Invalid identifier name %q is already in use", component.GetIdentifier()))
		}
	}

	instance := componentInstance{
		component: component,
		id:        config.Identifier,
		config:    &config,
	}

	bar.components = append(bar.components, &instance)

	return nil
}

func (bar *StatusBar) Init() error {
	for _, v := range bar.components {
		err := v.component.Init()
		if err != nil {
			return err
		}
	}
	return nil
}

func (bar *StatusBar) RenderTerminal() error {
	for _, v := range bar.components {
		defer v.component.Stop()
	}
	oldlen := 0
	for {
		// Remove old output
		fmt.Printf("\r")
		for i := 0; i < oldlen; i++ {
			fmt.Printf(" ")
		}
		// Print new Output
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
		fmt.Printf("\r%s", resultString)
		oldlen = len(resultString)
		time.Sleep(bar.RefreshInterval)
	}
}

func (c *componentInstance) GetIdentifier() string {
	return c.component.GetIdentifier()
}
