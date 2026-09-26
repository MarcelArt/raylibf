package scenes

import (
	"github.com/MarcelArt/raylibf/internal/entities"
	"github.com/MarcelArt/raylibf/pkg/engine"
	rl "github.com/gen2brain/raylib-go/raylib"
)

type LevelScene struct {
	Player *entities.PlayerEntity

	bullets []*entities.BulletEntity
}

// Draw implements [scene.IScene].
func (l *LevelScene) Draw() {
	rl.ClearBackground(rl.Black)

	l.Player.Draw()

	for _, bullet := range l.bullets {
		bullet.Draw()
	}
}

// GetID implements [scene.IScene].
func (l *LevelScene) GetID() string {
	return Level
}

// Update implements [scene.IScene].
func (l *LevelScene) Update() engine.SceneResult {
	var res engine.SceneResult
	dt := rl.GetFrameTime()

	playerCMD := l.Player.Update(dt)
	l.handlePlayerCMD(playerCMD)

	for _, bullet := range l.bullets {
		bullet.Update(dt)
	}

	return res
}

func (l *LevelScene) handlePlayerCMD(cmd entities.PlayerCommand) {
	if cmd.IsShooting {
		bullet := &entities.BulletEntity{
			Position: l.Player.Nose(),
			IsActive: true,
			Rotation: l.Player.Rotation,
		}

		l.bullets = append(l.bullets, bullet)
	}

}

var _ engine.IScene = &LevelScene{}
