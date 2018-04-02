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
	"github.com/c-mueller/statusbar/components/clock"
	"github.com/c-mueller/statusbar/components/cpu"
	"github.com/c-mueller/statusbar/components/date"
	"github.com/c-mueller/statusbar/components/hostname"
	"github.com/c-mueller/statusbar/components/mem"
	"github.com/c-mueller/statusbar/components/net"
	"github.com/c-mueller/statusbar/components/text"
	"github.com/c-mueller/statusbar/components/uptime"
)

const DefaultSeparator = " | "

var builders = []statusbarlib.ComponentBuilder{
	statusbarlib.ComponentBuilder(&cpu.Builder),
	statusbarlib.ComponentBuilder(&text.Builder),
	statusbarlib.ComponentBuilder(&date.Builder),
	statusbarlib.ComponentBuilder(&clock.Builder),
	statusbarlib.ComponentBuilder(&hostname.Builder),
	statusbarlib.ComponentBuilder(&mem.Builder),
	statusbarlib.ComponentBuilder(&uptime.Builder),
	statusbarlib.ComponentBuilder(&net.Builder),
}

var renderHandlers = []RenderHandler{
	NewTerminalRenderer(false),
	NewTerminalRenderer(true),
	&I3BarRenderer{},
}
