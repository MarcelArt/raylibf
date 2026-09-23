package entities

import (
	rl "github.com/gen2brain/raylib-go/raylib"

	"github.com/MarcelArt/raylibf/pkg/engine"
)

const (
	playerRotationSpeed = float32(4)
	playerThrust        = float32(350) // px/s² — how hard the engine pushes
	playerDrag          = float32(0.5) // fraction of velocity lost per second
	playerMaxSpeed      = float32(400) // px/s cap
)

type PlayerEntity struct {
	engine.Entity
	Position rl.Vector2
	Rotation float32
	Velocity rl.Vector2
}

// Draw implements [engine.IEntity].
func (p *PlayerEntity) Draw() {
	p.drawPlayer()
}

// Update implements [engine.IEntity].
func (p *PlayerEntity) Update() {
	dt := rl.GetFrameTime()

	p.playerMovement(dt)
}

// drawPlayer draws the classic Asteroids ship: a wireframe triangle pointing up.
func (p *PlayerEntity) drawPlayer() {
	const size = float32(20)

	nose := rl.NewVector2(0, -size)
	l := rl.NewVector2(-size*0.6, size)
	r := rl.NewVector2(size*0.6, size)

	// Rotate around origin, then translate into world space
	world := func(v rl.Vector2) rl.Vector2 {
		v = rl.Vector2Rotate(v, p.Rotation)
		return rl.NewVector2(p.Position.X+v.X, p.Position.Y+v.Y)
	}

	rl.DrawTriangleLines(world(nose), world(l), world(r), rl.White)

}

func (p *PlayerEntity) playerMovement(dt float32) {
	if rl.IsKeyDown(rl.KeyRight) || rl.IsKeyDown(rl.KeyD) {
		p.Rotation += playerRotationSpeed * dt
	}
	if rl.IsKeyDown(rl.KeyLeft) || rl.IsKeyDown(rl.KeyA) {
		p.Rotation -= playerRotationSpeed * dt
	}

	if rl.IsKeyDown(rl.KeyUp) || rl.IsKeyDown(rl.KeyW) {
		dir := rl.Vector2Rotate(rl.NewVector2(0, -1), p.Rotation)
		p.Velocity = rl.Vector2Add(p.Velocity, rl.Vector2Scale(dir, playerThrust*dt))
	}

	// --- 2. drag: exponential decay (see tuning notes) ---
	p.Velocity = rl.Vector2Scale(p.Velocity, 1-playerDrag*dt)

	// --- 3. clamp to max speed ---
	if speed := rl.Vector2Length(p.Velocity); speed > playerMaxSpeed {
		p.Velocity = rl.Vector2Scale(p.Velocity, playerMaxSpeed/speed)
	}

	// --- 4. integrate: position += velocity * dt ---
	p.Position = rl.Vector2Add(p.Position, rl.Vector2Scale(p.Velocity, dt))
}

var _ engine.IEntity = &PlayerEntity{}
