// Copyright 2018 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build js

package gl

import "syscall/js"

type Enum int

type Attrib struct {
	Value int
}

type Program struct {
	Value js.Value
}

type Shader struct {
	Value js.Value
}

type Buffer struct {
	Value js.Value
}

type Framebuffer struct {
	Value js.Value
}

type Renderbuffer struct {
	Value js.Value
}

type Texture struct {
	Value js.Value
}

type Uniform struct {
	Value js.Value
}

type VertexArray struct {
	Value js.Value
}

// ProgramNone helper for unbind purpose
var ProgramNone = Program{
	Value: js.Null(),
}

// ShaderNone helper for unbind purpose
var ShaderNone = Shader{
	Value: js.Null(),
}

// BufferNone helper for unbind purpose
var BufferNone = Buffer{
	Value: js.Null(),
}

// FramebufferNone helper for unbind purpose
var FramebufferNone = Framebuffer{
	Value: js.Null(),
}

// RenderbufferNone helper for unbind purpose
var RenderbufferNone = Renderbuffer{
	Value: js.Null(),
}

// TextureNone helper for unbind purpose
var TextureNone = Texture{
	Value: js.Null(),
}

// UniformNone helper for unbind purpose
var UniformNone = Uniform{
	Value: js.Null(),
}

// VertexArrayNone helper for unbind purpose
var VertexArrayNone = VertexArray{
	Value: js.Null(),
}

func (v Attrib) Valid() bool       { return v.Value != -1 }
func (v Program) Valid() bool      { return v.Value != js.Null() }
func (v Shader) Valid() bool       { return v.Value != js.Null() }
func (v Buffer) Valid() bool       { return v.Value != js.Null() }
func (v Framebuffer) Valid() bool  { return v.Value != js.Null() }
func (v Renderbuffer) Valid() bool { return v.Value != js.Null() }
func (v Texture) Valid() bool      { return v.Value != js.Null() }
func (v Uniform) Valid() bool      { return v.Value != js.Null() }
func (v VertexArray) Valid() bool  { return v.Value != js.Null() }
