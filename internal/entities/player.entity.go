package entities

import (
	"log"

	rl "github.com/gen2brain/raylib-go/raylib"

	"github.com/MarcelArt/raylibf/pkg/engine"
)

const (
	PlayerCollisionRadius = size * 0.7

	playerRotationSpeed = float32(4)
	playerThrust        = float32(350) // px/s² — how hard the engine pushes
	playerDrag          = float32(0.5) // fraction of velocity lost per second
	playerMaxSpeed      = float32(400) // px/s cap
	size                = float32(20)
)

type PlayerCommand struct {
	IsShooting bool
}

type PlayerEntity struct {
	engine.Entity
	Position rl.Vector2
	Rotation float32
	Velocity rl.Vector2
}

func (e *PlayerEntity) Draw() {
	e.drawPlayer()
}

func (e *PlayerEntity) Update(dt float32) PlayerCommand {
	var cmd PlayerCommand

	e.playerMovement(dt)
	e.shoot(&cmd)

	return cmd
}

func (e *PlayerEntity) Nose() rl.Vector2 {
	return e.localToWorld(rl.NewVector2(0, -size))
}

func (e *PlayerEntity) localToWorld(v rl.Vector2) rl.Vector2 {
	return rl.Vector2Add(e.Position, rl.Vector2Rotate(v, e.Rotation))
}

// drawPlayer draws the classic Asteroids ship: a wireframe triangle pointing up.
func (e *PlayerEntity) drawPlayer() {
	nose := rl.NewVector2(0, -size)
	l := rl.NewVector2(-size*0.6, size)
	r := rl.NewVector2(size*0.6, size)

	// Rotate around origin, then translate into world space
	// world := func(v rl.Vector2) rl.Vector2 {
	// 	v = rl.Vector2Rotate(v, e.Rotation)
	// 	return rl.NewVector2(e.Position.X+v.X, e.Position.Y+v.Y)
	// }

	rl.DrawTriangleLines(e.localToWorld(nose), e.localToWorld(l), e.localToWorld(r), rl.White)

}

func (e *PlayerEntity) playerMovement(dt float32) {
	if rl.IsKeyDown(rl.KeyRight) || rl.IsKeyDown(rl.KeyD) {
		e.Rotation += playerRotationSpeed * dt
	}
	if rl.IsKeyDown(rl.KeyLeft) || rl.IsKeyDown(rl.KeyA) {
		e.Rotation -= playerRotationSpeed * dt
	}

	if rl.IsKeyDown(rl.KeyUp) || rl.IsKeyDown(rl.KeyW) {
		dir := rl.Vector2Rotate(rl.NewVector2(0, -1), e.Rotation)
		e.Velocity = rl.Vector2Add(e.Velocity, rl.Vector2Scale(dir, playerThrust*dt))
	}

	// --- 2. drag: exponential decay (see tuning notes) ---
	e.Velocity = rl.Vector2Scale(e.Velocity, 1-playerDrag*dt)

	// --- 3. clamp to max speed ---
	if speed := rl.Vector2Length(e.Velocity); speed > playerMaxSpeed {
		e.Velocity = rl.Vector2Scale(e.Velocity, playerMaxSpeed/speed)
	}

	// --- 4. integrate: position += velocity * dt ---
	e.Position = rl.Vector2Add(e.Position, rl.Vector2Scale(e.Velocity, dt))
}

func (e *PlayerEntity) shoot(cmd *PlayerCommand) {
	cmd.IsShooting = false
	if rl.IsKeyPressed(rl.KeySpace) {
		cmd.IsShooting = true
		log.Println("shot")
	}
}
