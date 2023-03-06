package engine

import (
	"fmt"
	"github.com/go-gl/gl/v4.1-core/gl"
	"github.com/go-gl/glfw/v3.3/glfw"
	"sync"
)

type OpenGlContext struct {
	Window *glfw.Window
}

func (c *OpenGlContext) check() {
	if c.Window.ShouldClose() {
		panic("Window closed")
	}
}

var (
	instance *OpenGlContext
	once     sync.Once
)

func GLContext() *OpenGlContext {
	once.Do(func() {
		ctx, err := NewGLContext()
		if err != nil {
			panic(err)
		} else {
			instance = ctx
		}
	})
	return instance
}

func NewGLContext() (*OpenGlContext, error) {
	if err := InitGlfw(); err != nil {
		return nil, err
	}
	window, err := glfw.CreateWindow(1280, 800, "Physics", nil, nil)
	if err != nil {
		return nil, err
	}

	mVer := window.GetAttrib(glfw.ContextVersionMajor)
	mRev := window.GetAttrib(glfw.ContextVersionMinor)
	fmt.Printf("OpenGL version: %d.%d\n", mVer, mRev)

	window.MakeContextCurrent()

	if err := gl.Init(); err != nil {
		panic(err)
	}

	return &OpenGlContext{window}, nil
}

func InitGlfw() error {
	if err := glfw.Init(); err != nil {
		return err
	}
	glfw.WindowHint(glfw.Resizable, glfw.True)
	glfw.WindowHint(glfw.ContextVersionMajor, 4)
	glfw.WindowHint(glfw.ContextVersionMinor, 1)
	glfw.WindowHint(glfw.OpenGLProfile, glfw.OpenGLCoreProfile)
	glfw.WindowHint(glfw.OpenGLForwardCompatible, glfw.True)
	return nil
}
