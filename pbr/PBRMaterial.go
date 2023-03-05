package pbr

import (
	"github.com/go-gl/mathgl/mgl32"
	. "physics/engine"
)

type PBRMaterial struct {
	Shader *ShaderProgram
	//texture       *gl.Texture
	AlbedoColor mgl32.Vec3
	Metallic    float32
	Roughness   float32
}

var DefaultPbrShaderProgram = LoadShader("shaders/pbr.vert", "shaders/pbr.frag")

var DefaultAlbedoColor = mgl32.Vec3{0.5, 0.0, 0.0}
var DefaultMetallic = float32(0.0)
var DefaultRoughness = float32(0.5)

func NewPBRMaterial() *PBRMaterial {
	return &PBRMaterial{
		//texture:       texture,
		Shader:      DefaultPbrShaderProgram,
		AlbedoColor: DefaultAlbedoColor,
		Metallic:    DefaultMetallic,
		Roughness:   DefaultRoughness,
	}
}

func (m *PBRMaterial) GetShader() *ShaderProgram {
	if m.Shader == nil {
		return DefaultPbrShaderProgram
	}
	return m.Shader
}

func (m *PBRMaterial) GetAttributeMap() map[string]string {
	return map[string]string{
		"Position": "position",
		"Normal":   "normal",
		"TexCoord": "texCoord",
	}
}

func (m *PBRMaterial) BindShaderProperties(shader *ShaderProgram) error {
	shader.SetFloat("albedoColor", m.AlbedoColor[0])
	shader.SetFloat("metallic", m.Metallic)
	shader.SetFloat("roughness", m.Roughness)
	return nil
}
