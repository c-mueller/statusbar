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

import (
	"errors"
	"fmt"
)

func (c *ComponentInstance) GetIdentifier() string {
	return c.Component.GetIdentifier()
}

func (i *ComponentInstances) AddComponent(component BarComponent, config Component, context string) error {
	err := i.CheckIdentifierValidity(component.GetIdentifier())
	if err != nil {
		return err
	}

	instance := ComponentInstance{
		Component:              component,
		Identifier:             config.Identifier,
		ComponentConfiguration: &config,
	}

	*i = append(*i, &instance)

	log.Debugf("Block %q: Added Component %q of type %q", context, config.Identifier, config.Type)

	return nil
}

func (i *ComponentInstances) InsertFromComponentList(components *Components, context string, builders []ComponentBuilder) error {
	for _, v := range *components {
		componentFound := false
		for _, builder := range builders {
			if v.Type == builder.GetDescriptor() {
				componentFound = true
				component, err := builder.BuildComponent(v.Identifier, v.Spec, builders)
				if err != nil {
					return err
				}
				err = i.AddComponent(component, v, context)
				if err != nil {
					return err
				}
			}
		}
		if !componentFound {
			return errors.New(fmt.Sprintf("Block %q: No Component of type %q found", context, v.Type))
		}
	}

	log.Debugf("Block %q: Added %d components", context, len(*i))

	return nil
}

func (i *ComponentInstances) InitializeComponents(context string) error {
	for _, v := range *i {
		log.Debugf("Block %q: Initializing Component %q", context, v.ComponentConfiguration.Identifier)
		err := v.Component.Init()
		if err != nil {
			return err
		}
	}
	return nil
}

func (i *ComponentInstances) Stop() error {
	for _, v := range *i {
		err := v.Component.Stop()
		if err != nil {
			return err
		}
	}
	return nil
}

func (i ComponentInstances) CheckIdentifierValidity(name string) error {
	for _, v := range i {
		if v.GetIdentifier() == name {
			return errors.New(fmt.Sprintf("Invalid identifier name %q is already in use", v.GetIdentifier()))
		}
	}
	return nil
}
