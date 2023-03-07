package engine

import "github.com/go-gl/mathgl/mgl32"

// Field represents a Field that applies a force to particles within its Radius
type Force interface {
	GetForce(particle *GameObject, dt float32) mgl32.Vec3
}

type GravityForce struct {
	Acceleration mgl32.Vec3
}

func (f *GravityForce) GetForce(gameObject *GameObject, dt float32) mgl32.Vec3 {
	return f.Acceleration.Mul(gameObject.Mass)
}

type ElectricForce struct {
	Charge float32
	Field  mgl32.Vec3
}

func (f *ElectricForce) GetForce(gameObject *GameObject, dt float32) mgl32.Vec3 {
	return f.Field.Mul(f.Charge * gameObject.Charge)
}

type SpringForce struct {
	Anchor        mgl32.Vec3
	Stiffness     float32
	RestingLength float32
}

func (f *SpringForce) GetForce(particle *GameObject, dt float32) mgl32.Vec3 {
	displacement := particle.Position.Sub(f.Anchor)
	magnitude := f.Stiffness * (displacement.Len() - f.RestingLength)
	return displacement.Normalize().Mul(magnitude)
}
