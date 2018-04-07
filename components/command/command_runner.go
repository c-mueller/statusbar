package command

import (
	"os/exec"
	"strings"
	"time"
)

type CommandRunner struct {
	UpdateInterval time.Duration
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
		c.lastValue = err.Error()
		return
	}

	c.lastValue = strings.Split(string(data), "\n")[0]
}
