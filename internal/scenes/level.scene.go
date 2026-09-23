package scenes

import (
	"github.com/MarcelArt/raylibf/internal/entities"
	"github.com/MarcelArt/raylibf/pkg/engine"
	rl "github.com/gen2brain/raylib-go/raylib"
)

type LevelScene struct {
	Player *entities.PlayerEntity
}

// Draw implements [scene.IScene].
func (l *LevelScene) Draw() {
	rl.ClearBackground(rl.Black)

	l.Player.Draw()
}

// GetID implements [scene.IScene].
func (l *LevelScene) GetID() string {
	return Level
}

// Update implements [scene.IScene].
func (l *LevelScene) Update() engine.SceneResult {
	var res engine.SceneResult

	l.Player.Update()

	return res
}

var _ engine.IScene = &LevelScene{}
