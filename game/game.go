package game

import (
	"image/color"
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
		if outOfBounds(g.Stars[i].x, g.Stars[i].y) {
			g.Stars[i].x = utils.STAR_RADIUS + rand.Float32()*(utils.SCREEN_WIDTH-utils.STAR_RADIUS)
			g.Stars[i].y = utils.STAR_RADIUS + rand.Float32()*(utils.SCREEN_HEIGHT-utils.STAR_RADIUS)
			g.Stars[i].z = utils.MAX_Z
		}

		g.Stars[i].z -= 1.3
		if g.Stars[i].z < utils.MIN_Z {
			g.Stars[i].z = utils.MAX_Z
		}
	}
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	for _, star := range g.Stars {
		vector.DrawFilledCircle(screen, float32(star.x), float32(star.y), float32(utils.STAR_RADIUS), color.White, false)
	}
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return utils.SCREEN_WIDTH, utils.SCREEN_HEIGHT
}

func translateCoordinates(x, y float32) (float32, float32) {
	return x - utils.CENTER_WIDTH, y - utils.CENTER_HEIGHT
}

func outOfBounds(x, y float32) bool {
	if x < -utils.STAR_RADIUS || x > utils.SCREEN_WIDTH+utils.STAR_RADIUS {
		return true
	}
	if y < -utils.STAR_RADIUS || y > utils.SCREEN_HEIGHT+utils.STAR_RADIUS {
		return true
	}
	return false
}
