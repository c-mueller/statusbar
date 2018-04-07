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
	"fmt"
	"github.com/c-mueller/statusbar/bar/statusbarlib"
	"github.com/c-mueller/statusbar/components/block"
	"github.com/c-mueller/statusbar/components/text"
	"github.com/op/go-logging"
	"github.com/stretchr/testify/assert"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
	"strings"
	"testing"
)

func TestParse_Valid(t *testing.T) {
	initLogger()
	file, err := os.Open("testdata/config_valid.yml")
	defer file.Close()
	assert.NoError(t, err)

	content, err := ioutil.ReadAll(file)
	assert.NoError(t, err)

	sb, err := BuildFromConfig(content)
	assert.NoError(t, err)

	assert.Equal(t, 4, len(sb.Components))
	assert.Equal(t, "hostname", sb.Components[0].Identifier)
}

func TestParse_Invalid_UnknownType(t *testing.T) {
	initLogger()
	file, err := os.Open("testdata/config_invalid_unknown_type.yml")
	defer file.Close()
	assert.NoError(t, err)

	content, err := ioutil.ReadAll(file)
	assert.NoError(t, err)

	_, err = BuildFromConfig(content)

	assert.Error(t, err)
	assert.True(t, strings.Contains(err.Error(), "InvalidTestType"))
}

func TestParse_Invalid_SameID(t *testing.T) {
	initLogger()
	file, err := os.Open("testdata/config_invalid_equal_id.yml")
	defer file.Close()
	assert.NoError(t, err)

	content, err := ioutil.ReadAll(file)
	assert.NoError(t, err)

	_, err = BuildFromConfig(content)
	assert.Error(t, err)
	assert.True(t, strings.Contains(err.Error(), "Invalid identifier name \"hostname\" is already in use"))
}

func TestBlock_Build(t *testing.T) {
	initLogger()
	file, err := os.Open("testdata/config_block.yml")
	defer file.Close()
	assert.NoError(t, err)

	content, err := ioutil.ReadAll(file)
	assert.NoError(t, err)

	_, err = BuildFromConfig(content)
}

func TestMarshall_Block(t *testing.T) {
	initLogger()
	block := statusbarlib.Component{
		Identifier: "block",
		Type:       "Block",
		Spec: block.BlockConfig{
			Components: []statusbarlib.Component{
				{
					Identifier: "label",
					Type:       "Text",
					Spec: text.ComponentConfig{
						Text: "Test String A",
					},
				},
				{
					Identifier: "label",
					Type:       "Text",
					Spec: text.ComponentConfig{
						Text: "Test String B",
					},
				},
			},
		},
	}

	data, _ := yaml.Marshal(block)
	fmt.Println(string(data))
}

func initLogger() {
	var format = logging.MustStringFormatter(
		`[%{time:15:04:05} - %{level}] - %{module}: %{message}`,
	)
	backend := logging.NewLogBackend(os.Stderr, "", 0)
	backendFormatter := logging.NewBackendFormatter(backend, format)
	leveledBackend := logging.AddModuleLevel(backendFormatter)
	leveledBackend.SetLevel(logging.DEBUG, "")
	logging.SetBackend(leveledBackend)
	log.Debug("Parsed Command Line arguments")

}
