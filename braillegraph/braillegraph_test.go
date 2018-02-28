package braillegraph

import (
	"github.com/magiconair/properties/assert"
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

func TestBraille_Graph(t *testing.T) {

}
