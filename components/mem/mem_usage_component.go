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

package mem

import (
	"code.cloudfoundry.org/bytefmt"
	"fmt"
	"github.com/c-mueller/statusbar/bar/statusbarlib"
	"github.com/c-mueller/statusbar/util"
	"github.com/shirou/gopsutil/mem"
	"strings"
)

func (c *Component) Init() error {
	return nil
}

func (c *Component) Render() (*statusbarlib.RenderingOutput, error) {
	outputString := c.renderMemoryPercentage("MEM: ", getAvailableMemoryPercentage)

	if c.Config.ShowSwap {
		outputString += c.renderMemoryPercentage("| SWP: ", getSwapMemoryPercentage)
	}
	return &statusbarlib.RenderingOutput{LongText: outputString, ShortText: outputString}, nil
}

func (c *Component) Stop() error {
	return nil
}

func (c *Component) GetIdentifier() string {
	return c.id
}

func (c *Component) renderMemoryPercentage(prefix string, f func(bool) (float64, uint64)) string {
	outputString := ""
	percentage, bytes := f(c.Config.InvertValues)
	outputString += fmt.Sprintf("%s%02d%%", prefix, int(percentage))
	if c.Config.ShowBytes && bytes != 0 {
		value := fmt.Sprintf("(%6s)", bytefmt.ByteSize(bytes))
		value = strings.Replace(value, " ", "-", -1)
		outputString += " " + value
	}
	return outputString
}

func getAvailableMemoryPercentage(invert bool) (float64, uint64) {
	virtualMemory, _ := mem.VirtualMemory()
	return getMemoryPercentage(virtualMemory.Available, virtualMemory.Total, invert)
}

func getSwapMemoryPercentage(invert bool) (float64, uint64) {
	swapMemory, _ := mem.SwapMemory()
	return getMemoryPercentage(swapMemory.Free, swapMemory.Total, invert)
}

func getMemoryPercentage(free, total uint64, invert bool) (float64, uint64) {
	percentage := (float64(free) / float64(total)) * 100
	byteCount := free
	if invert {
		percentage = 100 - percentage
		byteCount = total - byteCount
	}
	return util.Round(percentage, 0.5, 0), byteCount
}
