//statusbar - (https://github.com/c-mueller/statusbar)
//Copyright (c) 2018 Christian Müller <cmueller.dev@gmail.com>.
//
//This program is free software: you can redistribute it and/or modify
//it under the terms of the GNU General Public License as published by
//the Free Software Foundation, either version 3 of the License, or
//(at your option) any later version.
//
//This program is distributed in the hope that it will be useful,
//but WITHOUT ANY WARRANTY; without even the implied warranty of
//MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
//GNU General Public License for more details.
//
//You should have received a copy of the GNU General Public License
//along with this program.  If not, see <http://www.gnu.org/licenses/>.

package braillechart

import (
	"math"
)

type GraphChar struct {
	LeftBar  []bool
	RightBar []bool
}

func NewChartChar(left, right float64) *GraphChar {
	//if !isBetweenZeroAndOne(left) || !isBetweenZeroAndOne(right) {
	//	fmt.Printf("L: %f, R: %f\n", left, right)
	//	return nil
	//}
	gc := GraphChar{
		LeftBar:  make([]bool, 4),
		RightBar: make([]bool, 4),
	}

	leftValue := uint64(math.Floor(left * 5))
	if leftValue == 5 {
		leftValue = 4
	}
	leftValue = encodeUnary(uint8(leftValue))
	gc.LeftBar = toBoolArray(uint(leftValue), 4)

	rightValue := uint64(math.Floor(right * 5))
	if rightValue == 5 {
		rightValue = 4
	}
	rightValue = encodeUnary(uint8(rightValue))
	gc.RightBar = toBoolArray(uint(rightValue), 4)

	return &gc
}

func (gc *GraphChar) ToBrailleChar() *BrailleCharacter {
	value := uint8(0)

	for i := 4; i > 1; i-- {
		if gc.LeftBar[i-1] {
			value = value | (1 << uint8(4-i))
		}
	}
	if gc.LeftBar[0] {
		value = value | (1 << 6)
	}

	for i := 4; i > 1; i-- {
		if gc.RightBar[i-1] {
			value = value | (1 << uint8(3+(4-i)))
		}
	}
	if gc.RightBar[0] {
		value = value | (1 << 7)
	}

	return NewBrailleChar(value)
}
