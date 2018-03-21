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
	"github.com/c-mueller/statusbar/bar/statusbarlib"
	"os"
	"os/user"
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
	return "Hostname"
}

func (c *Component) Init() error {
	return nil
}

func (c *Component) Render() (string, error) {
	hostname, err := os.Hostname()
	if err != nil {
		return "", err
	}
	currentUser, err := user.Current()
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%s@%s", currentUser.Username, hostname), nil
}

func (c *Component) Stop() error {
	return nil
}

func (c *Component) GetIdentifier() string {
	return c.id
}
