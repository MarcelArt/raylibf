package scenes

import (
	"math/rand/v2"

	"github.com/MarcelArt/raylibf/internal/entities"
	"github.com/MarcelArt/raylibf/pkg/engine"
	rl "github.com/gen2brain/raylib-go/raylib"
)

const (
	asteroidSpawnRate = 2
)

type LevelScene struct {
	Player                   *entities.PlayerEntity
	asteroidSpawnTimeCounter float32

	bullets   []*entities.BulletEntity
	asteroids []*entities.AsteroidEntity
}

func (s *LevelScene) Draw() {
	rl.ClearBackground(rl.Black)

	s.Player.Draw()

	for _, bullet := range s.bullets {
		bullet.Draw()
	}

	for _, asteroid := range s.asteroids {
		asteroid.Draw()
	}
}

func (s *LevelScene) GetID() string {
	return Level
}

func (s *LevelScene) Update() engine.SceneResult {
	var res engine.SceneResult
	dt := rl.GetFrameTime()

	playerCMD := s.Player.Update(dt)
	s.handlePlayerCMD(playerCMD)

	for _, bullet := range s.bullets {
		bullet.Update(dt)
	}

	s.spawnAsteroid(dt)
	for _, asteroid := range s.asteroids {
		asteroid.Update(dt)
	}

	return res
}

func (s *LevelScene) handlePlayerCMD(cmd entities.PlayerCommand) {
	if cmd.IsShooting {
		bullet := &entities.BulletEntity{
			Position: s.Player.Nose(),
			IsActive: true,
			Rotation: s.Player.Rotation,
		}

		s.bullets = append(s.bullets, bullet)
	}

}

func (s *LevelScene) spawnAsteroid(dt float32) {
	if s.asteroidSpawnTimeCounter < asteroidSpawnRate {
		s.asteroidSpawnTimeCounter += dt
		return
	}

	minX := float32(-10)
	minY := float32(-10)
	maxX := float32(rl.GetScreenWidth() + 10)
	maxY := float32(rl.GetScreenHeight() + 10)

	var spawnPoints = []rl.Vector2{
		{X: minX, Y: minY},
		{X: minX, Y: maxY},
		{X: maxX, Y: minY},
		{X: maxX, Y: maxY},
	}

	i := rand.IntN(len(spawnPoints))

	asteroidSpawn := spawnPoints[i]

	direction := rl.Vector2Subtract(s.Player.Position, asteroidSpawn)
	direction = rl.Vector2Normalize(direction)

	s.asteroids = append(s.asteroids, &entities.AsteroidEntity{
		Position: asteroidSpawn,
		IsActive: true,
		Velocity: rl.Vector2Scale(direction, 80),
		Health:   1,
	})
	s.asteroidSpawnTimeCounter = 0
}

var _ engine.IScene = &LevelScene{}
