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

package netlib

import "time"

type ThroughputLogger struct {
	UpdateInterval     time.Duration
	Global             bool
	InterfaceName      string
	StoredValueCount   int
	LastMeasurement    NetworkThroughput
	RecentMeasurements ThroughputList
	updateTicker       *time.Ticker
}

type NetworkThroughput struct {
	In  uint64 `json:"in" yaml:"in" xml:"in"`
	Out uint64 `json:"out" json:"out" xml:"out"`
}

type ThroughputList []NetworkThroughput
