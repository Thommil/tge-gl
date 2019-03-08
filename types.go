// Copyright 2014 The Go Authors.  All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package gl

// Enum is equivalent to GLenum, and is normally used with one of the
// constants defined in this package.
type Enum uint32

// Types are defined a structs so that in debug mode they can carry
// extra information, such as a string name. See typesdebug.go.

// Attrib identifies the location of a specific attribute variable.
type Attrib int32

// Program identifies a compiled shader program.
type Program uint32

// Shader identifies a GLSL shader.
type Shader uint32

// Buffer identifies a GL buffer object.
type Buffer uint32

// Framebuffer identifies a GL framebuffer.
type Framebuffer uint32

// A Renderbuffer is a GL object that holds an image in an internal format.
type Renderbuffer uint32

// A Texture identifies a GL texture unit.
type Texture uint32

// Uniform identifies the location of a specific uniform variable.
type Uniform int32

// A VertexArray is a GL object that holds vertices in an internal format.
type VertexArray uint32

// Valid indicates if attrib is valid in OpenGL context
func (v Attrib) Valid() bool { return v >= 0 }

// Valid indicates if program is valid in OpenGL context
func (v Program) Valid() bool { return v > 0 }

// Valid indicates if shader is valid in OpenGL context
func (v Shader) Valid() bool { return v > 0 }

// Valid indicates if buffer is valid in OpenGL context
func (v Buffer) Valid() bool { return v > 0 }

// Valid indicates if framebuffer is valid in OpenGL context
func (v Framebuffer) Valid() bool { return v > 0 }

// Valid indicates if renderbuffer is valid in OpenGL context
func (v Renderbuffer) Valid() bool { return v > 0 }

// Valid indicates if texture is valid in OpenGL context
func (v Texture) Valid() bool { return v > 0 }

// Valid indicates if uniform is valid in OpenGL context
func (v Uniform) Valid() bool { return v >= 0 }

// Valid indicates if VAO is valid in OpenGL context
func (v VertexArray) Valid() bool { return v > 0 }
