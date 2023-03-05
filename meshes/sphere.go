package meshes

import "math"
import . "physics/engine"

func NewSphere(radius float64, slices, stacks int) *Mesh {
	var vertices []float32
	var indices []uint32

	for i := 0; i <= stacks; i++ {
		v := float64(i) / float64(stacks)
		phi := v * math.Pi

		for j := 0; j <= slices; j++ {
			u := float64(j) / float64(slices)
			theta := u * 2 * math.Pi

			x := radius * math.Sin(phi) * math.Cos(theta)
			y := radius * math.Sin(phi) * math.Sin(theta)
			z := radius * math.Cos(phi)

			vertices = append(vertices, float32(x), float32(y), float32(z), float32(x), float32(y), float32(z), float32(u), float32(v))

			if i < stacks && j < slices {
				first := uint32((i * (slices + 1)) + j)
				second := first + uint32(slices) + uint32(1)

				indices = append(indices, first, second, first+1)
				indices = append(indices, second, second+1, first+1)
			}
		}
	}

	mesh := NewMesh(vertices, indices)
	return mesh
}
