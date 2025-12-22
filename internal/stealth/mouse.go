package stealth

import (
	"math"
	"math/rand"
	"time"

	"github.com/go-rod/rod"
	"github.com/go-rod/rod/lib/proto"
)

// MoveMouseHumanLike moves the mouse cursor along
// a Bezier curve to simulate natural human motion.
func MoveMouseHumanLike(page *rod.Page, startX, startY, endX, endY float64) {
	// Random control points for curve
	ctrlX := (startX+endX)/2 + rand.Float64()*100 - 50
	ctrlY := (startY+endY)/2 + rand.Float64()*100 - 50

	steps := rand.Intn(20) + 30 // 30–50 steps

	for i := 0; i <= steps; i++ {
		t := float64(i) / float64(steps)

		// Quadratic Bezier curve formula
		x := math.Pow(1-t, 2)*startX +
			2*(1-t)*t*ctrlX +
			math.Pow(t, 2)*endX

		y := math.Pow(1-t, 2)*startY +
			2*(1-t)*t*ctrlY +
			math.Pow(t, 2)*endY

		page.Mouse.MoveTo(proto.Point{
			X: x,
			Y: y,
		})

		time.Sleep(time.Duration(rand.Intn(15)+5) * time.Millisecond)
	}
}
