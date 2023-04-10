package engine

import "github.com/go-gl/mathgl/mgl32"

type Material struct {
    Ambient       mgl32.Vec3
    Diffuse       mgl32.Vec3
    Specular      mgl32.Vec3
    Shininess     float32
    TextureHandle uint32

    shader *ShaderProgram
}

func NewDefaultMaterial() *Material {
    // A default material using the default shader
    return &Material{
        Ambient:       mgl32.Vec3{0.1, 0.1, 0.1},
        Diffuse:       mgl32.Vec3{0.5, 0.5, 0.5},
        Specular:      mgl32.Vec3{0.5, 0.5, 0.5},
        Shininess:     32.0,
        TextureHandle: 0,
        shader:        LoadShader("shaders/default.vert", "shaders/default.frag"),
    }

}

func (m *Material) GetShader() *ShaderProgram {
    return m.shader
}

func (m *Material) BindShaderProperties(shader *ShaderProgram) error {
    // TODO
    return nil
}

func (m *Material) GetAttributeMap() map[uint32]string {
    // TODO
    return map[uint32]string{}
}
