package main

import (
	"fmt"
)

type Projectile struct {
	position Tuple
	velocity Tuple
}

type Environment struct {
	gravity Tuple
	wind    Tuple
}

func tick(env Environment, proj Projectile) Projectile {
	position := proj.position.Add(proj.velocity)
	velocity := proj.velocity.Add(env.gravity).Add(env.wind)
	result := Projectile{
		position,
		velocity,
	}
	return result
}

func main() {
	p := Projectile{
		position: CreatePoint(0, 1, 0),
		velocity: CreateVector(1, 1, 0).Normalize(),
	}
	e := Environment{
		gravity: CreateVector(0, -0.1, 0),
		wind:    CreateVector(-0.01, 0, 0),
	}
	for p.position.y >= 0 {
		p = tick(e, p)
		fmt.Printf("x: %f, y: %f\n", p.position.x, p.position.y)
	}
}
