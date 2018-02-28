package bar

import "time"

type BarComponent interface {
	Init() error
	Render() (string, error)
	IsLatest(date time.Time) bool
	Stop() error
}

type ComponentBuilder interface {
	BuildComponent(data *interface{}) (BarComponent, error)
	GetDescriptor() string
}
