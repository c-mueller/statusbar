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
	"github.com/op/go-logging"
	"gopkg.in/yaml.v2"
	"time"
)

var log = logging.MustGetLogger("sb_builder")

func BuildFromConfig(config []byte) (*StatusBar, error) {
	log.Debug("Building Statusbar...")

	var cfg *Config
	yaml.Unmarshal(config, &cfg)

	sb := newStatusBar()

	sb.components.insertFromComponentList(&cfg.Components, statusbarRootContext)

	return sb, nil
}

func newStatusBar() *StatusBar {
	return &StatusBar{
		components:      make(instantiatedComponents, 0),
		RefreshInterval: 500 * time.Millisecond,
	}
}

func (bar *StatusBar) addComponent(component statusbarlib.BarComponent, config Component) error {
	return bar.components.addComponent(component, config, statusbarRootContext)
}

func (bar *StatusBar) Init() error {
	return bar.components.init(statusbarRootContext)
}

func (bar *StatusBar) RenderTerminal(short bool) error {
	return bar.Render(&TerminalRenderer{
		ShortMode: short,
	})
}

func (bar *StatusBar) RenderI3() error {
	return bar.Render(&I3BarRenderer{})
}

func (bar *StatusBar) Render(renderer RenderHandler) error {
	err := renderer.Init(bar)
	if err != nil {
		return err
	}
	return renderer.Render(bar)
}

func (c *componentInstance) GetIdentifier() string {
	return c.component.GetIdentifier()
}

func GetComponents() []statusbarlib.ComponentBuilder {
	return builders
}

func GetRenderer() []RenderHandler {
	return renderHandlers
}
