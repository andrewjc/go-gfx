package engine

import "github.com/go-gl/gl/v4.1-core/gl"

type ComplexMesh struct {
	Vertices      []float32
	Indices       []uint32
	IndicesCount  int32
	TexCoords     []uint32
	Vao           uint32
	Vbo           uint32
	Ebo           uint32
	TextureHandle uint32
	VertexStride  int
}

func NewComplexMesh(vertices []float32, indices []uint32, texCoords []uint32, textureHandle uint32) *ComplexMesh {
	var vao, vbo, ebo uint32

	// Create and bind a VAO (Vertex Array Object)
	gl.GenVertexArrays(1, &vao)
	gl.BindVertexArray(vao)

	// Create and bind a VBO (Vertex Buffer Object)
	gl.GenBuffers(1, &vbo)
	gl.BindBuffer(gl.ARRAY_BUFFER, vbo)

	// Set the vertex data in the VBO
	gl.BufferData(gl.ARRAY_BUFFER, len(vertices)*4, gl.Ptr(vertices), gl.STATIC_DRAW)

	// Set up the vertex attribute pointers
	gl.EnableVertexAttribArray(0)
	gl.VertexAttribPointer(0, 3, gl.FLOAT, false, 0, nil)

	gl.EnableVertexAttribArray(1)
	gl.VertexAttribPointer(1, 3, gl.FLOAT, false, 0, nil)

	gl.EnableVertexAttribArray(2)
	gl.VertexAttribPointer(2, 2, gl.FLOAT, false, 0, nil)

	// Create and bind an EBO (Element Buffer Object)
	gl.GenBuffers(1, &ebo)
	gl.BindBuffer(gl.ELEMENT_ARRAY_BUFFER, ebo)

	// Set the index data in the EBO
	gl.BufferData(gl.ELEMENT_ARRAY_BUFFER, len(indices)*4, gl.Ptr(indices), gl.STATIC_DRAW)

	// Unbind the VBO, VAO, and EBO
	gl.BindBuffer(gl.ARRAY_BUFFER, 0)
	gl.BindVertexArray(0)
	gl.BindBuffer(gl.ELEMENT_ARRAY_BUFFER, 0)

	return &ComplexMesh{
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

func (m *ComplexMesh) SetNormals(normals []float32, vertexNormal int32) {
	gl.BindVertexArray(m.Vao)
	gl.BindBuffer(gl.ARRAY_BUFFER, m.Vbo)

	gl.BufferData(gl.ARRAY_BUFFER, len(m.Vertices)*4, gl.Ptr(m.Vertices), gl.STATIC_DRAW)

	gl.VertexAttribPointer(uint32(vertexNormal), 3, gl.FLOAT, false, int32(m.VertexStride), gl.PtrOffset(8*4))
	gl.EnableVertexAttribArray(uint32(vertexNormal))

	gl.BindVertexArray(0)
	gl.BindBuffer(gl.ARRAY_BUFFER, 0)

	copy(m.Vertices[8:], normals)

	m.VertexStride = 11 * 4
}

func (m *ComplexMesh) SetTextureCoordinates(texCoords []float32) {
	// Call this function after creating the mesh
	gl.BindVertexArray(m.Vao)
	gl.BindBuffer(gl.ARRAY_BUFFER, m.Vbo)
	var offset int = len(m.Vertices) * 4 // calculate offset for texture coordinates
	gl.BufferSubData(gl.ARRAY_BUFFER, offset, len(texCoords)*4, gl.Ptr(texCoords))
	gl.BindVertexArray(0)
}
