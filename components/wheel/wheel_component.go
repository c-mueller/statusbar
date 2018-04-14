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

package wheel

import (
	"errors"
	"fmt"
	"github.com/c-mueller/statusbar/bar/statusbarlib"
	"github.com/c-mueller/statusbar/components/text"
	"github.com/mitchellh/mapstructure"
)

var Builder = BuilderStruct{}

var DefaultConfig = WheelConfig{
	Width: 10,
	Component: &statusbarlib.Component{
		Identifier: "text",
		Type:       "Text",
		Spec: text.ComponentConfig{
			Text: "Hello World!",
		},
	},
}

type BuilderStruct struct {
}

type Wheel struct {
	config     WheelConfig
	identifier string
	children   statusbarlib.ComponentInstances
	idx        int
	lastString string
}

type WheelConfig struct {
	Width     int                     `yaml:"width" mapstructure:"width"`
	Component *statusbarlib.Component `yaml:"component" mapstructure:"component"`
}

func (b *Wheel) GetIdentifier() string {
	return b.identifier
}

func (b *Wheel) Init() error {
	return b.children.InitializeComponents(b.identifier)
}

func (b *Wheel) Render() (*statusbarlib.RenderingOutput, error) {
	l, _, err := b.children.RenderComponentsAsString()

	if err != nil {
		return nil, err
	}

	if len(l) != len(b.lastString) || len(l) <= b.config.Width || b.idx >= len(l)+3 {
		b.idx = 0
		b.lastString = l
	}

	result := ""

	if len(l) > b.config.Width {
		l = l + " | "
		runes := []rune(l)
		left, right := b.idx, b.idx+b.config.Width

		if right >= len(l) {
			result = string(runes[left:]) + string(runes[:(right%len(l))])
		} else {
			result = string(runes[left:right])
		}
	} else {
		result = l
	}

	format := "%-" + fmt.Sprintf("%d", b.config.Width) + "s"

	result = fmt.Sprintf(format, result)

	b.idx++

	return &statusbarlib.RenderingOutput{LongText: result, ShortText: result}, nil
}

func (b *Wheel) Stop() error {
	return b.children.Stop()
}

func (b *BuilderStruct) BuildComponent(identifier string, data interface{}, builders []statusbarlib.ComponentBuilder) (statusbarlib.BarComponent, error) {
	wheel := Wheel{
		identifier: identifier,
		children:   make(statusbarlib.ComponentInstances, 0),
	}

	var childComponentConfig *WheelConfig
	err := mapstructure.Decode(data, &childComponentConfig)
	if err != nil {
		return nil, err
	}

	components := make(statusbarlib.Components, 1)
	components[0] = *childComponentConfig.Component

	err = wheel.children.InsertFromComponentList(&components, identifier, builders)
	if err != nil {
		return nil, err
	}

	if len(wheel.children) != 1 {
		return nil, errors.New("wheel: Invalid child count (has to be 1)")
	}

	wheel.config = *childComponentConfig

	return &wheel, nil
}

func (b *BuilderStruct) GetDescriptor() string {
	return "Wheel"
}

func (b *BuilderStruct) GetDefaultConfig() interface{} {
	return DefaultConfig
}
