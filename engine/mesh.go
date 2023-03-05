package engine

import "github.com/go-gl/gl/v4.1-core/gl"

type Mesh struct {
    Vertices     []float32
    Indices      []uint32
    IndicesCount int32
    Vao          uint32
    Vbo          uint32
    Ebo          uint32
}

func NewMesh(vertices []float32, indices []uint32) *Mesh {
    var vao, vbo, ebo uint32
    gl.GenVertexArrays(1, &vao)
    gl.BindVertexArray(vao)

    // Set up vertex buffer object
    gl.GenBuffers(1, &vbo)
    gl.BindBuffer(gl.ARRAY_BUFFER, vbo)
    gl.BufferData(gl.ARRAY_BUFFER, len(vertices)*4, gl.Ptr(vertices), gl.STATIC_DRAW)

    // Set up element buffer object
    gl.GenBuffers(1, &ebo)
    gl.BindBuffer(gl.ELEMENT_ARRAY_BUFFER, ebo)
    gl.BufferData(gl.ELEMENT_ARRAY_BUFFER, len(indices)*4, gl.Ptr(indices), gl.STATIC_DRAW)

    // Set up vertex attributes
    gl.EnableVertexAttribArray(0)
    gl.VertexAttribPointer(0, 3, gl.FLOAT, false, 8*4, gl.PtrOffset(0))

    gl.EnableVertexAttribArray(1)
    gl.VertexAttribPointer(1, 3, gl.FLOAT, false, 8*4, gl.PtrOffset(3*4))

    gl.EnableVertexAttribArray(2)
    gl.VertexAttribPointer(2, 2, gl.FLOAT, false, 8*4, gl.PtrOffset(6*4))

    // Unbind VAO
    gl.BindVertexArray(0)
    gl.BindBuffer(gl.ARRAY_BUFFER, 0)
    gl.BindBuffer(gl.ELEMENT_ARRAY_BUFFER, 0)

    return &Mesh{
        Vao:          vao,
        Vbo:          vbo,
        Ebo:          ebo,
        IndicesCount: int32(len(indices)),
    }
}

func (m *Mesh) Draw(shader *ShaderProgram) {
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
