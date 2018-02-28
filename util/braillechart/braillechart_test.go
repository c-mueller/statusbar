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
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_EncodeUnary(t *testing.T) {
	for i := 1; i <= 63; i++ {
		encoded := encodeUnary(uint8(i))
		expected := (1 << uint64(i)) - 1
		t.Logf("i=%d: Expected: %d Got: %d", i, uint64(expected), encoded)
		assert.Equal(t, encoded, uint64(expected))
	}
}

func Test_Between_Zero_And_One(t *testing.T) {
	//True Cases
	assert.Equal(t, isBetweenZeroAndOne(0), true)
	assert.Equal(t, isBetweenZeroAndOne(1), true)
	assert.Equal(t, isBetweenZeroAndOne(0.5), true)
	//False Cases
	assert.Equal(t, isBetweenZeroAndOne(10), false)
	assert.Equal(t, isBetweenZeroAndOne(-1), false)
}

func TestBraille_Char_Empty(t *testing.T) {
	bc := NewBrailleChar(0)
	assert.NotNil(t, bc)
	chr, err := bc.MapToBrailleChar()
	logChar(t, chr)
	assert.NoError(t, err)
	assert.Equal(t, rune(0x2800), chr)
}

func TestBraille_Char_Full(t *testing.T) {
	bc := NewBrailleChar(0xFF)
	assert.NotNil(t, bc)
	chr, err := bc.MapToBrailleChar()
	logChar(t, chr)
	assert.NoError(t, err)
	assert.Equal(t, rune(0x28FF), chr)
}

func TestBraille_Char_6_Full(t *testing.T) {
	bc := NewBrailleChar(0x3F)
	assert.NotNil(t, bc)
	chr, err := bc.MapToBrailleChar()
	logChar(t, chr)
	assert.NoError(t, err)
	assert.Equal(t, rune(0x283F), chr)
}

func TestBraille_Char_C0(t *testing.T) {
	bc := NewBrailleChar(0xC0)
	assert.NotNil(t, bc)
	chr, err := bc.MapToBrailleChar()
	logChar(t, chr)
	assert.NoError(t, err)
	assert.Equal(t, rune(0x28C0), chr)
}

func TestBraille_Chart_Right_Complete(t *testing.T) {
	g := NewChartChar(0, 1)
	assert.NotNil(t, g)
	bc := g.ToBrailleChar()
	chr := processBrailleChar(t, bc)
	assert.Equal(t, rune(0x28B8), chr)
}

func TestBraille_Chart_Right_3Quarter(t *testing.T) {
	g := NewChartChar(0, 0.75)
	assert.NotNil(t, g)
	bc := g.ToBrailleChar()
	chr := processBrailleChar(t, bc)
	assert.Equal(t, rune(0x28B0), chr)
}

func TestBraille_Chart_Right_Half(t *testing.T) {
	g := NewChartChar(0, 0.5)
	assert.NotNil(t, g)
	bc := g.ToBrailleChar()
	chr := processBrailleChar(t, bc)
	assert.Equal(t, rune(0x28A0), chr)
}

func TestBraille_Chart_Right_Quarter(t *testing.T) {
	g := NewChartChar(0, 0.25)
	assert.NotNil(t, g)
	bc := g.ToBrailleChar()
	chr := processBrailleChar(t, bc)
	assert.Equal(t, rune(0x2880), chr)
}

func TestBraille_Chart_Left_Complete(t *testing.T) {
	g := NewChartChar(1, 0)
	assert.NotNil(t, g)
	bc := g.ToBrailleChar()
	chr := processBrailleChar(t, bc)
	assert.Equal(t, rune(0x2847), chr)
}

func TestBraille_Chart_Left_3Quarter(t *testing.T) {
	g := NewChartChar(0.75, 0)
	assert.NotNil(t, g)
	bc := g.ToBrailleChar()
	chr := processBrailleChar(t, bc)
	assert.Equal(t, rune(0x2846), chr)
}

func TestBraille_Chart_Left_Half(t *testing.T) {
	g := NewChartChar(0.5, 0)
	assert.NotNil(t, g)
	bc := g.ToBrailleChar()
	chr := processBrailleChar(t, bc)
	assert.Equal(t, rune(0x2844), chr)
}

func TestBraille_Chart_Left_Quarter(t *testing.T) {
	g := NewChartChar(0.25, 0)
	assert.NotNil(t, g)
	bc := g.ToBrailleChar()
	chr := processBrailleChar(t, bc)
	assert.Equal(t, rune(0x2840), chr)
}

func TestBraille_Chart_Zero(t *testing.T) {
	g := NewChartChar(0, 0)
	assert.NotNil(t, g)
	bc := g.ToBrailleChar()
	chr := processBrailleChar(t, bc)
	assert.Equal(t, rune(0x2800), chr)
}

func processBrailleChar(t *testing.T, bc *BrailleCharacter) rune {
	assert.NotNil(t, bc)
	chr, err := bc.MapToBrailleChar()
	logChar(t, chr)
	assert.NoError(t, err)
	return chr
}

func logChar(t *testing.T, chr rune) {
	t.Logf("Got Char: %U - %q", chr, chr)
}
