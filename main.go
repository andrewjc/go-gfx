package main

import (
	"github.com/go-gl/gl/v4.1-core/gl"
	"github.com/go-gl/glfw/v3.3/glfw"
	"github.com/go-gl/mathgl/mgl32"
	"log"
	. "physics/engine"
	"runtime"
	"sort"
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
	// Initialize GLFW
	err := glfw.Init()
	if err != nil {
		log.Fatal(err)
	}
	defer glfw.Terminate()

	// Set the OpenGL version and profile to use
	glfw.WindowHint(glfw.ContextVersionMajor, 4)
	glfw.WindowHint(glfw.ContextVersionMinor, 1)
	glfw.WindowHint(glfw.OpenGLProfile, glfw.OpenGLCoreProfile)
	glfw.WindowHint(glfw.OpenGLForwardCompatible, glfw.True)

	// Create a window
	window, err := glfw.CreateWindow(800, 600, "OpenGL Window", nil, nil)
	if err != nil {
		log.Fatal(err)
	}
	window.MakeContextCurrent()

	// Initialize OpenGL
	if err := gl.Init(); err != nil {
		log.Fatal(err)
	}

	// Set up the viewport
	width, height := window.GetSize()
	gl.Viewport(0, 0, int32(width), int32(height))

	camera := NewCamera(
		window,
		mgl32.Vec3{0.0, 0.0, 3.0},
		mgl32.Vec3{0.0, 1.0, 0.0},
		mgl32.Vec3{0.0, 0.0, -1.0},
	)

	var lights = []Light{
		{mgl32.Vec3{0.0, 0.0, 2.0}, mgl32.Vec3{1.0, 1.0, 1.0}},
		{mgl32.Vec3{2.0, 0.0, 0.0}, mgl32.Vec3{1.0, 0.0, 0.0}},
		{mgl32.Vec3{0.0, 2.0, 0.0}, mgl32.Vec3{0.0, 1.0, 0.0}},
		{mgl32.Vec3{0.0, 0.0, -2.0}, mgl32.Vec3{0.0, 0.0, 1.0}},
	}

	// Create an Objects array and populate it with objects to be rendered
	objects := []Mesh{
		{
			Vertices: []float32{
				// Define object vertices here
			},
			Material: Material{
				Ambient:   mgl32.Vec3{0.1, 0.1, 0.1},
				Diffuse:   mgl32.Vec3{0.5, 0.5, 0.5},
				Specular:  mgl32.Vec3{1.0, 1.0, 1.0},
				Shininess: 32.0,
			},
			Depth: 0.5,
		},
		// Add more objects here
	}

	sort.Slice(objects, func(i, j int) bool {
		return objects[i].Depth < objects[j].Depth
	})

	// Compile and link the PBR shader program
	shaderProgram, err := NewShaderProgram("shaders/default.vert", "shaders/default.frag")
	if err != nil {
		log.Fatal(err)
	}

	// Use the PBR shader program
	shaderProgram.Use()

	// Get the uniform locations in the PBR shader program
	modelLoc := gl.GetUniformLocation(shaderProgram.Handle(), gl.Str("model\x00"))
	viewLoc := gl.GetUniformLocation(shaderProgram.Handle(), gl.Str("view\x00"))
	projLoc := gl.GetUniformLocation(shaderProgram.Handle(), gl.Str("projection\x00"))
	albedoLoc := gl.GetUniformLocation(shaderProgram.Handle(), gl.Str("albedo\x00"))
	metallicLoc := gl.GetUniformLocation(shaderProgram.Handle(), gl.Str("metallic\x00"))
	roughnessLoc := gl.GetUniformLocation(shaderProgram.Handle(), gl.Str("roughness\x00"))
	aoLoc := gl.GetUniformLocation(shaderProgram.Handle(), gl.Str("ao\x00"))

	lightPositionsLoc := gl.GetUniformLocation(shaderProgram.Handle(), gl.Str("lightPositions\x00"))
	lightColorsLoc := gl.GetUniformLocation(shaderProgram.Handle(), gl.Str("lightColors\x00"))

	// Create arrays to hold the light positions and colors
	lightPositions := make([]mgl32.Vec3, len(lights))
	lightColors := make([]mgl32.Vec3, len(lights))

	// Populate the arrays with the light positions and colors
	for i, light := range lights {
		lightPositions[i] = light.Position
		lightColors[i] = light.Color
	}

	// Pass the light positions and colors to the shader
	gl.Uniform3fv(lightPositionsLoc, int32(len(lightPositions)), &lightPositions[0][0])
	gl.Uniform3fv(lightColorsLoc, int32(len(lightColors)), &lightColors[0][0])

	// Main loop
	for !window.ShouldClose() {
		// Clear the screen
		gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)

		// Draw your scene here
		for _, obj := range objects {
			// Bind the material for the current object
			gl.BindTexture(gl.TEXTURE_2D, obj.Material.TextureHandle)

			// Set the appropriate uniform variables for the PBR shader program
			gl.Uniform1i(albedoLoc, obj.Material.)
			gl.Uniform1f(gl.GetUniformLocation(shaderProgram.Handle(), gl.Str("material.shininess\x00")), obj.Material.Shininess)
			gl.Uniform3fv(gl.GetUniformLocation(shaderProgram.Handle(), gl.Str("camPos\x00")), 1, &camera.Position[0])

			// Render the current object using the appropriate OpenGL draw call
			gl.BindVertexArray(obj.Vao)
			gl.DrawElements(gl.TRIANGLES, obj.IndexCount, gl.UNSIGNED_INT, nil)
			gl.BindVertexArray(0)
		}

		// Swap buffers
		window.SwapBuffers()

		// Poll for events
		glfw.PollEvents()
	}
}
