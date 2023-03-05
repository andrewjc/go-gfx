package main

import (
    "github.com/go-gl/gl/v4.1-core/gl"
    "github.com/go-gl/glfw/v3.3/glfw"
    "log"
    "math/rand"
    . "physics/engine"
    "runtime"
    "time"
)

const (
    // Constants for the pocket universe simulation
    particleCount   = 5
    particleRadius  = 0.02
    forceStrength   = 0.01
    fieldStrength   = 0.02
    timeStep        = 0.01
    initialVelocity = 0.1
    windowWidth     = 800
    windowHeight    = 600
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
        Vec3{X: 0, Y: 0, Z: -5},
        Vec3{X: 0, Y: 0, Z: 0},
        Vec3{X: 0, Y: 1, Z: 0},
    )
    scene.Camera = camera

    mesh, err := LoadObjFile("mapdata/doom_E1M1.obj")
    scene.AddObject(&GameObject{
        Position: Vec3{0.0, 0.0, 0.0},
        Rotation: QuatIdent,
        Velocity: Vec3{},
        Scale:    1.0,
        Mass:     3,
        Mesh:     mesh,
        Material: NewBasicMaterial(),
    })

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
        camera.Update(window, dt)

        // update all objects in the scene
        scene.Update(dt)

        // Render the objects in the scene
        scene.Render(camera)

        // Swap buffers
        window.SwapBuffers()
    }
}
