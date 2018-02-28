package braillegraph

import (
	"testing"
	"github.com/stretchr/testify/assert"
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
	logChar(t,chr)
	assert.NoError(t, err)
	assert.Equal(t, rune(0x2800), chr)
}

func TestBraille_Char_Full(t *testing.T) {
	bc := NewBrailleChar(0xFF)
	assert.NotNil(t, bc)
	chr, err := bc.MapToBrailleChar()
	logChar(t,chr)
	assert.NoError(t, err)
	assert.Equal(t, rune(0x28FF), chr)
}

func TestBraille_Char_6_Full(t *testing.T) {
	bc := NewBrailleChar(0x3F)
	assert.NotNil(t, bc)
	chr, err := bc.MapToBrailleChar()
	logChar(t,chr)
	assert.NoError(t, err)
	assert.Equal(t, rune(0x283F), chr)
}

func TestBraille_Char_C0(t *testing.T) {
	bc := NewBrailleChar(0xC0)
	assert.NotNil(t, bc)
	chr, err := bc.MapToBrailleChar()
	logChar(t,chr)
	assert.NoError(t, err)
	assert.Equal(t, rune(0x28C0), chr)
}

func TestBraille_Graph_Right_Complete(t *testing.T) {
	g := NewGraphChar(0, 1)
	assert.NotNil(t, g)
	bc := g.ToBrailleChar()
	chr := processBrailleChar(t, bc)
	assert.Equal(t, rune(0x28B8), chr)
}

func TestBraille_Graph_Right_3Quarter(t *testing.T) {
	g := NewGraphChar(0, 0.75)
	assert.NotNil(t, g)
	bc := g.ToBrailleChar()
	chr := processBrailleChar(t, bc)
	assert.Equal(t, rune(0x28B0), chr)
}

func TestBraille_Graph_Right_Half(t *testing.T) {
	g := NewGraphChar(0, 0.5)
	assert.NotNil(t, g)
	bc := g.ToBrailleChar()
	chr := processBrailleChar(t, bc)
	assert.Equal(t, rune(0x28A0), chr)
}

func TestBraille_Graph_Right_Quarter(t *testing.T) {
	g := NewGraphChar(0, 0.25)
	assert.NotNil(t, g)
	bc := g.ToBrailleChar()
	chr := processBrailleChar(t, bc)
	assert.Equal(t, rune(0x2880), chr)
}

func TestBraille_Graph_Left_Complete(t *testing.T) {
	g := NewGraphChar(1, 0)
	assert.NotNil(t, g)
	bc := g.ToBrailleChar()
	chr := processBrailleChar(t, bc)
	assert.Equal(t, rune(0x2847), chr)
}

func TestBraille_Graph_Left_3Quarter(t *testing.T) {
	g := NewGraphChar(0.75, 0)
	assert.NotNil(t, g)
	bc := g.ToBrailleChar()
	chr := processBrailleChar(t, bc)
	assert.Equal(t, rune(0x2846), chr)
}

func TestBraille_Graph_Left_Half(t *testing.T) {
	g := NewGraphChar(0.5, 0)
	assert.NotNil(t, g)
	bc := g.ToBrailleChar()
	chr := processBrailleChar(t, bc)
	assert.Equal(t, rune(0x2844), chr)
}

func TestBraille_Graph_Left_Quarter(t *testing.T) {
	g := NewGraphChar(0.25, 0)
	assert.NotNil(t, g)
	bc := g.ToBrailleChar()
	chr := processBrailleChar(t, bc)
	assert.Equal(t, rune(0x2840), chr)
}

func TestBraille_Graph_Zero(t *testing.T) {
	g := NewGraphChar(0, 0)
	assert.NotNil(t, g)
	bc := g.ToBrailleChar()
	chr := processBrailleChar(t, bc)
	assert.Equal(t, rune(0x2800), chr)
}

func processBrailleChar(t *testing.T, bc *BrailleChar) rune {
	assert.NotNil(t, bc)
	chr, err := bc.MapToBrailleChar()
	logChar(t, chr)
	assert.NoError(t, err)
	return chr
}

func logChar(t *testing.T, chr rune) {
	t.Logf("Got Char: %U - %q", chr, chr)
}
