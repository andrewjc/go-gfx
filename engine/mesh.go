package engine

type Mesh struct {
	Vertices   []float32
	TexCoords  []float32
	Normals    []float32
	Indices    []uint32
	Depth      float32
	Vao        uint32
	IndexCount int32
	Material   Material
}

func (m *Mesh) GetAttribSize(attr uint32) int32 {
	return 16
}

func NewMesh(vertices []Vertex, coords []TexCoord, normals []Normal, indices []FaceVertex) *Mesh {
	mesh := &Mesh{
		Vertices:  make([]float32, len(vertices)*3),
		TexCoords: make([]float32, len(coords)*2),
		Normals:   make([]float32, len(normals)*3),
		Indices:   make([]uint32, len(indices)*3),

		Depth: 0.0,

		Vao:        0,
		IndexCount: int32(len(indices) * 3),

		Material: Material{},
	}

	for i, v := range vertices {
		mesh.Vertices[i*3] = v.X
		mesh.Vertices[i*3+1] = v.Y
		mesh.Vertices[i*3+2] = v.Z
	}

	for i, t := range coords {
		mesh.TexCoords[i*2] = t.Y
		mesh.TexCoords[i*2+1] = t.V
	}

	for i, n := range normals {
		mesh.Normals[i*3] = n.X
		mesh.Normals[i*3+1] = n.Y
		mesh.Normals[i*3+2] = n.Z
	}

	for i, f := range indices {
		mesh.Indices[i*3] = uint32(f.VertexIndex)
		mesh.Indices[i*3+1] = uint32(f.TexCoordIndex)
		mesh.Indices[i*3+2] = uint32(f.NormalIndex)
	}

	return mesh
}
