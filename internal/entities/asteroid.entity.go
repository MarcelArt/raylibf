package entities

import (
	"math"

	"github.com/MarcelArt/raylibf/pkg/engine"
	rl "github.com/gen2brain/raylib-go/raylib"
)

const (
	asteroidDefaultRadius = float32(40)
	asteroidVertices      = 10
)

type AsteroidEntity struct {
	engine.Entity
	Position rl.Vector2
	Velocity rl.Vector2
	Health   uint

	Radius   float32
	Rotation float32
	// shape holds a per-vertex radius multiplier in (0, 1] so every asteroid
	// gets its own jagged silhouette. Generated once, lazily.
	shape []float32
}

func (e *AsteroidEntity) Draw() {
	rl.DrawLineStrip(e.worldOutline(), rl.White)
}

func (e *AsteroidEntity) Update(dt float32) {
	e.Position = rl.Vector2Add(e.Position, rl.Vector2Scale(e.Velocity, dt))
}

// worldOutline returns the asteroid's jagged silhouette in world space,
// with the first point repeated so DrawLineStrip renders a closed loop.
func (e *AsteroidEntity) worldOutline() []rl.Vector2 {
	e.ensureShape()

	points := make([]rl.Vector2, 0, asteroidVertices+1)
	for i := 0; i <= asteroidVertices; i++ {
		angle := e.Rotation + 2*math.Pi*float32(i%asteroidVertices)/asteroidVertices
		r := e.Radius * e.shape[i%asteroidVertices]
		points = append(points, rl.NewVector2(
			e.Position.X+float32(math.Sin(float64(angle)))*r,
			e.Position.Y+float32(math.Cos(float64(angle)))*r,
		))
	}
	return points
}

// ensureShape lazily picks a default radius and rolls the per-vertex
// jaggedness, so asteroids created as plain literals still render.
func (e *AsteroidEntity) ensureShape() {
	if e.Radius <= 0 {
		e.Radius = asteroidDefaultRadius
	}
	if e.shape == nil {
		e.shape = make([]float32, asteroidVertices)
		for i := range e.shape {
			e.shape[i] = 0.7 + float32(rl.GetRandomValue(0, 30))/100 // 0.7..1.0
		}
	}
}
