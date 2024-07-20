package game

import (
	"image/color"

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
		g.Stars[i].x += utils.STAR_SPEED
		if g.Stars[i].x > utils.SCREEN_WIDTH+utils.STAR_RADIUS {
			g.Stars[i].x = -utils.STAR_RADIUS
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
