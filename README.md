# TGE-GL plugin
OpenGL support for TGE runtime - [TGE](https://github.com/thommil/tge)

Targets:
 * OpenGL 3.3 core on Desktop
 * WebGL2 on Browsers
 * OpenGLES 3 on Mobile

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