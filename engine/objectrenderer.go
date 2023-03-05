package engine

import "github.com/go-gl/gl/v4.1-core/gl"

type ObjectRenderer struct {
    vao       uint32
    vertexVBO uint32
    indexVBO  uint32
    vertices  []float32
    indices   []uint32

    count int32
}

// implement init method for RenderableObject
func (r *ObjectRenderer) init() {
    gl.GenVertexArrays(1, &r.vao)
    gl.BindVertexArray(r.vao)

    gl.GenBuffers(1, &r.vertexVBO)
    gl.BindBuffer(gl.ARRAY_BUFFER, r.vertexVBO)
    gl.BufferData(gl.ARRAY_BUFFER, len(r.vertices)*4, gl.Ptr(r.vertices), gl.STATIC_DRAW)
    //defer gl.DeleteBuffers(1, &r.vertexVBO)

    gl.GenBuffers(1, &r.indexVBO)
    gl.BindBuffer(gl.ELEMENT_ARRAY_BUFFER, r.indexVBO)
    gl.BufferData(gl.ELEMENT_ARRAY_BUFFER, len(r.indices)*4, gl.Ptr(r.indices), gl.STATIC_DRAW)
    //defer gl.DeleteBuffers(1, &r.indexVBO)

    // Position attribute
    gl.VertexAttribPointerWithOffset(0, 3, gl.FLOAT, false, 6*4, 0)
    gl.EnableVertexAttribArray(0)

    // Normal attribute
    gl.VertexAttribPointerWithOffset(1, 3, gl.FLOAT, false, 6*4, 3*4)
    gl.EnableVertexAttribArray(1)

    gl.BindVertexArray(0)
    gl.BindBuffer(gl.ARRAY_BUFFER, 0)
    gl.BindBuffer(gl.ELEMENT_ARRAY_BUFFER, 0)
    
    r.count = int32(len(r.indices))
}

func (r *ObjectRenderer) Render(camera Camera) {
    gl.BindVertexArray(r.vao)
    gl.DrawElementsWithOffset(gl.TRIANGLES, int32(len(r.indices)), gl.UNSIGNED_INT, 0)

    gl.BindVertexArray(0)
}
