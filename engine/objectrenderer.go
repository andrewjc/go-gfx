package engine

import "github.com/go-gl/mathgl/mgl32"

type ObjectRenderer interface {
	Render(mesh *Mesh, material Material, model mgl32.Mat4, proj mgl32.Mat4, view mgl32.Mat4)
}
