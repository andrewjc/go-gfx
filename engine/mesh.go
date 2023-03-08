package engine

type Mesh struct {
	Vertices []float32
	Material Material
	Depth    float32

	Vao        uint32
	IndexCount int32
}
