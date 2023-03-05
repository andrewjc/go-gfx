package engine

import (
    "github.com/go-gl/mathgl/mgl32"
)

type GameObject struct {
    ObjectRenderer
    Position Vec3
    Rotation mgl32.Quat
    Velocity Vec3
    Scale    float32
    Mass     float64
    Charge   float64
    Fields   []Force
    Material Material
    Mesh     *Mesh

    Scene *Scene
}

func (g *GameObject) Equals(other *GameObject) bool {
    return g.Position.Equals(other.Position) &&
        g.Rotation.OrientationEqual(other.Rotation) &&
        g.Velocity.Equals(other.Velocity) &&
        g.Mass == other.Mass &&
        g.Charge == other.Charge
}

func (g *GameObject) Update(dt float64) {
    // update position and velocity based on forces
    for _, f := range g.Fields {
        g.Velocity = g.Velocity.Add(f.GetForce(g, dt).Multiply(dt / g.Mass))
    }
    g.Position = g.Position.Add(g.Velocity.Multiply(dt))
}

func (g *GameObject) Render(camera *Camera) {
    // Set the model matrix based on the object's position and orientation
    translation := mgl32.Translate3D(float32(g.Position.X), float32(g.Position.Y), float32(g.Position.Z))
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

    shader.SetMat4("model", &model)
    shader.SetMat4("projection", &projectionMatrix)
    shader.SetMat4("view", &viewMatrix)

    e := g.Material.BindShaderProperties(shader)
    if e != nil {
        panic(e)
    }

    // Draw the mesh
    g.Mesh.Draw(shader)

    shader.Unuse()
}
