package engine

import (
	"github.com/go-gl/gl/v4.1-core/gl"
)

type ComplexMesh struct {
	Vao, Vbo, Ebo uint32
	VertexCount   int32
}

func (c ComplexMesh) Draw(shader *ShaderProgram) {
	//TODO implement me
	panic("implement me")
}

func NewComplexMesh(vertices []Vertex, texCoords []TexCoord, normals []Normal, indices []FaceVertex) *ComplexMesh {
	var vao, vbo, ebo uint32
	var vertexCount int32

	gl.GenVertexArrays(1, &vao)
	gl.BindVertexArray(vao)

	gl.GenBuffers(1, &vbo)
	gl.BindBuffer(gl.ARRAY_BUFFER, vbo)

	// Interleave vertex data
	var interleavedData []float32
	for _, face := range indices {
		vertex := vertices[face.vertexIndex]
		texCoord := texCoords[face.texCoordIndex]
		normal := normals[face.normalIndex]

		interleavedData = append(interleavedData, vertex.x, vertex.y, vertex.z)
		interleavedData = append(interleavedData, texCoord.u, texCoord.v)
		interleavedData = append(interleavedData, normal.x, normal.y, normal.z)
	}

	gl.BufferData(gl.ARRAY_BUFFER, len(interleavedData)*4, gl.Ptr(interleavedData), gl.STATIC_DRAW)

	gl.GenBuffers(1, &ebo)
	gl.BindBuffer(gl.ELEMENT_ARRAY_BUFFER, ebo)

	var indicesData []uint32
	for _, face := range indices {
		indicesData = append(indicesData, uint32(face.vertexIndex))
	}

	gl.BufferData(gl.ELEMENT_ARRAY_BUFFER, len(indicesData)*4, gl.Ptr(indicesData), gl.STATIC_DRAW)

	// Position attribute
	gl.EnableVertexAttribArray(0)
	gl.VertexAttribPointer(0, 3, gl.FLOAT, false, 8*4, gl.PtrOffset(0))

	// TexCoord attribute
	gl.EnableVertexAttribArray(1)
	gl.VertexAttribPointer(1, 2, gl.FLOAT, false, 8*4, gl.PtrOffset(3*4))

	// Normal attribute
	gl.EnableVertexAttribArray(2)
	gl.VertexAttribPointer(2, 3, gl.FLOAT, false, 8*4, gl.PtrOffset(5*4))

	gl.BindVertexArray(0)

	vertexCount = int32(len(indices))

	return &ComplexMesh{
		Vao:         vao,
		Vbo:         vbo,
		Ebo:         ebo,
		VertexCount: vertexCount,
	}
}
