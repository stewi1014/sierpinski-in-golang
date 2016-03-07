package main

import (
	"fmt"
	"runtime"
	"math"

	"github.com/go-gl/gl/v2.1/gl"
	"github.com/go-gl/glfw/v3.1/glfw"
)

var SCALE = 1.20
var POSX float64 = -1920
var POSY float64 = 1080

var h = -math.Sqrt(3) / 2

func init() {
	runtime.LockOSThread()
}

func main() {
	if err := glfw.Init(); err != nil {
		fmt.Println("Couldn't start GLFW!")
		panic(err)
	}
	defer glfw.Terminate()

	monitor := glfw.GetPrimaryMonitor()
	monitorRes := monitor.GetVideoMode()

	glfw.WindowHint(glfw.Resizable, glfw.False)
	glfw.WindowHint(glfw.ContextVersionMajor, 2)
	glfw.WindowHint(glfw.ContextVersionMinor, 1)
	glfw.WindowHint(glfw.Samples, 16)
	window, err := glfw.CreateWindow(monitorRes.Width, monitorRes.Height, "Fractal", monitor, nil)
	if err != nil {
		fmt.Println("Couldn't make window!")
		panic(err)
	}
	window.MakeContextCurrent()

	if err := gl.Init(); err != nil {
		fmt.Println("Couldn't init OpenGL")
		print(err)
	}

	gl.Enable(gl.BLEND)

	gl.Viewport(0, 0, int32(monitorRes.Width), int32(monitorRes.Height))
	gl.MatrixMode(gl.PROJECTION)
	gl.LoadIdentity()
	gl.Ortho(-float64(monitorRes.Width)/2.0, float64(monitorRes.Width)/2.0, float64(monitorRes.Height)/2.0, -float64(monitorRes.Height)/2.0, -1, 1)
	gl.MatrixMode(gl.MODELVIEW)
	gl.LoadIdentity()

	gl.ClearColor(0, 0, 0, 0)
	gl.Clear(gl.COLOR_BUFFER_BIT)
	gl.PushMatrix()
	gl.Translated(POSX, POSY, 0)
	gl.Scaled(SCALE, SCALE, 0)
	drawTriangle(11)
	gl.PopMatrix()
	window.SwapBuffers()

	for !window.ShouldClose() {
		glfw.PollEvents()
	}
}

func triangleBase() {
	gl.Color4f(1, 1, 1, 1)
	gl.Begin(gl.TRIANGLES)
	gl.Vertex2d(0, 0)
	gl.Vertex2d(0.5, h)
	gl.Vertex2d(1, 0)
	gl.End()
}

func drawTriangle(x int) {
	if x == 0 {
		triangleBase() //We've hit our recursion limit. Draw a triangle and bail.
		return
	}
	gl.PushMatrix()
	drawTriangle(x - 1)
	gl.Translated(math.Pow(2, float64(x))*.5, 0, 0)
	drawTriangle(x - 1)
	gl.Translated(-math.Pow(2, float64(x))*.25, math.Pow(2, float64(x))*h*.5, 0)
	drawTriangle(x - 1)
	gl.PopMatrix()
	return
}
