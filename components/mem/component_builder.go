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

package mem

import (
	"github.com/c-mueller/statusbar/bar/statusbarlib"
	"github.com/mitchellh/mapstructure"
)

var Builder = ComponentBuilder{}

func (c *ComponentBuilder) BuildComponent(identifier string, i interface{}) (statusbarlib.BarComponent, error) {
	cfg := &Configuration{}
	if i == nil {
		cfg = &DefaultConfig
	} else {
		var ic *Configuration
		err := mapstructure.Decode(i, &ic)
		if err != nil {
			return nil, err
		}
		cfg = ic
	}
	component := &Component{
		Config: cfg,
		id:     identifier,
	}

	return statusbarlib.BarComponent(component), nil
}

func (c *ComponentBuilder) GetDefaultConfig() interface{} {
	return DefaultConfig
}

func (c *ComponentBuilder) GetDescriptor() string {
	return "Memory"
}
