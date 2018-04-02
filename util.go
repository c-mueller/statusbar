package main

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

import (
	"github.com/c-mueller/statusbar/bar"
	"github.com/op/go-logging"
	"os"
	"strings"
)

var format = logging.MustStringFormatter(
	`%{color}[%{time:15:04:05} - %{level}] - %{module}:%{color:reset} %{message}`,
)

var log = logging.MustGetLogger("sb_main")

func exitOnErr(err error, code int, format string, values ...interface{}) {
	if err != nil {
		log.Errorf(format, values...)
		log.Errorf("Error Message: %q", err.Error())
		os.Exit(code)
	}
}

func initializeLogger() {
	backend := logging.NewLogBackend(os.Stderr, "", 0)
	backendFormatter := logging.NewBackendFormatter(backend, format)
	leveledBackend := logging.AddModuleLevel(backendFormatter)
	if *debug {
		leveledBackend.SetLevel(logging.DEBUG, "")
	} else if *verbose {
		leveledBackend.SetLevel(logging.INFO, "")
	} else {
		leveledBackend.SetLevel(logging.ERROR, "")
	}
	logging.SetBackend(leveledBackend)
	log.Debug("Parsed Command Line arguments")
}

func findRenderer(name string) bar.RenderHandler {
	for _, v := range bar.GetRenderer() {
		if strings.ToLower(name) == strings.ToLower(v.GetName()) {
			return v
		}
	}
	return nil
}
