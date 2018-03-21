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

package net

import (
	"github.com/c-mueller/statusbar/components/net/netlib"
	"github.com/op/go-logging"
	"time"
)

var log = logging.MustGetLogger("sb_builder")

var DefaultConfig = Configuration{
	InterfaceName:  "eth0",
	RecentCount:    20,
	UpdateInterval: 500,
	Global:         true,
	ShowTotal:      false,
}

type ComponentBuilder struct {
}

type Component struct {
	Config            Configuration
	updateTicker      *time.Ticker
	totalThroughput   *netlib.NetworkThroughput
	recentThroughputs netlib.ThroughputList
	id                string
}

type Configuration struct {
	InterfaceName  string `yaml:"interface_name" mapstructure:"interface_name"`
	UpdateInterval int    `yaml:"update_interval" mapstructure:"update_interval"`
	RecentCount    int    `yaml:"recent_count" mapstructure:"recent_count"`
	Global         bool   `yaml:"global" mapstructure:"global"`
	ShowTotal      bool   `yaml:"show_total" mapstructure:"show_total"`
}
