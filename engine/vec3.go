package engine

import (
    "github.com/go-gl/mathgl/mgl32"
    "math"
)

type Vec3 struct {
    X, Y, Z float64
}

func (v Vec3) Add(other Vec3) Vec3 {
    return Vec3{v.X + other.X, v.Y + other.Y, v.Z + other.Z}
}

func (v Vec3) Subtract(other Vec3) Vec3 {
    return Vec3{v.X - other.X, v.Y - other.Y, v.Z - other.Z}
}

func (v Vec3) Multiply(scalar float64) Vec3 {
    return Vec3{v.X * scalar, v.Y * scalar, v.Z * scalar}
}

func (v Vec3) Divide(scalar float64) Vec3 {
    return Vec3{v.X / scalar, v.Y / scalar, v.Z / scalar}
}

func (v Vec3) Dot(other Vec3) float64 {
    return v.X*other.X + v.Y*other.Y + v.Z*other.Z
}

func (v Vec3) Cross(other Vec3) Vec3 {
    return Vec3{v.Y*other.Z - v.Z*other.Y, v.Z*other.X - v.X*other.Z, v.X*other.Y - v.Y*other.X}
}

func (v Vec3) Magnitude() float64 {
    return math.Sqrt(v.X*v.X + v.Y*v.Y + v.Z*v.Z)
}

func (v Vec3) Normalize() Vec3 {
    mag := v.Magnitude()
    if mag == 0 {
        return v
    }
    return v.Divide(mag)
}

func (v Vec3) ToMGl32() mgl32.Vec3 {
    return mgl32.Vec3{float32(v.X), float32(v.Y), float32(v.Z)}
}

func (v Vec3) Equals(other Vec3) bool {
    return v.X == other.X && v.Y == other.Y && v.Z == other.Z
}
