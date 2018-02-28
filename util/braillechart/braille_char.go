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

package braillechart

type BrailleCharacter struct {
	Dots []bool
}

func NewBrailleChar(i uint8) *BrailleCharacter {
	dotValues := make([]bool, 8)
	for k := range dotValues {
		v := (i >> uint(k)) & 0x1
		boolv := false
		if v == 1 {
			boolv = true
		}
		dotValues[k] = boolv
	}
	return &BrailleCharacter{
		Dots: dotValues,
	}
}

func (c BrailleCharacter) MapToBrailleChar() (rune, error) {
	if len(c.Dots) != 8 {
		return 0, InvalidDotLength
	}
	value := uint(0)
	for k, v := range c.Dots {
		i := uint(0)
		if v {
			i = 1
		}
		value = value | (i << uint(k))
	}
	return rune(0x2800 | (value & 0xFF)), nil
}
