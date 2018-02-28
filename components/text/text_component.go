package text

import (
	"github.com/c-mueller/statusbar/bar"
	"time"
)

var DefaultConfig = TextComponentConfig{
	Text: "My Message",
}

var Builder = TextComponentBuilder{}

type TextComponentBuilder struct{}

type TextComponentConfig struct {
	Text string `yaml:"text"`
}

type TextComponent struct {
	Config TextComponentConfig
}

func (b *TextComponentBuilder) BuildComponent(i *interface{}) (bar.BarComponent, error) {
	component := &TextComponent{
		Config: DefaultConfig,
	}

	return bar.BarComponent(component), nil
}

func (b *TextComponentBuilder) GetDescriptor() string {
	return "Text"
}

func (c *TextComponent) Init() error {
	return nil
}

func (c *TextComponent) Render() (string, error) {
	return c.Config.Text, nil
}

func (c *TextComponent) IsLatest(t time.Time) bool {
	return true
}

func (c *TextComponent) Stop() error {
	return nil
}
