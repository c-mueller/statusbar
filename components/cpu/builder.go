//statusbar - (https://github.com/c-mueller/statusbar)
//Copyright (c) 2018 Christian Müller <cmueller.dev@gmail.com>.
//
//This program is free software: you can redistribute it and/or modify
//it under the terms of the GNU General Public License as published by
//the Free Software Foundation, either version 3 of the License, or
//(at your option) any later version.
//
//This program is distributed in the hope that it will be useful,
//but WITHOUT ANY WARRANTY; without even the implied warranty of
//MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
//GNU General Public License for more details.
//
//You should have received a copy of the GNU General Public License
//along with this program.  If not, see <http://www.gnu.org/licenses/>.

package cpu

import (
	"github.com/c-mueller/statusbar/bar/bi"
	"github.com/mitchellh/mapstructure"
)

var Builder = CPUComponentBuilder{}

func (c *CPUComponentBuilder) BuildComponent(identifier string, i interface{}) (bi.BarComponent, error) {
	cfg := &CPULoadConfiguration{}
	if i == nil {
		cfg = &DefaultConfiguration
	} else {
		var ic *CPULoadConfiguration
		err := mapstructure.Decode(i, &ic)
		if err != nil {
			return nil, err
		}
		cfg = ic
	}
	component := &CPULoadComponent{
		Config: cfg,
		id:     identifier,
	}

	return bi.BarComponent(component), nil
}

func (c *CPUComponentBuilder) GetDescriptor() string {
	return "CPULoadDisplay"
}
