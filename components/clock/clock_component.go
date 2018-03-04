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

package clock

import (
	"fmt"
	"github.com/c-mueller/statusbar/bar/statusbarlib"
	"github.com/mitchellh/mapstructure"
	"time"
)

var DefaultConfig = ClockConfig{
	Blink: true,
}

var Builder = ClockComponentBuilder{}

type ClockComponentBuilder struct{}

type ClockConfig struct {
	Blink bool `yaml:"blink" mapstructure:"blink"`
}

type ClockComponent struct {
	Config ClockConfig
	id     string
}

func (b *ClockComponentBuilder) BuildComponent(identifier string, i interface{}) (statusbarlib.BarComponent, error) {
	cfg := ClockConfig{}
	if i == nil {
		cfg = DefaultConfig
	} else {
		var ic *ClockConfig
		err := mapstructure.Decode(i, &ic)
		if err != nil {
			return nil, err
		}
		cfg = *ic
	}

	component := &ClockComponent{
		Config: cfg,
		id:     identifier,
	}

	return statusbarlib.BarComponent(component), nil
}

func (b *ClockComponentBuilder) GetDescriptor() string {
	return "Clock"
}

func (c *ClockComponent) Init() error {
	return nil
}

func (c *ClockComponent) Render() (string, error) {
	format := "%02d:%02d:%02d"
	h, m, s := time.Now().Clock()
	if (s%2) == 0 && c.Config.Blink {
		format = "%02d %02d %02d"
	}
	return fmt.Sprintf(format, h, m, s), nil
}

func (c *ClockComponent) Stop() error {
	return nil
}

func (c *ClockComponent) GetIdentifier() string {
	return c.id
}
