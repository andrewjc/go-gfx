package engine

import "math"

type Vec2 struct {
	X float64
	Y float64
}

func (v Vec2) Add(other Vec2) Vec2 {
	return Vec2{v.X + other.X, v.Y + other.Y}
}

func (v Vec2) Subtract(other Vec2) Vec2 {
	return Vec2{v.X - other.X, v.Y - other.Y}
}

func (v Vec2) Multiply(scalar float64) Vec2 {
	return Vec2{v.X * scalar, v.Y * scalar}
}

func (v Vec2) Magnitude() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func (v Vec2) Normalize() Vec2 {
	magnitude := v.Magnitude()
	if magnitude == 0 {
		return Vec2{0, 0}
	}
	return Vec2{v.X / magnitude, v.Y / magnitude}
}
