package main

import (
	"image/color"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

// Game struct contains entities and logic
type Game struct {
	entities      []*Entity
	bullets       []*Bullet // Keep track of live bullets
}

// Update handles game logic (movement system)
func (g *Game) Update() error {
	// Handle input and update velocity
	for _, e := range g.entities {
		if e.Playable != nil {
			g.PlayableSystem(e)
		}
	}

	// Update bullets
	for i := 0; i < len(g.bullets); {
		b := g.bullets[i]

		// Update position
		b.Position.X += b.Velocity.DX
		b.Position.Y += b.Velocity.DY

		// Reduce lifetime
		b.LifeTime -= 1.0 / ebiten.ActualFPS()

		// Remove bullets that go off-screen or have no lifetime left
		if b.LifeTime <= 0 {
			g.bullets = append(g.bullets[:i], g.bullets[i+1:]...)
		} else {
			i++
		}
	}

	return nil
}


// Draw renders the game screen
func (g *Game) Draw(screen *ebiten.Image) {
	screen.Fill(color.RGBA{0, 0, 0, 255}) // Clear the screen with black

	// Draw entities
	for _, e := range g.entities {
		if e.Playable != nil { // Player draw options
			op := &ebiten.DrawImageOptions{}
			w, h := e.Sprite.Image.Bounds().Dx(), e.Sprite.Image.Bounds().Dy()
			op.GeoM.Translate(-float64(w)/2, -float64(h)/2) // Align center of image to player

			op.GeoM.Rotate(e.Orientation.Angle) // Apply rotation

			op.GeoM.Translate(e.Position.X, e.Position.Y) // Set the position of the image on the screen

			screen.DrawImage(e.Sprite.Image, op) // Draw the image on the screen
		}
	}

	// Draw bullets
	for _, b := range g.bullets {
		op := &ebiten.DrawImageOptions{}
		w, h := b.Sprite.Image.Bounds().Dx(), b.Sprite.Image.Bounds().Dy() // Center and position the bullet
		op.GeoM.Translate(-float64(w)/2, -float64(h)/2) // Center of bullet will be center of image

		op.GeoM.Translate(b.Position.X, b.Position.Y) // Set the position of the image on the screen

		screen.DrawImage(b.Sprite.Image, op)
	}
}

// Layout specifies the game's screen size
func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return 640, 480 // Width x Height
}

func main() {
	// Load the player sprite
	player, _, err := ebitenutil.NewImageFromFile("media/spaceship_normal.png")
	if err != nil {
		log.Fatal(err)
	}
	// Load the bullet sprite
	bullet, _, err := ebitenutil.NewImageFromFile("media/bullet.png")
	if err != nil {
		log.Fatal(err)
	}

	// Initialize the game
	game := &Game{
		entities: []*Entity{
			{
				Position:     &Position{X: 160, Y: 120}, // Initial position in the center of the screen
				Velocity:     &Velocity{DX: 0, DY: 0, MaxSpeed: 3.0, StopThreshold: 0.01},
				Acceleration: &Acceleration{A: 0.1, DA: 0.05},
				Playable:     &Playable{},
				Orientation:  &Orientation{Angle: 0.0, RotationSpeed: 0.1},
				Sprite:       &Sprite{Image: player},
				BulletImage:  &Sprite{Image: bullet},
			},
		},
		bullets: []*Bullet{},
	}

	// Start the game loop
	ebiten.SetWindowSize(1280, 960)    // Window size
	ebiten.SetWindowTitle("Astroids") // Window title
	if err := ebiten.RunGame(game); err != nil {
		panic(err)
	}
}
