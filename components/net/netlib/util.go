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

import (
	"code.cloudfoundry.org/bytefmt"
	"fmt"
	"github.com/shirou/gopsutil/net"
	"strings"
)

func FromNetworkStats(stats net.IOCountersStat) *NetworkThroughput {
	return &NetworkThroughput{
		In:  stats.BytesRecv,
		Out: stats.BytesSent,
	}
}

func (a *NetworkThroughput) Subtract(b *NetworkThroughput) *NetworkThroughput {
	return &NetworkThroughput{
		In:  a.In - b.In,
		Out: a.Out - b.Out,
	}
}

func (a *NetworkThroughput) Add(b *NetworkThroughput) *NetworkThroughput {
	return &NetworkThroughput{
		In:  a.In + b.In,
		Out: a.Out + b.Out,
	}
}

func (a *NetworkThroughput) Divide(v uint64) *NetworkThroughput {
	return &NetworkThroughput{
		In:  a.In / v,
		Out: a.Out / v,
	}
}

func (a *NetworkThroughput) ToSpeedPerSecond(intervalMs int) *NetworkThroughput {
	multiplier := float64(1000) / float64(intervalMs)
	return &NetworkThroughput{
		In:  uint64(float64(a.In) * multiplier),
		Out: uint64(float64(a.Out) * multiplier),
	}
}

func (a *NetworkThroughput) FormatToString() string {
	// Format the String with '-' As padding '_' is used to define spaces in the result string
	result := fmt.Sprintf("D:_%7s_U:_%7s", bytefmt.ByteSize(a.In), bytefmt.ByteSize(a.Out))
	// Replace Spaces with '-' to get a output like --3.1G
	result = strings.Replace(result, " ", "-", -1)
	// Replace '_' with spaces to produce the end Result
	return strings.Replace(result, "_", " ", -1)
}

func (r ThroughputList) ComputeAverage() *NetworkThroughput {
	sum := &NetworkThroughput{In: 0, Out: 0}
	if len(r) == 0 {
		return sum
	}
	for _, v := range r {
		sum = sum.Add(&v)
	}
	return sum.Divide(uint64(len(r)))
}

func (r *NetworkThroughput) ToBit() NetworkThroughput {
	return NetworkThroughput{
		In:  r.In * 8,
		Out: r.Out * 8,
	}
}
