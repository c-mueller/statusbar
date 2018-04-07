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

package cpu

var DefaultLoadBarConfig = LoadBarConfiguration{
	UpdateInterval: 500,
}

var DefaultAvgChartConfig = AvgChartConfig{
	UpdateInterval: 5000,
	Width:          30,
}

type LoadBarComponentBuilder struct{}

type LoadBarComponent struct {
	Config        *LoadBarConfiguration
	id            string
	updateHandler *LoadMonitor
}

type LoadBarConfiguration struct {
	UpdateInterval int `yaml:"update_interval" mapstructure:"update_interval"`
}

type AvgChartBuilder struct {
}

type AvgChartComponent struct {
	id            string
	updateHandler *LoadMonitor
	Config        AvgChartConfig
}

type AvgChartConfig struct {
	UpdateInterval int `yaml:"update_interval" mapstructure:"update_interval"`
	Width          int `yaml:"width" mapstructure:"width"`
}
