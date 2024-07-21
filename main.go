package main

import (
	"log"
	"math/rand"
	"time"

	"github.com/arthurasanaliev/starfield-go/game"
	"github.com/arthurasanaliev/starfield-go/utils"
	"github.com/hajimehoshi/ebiten/v2"
)

func main() {
	g := game.NewGame()

	rand.Seed(time.Now().UnixNano())
	addStars(g)

	ebiten.SetWindowSize(utils.WINDOW_WIDTH, utils.WINDOW_HEIGHT)
	ebiten.SetWindowTitle("Starfield by Arthur Asanaliev")
	if err := ebiten.RunGame(g); err != nil {
		log.Fatal(err)
	}
}

func addStars(g *game.Game) {
	for i := 0; i < 300; i++ {
		g.Stars = append(g.Stars, *game.NewStar())
	}
}
