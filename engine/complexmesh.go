package engine

import (
	"github.com/go-gl/gl/v4.1-core/gl"
	"github.com/go-gl/mathgl/mgl32"
)

type ComplexMesh struct {
	Vertices      []mgl32.Vec3
	Indices       []uint16
	IndicesCount  int32
	TexCoords     []mgl32.Vec3
	Vao           uint32
	Vbo           uint32
	NormalVbo     uint32
	Ebo           uint32
	TextureHandle uint32
	VertexStride  int32
}

func NewComplexMesh(vertices []mgl32.Vec3, indices []uint16, normals []mgl32.Vec3, texCoords []mgl32.Vec3, textureHandle uint32) *ComplexMesh {
	var vao, vbo, ebo uint32

	newMesh := &ComplexMesh{
		Vao:           vao,
		Vbo:           vbo,
		Ebo:           ebo,
		Vertices:      vertices,
		Indices:       indices,
		IndicesCount:  int32(len(indices)),
		TexCoords:     texCoords,
		TextureHandle: textureHandle,
		VertexStride:  11 * 4,
	}

	newMesh.SetVertices(vertices, indices, vao, vbo, ebo)
	newMesh.SetNormals(normals)
	newMesh.SetTextureCoords(texCoords, vao, vbo)

	return newMesh
}

func (m *ComplexMesh) SetVertices(vertices []mgl32.Vec3, indices []uint16, vao uint32, vbo uint32, ebo uint32) {
	// Create and bind a VAO (Vertex Array Object)
	gl.GenVertexArrays(1, &vao)
	gl.BindVertexArray(vao)

	// Create and bind a VBO (Vertex Buffer Object) for the vertex positions
	gl.GenBuffers(1, &vbo)
	gl.BindBuffer(gl.ARRAY_BUFFER, vbo)
	gl.BufferData(gl.ARRAY_BUFFER, len(vertices)*4, gl.Ptr(vertices), gl.STATIC_DRAW)
	gl.EnableVertexAttribArray(0)
	gl.VertexAttribPointer(0, 3, gl.FLOAT, false, 0, nil)

	// Bind the index buffer
	gl.GenBuffers(1, &ebo)
	gl.BindBuffer(gl.ELEMENT_ARRAY_BUFFER, ebo)
	gl.BufferData(gl.ELEMENT_ARRAY_BUFFER, len(indices)*2, gl.Ptr(indices), gl.STATIC_DRAW)

	// Unbind the VBO, VAO, and EBO
	gl.BindBuffer(gl.ARRAY_BUFFER, 0)
	gl.BindVertexArray(0)
	gl.BindBuffer(gl.ELEMENT_ARRAY_BUFFER, 0)
}

func (m *ComplexMesh) Draw(shader *ShaderProgram) {
	gl.BindVertexArray(m.Vao)
	gl.BindBuffer(gl.ELEMENT_ARRAY_BUFFER, m.Ebo)
	gl.BindBuffer(gl.ARRAY_BUFFER, m.Vbo)
	shader.Use()

	gl.DrawElements(gl.TRIANGLES, m.IndicesCount, gl.UNSIGNED_INT, nil)

	gl.BindVertexArray(0)
	gl.BindBuffer(gl.ARRAY_BUFFER, 0)
	gl.BindBuffer(gl.ELEMENT_ARRAY_BUFFER, 0)

	shader.Unuse()
}

func (m *ComplexMesh) SetNormals(normals []mgl32.Vec3) {
	// Create and bind a new VBO for the normal data
	var normalVBO uint32
	gl.GenBuffers(1, &normalVBO)
	gl.BindBuffer(gl.ARRAY_BUFFER, normalVBO)

	// Set the normal data in the VBO
	gl.BufferData(gl.ARRAY_BUFFER, len(normals)*4, gl.Ptr(normals), gl.STATIC_DRAW)

	// Set up the vertex attribute pointer for the normal data
	gl.EnableVertexAttribArray(1)
	gl.VertexAttribPointer(1, 3, gl.FLOAT, false, 3*4, nil)

	// Unbind the normal VBO
	gl.BindBuffer(gl.ARRAY_BUFFER, 0)

	m.NormalVbo = normalVBO
}

func (m *ComplexMesh) SetTextureCoords(texCoords []mgl32.Vec3, vao uint32, vbo uint32) {
	gl.BindVertexArray(vao)

	// Create and bind a VBO (Vertex Buffer Object) for texture coordinates
	gl.GenBuffers(1, &vbo)
	gl.BindBuffer(gl.ARRAY_BUFFER, vbo)

	// Set the texture coordinate data in the VBO
	gl.BufferData(gl.ARRAY_BUFFER, len(texCoords)*12, gl.Ptr(texCoords), gl.STATIC_DRAW)

	// Set up the vertex attribute pointer for texture coordinates
	gl.EnableVertexAttribArray(1)
	gl.VertexAttribPointer(1, 3, gl.FLOAT, false, 0, nil)

	// Unbind the VBO and VAO
	gl.BindBuffer(gl.ARRAY_BUFFER, 0)
	gl.BindVertexArray(0)
}
