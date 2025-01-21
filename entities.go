package main

// Defalt entity
type Entity struct {
	Playable *Playable
	Sprite *Sprite 
	Position *Position
	Velocity *Velocity
	Acceleration *Acceleration
	Orientation *Orientation
	BulletImage *Sprite
	Cooldown float64
}

// Holds bullet information
type Bullet struct {
	Position *Position
	Velocity *Velocity
	LifeTime float64 // Remaining time for the bullet to exist
	Sprite   *Sprite
}
