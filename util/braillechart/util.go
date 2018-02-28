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

func encodeUnary(i uint8) uint64 {
	if i > 64 {
		return 0
	}
	result := uint64(0)
	for idx := 0; uint8(idx) < i; idx++ {
		result = (uint64(1) << uint(idx)) | result
	}
	return result
}

func toBoolArray(v, len uint) []bool {
	data := make([]bool, len)
	for i := uint(0); i < len; i++ {
		value := (v >> i) & 0x1
		data[i] = value == 1
	}
	return data
}

func isBetweenZeroAndOne(value float64) bool {
	return isInRange(value, 0, 1)
}

func isInRange(value, lower, upper float64) bool {
	return value >= lower && value <= upper
}
