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

package uptime

import (
	"fmt"
	"github.com/c-mueller/statusbar/bar/statusbarlib"
	"github.com/shirou/gopsutil/host"
	"time"
)

var Builder = ComponentBuilder{}

type ComponentBuilder struct{}

type Component struct {
	id string
}

func (b *ComponentBuilder) BuildComponent(identifier string, i interface{}) (statusbarlib.BarComponent, error) {
	component := &Component{
		id: identifier,
	}

	return statusbarlib.BarComponent(component), nil
}

func (b *ComponentBuilder) GetDescriptor() string {
	return "Uptime"
}

func (c *Component) Init() error {
	return nil
}

func (c *Component) Render() (string, error) {
	ut, err := getUptimeDuration()
	if err != nil {
		return "", err
	}
	return formatDuration(ut), nil
}

func (c *Component) Stop() error {
	return nil
}

func (c *Component) GetIdentifier() string {
	return c.id
}

func getUptimeDuration() (time.Duration, error) {
	uptime, err := host.Uptime()
	if err != nil {
		return 0, err
	}
	return time.Duration(uptime) * time.Second, nil
}

func formatDuration(d time.Duration) string {
	days := int(d.Hours() / 24)
	hours := int(float64(d.Hours()) - float64(days)*24)
	minutes := int(d.Minutes()) - (days*24+hours)*60
	if days == 0 {
		return fmt.Sprintf("%02dh %02dm", hours, minutes)
	} else {
		return fmt.Sprintf("%02dd %02dh %02dm", days, hours, minutes)
	}
}
