package main

import (
	"github.com/go-gl/gl/v4.1-core/gl"
	"github.com/go-gl/glfw/v3.3/glfw"
	"github.com/go-gl/mathgl/mgl32"
	"log"
	"math/rand"
	. "physics/engine"
	"runtime"
	"time"
)

const (
	width  = 800
	height = 600
)

func init() {
	// GLFW event handling must be run on the main OS thread
	runtime.LockOSThread()
}

func main() {
	rand.Seed(time.Now().UnixNano())

	defer glfw.Terminate()

	window := GLContext().Window

	// Enable depth testing
	gl.Enable(gl.DEPTH_TEST)
	gl.DepthFunc(gl.LESS)

	scene, err := NewScene(window)
	if err != nil {
		print("Failed creating scene")
		log.Fatal(err)
	}

	camera := scene.CreateCamera(
		mgl32.Vec3{0, 0, -15},
		mgl32.Vec3{0, 0, 0},
		mgl32.Vec3{0, 1, 0},
	)
	scene.Camera = camera

	objFileMeshes, err := LdrParseObj("mapdata/doom_E1M1.obj")
	if err != nil {
		print("Failed loading obj file")
		log.Fatal(err)
	}

	if objFileMeshes != nil {
		print("hello")
	}

	for _, mesh := range objFileMeshes.Objects {
		basicMesh := NewComplexMesh(mesh.Vertices, mesh.TexCoords, mesh.Normals, mesh.FaceIndices)
		scene.AddObject(&GameObject{
			Position: mgl32.Vec3{0.0, 0.0, 10.0},
			Rotation: QuatIdent,
			Velocity: mgl32.Vec3{},
			Scale:    5.0,
			Mass:     3,
			Mesh:     basicMesh,
			Material: NewBasicMaterial(),
		})
	}

	//CenterCameraOnModel(objFileMeshes[0].Positions, camera)

	/*scene.AddObject(&GameObject{
		Position: mgl32.Vec3{0.0, 0.0, 0.0},
		Rotation: QuatIdent,
		Velocity: mgl32.Vec3{},
		Scale:    1.0,
		Mass:     3,
		Mesh:     meshes.NewCube(5),
		Material: NewBasicMaterial(),
	})*/

	/*
	   // Add the cube object to the scene
	   scene.AddObject(&GameObject{
	       Position: Vec3{0.0, 0.0, 0.0},
	       Rotation: QuatIdent,
	       Velocity: Vec3{},
	       Scale:    1.0,
	       Mass:     3,
	       Mesh:     meshes.NewCube(5),
	       Material: NewBasicMaterial(),
	   })

	   // Add the sphere object to the scene
	   scene.AddObject(&GameObject{
	       Position: Vec3{5.0, 0.0, 0.0},
	       Rotation: QuatIdent,
	       Velocity: Vec3{},
	       Scale:    1.0,
	       Mass:     3,
	       Mesh:     meshes.NewSphere(5, 32, 16),
	       Material: NewBasicMaterial(),
	   })*/

	// Create a PBR material for the sphere
	/*sphereMaterial := pbr.NewPBRMaterial()
	  sphereMaterial.AlbedoColor = mgl32.Vec3{1.0, 0.0, 0.0}
	  sphereMaterial.Roughness = 0.2
	  sphereMaterial.Metallic = 0.8*/

	// Initialize the last frame time
	var lastFrameTime float64 = 0.0
	for !window.ShouldClose() {
		currentFrameTime := glfw.GetTime()
		dt := currentFrameTime - lastFrameTime
		lastFrameTime = currentFrameTime

		// Poll events
		glfw.PollEvents()

		// update camera position based on user input
		camera.Update(window, float32(dt))

		// update all objects in the scene
		scene.Update(float32(dt))

		// Render the objects in the scene
		scene.Render(camera)

		// Swap buffers
		window.SwapBuffers()
	}
}
