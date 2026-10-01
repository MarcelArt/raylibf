package scenes

import (
	"encoding/json/v2"
	"math/rand/v2"
	"slices"
	"strconv"

	"github.com/MarcelArt/raylibf/internal/entities"
	"github.com/MarcelArt/raylibf/internal/models"
	"github.com/MarcelArt/raylibf/pkg/engine"
	rl "github.com/gen2brain/raylib-go/raylib"
)

const (
	asteroidSpawnRate = 2
	saveFileDir       = "asteroid/save.json"
)

type gameState uint8

const (
	play gameState = iota
	gameOver
)

type LevelScene struct {
	Player *entities.PlayerEntity

	asteroidSpawnTimeCounter float32
	bullets                  []*entities.BulletEntity
	asteroids                []*entities.AsteroidEntity
	score                    uint
	gameState                gameState
	highScore                uint
}

func NewLevelScene(player *entities.PlayerEntity) *LevelScene {
	var highScore uint
	saveFileData, err := engine.LoadGame(saveFileDir)
	if err == nil {
		var saveFile models.SaveFile
		json.Unmarshal(saveFileData, &saveFile)
		highScore = saveFile.HighScore
	}

	return &LevelScene{
		Player: player,

		asteroidSpawnTimeCounter: 0,
		bullets:                  make([]*entities.BulletEntity, 0),
		asteroids:                make([]*entities.AsteroidEntity, 0),
		score:                    0,
		gameState:                play,
		highScore:                highScore,
	}
}

func (s *LevelScene) Draw() {
	rl.ClearBackground(rl.Black)

	if s.gameState == play {
		s.Player.Draw()

		for _, bullet := range s.bullets {
			bullet.Draw()
		}

		for _, asteroid := range s.asteroids {
			asteroid.Draw()
		}
		s.drawScore()
	} else {
		width := rl.GetScreenWidth()
		height := rl.GetScreenHeight()
		rl.DrawText("Game Over", int32(width/2-100), int32(height/2-50), 36, rl.White)
		rl.DrawText("Press R to restart", int32(width/2-150), int32(height/2), 24, rl.White)
	}

}

func (s *LevelScene) GetID() string {
	return Level
}

func (s *LevelScene) Update() engine.SceneResult {
	var res engine.SceneResult
	dt := rl.GetFrameTime()

	if s.gameState == play {
		playerCMD := s.Player.Update(dt)
		s.handlePlayerCMD(playerCMD)

		for _, bullet := range s.bullets {
			bullet.Update(dt)
		}

		s.spawnAsteroid(dt)
		for _, asteroid := range s.asteroids {
			asteroid.Update(dt)
		}

		s.handleCollisions()
	} else {
		if rl.IsKeyPressed(rl.KeyR) {
			s.restart()
		}
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
	if s.Player == nil {
		return
	}

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

func (s *LevelScene) checkBulletAsteroidCollisions() {
	for _, bullet := range s.bullets {
		if !bullet.IsActive {
			continue
		}
		for _, asteroid := range s.asteroids {
			if !asteroid.IsActive {
				continue
			}
			if rl.CheckCollisionCircles(bullet.Position, entities.BulletRadius, asteroid.Position, asteroid.Radius) {
				bullet.IsActive = false
				asteroid.IsActive = false
				s.score++
				break
			}
		}
	}

	s.bullets = slices.DeleteFunc(s.bullets, func(b *entities.BulletEntity) bool {
		return !b.IsActive
	})
	s.asteroids = slices.DeleteFunc(s.asteroids, func(a *entities.AsteroidEntity) bool {
		return !a.IsActive
	})
}

func (s *LevelScene) checkPlayerAsteroidCollisions() {
	for _, asteroid := range s.asteroids {
		if rl.CheckCollisionCircles(asteroid.Position, asteroid.Radius, s.Player.Position, entities.PlayerCollisionRadius) {
			s.loseGame()
			return
		}
	}
}

func (s *LevelScene) handleCollisions() {
	s.checkBulletAsteroidCollisions()
	s.checkPlayerAsteroidCollisions()
}

func (s *LevelScene) drawScore() {
	width := rl.GetScreenWidth()
	// height := rl.GetScreenHeight()

	score := strconv.Itoa(int(s.score))

	rl.DrawText(score, int32(width)/2, 0, 36, rl.White)
}

func (s *LevelScene) loseGame() {
	s.gameState = gameOver
	s.Player = nil
	s.asteroids = make([]*entities.AsteroidEntity, 0)
	s.bullets = make([]*entities.BulletEntity, 0)

	if s.score > s.highScore {
		saveFile := models.SaveFile{
			HighScore: s.score,
		}
		saveFileData, err := json.Marshal(saveFile)
		if err == nil {
			engine.SaveGame(saveFileDir, saveFileData)
		}
		s.highScore = s.score
	}
}

func (s *LevelScene) restart() {
	width := rl.GetScreenWidth()
	height := rl.GetScreenHeight()

	s.gameState = play
	s.score = 0
	s.Player = &entities.PlayerEntity{
		Position: rl.NewVector2(float32(width)/2, float32(height)/2),
		IsActive: true,
	}
	s.asteroids = make([]*entities.AsteroidEntity, 0)
	s.bullets = make([]*entities.BulletEntity, 0)
}

var _ engine.IScene = &LevelScene{}
