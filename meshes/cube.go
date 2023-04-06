package meshes

import . "physics/engine"

func NewCube(size int32) *ComplexMesh {

	// Calculate half of the size to simplify vertex calculations
	halfSize := float32(size / 2.0)

	// Define the vertices of the cube
	vertices := []Vertex{
		{-halfSize, -halfSize, halfSize},
		{halfSize, -halfSize, halfSize},
		{halfSize, halfSize, halfSize},
		{-halfSize, halfSize, halfSize},
		{-halfSize, -halfSize, -halfSize},
		{halfSize, -halfSize, -halfSize},
		{halfSize, halfSize, -halfSize},
		{-halfSize, halfSize, -halfSize},
	}

	// Define the texture coordinates of the cube (not used in this case)
	texCoords := []TexCoord{
		{0, 0},
		{1, 0},
		{1, 1},
		{0, 1},
	}

	// Define the normals of the cube (uniformly facing outwards)
	normals := []Normal{
		{0, 0, 1},
		{1, 0, 0},
		{0, 0, -1},
		{-1, 0, 0},
		{0, 1, 0},
		{0, -1, 0},
	}

	// Define the indices of the cube, referencing the vertices, texCoords and normals positions in the other arrays
	indices := []FaceVertex{
		{0, 0, 0}, {1, 0, 0}, {2, 0, 0}, {2, 0, 0}, {3, 0, 0}, {0, 0, 0},
		{1, 0, 1}, {5, 0, 1}, {6, 0, 1}, {6, 0, 1}, {2, 0, 1}, {1, 0, 1},
		{7, 0, 2}, {6, 0, 2}, {5, 0, 2}, {5, 0, 2}, {4, 0, 2}, {7, 0, 2},
		{4, 0, 3}, {0, 0, 3}, {3, 0, 3}, {3, 0, 3}, {7, 0, 3}, {4, 0, 3},
		{4, 0, 4}, {5, 0, 4}, {1, 0, 4}, {1, 0, 4}, {0, 0, 4}, {4, 0, 4},
		{3, 0, 5}, {2, 0, 5}, {6, 0, 5}, {6, 0, 5}, {7, 0, 5}, {3, 0, 5},
	}

	// Create a new ComplexMesh object
	mesh := NewComplexMesh(vertices, texCoords, normals, indices)

	return mesh

}
