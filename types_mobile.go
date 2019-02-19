// Copyright 2014 The Go Authors.  All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build android ios

package gl

// Enum is equivalent to GLenum, and is normally used with one of the
// constants defined in this package.
type Enum uint32

// Types are defined a structs so that in debug mode they can carry
// extra information, such as a string name. See typesdebug.go.

// Attrib identifies the location of a specific attribute variable.
type Attrib struct {
	Value uint
}

// Program identifies a compiled shader program.
type Program struct {
	Value uint32
}

// Shader identifies a GLSL shader.
type Shader struct {
	Value uint32
}

// Buffer identifies a GL buffer object.
type Buffer struct {
	Value uint32
}

// Framebuffer identifies a GL framebuffer.
type Framebuffer struct {
	Value uint32
}

// A Renderbuffer is a GL object that holds an image in an internal format.
type Renderbuffer struct {
	Value uint32
}

// A Texture identifies a GL texture unit.
type Texture struct {
	Value uint32
}

// Uniform identifies the location of a specific uniform variable.
type Uniform struct {
	Value int32
}

// A VertexArray is a GL object that holds vertices in an internal format.
type VertexArray struct {
	Value uint32
}

// AttribNone helper for unbind purpose
var AttribNone = Attrib{
	Value: NONE,
}

// ProgramNone helper for unbind purpose
var ProgramNone = Program{
	Value: NONE,
}

// ShaderNone helper for unbind purpose
var ShaderNone = Shader{
	Value: NONE,
}

// BufferNone helper for unbind purpose
var BufferNone = Buffer{
	Value: NONE,
}

// FramebufferNone helper for unbind purpose
var FramebufferNone = Framebuffer{
	Value: NONE,
}

// RenderbufferNone helper for unbind purpose
var RenderbufferNone = Renderbuffer{
	Value: NONE,
}

// TextureNone helper for unbind purpose
var TextureNone = Texture{
	Value: NONE,
}

// UniformNone helper for unbind purpose
var UniformNone = Uniform{
	Value: NONE,
}

// VertexArrayNone helper for unbind purpose
var VertexArrayNone = VertexArray{
	Value: NONE,
}

func (v Attrib) c() uintptr       { return uintptr(v.Value) }
func (v Enum) c() uintptr         { return uintptr(v) }
func (v Program) c() uintptr      { return uintptr(v.Value) }
func (v Shader) c() uintptr       { return uintptr(v.Value) }
func (v Buffer) c() uintptr       { return uintptr(v.Value) }
func (v Framebuffer) c() uintptr  { return uintptr(v.Value) }
func (v Renderbuffer) c() uintptr { return uintptr(v.Value) }
func (v Texture) c() uintptr      { return uintptr(v.Value) }
func (v Uniform) c() uintptr      { return uintptr(v.Value) }
func (v VertexArray) c() uintptr  { return uintptr(v.Value) }

func (v Program) Valid() bool      { return v.Value != NONE }
func (v Shader) Valid() bool       { return v.Value != NONE }
func (v Buffer) Valid() bool       { return v.Value != NONE }
func (v Framebuffer) Valid() bool  { return v.Value != NONE }
func (v Renderbuffer) Valid() bool { return v.Value != NONE }
func (v Texture) Valid() bool      { return v.Value != NONE }
func (v Uniform) Valid() bool      { return v.Value != NONE }
func (v VertexArray) Valid() bool  { return v.Value != NONE }
