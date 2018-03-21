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

var DefaultConfig = Configuration{
	ShowSwap: false,
}

type ComponentBuilder struct {
}

type Component struct {
	Config *Configuration
	id     string
}

type Configuration struct {
	ShowSwap     bool `yaml:"show_swap" mapstructure:"show_swap"`
	ShowBytes    bool `yaml:"show_bytes" mapstructure:"show_bytes"`
	InvertValues bool `yaml:"invert" mapstructure:"invert"`
}
