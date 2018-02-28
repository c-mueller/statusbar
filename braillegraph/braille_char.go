package braillegraph

type BrailleChar struct {
	Dots []bool
}

func NewBrailleChar(i uint8) *BrailleChar {
	dotValues := make([]bool, 8)
	for k := range dotValues {
		v := (i >> uint(k)) & 0x1
		boolv := false
		if v == 1 {
			boolv = true
		}
		dotValues[k] = boolv
	}
	return &BrailleChar{
		Dots: dotValues,
	}
}

func (c BrailleChar) MapToBrailleChar() (rune, error) {
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
