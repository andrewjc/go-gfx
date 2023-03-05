package meshes

import . "physics/engine"

func NewCube(size int32) *Mesh {
	halfSize := float32(size / 2.0)

	vertices := []float32{
		// front face
		-halfSize, -halfSize, halfSize, 1.0, 0.0, 0.0, 0.0, 0.0,
		halfSize, -halfSize, halfSize, 1.0, 0.0, 0.0, 1.0, 0.0,
		halfSize, halfSize, halfSize, 1.0, 0.0, 0.0, 1.0, 1.0,
		-halfSize, halfSize, halfSize, 1.0, 0.0, 0.0, 0.0, 1.0,

		// back face
		-halfSize, -halfSize, -halfSize, 0.0, 1.0, 0.0, 0.0, 0.0,
		halfSize, -halfSize, -halfSize, 0.0, 1.0, 0.0, 1.0, 0.0,
		halfSize, halfSize, -halfSize, 0.0, 1.0, 0.0, 1.0, 1.0,
		-halfSize, halfSize, -halfSize, 0.0, 1.0, 0.0, 0.0, 1.0,

		// top face
		-halfSize, halfSize, halfSize, 0.0, 0.0, 1.0, 0.0, 0.0,
		halfSize, halfSize, halfSize, 0.0, 0.0, 1.0, 1.0, 0.0,
		halfSize, halfSize, -halfSize, 0.0, 0.0, 1.0, 1.0, 1.0,
		-halfSize, halfSize, -halfSize, 0.0, 0.0, 1.0, 0.0, 1.0,

		// bottom face
		-halfSize, -halfSize, halfSize, 1.0, 1.0, 0.0, 0.0, 0.0,
		halfSize, -halfSize, halfSize, 1.0, 1.0, 0.0, 1.0, 0.0,
		halfSize, -halfSize, -halfSize, 1.0, 1.0, 0.0, 1.0, 1.0,
		-halfSize, -halfSize, -halfSize, 1.0, 1.0, 0.0, 0.0, 1.0,

		// right face
		halfSize, -halfSize, halfSize, 0.0, 1.0, 1.0, 0.0, 0.0,
		halfSize, halfSize, halfSize, 0.0, 1.0, 1.0, 1.0, 0.0,
		halfSize, halfSize, -halfSize, 0.0, 1.0, 1.0, 1.0, 1.0,
		halfSize, -halfSize, -halfSize, 0.0, 1.0, 1.0, 0.0, 1.0,

		// left face
		-halfSize, -halfSize, halfSize, 1.0, 0.0, 1.0, 0.0, 0.0,
		-halfSize, halfSize, halfSize, 1.0, 0.0, 1.0, 1.0, 0.0,
		-halfSize, halfSize, -halfSize, 1.0, 0.0, 1.0, 1.0, 1.0,
		-halfSize, -halfSize, -halfSize, 1.0, 0.0, 1.0, 0.0, 1.0,
	}

	indices := []uint32{
		// front face
		0, 1, 2,
		2, 3, 0,

		// back face
		4, 5, 6,
		6, 7, 4,

		// top face
		8, 9, 10,
		10, 11, 8,

		// bottom face
		12, 13, 14,
		14, 15, 12,

		// right face
		16, 17, 18,
		18, 19, 16,

		// left face
		20, 21, 22,
		22, 23, 20,
	}

	// Create a new mesh from the sphere data
	mesh := NewBasicMesh(vertices, indices)

	return mesh
}
