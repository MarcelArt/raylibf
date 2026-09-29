package entities

import (
	"github.com/MarcelArt/raylibf/pkg/engine"
	rl "github.com/gen2brain/raylib-go/raylib"
)

const (
	bulletSpeed  float32 = 300
	BulletRadius float32 = 3
)

type BulletEntity struct {
	engine.Entity
	// Velocity rl.Vector2
	Position rl.Vector2
	Rotation float32
}

func (e *BulletEntity) Draw() {
	rl.DrawCircleV(e.Position, BulletRadius, rl.White)
}

func (e *BulletEntity) Update(dt float32) {
	dir := rl.Vector2Rotate(rl.NewVector2(0, -1), e.Rotation)
	e.Position = rl.Vector2Add(e.Position, rl.Vector2Scale(dir, bulletSpeed*dt))
}

// var _ engine.IEntity = &BulletEntity{}
