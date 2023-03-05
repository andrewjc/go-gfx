package engine

type Material interface {
	GetShader() *ShaderProgram
	GetAttributeMap() map[string]string
	BindShaderProperties(shader *ShaderProgram) error
}

type BasicMaterial struct {
	Shader *ShaderProgram
}

func NewBasicMaterial() *BasicMaterial {
	return &BasicMaterial{}
}

var DefaultBasicShaderProgram = LoadShader("shaders/default.vert", "shaders/default.frag")

func (m *BasicMaterial) GetShader() *ShaderProgram {
	if m.Shader == nil {
		return DefaultBasicShaderProgram
	}
	return m.Shader
}

func (m *BasicMaterial) GetAttributeMap() map[string]string {
	return map[string]string{
		"Position": "position",
		"Normal":   "normal",
		"TexCoord": "texCoord",
	}
}

func (m *BasicMaterial) BindShaderProperties(shader *ShaderProgram) error {
	/*	shader.SetFloat("albedoColor", m.AlbedoColor[0])
		shader.SetFloat("metallic", m.Metallic)
		shader.SetFloat("roughness", m.Roughness)*/
	return nil
}
