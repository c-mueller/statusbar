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

import (
	"errors"
	"github.com/c-mueller/statusbar/bar/statusbarlib"
	"github.com/mitchellh/mapstructure"
)

var ChartBuilder = AvgChartBuilder{}

func (c *AvgChartBuilder) BuildComponent(identifier string, i interface{}) (statusbarlib.BarComponent, error) {
	cfg := &DefaultAvgChartConfig
	if i != nil {
		var ic *AvgChartConfig
		err := mapstructure.Decode(i, &ic)
		if err != nil {
			return nil, err
		}

		if (ic.Width % 2) != 0 {
			return nil, errors.New("cpu_avg_chart: Width has to be a multiple of 2")
		}

		cfg = ic
	}
	component := &AvgChartComponent{
		Config: *cfg,
		id:     identifier,
	}

	return statusbarlib.BarComponent(component), nil
}

func (c *AvgChartBuilder) GetDescriptor() string {
	return "CPUAvgChart"
}

func (c *AvgChartBuilder) GetDefaultConfig() interface{} {
	return DefaultAvgChartConfig
}

func (c *AvgChartComponent) Init() error {
	c.updateHandler = &LoadMonitor{
		UpdateInterval: c.Config.UpdateInterval,
		LoadAvgCount:   c.Config.Width,
		StoreAvg:       true,
	}

	return c.updateHandler.init()
}

func (c *AvgChartComponent) Render() (*statusbarlib.RenderingOutput, error) {
	value := c.updateHandler.getBrailleAverageLoadChart()
	return &statusbarlib.RenderingOutput{LongText: value, ShortText: value}, nil
}

func (c *AvgChartComponent) Stop() error {
	c.updateHandler.stop()
	return nil
}

func (c *AvgChartComponent) GetIdentifier() string {
	return c.id
}
