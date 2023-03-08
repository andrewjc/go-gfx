package engine

import (
	"github.com/go-gl/mathgl/mgl32"
)

type GameObject struct {
	Position mgl32.Vec3
	Rotation mgl32.Quat
	Velocity mgl32.Vec3
	Scale    float32
	Mass     float32
	Charge   float32
	Fields   []Force
	Material Material
	Mesh     *Mesh
	Scene    *Scene

	Renderer ObjectRenderer
}

func (g *GameObject) Equals(other *GameObject) bool {
	return g.Position == (other.Position) &&
		g.Rotation.OrientationEqual(other.Rotation) &&
		g.Velocity == (other.Velocity) &&
		g.Mass == other.Mass &&
		g.Charge == other.Charge
}

func (g *GameObject) Update(dt float32) {
	// update position and velocity based on forces
	for _, f := range g.Fields {
		g.Velocity = g.Velocity.Add(f.GetForce(g, dt).Mul(dt / g.Mass))
	}
	g.Position = g.Position.Add(g.Velocity.Mul(dt))
}

func (g *GameObject) Render(projMatrix mgl32.Mat4, viewMatrix mgl32.Mat4) {
	modelMatrix := g.getModelMatrix()

	g.Renderer.Render(g.Mesh, g.Material, modelMatrix, projMatrix, viewMatrix)

}

func (g *GameObject) getModelMatrix() mgl32.Mat4 {
	modelMatrix := mgl32.Translate3D(g.Position.X(), g.Position.Y(), g.Position.Z())
	modelMatrix = modelMatrix.Mul4(mgl32.Scale3D(g.Scale, g.Scale, g.Scale))
	modelMatrix = modelMatrix.Mul4(g.Rotation.Mat4())
	return modelMatrix
}
