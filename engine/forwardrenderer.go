package engine

import (
	"github.com/go-gl/gl/v4.1-core/gl"
	"github.com/go-gl/glfw/v3.3/glfw"
	"github.com/go-gl/mathgl/mgl32"
)

type ForwardRenderer struct {
	window *glfw.Window
	camera *Camera
}

func NewForwardRenderer(window *glfw.Window, camera *Camera) ForwardRenderer {
	return ForwardRenderer{
		window: window,
		camera: camera,
	}
}

func (fr *ForwardRenderer) Render(mesh *ComplexMesh, material Material, model mgl32.Mat4, proj mgl32.Mat4, view mgl32.Mat4) {
	shaderProgram := material.GetShader()

	shaderProgram.Use()

	// Bind the material properties to the shader program
	if err := material.BindShaderProperties(shaderProgram); err != nil {
		panic(err)
	}

	// Set the model, view, and projection matrices for the shader program
	shaderProgram.SetMat4UniformLocation("model", &model)
	shaderProgram.SetMat4UniformLocation("view", &view)
	shaderProgram.SetMat4UniformLocation("projection", &proj)

	// Bind the mesh data
	gl.BindVertexArray(mesh.Vao)
	gl.DrawElements(gl.TRIANGLES, mesh.VertexCount, gl.UNSIGNED_INT, nil)
	gl.BindVertexArray(0)

	shaderProgram.Unuse()
}
