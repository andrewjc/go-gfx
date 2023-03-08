package engine

import "github.com/go-gl/gl/v4.1-core/gl"

type BasicMesh struct {
	Vertices     []float32
	Indices      []uint32
	IndicesCount int32
	Vao          uint32
	Vbo          uint32
	Ebo          uint32
}

func NewBasicMesh(vertices []float32, indices []uint32) *BasicMesh {
	var vao, vbo, ebo uint32

	gl.GenVertexArrays(1, &vao)

	gl.GenBuffers(1, &vbo)

	gl.GenBuffers(1, &ebo)

	gl.BindVertexArray(vao)

	gl.BindBuffer(gl.ARRAY_BUFFER, vbo)
	gl.BufferData(gl.ARRAY_BUFFER, len(vertices)*4, gl.Ptr(vertices), gl.STATIC_DRAW)

	gl.BindBuffer(gl.ELEMENT_ARRAY_BUFFER, ebo)
	gl.BufferData(gl.ELEMENT_ARRAY_BUFFER, len(indices)*4, gl.Ptr(indices), gl.STATIC_DRAW)

	var stride int32 = 3*4 + 3*4 + 2*4
	var offset int = 0

	// position
	gl.VertexAttribPointer(0, 3, gl.FLOAT, false, stride, gl.PtrOffset(offset))
	gl.EnableVertexAttribArray(0)
	offset += 3 * 4

	// color
	gl.VertexAttribPointer(1, 3, gl.FLOAT, false, stride, gl.PtrOffset(offset))
	gl.EnableVertexAttribArray(1)
	offset += 3 * 4

	// texture position
	gl.VertexAttribPointer(2, 2, gl.FLOAT, false, stride, gl.PtrOffset(offset))
	gl.EnableVertexAttribArray(2)
	offset += 2 * 4

	// unbind the VAO (safe practice so we don't accidentally (mis)configure it later)
	gl.BindVertexArray(0)

	return &BasicMesh{
		Vao:          vao,
		Vbo:          vbo,
		Ebo:          ebo,
		IndicesCount: int32(len(indices)),
	}
}

func (m *BasicMesh) Draw(shader *ShaderProgram) {
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
