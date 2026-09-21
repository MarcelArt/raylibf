package main

import (
	"github.com/MarcelArt/raylibf/internal/scenes"
	"github.com/MarcelArt/raylibf/pkg/engine/game"
)

func main() {
	g := game.New("Raylibf", 800, 540, 60)
	g.SetActiveScene(&scenes.DummyScene{})
	g.Start()
}
