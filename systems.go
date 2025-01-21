package main

import (
	"math"
	"github.com/hajimehoshi/ebiten/v2"
)

func (g *Game) PlayableSystem(e *Entity) {
	adjustedAngle := e.Orientation.Angle - math.Pi/2 // Rotate due to sprite image being used...

	// Update velocity based on key presses
	switch {
	case ebiten.IsKeyPressed(ebiten.KeyA): // Turn left
		e.Orientation.Angle -= e.Orientation.RotationSpeed

	case ebiten.IsKeyPressed(ebiten.KeyD): // Turn right
		e.Orientation.Angle += e.Orientation.RotationSpeed

	case ebiten.IsKeyPressed(ebiten.KeyW): // Go forward
		e.Velocity.DY += e.Acceleration.A * math.Sin(adjustedAngle)
		e.Velocity.DX += e.Acceleration.A * math.Cos(adjustedAngle)
		e.Velocity.ClampVelocity()

	case ebiten.IsKeyPressed(ebiten.KeyS): // Go backward
		e.Velocity.DY -= e.Acceleration.A * math.Sin(adjustedAngle)
		e.Velocity.DX -= e.Acceleration.A * math.Cos(adjustedAngle)
		e.Velocity.ClampVelocity()

	default: // Apply deceleration
		e.Velocity.ApplyDecelleration(e.Acceleration.DA)
	}

	// Update position using velocity
	if e.Position != nil && e.Velocity != nil {
		e.Position.X += e.Velocity.DX
		e.Position.Y += e.Velocity.DY
	}

	// Spawn a bullet if the spacebar is pressed
	if e.Cooldown <= 0 && ebiten.IsKeyPressed(ebiten.KeySpace) {
		bulletSpeed := 5.0 // Speed of the bullet
		bullet := &Bullet{
			Position: &Position{
				X: e.Position.X,
				Y: e.Position.Y,
			},
			Velocity: &Velocity{
				DX: bulletSpeed * math.Cos(adjustedAngle),
				DY: bulletSpeed * math.Sin(adjustedAngle),
			},
			LifeTime: 3.0,
			Sprite: &Sprite{
				Image: e.BulletImage.Image, // Use a smaller version of the player's sprite or a new sprite
			},
		}
		g.bullets = append(g.bullets, bullet)
		e.Cooldown = 0.5 // 500ms cooldown
	}
	// Reduce bullet cooldown timer
	if e.Cooldown > 0 {
		e.Cooldown -= 1.0 / ebiten.ActualFPS()
	}
}
