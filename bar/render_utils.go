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

import "fmt"

func renderComponent(index int, bar *StatusBar, component *componentInstance) (string, string, error) {
	shortString := ""
	longString := ""

	r, err := component.component.Render()
	if err != nil {
		return "", "", err
	}
	if !component.config.HideInShortMode {
		shortString = getResultString(r, index, bar, component)
	}

	longString = getResultString(r, index, bar, component)

	return longString, shortString, nil
}

func getResultString(r string, i int, bar *StatusBar, v *componentInstance) string {
	renderString := r
	if i < len(bar.components)-1 {
		if v.config.CustomSeparator {
			renderString += v.config.CustomSeparatorValue
		} else {
			renderString += DefaultSeparator
		}
	}
	return renderString
}

func writeBlanksOnLine(count int) {
	fmt.Printf("\r")
	for i := 0; i < count; i++ {
		fmt.Printf(" ")
	}
}
