package engine

// Field represents a Field that applies a force to particles within its Radius
type Force interface {
	GetForce(particle *GameObject, dt float64) Vec3
}

type GravityForce struct {
	Acceleration Vec3
}

func (f *GravityForce) GetForce(gameObject *GameObject, dt float64) Vec3 {
	return f.Acceleration.Multiply(gameObject.Mass)
}

type ElectricForce struct {
	Charge float64
	Field  Vec3
}

func (f *ElectricForce) GetForce(gameObject *GameObject, dt float64) Vec3 {
	return f.Field.Multiply(f.Charge * gameObject.Charge)
}

type SpringForce struct {
	Anchor        Vec3
	Stiffness     float64
	RestingLength float64
}

func (f *SpringForce) GetForce(particle *GameObject, dt float64) Vec3 {
	displacement := particle.Position.Subtract(f.Anchor)
	magnitude := f.Stiffness * (displacement.Magnitude() - f.RestingLength)
	return displacement.Normalize().Multiply(magnitude)
}
