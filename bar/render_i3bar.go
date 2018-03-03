package bar

import (
	"fmt"
	"github.com/gin-gonic/gin/json"
)

type I3BarRenderer struct {
}

type i3BarHeader struct {
	Version     int  `json:"version"`
	ClickEvents bool `json:"click_events"`
}

func (r *I3BarRenderer) writeHeader() {
	header := i3BarHeader{
		Version:     1,
		ClickEvents: false,
	}
	data, _ := json.Marshal(header)
	fmt.Println(string(data))
}

func (r *I3BarRenderer) Render(sb *StatusBar) error {
	return nil
}

func (r *I3BarRenderer) Init(sb *StatusBar) error {
	r.writeHeader()
	return nil
}
