package engine

import (
	"github.com/go-gl/mathgl/mgl32"
)

type GameObject struct {
	ObjectRenderer
	Position mgl32.Vec3
	Rotation mgl32.Quat
	Velocity mgl32.Vec3
	Scale    float32
	Mass     float32
	Charge   float32
	Fields   []Force
	Material Material
	Mesh     Mesh

	Scene *Scene
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

func (g *GameObject) Render(camera *Camera) {
	// Set the model matrix based on the object's position and orientation
	translation := mgl32.Translate3D(g.Position.X(), g.Position.Y(), g.Position.Z())
	rotation := g.Rotation.Mat4()
	scale := mgl32.Scale3D(g.Scale, g.Scale, g.Scale)
	model := translation.Mul4(rotation).Mul4(scale)

	width, height := camera.Window.GetFramebufferSize()

	var shader *ShaderProgram = nil
	if g.Material != nil && g.Material.GetShader() != nil {
		shader = g.Material.GetShader()
	} else {
		shader = g.Scene.DefaultShaderProgram
	}

	shader.Use()

	projectionMatrix := camera.ProjectionMatrix(float32(width), float32(height))
	viewMatrix := camera.ViewMatrix()

	shader.SetMat4UniformLocation("model", &model)
	shader.SetMat4UniformLocation("projection", &projectionMatrix)
	shader.SetMat4UniformLocation("view", &viewMatrix)

	e := g.Material.BindShaderProperties(shader)
	if e != nil {
		panic(e)
	}

	// Draw the mesh
	g.Mesh.Draw(shader)

	shader.Unuse()
}
