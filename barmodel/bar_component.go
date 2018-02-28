package barmodel

import "time"

type BarComponent interface {
	Init() error
	Render() (string, error)
	IsLatest(date *time.Time) bool
}

type ComponentBuilder interface {
	BuildComponent(data *interface{}) (*BarComponent, error)
}
