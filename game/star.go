package game

import (
	"math/rand"

	"github.com/arthurasanaliev/starfield-go/utils"
)

type Star struct {
	x float32
	y float32
	z float32
	r float32
}

func NewStar() *Star {
	r := utils.STAR_RADIUS
	x := r + rand.Float32()*(utils.SCREEN_WIDTH-r)
	y := r + rand.Float32()*(utils.SCREEN_HEIGHT-r)
	z := utils.MAX_Z
	return &Star{x: x, y: y, z: z, r: r}
}
