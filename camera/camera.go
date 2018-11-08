// +build !wasm

package camera

import (
	"github.com/gopherjs/gopherjs/js"
	"github.com/gopherjs/vecty"
	"github.com/gopherjs/vecty/elem"
	"github.com/gopherjs/vecty/prop"
)

// Camera ...
type Camera struct {
	vecty.Core
	ID         string
	Constraint js.M
	stream     *js.Object
}

// Render ...
func (c *Camera) Render() vecty.ComponentOrHTML {
	return elem.Video(
		vecty.Markup(
			prop.ID(c.ID),
			vecty.Attribute("autoplay", true),
			vecty.Attribute("playsinline", true),
		),
	)
}

// Mount ...
func (c *Camera) Mount() {
	if c.Constraint == nil {
		c.Constraint = js.M{
			"video": true,
			"audio": true,
		}
	}
	mediaDevices := js.Global.Get("navigator").Get("mediaDevices")
	mediaDevices.Call("getUserMedia", c.Constraint).Call(
		"then",
		func(stream *js.Object) {
			c.stream = stream
			video := js.Global.Get("document").Call("getElementById", c.ID)
			video.Set("srcObject", c.stream)
		},
	)
}

// Unmount ...
func (c *Camera) Unmount() {
	if c.stream != nil {
		c.stream.Call("getTracks").Call("forEach", func(s *js.Object) {
			s.Call("stop")
		})
	}
}
