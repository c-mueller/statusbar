package bar

import (
	"time"
)

type StatusBar struct {
	Components      []BarComponent
	RefreshInterval time.Duration
}

type StatusBarConfig struct {
	RefreshInterval int                        `yaml:"refresh_interval"`
	Components      []StatusBarComponentConfig `yaml:"components"`
}

type StatusBarComponentConfig struct {
	Identifier string      `yaml:"identifier"`
	Type       string      `yaml:"type"`
	Spec       interface{} `yaml:"spec"`
}
