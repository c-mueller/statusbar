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

package command

import (
	"os/exec"
	"strings"
	"time"
)

type CommandRunner struct {
	UpdateInterval time.Duration
	ErrorMessage   string
	Command        string
	lastValue      string
	updateTicker   *time.Ticker
}

func NewCommandRunner(cmd string, interval time.Duration) *CommandRunner {
	return &CommandRunner{
		UpdateInterval: interval,
		Command:        cmd,
		updateTicker:   time.NewTicker(interval),
		lastValue:      "Not executed!",
	}
}

func (c *CommandRunner) Start() {
	go c.commandRunnerTickerLoop()
}

func (c *CommandRunner) Stop() {
	go c.updateTicker.Stop()
}

func (c *CommandRunner) commandRunnerTickerLoop() {
	c.runCommand()
	for range c.updateTicker.C {
		c.runCommand()
	}
}

func (c *CommandRunner) runCommand() {
	splitCmd := strings.Split(c.Command, " ")
	remaining := make([]string, 0)
	if len(splitCmd) > 1 {
		remaining = splitCmd[1:]
	}
	cmd := exec.Command(splitCmd[0], remaining...)

	data, err := cmd.Output()
	if err != nil {
		if c.ErrorMessage != "" {
			c.lastValue = c.ErrorMessage
		} else  {
			c.lastValue = err.Error()
		}
		return
	}

	c.lastValue = strings.Split(string(data), "\n")[0]
}
