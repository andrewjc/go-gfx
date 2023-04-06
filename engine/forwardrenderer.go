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

func NewForwardRenderer(window *glfw.Window, camera *Camera) *ForwardRenderer {
	return &ForwardRenderer{
		window: window,
		camera: camera,
	}
}

func (r *ForwardRenderer) RenderObject(obj ObjectRenderer, mesh *ComplexMesh, material Material, model mgl32.Mat4, proj mgl32.Mat4, view mgl32.Mat4) {
	shader := material.GetShader()
	shader.Use()

	// Set shader properties
	err := material.BindShaderProperties(shader)
	if err != nil {
		panic(err)
	}

	// Set matrices
	shader.SetMat4UniformLocation("model", &model)
	shader.SetMat4UniformLocation("view", &view)
	shader.SetMat4UniformLocation("projection", &proj)

	// Bind vertex array
	gl.BindVertexArray(mesh.Vao)

	// Enable attribute pointers
	attribMap := material.GetAttributeMap()
	for attr, attrName := range attribMap {
		attrLoc := shader.GetAttribLocation(attrName)
		gl.EnableVertexAttribArray(attrLoc)
	}

	// Render the mesh
	obj.Render(mesh, material, model, proj, view)

	// Disable attribute pointers
	for attr, attrName := range attribMap {
		attrLoc := shader.GetAttribLocation(attrName)
		gl.DisableVertexAttribArray(attrLoc)
	}

	// Unbind vertex array
	gl.BindVertexArray(0)

	// Unuse the shader
	shader.Unuse()
}
