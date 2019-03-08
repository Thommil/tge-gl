<h1 align="center">TGE-GL - OpenGL plugin for TGE</h1>

 <p align="center">
    <a href="https://godoc.org/github.com/thommil/tge-gl"><img src="https://godoc.org/github.com/thommil/tge-gl?status.svg" alt="Godoc"></img></a>
    <a href="https://goreportcard.com/report/github.com/thommil/tge-gl"><img src="https://goreportcard.com/badge/github.com/thommil/tge-gl"  alt="Go Report Card"/></a>
</p>

OpenGL support for TGE runtime - [TGE](https://github.com/thommil/tge)

## Targets
 * OpenGL 3.3 core on Desktop
 * WebGL2 on Browsers
 * OpenGLES 3 on Mobile

## Dependencies
 * [TGE core](https://github.com/thommil/tge)

## Limitations
### Not implemented
 * glUniformMatrix2x3fv
 * glUniformMatrix3x2fv
 * glUniformMatrix2x4fv
 * glUniformMatrix4x2fv
 * glUniformMatrix3x4fv
 * glUniformMatrix4x3fv
 * glBlitFramebuffer
 * PolygonMode on Mobile/Browser 

## Implementation
See example at [OpenGL example](https://github.com/Thommil/tge-examples/tree/master/plugins/tge-gl)

Just import package and OpenGL API is available in all methods of your App except OnCreate():

```golang
package main

import (
	tge "github.com/thommil/tge"
	gl "github.com/thommil/tge-gl"
)

type MyApp struct {
}

func (app *MyApp) OnStart(runtime tge.Runtime) error {
	runtime.Subscribe(tge.ResizeEvent{}.Channel(), app.OnResize)
	gl.ClearColor(0.15, 0.04, 0.15, 1)
	return nil
}

func (app *MyApp) OnResize(event tge.Event) bool {
	gl.Viewport(0, 0, int(event.(tge.ResizeEvent).Width), int(event.(tge.ResizeEvent).Height))
	return false
}

...

```