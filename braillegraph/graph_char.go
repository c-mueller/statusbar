package braillegraph

import (
	"fmt"
	"math"
)

type GraphChar struct {
	LeftBar  []bool
	RightBar []bool
}

func NewGraphChar(left, right float64) *GraphChar {
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

func (gc *GraphChar) ToBrailleChar() *BrailleChar {
	value := uint8(0)

	for i := 4; i > 1; i-- {
		if gc.RightBar[i-1] {
			value = value | (1 << uint8(4-i))
		}
	}
	if gc.RightBar[0] {
		value = value | (1 << 6)
	}

	for i := 4; i > 1; i-- {
		if gc.LeftBar[i-1] {
			value = value | (1 << uint8(3+(4-i)))
		}
	}
	if gc.LeftBar[0] {
		value = value | (1 << 7)
	}

	fmt.Printf("%b", value)

	return NewBrailleChar(value)
}
