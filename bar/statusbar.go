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

var log = logging.MustGetLogger("bar")

func BuildFromConfig(config []byte) (*StatusBar, error) {
	log.Debug("Building 'statusbar'...")

	var cfg *statusbarlib.Config
	yaml.Unmarshal(config, &cfg)

	sb := newStatusBar()

	err := sb.Components.InsertFromComponentList(&cfg.Components, statusbarRootContext, ComponentBuilders)

	if err != nil {
		return nil, err
	}

	return sb, nil
}

func newStatusBar() *StatusBar {
	return &StatusBar{
		Components:      make(statusbarlib.ComponentInstances, 0),
		RefreshInterval: 500 * time.Millisecond,
	}
}

func (bar *StatusBar) GetComponents() statusbarlib.ComponentInstances {
	return bar.Components
}

func (bar *StatusBar) GetRefreshInterval() time.Duration {
	return bar.RefreshInterval
}

func (bar *StatusBar) Init() error {
	return bar.Components.InitializeComponents(statusbarRootContext)
}

func (bar *StatusBar) Render(renderer statusbarlib.RenderHandler) error {
	err := renderer.Init(bar)
	if err != nil {
		return err
	}
	return renderer.Render(bar)
}
