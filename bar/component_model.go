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
	"github.com/c-mueller/statusbar/bar/statusbarlib"
	"time"
)

type StatusBar struct {
	components      instantiatedComponents
	RefreshInterval time.Duration
}

type Config struct {
	RefreshInterval int        `yaml:"refresh_interval"`
	Components      Components `yaml:"components"`
}

type Component struct {
	Identifier           string      `yaml:"identifier" mapstructure:"identifier"`
	Type                 string      `yaml:"type" mapstructure:"type"`
	CustomSeparator      bool        `yaml:"custom_separator" mapstructure:"custom_separator"`
	CustomSeparatorValue string      `yaml:"separator" mapstructure:"separator"`
	HideInShortMode      bool        `yaml:"short_mode_hidden" mapstructure:"short_mode_hidden"`
	Spec                 interface{} `yaml:"spec" mapstructure:"spec"`
}

type componentInstance struct {
	config    *Component
	component statusbarlib.BarComponent
	id        string
}

type instantiatedComponents []*componentInstance
type Components []Component
