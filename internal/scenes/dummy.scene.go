package scenes

import (
	"github.com/MarcelArt/raylibf/pkg/engine/scene"
	rl "github.com/gen2brain/raylib-go/raylib"
)

type DummyScene struct {
}

// GetID implements [scene.IScene].
func (s *DummyScene) GetID() string {
	return Dummy
}

// Draw implements [scene.IScene].
func (s *DummyScene) Draw() {
	rl.DrawText("Dummy 1", 190, 200, 20, rl.LightGray)
}

// Update implements [scene.IScene].
func (s *DummyScene) Update() scene.SceneResult {
	var result scene.SceneResult

	if rl.IsKeyPressed(rl.KeySpace) {
		result.NextScene = &Dummy2Scene{}
		return result
	}

	return result
}

var _ scene.IScene = &DummyScene{}
