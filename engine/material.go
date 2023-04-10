package engine

import (
	"github.com/go-gl/gl/v4.1-core/gl"
	"github.com/go-gl/mathgl/mgl32"
)

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
	shader.Use()

	// Bind material properties to the shader
	ambientUniform := shader.GetUniformLocation("material.ambient")
	shader.SetUniform3f(ambientUniform, m.Ambient.X(), m.Ambient.Y(), m.Ambient.Z())

	diffuseUniform := shader.GetUniformLocation("material.diffuse")
	shader.SetUniform3f(diffuseUniform, m.Diffuse.X(), m.Diffuse.Y(), m.Diffuse.Z())

	specularUniform := shader.GetUniformLocation("material.specular")
	shader.SetUniform3f(specularUniform, m.Specular.X(), m.Specular.Y(), m.Specular.Z())

	shininessUniform := shader.GetUniformLocation("material.shininess")
	shader.SetUniform1f(shininessUniform, m.Shininess)

	// Bind texture if available
	if m.TextureHandle != 0 {
		gl.ActiveTexture(gl.TEXTURE0)
		gl.BindTexture(gl.TEXTURE_2D, m.TextureHandle)
		textureUniform := shader.GetUniformLocation("material.texture")
		shader.SetUniform1i(textureUniform, 0)
	}

	return nil
}

func (m *Material) GetAttributeMap() map[uint32]string {
	attributeMap := make(map[uint32]string)

	attributeMap[0] = "aPos"      // Vertex positions
	attributeMap[1] = "aTexCoord" // Texture coordinates
	attributeMap[2] = "aNormal"   // Normals

	return attributeMap
}
