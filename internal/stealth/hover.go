package stealth

import (
	"math/rand"
	"time"

	"github.com/go-rod/rod"
	"github.com/go-rod/rod/lib/proto"
)

// HoverHumanLike moves the mouse to an element,
// pauses briefly, and adds small jitter to simulate
// natural hover behavior.
func HoverHumanLike(page *rod.Page, el *rod.Element) {
	// Get element center via browser JS (version-safe)
	result, err := el.Eval(`() => {
		const r = this.getBoundingClientRect();
		return { x: r.left + r.width / 2, y: r.top + r.height / 2 };
	}`)
	if err != nil {
		return
	}

	x := result.Value.Get("x").Num()
	y := result.Value.Get("y").Num()

	// Move mouse to element center
	page.Mouse.MoveTo(proto.Point{
		X: x,
		Y: y,
	})

	// Pause as if reading
	time.Sleep(time.Duration(rand.Intn(800)+400) * time.Millisecond)

	// Small jitter movements while hovering
	for i := 0; i < rand.Intn(3)+2; i++ {
		jx := x + rand.Float64()*10 - 5
		jy := y + rand.Float64()*10 - 5

		page.Mouse.MoveTo(proto.Point{
			X: jx,
			Y: jy,
		})

		time.Sleep(time.Duration(rand.Intn(200)+100) * time.Millisecond)
	}
}
