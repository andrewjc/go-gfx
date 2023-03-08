package engine

type Mesh interface {
	Draw(shader *ShaderProgram)
}
