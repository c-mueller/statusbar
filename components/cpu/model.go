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

import "time"

var DefaultConfiguration = CPULoadConfiguration{
	UpdateInterval:   1,
	ShowAverageLoad:  true,
	LoadAverageCount: 120,
}

type CPUComponentBuilder struct{}

type CPULoadComponent struct {
	Config          *CPULoadConfiguration
	id              string
	cpuUpdateTicker *time.Ticker
	cpuLoads        []float64
	currentValue    string
	updateTimestamp time.Time
	recentAverages  []float64
	currentAverage  float64
}

type CPULoadConfiguration struct {
	UpdateInterval   int  `yaml:"update_interval" mapstructure:"update_interval"`
	ShowAverageLoad  bool `yaml:"show_average_load" mapstructure:"show_average_load"`
	LoadAverageCount int  `yaml:"load_average_count" mapstructure:"load_average_count"`
}
