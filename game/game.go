package game

import (
	"image/color"
	"math"
	"math/rand"

	"github.com/arthurasanaliev/starfield-go/utils"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

type Game struct {
	Stars []Star
}

func NewGame() *Game {
	return &Game{}
}

func (g *Game) Update() error {
	for i := range g.Stars {
		translatedX, translatedY := translateCoordinates(g.Stars[i].x, g.Stars[i].y)

		stepX := translatedX / g.Stars[i].z
		stepY := translatedY / g.Stars[i].z

		g.Stars[i].x += stepX
		g.Stars[i].y += stepY
		g.Stars[i].r += float32(math.Max(math.Abs(float64(stepX)), math.Abs(float64(stepY))) / 40.0)
		if outOfBounds(g.Stars[i].x, g.Stars[i].y, g.Stars[i].r) {
			g.Stars[i].x = g.Stars[i].r + rand.Float32()*(utils.SCREEN_WIDTH-g.Stars[i].r)
			g.Stars[i].y = g.Stars[i].r + rand.Float32()*(utils.SCREEN_HEIGHT-g.Stars[i].r)
			g.Stars[i].z = utils.MAX_Z
			g.Stars[i].r = utils.STAR_RADIUS
		}

		g.Stars[i].z -= 1.2
		if g.Stars[i].z < utils.MIN_Z {
			g.Stars[i].z = 1.2
		}
	}
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	for _, star := range g.Stars {
		vector.DrawFilledCircle(screen, star.x, star.y, star.r, color.White, false)
	}
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return utils.SCREEN_WIDTH, utils.SCREEN_HEIGHT
}

func translateCoordinates(x, y float32) (float32, float32) {
	return x - utils.CENTER_WIDTH, y - utils.CENTER_HEIGHT
}

func outOfBounds(x, y, r float32) bool {
	if x < -r || x > utils.SCREEN_WIDTH+r {
		return true
	}
	if y < -r || y > utils.SCREEN_HEIGHT+r {
		return true
	}
	return false
}
