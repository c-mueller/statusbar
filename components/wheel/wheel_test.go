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

package wheel

import (
	"fmt"
	"github.com/c-mueller/statusbar/bar/statusbarlib"
	"github.com/c-mueller/statusbar/components/block"
	"github.com/c-mueller/statusbar/components/text"
	"github.com/stretchr/testify/assert"
	"gopkg.in/yaml.v2"
	"testing"
	"unicode/utf8"
)

func TestWheelLength_SpecialChars(t *testing.T) {
	testWheelLength(t, wheelOverflowTestTextConfig)
}

func TestWheelLength_RegularChars(t *testing.T) {
	testWheelLength(t, wheelIndexTestComponentDefinition)
}

func testWheelLength(t *testing.T, wc WheelConfig) {
	data, err := yaml.Marshal(wc)
	assert.NoError(t, err)
	fmt.Println(string(data))
	var parsedData interface{}
	err = yaml.Unmarshal(data, &parsedData)
	assert.NoError(t, err)
	component, err := Builder.BuildComponent("test", parsedData, wheelTestComponents)
	assert.NoError(t, err)
	assert.NoError(t, component.Init())
	for i := 0; i < len(wheelOverflowTestText)+10; i++ {
		output, err := component.Render()
		assert.NoError(t, err)

		assert.Equal(t, 10, utf8.RuneCountInString(output.LongText), fmt.Sprintf("Text: %q", output.LongText))
	}
}

const wheelOverflowTestText = "Playing | Brothers of Metal - Prophecy of Ragnarök"

var wheelTestComponents = statusbarlib.ComponentBuilders{
	statusbarlib.ComponentBuilder(&block.Builder),
	statusbarlib.ComponentBuilder(&text.Builder),
}

var wheelOverflowTestTextConfig = WheelConfig{
	Width: 10,
	Component: &statusbarlib.Component{
		Identifier: "test_block",
		Type:       "Block",
		Spec: block.BlockConfig{
			Components: statusbarlib.Components{
				statusbarlib.Component{
					Identifier: "test1",
					Type:       "Text",
					Spec: text.ComponentConfig{
						Text: wheelOverflowTestText,
					},
				},
			},
		},
	},
}

var wheelIndexTestComponentDefinition = WheelConfig{
	Width: 10,
	Component: &statusbarlib.Component{
		Identifier: "test_block",
		Type:       "Block",
		Spec: block.BlockConfig{
			Components: statusbarlib.Components{
				statusbarlib.Component{
					Identifier: "test1",
					Type:       "Text",
					Spec: text.ComponentConfig{
						Text: "1 ST MESSAGE 1",
					},
				},
				statusbarlib.Component{
					Identifier: "test2",
					Type:       "Text",
					Spec: text.ComponentConfig{
						Text: "TEST ME[]AGE 2",
					},
				},
				statusbarlib.Component{
					Identifier: "test3",
					Type:       "Text",
					Spec: text.ComponentConfig{
						Text: "TEST ME``AGE 3",
					},
				},
				statusbarlib.Component{
					Identifier: "test4",
					Type:       "Text",
					Spec: text.ComponentConfig{
						Text: "TEST ME.,AGE 4",
					},
				},
				statusbarlib.Component{
					Identifier: "test5",
					Type:       "Text",
					Spec: text.ComponentConfig{
						Text: "TEST MESSAGE 5",
					},
				},
				statusbarlib.Component{
					Identifier: "test6",
					Type:       "Text",
					Spec: text.ComponentConfig{
						Text: "TEST MESSAGE 6",
					},
				},
				statusbarlib.Component{
					Identifier: "test7",
					Type:       "Text",
					Spec: text.ComponentConfig{
						Text: "TEST MESSAGE 7",
					},
				},
				statusbarlib.Component{
					Identifier: "test8",
					Type:       "Text",
					Spec: text.ComponentConfig{
						Text: "TEST MESSAGE 8",
					},
				},
			},
		},
	},
}
