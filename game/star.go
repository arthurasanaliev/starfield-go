package game

import (
	"math/rand"

	"github.com/arthurasanaliev/starfield-go/utils"
)

type Star struct {
	x float32
	y float32
}

func NewStar() *Star {
	x := utils.STAR_RADIUS + rand.Float32()*(utils.SCREEN_WIDTH-utils.STAR_RADIUS)
	y := utils.STAR_RADIUS + rand.Float32()*(utils.SCREEN_HEIGHT-utils.STAR_RADIUS)
	return &Star{x: x, y: y}
}
