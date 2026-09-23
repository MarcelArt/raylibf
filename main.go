package main

import (
	"github.com/MarcelArt/raylibf/internal/entities"
	"github.com/MarcelArt/raylibf/internal/scenes"
	"github.com/MarcelArt/raylibf/pkg/engine"
	rl "github.com/gen2brain/raylib-go/raylib"
)

func main() {
	g := engine.NewGame("Raylibf", 800, 540, 60)
	g.SetActiveScene(&scenes.LevelScene{
		Player: &entities.PlayerEntity{
			Position: rl.NewVector2(float32(g.Width)/2, float32(g.Height)/2),
			IsActive: true,
		},
	})
	g.Start()
}
