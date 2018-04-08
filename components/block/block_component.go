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

package block

import (
	"github.com/c-mueller/statusbar/bar/statusbarlib"
	"github.com/c-mueller/statusbar/components/clock"
	"github.com/c-mueller/statusbar/components/text"
	"github.com/mitchellh/mapstructure"
)

var Builder = BlockBuilder{}

var DefaultConfig = BlockConfig{
	Components: statusbarlib.Components{
		{
			Identifier:           "time_label",
			Type:                 "Text",
			CustomSeparator:      true,
			CustomSeparatorValue: " ",
			Spec: text.ComponentConfig{
				Text: "TIME:",
			},
		},
		{
			Identifier: "time",
			Type:       "Clock",
			Spec: clock.Configuration{
				Blink: true,
			},
		},
	},
}

type BlockBuilder struct {
}

type Block struct {
	identifier string
	children   statusbarlib.ComponentInstances
}

type BlockConfig struct {
	Components statusbarlib.Components `yaml:"components" mapstructure:"components"`
}

func (b *Block) GetIdentifier() string {
	return b.identifier
}

func (b *Block) Init() error {
	return b.children.InitializeComponents(b.identifier)
}

func (b *Block) Render() (*statusbarlib.RenderingOutput, error) {
	l, s, err := b.children.RenderComponentsAsString()
	if err != nil {
		return nil, err
	}

	return &statusbarlib.RenderingOutput{LongText: l, ShortText: s}, nil
}

func (b *Block) Stop() error {
	return b.children.Stop()
}

func (b *BlockBuilder) BuildComponent(identifier string, data interface{}, builders []statusbarlib.ComponentBuilder) (statusbarlib.BarComponent, error) {
	block := Block{
		identifier: identifier,
		children:   make(statusbarlib.ComponentInstances, 0),
	}

	var childComponentConfig *BlockConfig
	err := mapstructure.Decode(data, &childComponentConfig)
	if err != nil {
		return nil, err
	}

	err = block.children.InsertFromComponentList(&childComponentConfig.Components, identifier, builders)
	if err != nil {
		return nil, err
	}

	return &block, nil
}

func (b *BlockBuilder) GetDescriptor() string {
	return "Block"
}

func (b *BlockBuilder) GetDefaultConfig() interface{} {
	return DefaultConfig
}
