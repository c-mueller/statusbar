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

package hostname

import (
	"fmt"
	"github.com/c-mueller/statusbar/bar/bi"
	"os"
	"os/user"
	"time"
)

var Builder = HostnameComponentBuilder{}

type HostnameComponentBuilder struct{}

type HostnameComponent struct {
	id string
}

func (b *HostnameComponentBuilder) BuildComponent(identifier string, i interface{}) (bi.BarComponent, error) {
	component := &HostnameComponent{
		id: identifier,
	}

	return bi.BarComponent(component), nil
}

func (b *HostnameComponentBuilder) GetDescriptor() string {
	return "Hostname"
}

func (c *HostnameComponent) Init() error {
	return nil
}

func (c *HostnameComponent) Render() (string, error) {
	hostname, err := os.Hostname()
	if err != nil {
		return "", err
	}
	u, err := user.Current()
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%s@%s", u.Username, hostname), nil
}

func (c *HostnameComponent) IsLatest(t time.Time) bool {
	return true
}

func (c *HostnameComponent) Stop() error {
	return nil
}

func (c *HostnameComponent) GetIdentifier() string {
	return c.id
}
