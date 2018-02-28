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

package text

import (
	"github.com/c-mueller/statusbar/bar/bi"
	"github.com/mitchellh/mapstructure"
	"time"
)

var DefaultConfig = TextComponentConfig{
	Text: "My Message",
}

var Builder = TextComponentBuilder{}

type TextComponentBuilder struct{}

type TextComponentConfig struct {
	Text string `yaml:"text" mapstructure:"text"`
}

type TextComponent struct {
	Config TextComponentConfig
	id     string
}

func (b *TextComponentBuilder) BuildComponent(identifier string, i interface{}) (bi.BarComponent, error) {
	cfg := TextComponentConfig{}
	if i == nil {
		cfg = DefaultConfig
	} else {
		var ic *TextComponentConfig
		err := mapstructure.Decode(i, &ic)
		if err != nil {
			return nil, err
		}
		cfg = *ic
	}

	component := &TextComponent{
		Config: cfg,
		id:     identifier,
	}

	return bi.BarComponent(component), nil
}

func (b *TextComponentBuilder) GetDescriptor() string {
	return "Text"
}

func (c *TextComponent) Init() error {
	return nil
}

func (c *TextComponent) Render() (string, error) {
	return c.Config.Text, nil
}

func (c *TextComponent) IsLatest(t time.Time) bool {
	return true
}

func (c *TextComponent) Stop() error {
	return nil
}

func (c *TextComponent) GetIdentifier() string {
	return c.id
}
