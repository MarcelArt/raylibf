package scenes

import (
	"github.com/MarcelArt/raylibf/pkg/engine/scene"
	rl "github.com/gen2brain/raylib-go/raylib"
)

type Dummy2Scene struct {
}

func (s *Dummy2Scene) GetID() string {
	return Dummy2
}

// Draw implements [scene.IScene].
func (s *Dummy2Scene) Draw() {
	rl.DrawText("Dummy 2", 190, 200, 20, rl.LightGray)
}

// Update implements [scene.IScene].
func (s *Dummy2Scene) Update() scene.SceneResult {
	var result scene.SceneResult
	if rl.IsKeyPressed(rl.KeySpace) {
		result.NextScene = &DummyScene{}
		return result
	}

	return result
}

var _ scene.IScene = &Dummy2Scene{}
