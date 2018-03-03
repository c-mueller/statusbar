package bar

type RenderHandler interface {
	Init(bar *StatusBar) error
	Render(bar *StatusBar) error
}
