package engine

import "github.com/go-gl/mathgl/mgl32"

type Light struct {
	Ambient  mgl32.Vec3
	Diffuse  mgl32.Vec3
	Specular mgl32.Vec3
	Position mgl32.Vec3
	Velocity mgl32.Vec3
}
