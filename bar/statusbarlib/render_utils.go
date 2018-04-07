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

package statusbarlib

func (i *ComponentInstances) RenderComponentsAsString() (string, string, error) {
	longString, shortString := "", ""
	for idx, v := range *i {
		l, s, err := RenderComponentAsString(idx, i, v)
		if err != nil {
			return "", "", err
		}
		longString += l
		shortString += s
	}

	return longString, shortString, nil
}

func RenderComponentAsString(index int, components *ComponentInstances, component *ComponentInstance) (string, string, error) {
	shortString := ""
	longString := ""

	r, err := component.Component.Render()
	if err != nil {
		return "", "", err
	}
	if !component.ComponentConfiguration.HideInShortMode {
		shortString = AppendSeparator(r.ShortText, index, components, component)
	}

	longString = AppendSeparator(r.LongText, index, components, component)

	return longString, shortString, nil
}

func AppendSeparator(r string, i int, components *ComponentInstances, v *ComponentInstance) string {
	renderString := r

	if i < len(*components)-1 {
		if v.ComponentConfiguration.CustomSeparator {
			renderString += v.ComponentConfiguration.CustomSeparatorValue
		} else {
			renderString += DefaultSeparator
		}
	}
	return renderString
}
