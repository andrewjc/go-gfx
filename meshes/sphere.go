package meshes

import "math"
import . "physics/engine"

func NewSphere(radius float64, slices, stacks int) *Mesh {
	// Initialize the arrays to store the sphere data
	var vertices []float32
	var indices []uint32

	// Generate the sphere data
	for i := 0; i <= stacks; i++ {
		phi := float32(i) * math.Pi / float32(stacks)

		for j := 0; j <= slices; j++ {
			theta := float32(j) * 2 * math.Pi / float32(slices)

			x := radius * math.Cos(float64(theta)) * math.Sin(float64(phi))
			y := radius * math.Sin(float64(theta)) * math.Sin(float64(phi))
			z := radius * math.Cos(float64(phi))

			u := float32(j) / float32(slices)
			v := float32(i) / float32(stacks)

			vertices = append(vertices, float32(x), float32(y), float32(z), u, v)
		}
	}

	// Generate the sphere indices
	for i := 0; i < stacks; i++ {
		for j := 0; j < slices; j++ {
			p1 := uint32(i*(slices+1) + j)
			p2 := uint32((i+1)*(slices+1) + j)
			p3 := uint32((i+1)*(slices+1) + j + 1)
			p4 := uint32(i*(slices+1) + j + 1)

			indices = append(indices, p1, p2, p3, p1, p3, p4)
		}
	}

	// Create a new mesh from the sphere data
	mesh := NewMesh(vertices, indices)

	return mesh
}
