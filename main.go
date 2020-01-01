package main

import (
	"fmt"
	"goray/tuple"
)

type Projectile struct {
	position tuple.Tuple
	velocity tuple.Tuple
}

type Environment struct {
	gravity tuple.Tuple
	wind    tuple.Tuple
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
		position: tuple.CreatePoint(0, 1, 0),
		velocity: tuple.CreateVector(1, 1, 0).Normalize(),
	}
	e := Environment{
		gravity: tuple.CreateVector(0, -0.1, 0),
		wind:    tuple.CreateVector(-0.01, 0, 0),
	}
	for p.position.Y >= 0 {
		p = tick(e, p)
		fmt.Printf("x: %f, y: %f\n", p.position.X, p.position.Y)
	}
}
