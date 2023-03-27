package main

import (
	"github.com/municorn/chicky-chicky-go/blocks"
	"github.com/municorn/chicky-chicky-go/characters"
	"github.com/municorn/chicky-chicky-go/game"
	"github.com/municorn/chicky-chicky-go/input"
	"github.com/municorn/chicky-chicky-go/textures"
	"github.com/municorn/chicky-chicky-go/world"

	"fmt"
	"go/build"
	"log"
	"os"
	"runtime"
	"time"

	"github.com/go-gl/gl/v4.1-core/gl"
	"github.com/go-gl/glfw/v3.2/glfw"
)

func init() {
    // we have to change the working directory of chicky
    // chicky to its package path for imports and such
	p, err := build.Import("github.com/municorn/chicky-chicky-go", "", build.FindOnly)
	if err != nil {
		panic(err)
	}

	if err != nil {
		log.Fatalln("", err)
	}

	err = os.Chdir(p.Dir)
	if err != nil {
		log.Panicln("os.Chdir:", err)
	}
}

func main() {
	// I believe this ensures that our program always runs on the same process
	runtime.LockOSThread()

	if err := glfw.Init(); err != nil {
		log.Fatalln("failed to initialize glfw: ", err)
	}
	defer glfw.Terminate()

	glfw.WindowHint(glfw.Resizable, glfw.True)
	glfw.WindowHint(glfw.ContextVersionMajor, 4)
	glfw.WindowHint(glfw.ContextVersionMinor, 1)
	glfw.WindowHint(glfw.OpenGLProfile, glfw.OpenGLCoreProfile)
	glfw.WindowHint(glfw.OpenGLForwardCompatible, glfw.True)

	window, err := glfw.CreateWindow(800, 600, "Chicky Chicky Go", nil, nil)
	if err != nil {
		panic(err)
	}
	window.SetKeyCallback(input.KeyCallback)
	window.MakeContextCurrent()

	if err := gl.Init(); err != nil {
		panic(err)
	}

	initPackageGL()

	version := gl.GoStr(gl.GetString(gl.VERSION))
	fmt.Println("OpenGL version", version)

	gl.ClearColor(0, 0.5, 0.7, 1)

	lastUpdate := time.Now()

	for !window.ShouldClose() {
		deltaSeconds := float32(time.Now().Sub(lastUpdate).Seconds())

		game.Logic(deltaSeconds)

		gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)
		game.Render()

		window.SwapBuffers()
		glfw.PollEvents()
	}
}

func initPackageGL() {
	characters.InitGL()
	blocks.InitGL()
	textures.InitGL()
	world.InitGL()
}
