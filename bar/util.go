package bar

import "github.com/c-mueller/statusbar/bar/statusbarlib"

func (c *componentInstance) GetIdentifier() string {
	return c.component.GetIdentifier()
}

func GetComponents() []statusbarlib.ComponentBuilder {
	return builders
}

func GetRenderer() []RenderHandler {
	return renderHandlers
}
