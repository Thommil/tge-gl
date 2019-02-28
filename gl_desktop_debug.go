// Copyright (c) 2019 Thomas MILLET. All rights reserved.

// +build darwin freebsd linux windows
// +build !android
// +build !ios
// +build !js
// +build debug

package gl

import (
	fmt "fmt"
	strings "strings"
	unsafe "unsafe"

	gl "github.com/go-gl/gl/v3.3-core/gl"
	tge "github.com/thommil/tge"
)

type plugin struct {
	isInit bool
}

var _pluginInstance = &plugin{}

func (p *plugin) Init(runtime tge.Runtime) error {
	if !p.isInit {
		p.isInit = true
		return gl.Init()
	}
	return fmt.Errorf("Already initialized")
}

func (p *plugin) GetName() string {
	return Name
}

func (p *plugin) Dispose() {
	p.isInit = false
}

// GetPlugin returns plugin handler
func GetPlugin() tge.Plugin {
	return _pluginInstance
}

// GetGLSLVersion gives the glsl version ti put in #version ${VERSION}
func GetGLSLVersion() string {
	return "330 core"
}

// ActiveTexture sets the active texture unit.
//
// http://www.khronos.org/opengles/sdk/docs/man3/html/glActiveTexture.xhtml
func ActiveTexture(texture Enum) {
	fmt.Printf("ActiveTexture : %v\n", texture)
	gl.ActiveTexture(uint32(texture))
	if err := int(GetError()); err != NO_ERROR {
		fmt.Printf("ERROR: %v\n", err)
	}
}

// AttachShader attaches a shader to a program.
//
// http://www.khronos.org/opengles/sdk/docs/man3/html/glAttachShader.xhtml
func AttachShader(p Program, s Shader) {
	fmt.Printf("AttachShader : %v, %v\n", p, s)
	gl.AttachShader(p.Value, s.Value)
	if err := int(GetError()); err != NO_ERROR {
		fmt.Printf("ERROR: %v\n", err)
	}

}

// BindAttribLocation binds a vertex attribute index with a named
// variable.
//
// http://www.khronos.org/opengles/sdk/docs/man3/html/glBindAttribLocation.xhtml
func BindAttribLocation(p Program, a Attrib, name string) {
	fmt.Printf("BindAttribLocation : %v, %v, %v\n", p, a, name)
	gl.BindAttribLocation(p.Value, uint32(a.Value), gl.Str(name+"\x00"))
	if err := int(GetError()); err != NO_ERROR {
		fmt.Printf("ERROR: %v\n", err)
	}
}

// BindBuffer binds a buffer.
//
// http://www.khronos.org/opengles/sdk/docs/man3/html/glBindBuffer.xhtml
func BindBuffer(target Enum, b Buffer) {
	fmt.Printf("BindBuffer : %v, %v\n", target, b)
	gl.BindBuffer(uint32(target), b.Value)
	if err := int(GetError()); err != NO_ERROR {
		fmt.Printf("ERROR: %v\n", err)
	}
}

// BindFramebuffer binds a framebuffer.
//
// http://www.khronos.org/opengles/sdk/docs/man3/html/glBindFramebuffer.xhtml
func BindFramebuffer(target Enum, fb Framebuffer) {
	fmt.Printf("BindFramebuffer : %v, %v\n", target, fb)
	gl.BindFramebuffer(uint32(target), fb.Value)
	if err := int(GetError()); err != NO_ERROR {
		fmt.Printf("ERROR: %v\n", err)
	}
}

// BindRenderbuffer binds a render buffer.
//
// http://www.khronos.org/opengles/sdk/docs/man3/html/glBindRenderbuffer.xhtml
func BindRenderbuffer(target Enum, rb Renderbuffer) {
	fmt.Printf("BindRenderbuffer : %v; %v\n", target, rb)
	gl.BindRenderbuffer(uint32(target), rb.Value)
	if err := int(GetError()); err != NO_ERROR {
		fmt.Printf("ERROR: %v\n", err)
	}
}

// BindTexture binds a texture.
//
// http://www.khronos.org/opengles/sdk/docs/man3/html/glBindTexture.xhtml
func BindTexture(target Enum, t Texture) {
	fmt.Printf("BindTexture : %v, %v\n", target, t)
	gl.BindTexture(uint32(target), t.Value)
	if err := int(GetError()); err != NO_ERROR {
		fmt.Printf("ERROR: %v\n", err)
	}
}

// BindVertexArray binds a VAO.
//
// http://www.khronos.org/opengles/sdk/docs/man3/html/glBindVertexArray.xhtml
func BindVertexArray(vao VertexArray) {
	fmt.Printf("BindVertexArray : %v\n", vao)
	gl.BindVertexArray(vao.Value)
	if err := int(GetError()); err != NO_ERROR {
		fmt.Printf("ERROR: %v\n", err)
	}
}

// BlendColor sets the blend color.
//
// http://www.khronos.org/opengles/sdk/docs/man3/html/glBlendColor.xhtml
func BlendColor(red, green, blue, alpha float32) {
	fmt.Printf("BlendColor : %v, %v, %v, %v\n", red, green, blue, alpha)
	gl.BlendColor(red, green, blue, alpha)
	if err := int(GetError()); err != NO_ERROR {
		fmt.Printf("ERROR: %v\n", err)
	}
}

// BlendEquation sets both RGB and alpha blend equations.
//
// http://www.khronos.org/opengles/sdk/docs/man3/html/glBlendEquation.xhtml
func BlendEquation(mode Enum) {
	fmt.Printf("BlendEquation : %v\n", mode)
	gl.BlendEquation(uint32(mode))
	if err := int(GetError()); err != NO_ERROR {
		fmt.Printf("ERROR: %v\n", err)
	}
}

// BlendEquationSeparate sets RGB and alpha blend equations separately.
//
// http://www.khronos.org/opengles/sdk/docs/man3/html/glBlendEquationSeparate.xhtml
func BlendEquationSeparate(modeRGB, modeAlpha Enum) {
	fmt.Printf("BlendEquationSeparate : %v, %v\n", modeRGB, modeAlpha)
	gl.BlendEquationSeparate(uint32(modeRGB), uint32(modeAlpha))
	if err := int(GetError()); err != NO_ERROR {
		fmt.Printf("ERROR: %v\n", err)
	}
}

// BlendFunc sets the pixel blending factors.
//
// http://www.khronos.org/opengles/sdk/docs/man3/html/glBlendFunc.xhtml
func BlendFunc(sfactor, dfactor Enum) {
	fmt.Printf("BlendFunc : %v, %v\n", sfactor, dfactor)
	gl.BlendFunc(uint32(sfactor), uint32(dfactor))
	if err := int(GetError()); err != NO_ERROR {
		fmt.Printf("ERROR: %v\n", err)
	}
}

// BlendFunc sets the pixel RGB and alpha blending factors separately.
//
// http://www.khronos.org/opengles/sdk/docs/man3/html/glBlendFuncSeparate.xhtml
func BlendFuncSeparate(sfactorRGB, dfactorRGB, sfactorAlpha, dfactorAlpha Enum) {
	fmt.Printf("BlendFuncSeparate : %v, %v, %v, %v\n", sfactorRGB, dfactorRGB, sfactorAlpha, dfactorAlpha)
	gl.BlendFuncSeparate(uint32(sfactorRGB), uint32(dfactorRGB), uint32(sfactorAlpha), uint32(dfactorAlpha))
	if err := int(GetError()); err != NO_ERROR {
		fmt.Printf("ERROR: %v\n", err)
	}
}

// BufferData creates a new data store for the bound buffer object.
//
// http://www.khronos.org/opengles/sdk/docs/man3/html/glBufferData.xhtml
func BufferData(target Enum, src []byte, usage Enum) {
	if len(src) > 10 {
		fmt.Printf("BufferData : %v, %v, %v\n", target, src[:10], usage)
	} else {
		fmt.Printf("BufferData : %v, %v, %v\n", target, src, usage)
	}
	gl.BufferData(uint32(target), int(len(src)), gl.Ptr(&src[0]), uint32(usage))
	if err := int(GetError()); err != NO_ERROR {
		fmt.Printf("ERROR: %v\n", err)
	}
}

// BufferInit creates a new unitialized data store for the bound buffer object.
//
// http://www.khronos.org/opengles/sdk/docs/man3/html/glBufferData.xhtml
func BufferInit(target Enum, size int, usage Enum) {
	fmt.Printf("BufferInit : %v, %v, %v\n", target, size, usage)
	gl.BufferData(uint32(target), size, nil, uint32(usage))
	if err := int(GetError()); err != NO_ERROR {
		fmt.Printf("ERROR: %v\n", err)
	}
}

// BufferSubData sets some of data in the bound buffer object.
//
// http://www.khronos.org/opengles/sdk/docs/man3/html/glBufferSubData.xhtml
func BufferSubData(target Enum, offset int, data []byte) {
	if len(data) > 10 {
		fmt.Printf("BufferSubData : %v, %v, %v\n", target, offset, data[:10])
	} else {
		fmt.Printf("BufferSubData : %v, %v, %v\n", target, offset, data)
	}
	gl.BufferSubData(uint32(target), offset, int(len(data)), gl.Ptr(&data[0]))
	if err := int(GetError()); err != NO_ERROR {
		fmt.Printf("ERROR: %v\n", err)
	}
}

// CheckFramebufferStatus reports the completeness status of the
// active framebuffer.
//
// http://www.khronos.org/opengles/sdk/docs/man3/html/glCheckFramebufferStatus.xhtml
func CheckFramebufferStatus(target Enum) Enum {
	s := Enum(gl.CheckFramebufferStatus(uint32(target)))
	fmt.Printf("CheckFramebufferStatus : %v -> %v\n", target, s)
	if err := int(GetError()); err != NO_ERROR {
		fmt.Printf("ERROR: %v\n", err)
	}
	return s

}

// Clear clears the window.
//
// The behavior of Clear is influenced by the pixel ownership test,
// the scissor test, dithering, and the buffer writemasks.
//
// http://www.khronos.org/opengles/sdk/docs/man3/html/glClear.xhtml
func Clear(mask Enum) {
	fmt.Printf("Clear : %v\n", mask)
	gl.Clear(uint32(mask))
	if err := int(GetError()); err != NO_ERROR {
		fmt.Printf("ERROR: %v\n", err)
	}
}

// ClearColor specifies the RGBA values used to clear color buffers.
//
// http://www.khronos.org/opengles/sdk/docs/man3/html/glClearColor.xhtml
func ClearColor(red, green, blue, alpha float32) {
	fmt.Printf("ClearColor : %v, %v, %v, %v\n", red, green, blue, alpha)
	gl.ClearColor(red, green, blue, alpha)
	if err := int(GetError()); err != NO_ERROR {
		fmt.Printf("ERROR: %v\n", err)
	}
}

// ClearDepthf sets the depth value used to clear the depth buffer.
//
// http://www.khronos.org/opengles/sdk/docs/man3/html/glClearDepthf.xhtml
func ClearDepthf(d float32) {
	fmt.Printf("ClearDepthf : %v\n", d)
	gl.ClearDepthf(d)
	if err := int(GetError()); err != NO_ERROR {
		fmt.Printf("ERROR: %v\n", err)
	}
}

// ClearStencil sets the index used to clear the stencil buffer.
//
// http://www.khronos.org/opengles/sdk/docs/man3/html/glClearStencil.xhtml
func ClearStencil(s int) {
	fmt.Printf("ClearStencil : %v\n", s)
	gl.ClearStencil(int32(s))
	if err := int(GetError()); err != NO_ERROR {
		fmt.Printf("ERROR: %v\n", err)
	}
}

// ColorMask specifies whether color components in the framebuffer
// can be written.
//
// http://www.khronos.org/opengles/sdk/docs/man3/html/glColorMask.xhtml
func ColorMask(red, green, blue, alpha bool) {
	fmt.Printf("ColorMask : %v, %v, %v, %v\n", red, green, blue, alpha)
	gl.ColorMask(red, green, blue, alpha)
	if err := int(GetError()); err != NO_ERROR {
		fmt.Printf("ERROR: %v\n", err)
	}
}

// CompileShader compiles the source code of s.
//
// http://www.khronos.org/opengles/sdk/docs/man3/html/glCompileShader.xhtml
func CompileShader(s Shader) {
	fmt.Printf("CompileShader : %v\n", s)
	gl.CompileShader(s.Value)
	if err := int(GetError()); err != NO_ERROR {
		fmt.Printf("ERROR: %v\n", err)
	}
}

// CompressedTexImage2D writes a compressed 2D texture.
//
// http://www.khronos.org/opengles/sdk/docs/man3/html/glCompressedTexImage2D.xhtml
func CompressedTexImage2D(target Enum, level int, internalformat Enum, width, height, border int, data []byte) {
	fmt.Printf("CompressedTexImage2D : %v, %v, %v, %v, %v, %v, %v\n", target, level, internalformat, width, height, border, len(data))
	gl.CompressedTexImage2D(uint32(target), int32(level), uint32(internalformat), int32(width), int32(height), int32(border), int32(len(data)), gl.Ptr(data))
	if err := int(GetError()); err != NO_ERROR {
		fmt.Printf("ERROR: %v\n", err)
	}
}

// CompressedTexSubImage2D writes a subregion of a compressed 2D texture.
//
// http://www.khronos.org/opengles/sdk/docs/man3/html/glCompressedTexSubImage2D.xhtml
func CompressedTexSubImage2D(target Enum, level, xoffset, yoffset, width, height int, format Enum, data []byte) {
	fmt.Printf("CompressedTexSubImage2D : %v, %v, %v, %v, %v, %v, %v, %v\n", target, level, xoffset, yoffset, width, height, format, len(data))
	gl.CompressedTexSubImage2D(uint32(target), int32(level), int32(xoffset), int32(yoffset), int32(width), int32(height), uint32(format), int32(len(data)), gl.Ptr(data))
	if err := int(GetError()); err != NO_ERROR {
		fmt.Printf("ERROR: %v\n", err)
	}
}

// CopyTexImage2D writes a 2D texture from the current framebuffer.
//
// http://www.khronos.org/opengles/sdk/docs/man3/html/glCopyTexImage2D.xhtml
func CopyTexImage2D(target Enum, level int, internalformat Enum, x, y, width, height, border int) {
	fmt.Printf("CopyTexImage2D : %v, %v, %v, %v, %v, %v, %v, %v\n", target, level, internalformat, x, y, width, height, border)
	gl.CopyTexImage2D(uint32(target), int32(level), uint32(internalformat), int32(x), int32(y), int32(width), int32(height), int32(border))
	if err := int(GetError()); err != NO_ERROR {
		fmt.Printf("ERROR: %v\n", err)
	}
}

// CopyTexSubImage2D writes a 2D texture subregion from the
// current framebuffer.
//
// http://www.khronos.org/opengles/sdk/docs/man3/html/glCopyTexSubImage2D.xhtml
func CopyTexSubImage2D(target Enum, level, xoffset, yoffset, x, y, width, height int) {
	fmt.Printf("CopyTexSubImage2D : %v, %v, %v, %v, %v, %v, %v, %v\n", target, level, xoffset, yoffset, x, y, width, height)
	gl.CopyTexSubImage2D(uint32(target), int32(level), int32(xoffset), int32(yoffset), int32(x), int32(y), int32(width), int32(height))
	if err := int(GetError()); err != NO_ERROR {
		fmt.Printf("ERROR: %v\n", err)
	}
}

// CreateBuffer creates a buffer object.
//
// http://www.khronos.org/opengles/sdk/docs/man3/html/glGenBuffers.xhtml
func CreateBuffer() Buffer {
	var b Buffer
	gl.GenBuffers(1, &b.Value)
	fmt.Printf("CreateBuffer -> %v\n", b)
	if err := int(GetError()); err != NO_ERROR {
		fmt.Printf("ERROR: %v\n", err)
	}
	return b
}

// CreateFramebuffer creates a framebuffer object.
//
// http://www.khronos.org/opengles/sdk/docs/man3/html/glGenFramebuffers.xhtml
func CreateFramebuffer() Framebuffer {
	var b Framebuffer
	gl.GenFramebuffers(1, &b.Value)
	fmt.Printf("CreateFramebuffer -> %v\n", b)
	if err := int(GetError()); err != NO_ERROR {
		fmt.Printf("ERROR: %v\n", err)
	}
	return b
}

// CreateProgram creates a new empty program object.
//
// http://www.khronos.org/opengles/sdk/docs/man3/html/glCreateProgram.xhtml
func CreateProgram() Program {
	p := gl.CreateProgram()
	fmt.Printf("CreateProgram ->%v\n", p)
	if err := int(GetError()); err != NO_ERROR {
		fmt.Printf("ERROR: %v\n", err)
	}
	return Program{Value: uint32(p)}
}

// CreateRenderbuffer create a renderbuffer object.
//
// http://www.khronos.org/opengles/sdk/docs/man3/html/glGenRenderbuffers.xhtml
func CreateRenderbuffer() Renderbuffer {
	var b Renderbuffer
	gl.GenRenderbuffers(1, &b.Value)
	fmt.Printf("CreateRenderbuffer -> %v\n", b)
	if err := int(GetError()); err != NO_ERROR {
		fmt.Printf("ERROR: %v\n", err)
	}
	return b
}

// CreateShader creates a new empty shader object.
//
// http://www.khronos.org/opengles/sdk/docs/man3/html/glCreateShader.xhtml
func CreateShader(ty Enum) Shader {
	s := gl.CreateShader(uint32(ty))
	fmt.Printf("CreateShader : %v -> %v\n", ty, s)
	if err := int(GetError()); err != NO_ERROR {
		fmt.Printf("ERROR: %v\n", err)
	}
	return Shader{Value: uint32(s)}
}

// CreateTexture creates a texture object.
//
// http://www.khronos.org/opengles/sdk/docs/man3/html/glGenTextures.xhtml
func CreateTexture() Texture {
	var t Texture
	gl.GenTextures(1, &t.Value)
	fmt.Printf("CreateTexture -> %v\n", t)
	if err := int(GetError()); err != NO_ERROR {
		fmt.Printf("ERROR: %v\n", err)
	}
	return t
}

// CreateVertexArray creates a VAO.
//
// http://www.khronos.org/opengles/sdk/docs/man3/html/glGenVertexArrays.xhtml
func CreateVertexArray() VertexArray {
	var vao VertexArray
	gl.GenVertexArrays(1, &vao.Value)
	fmt.Printf("CreateVertexArray -> %v\n", vao)
	if err := int(GetError()); err != NO_ERROR {
		fmt.Printf("ERROR: %v\n", err)
	}
	return vao
}

// CullFace specifies which polygons are candidates for culling.
//
// Valid modes: FRONT, BACK, FRONT_AND_BACK.
//
// http://www.khronos.org/opengles/sdk/docs/man3/html/glCullFace.xhtml
func CullFace(mode Enum) {
	fmt.Printf("CullFace : %v\n", mode)
	gl.CullFace(uint32(mode))
	if err := int(GetError()); err != NO_ERROR {
		fmt.Printf("ERROR: %v\n", err)
	}
}

// DeleteBuffer deletes the given buffer object.
//
// http://www.khronos.org/opengles/sdk/docs/man3/html/glDeleteBuffers.xhtml
func DeleteBuffer(v Buffer) {
	fmt.Printf("DeleteBuffer : %v\n", v)
	gl.DeleteBuffers(1, &v.Value)
	if err := int(GetError()); err != NO_ERROR {
		fmt.Printf("ERROR: %v\n", err)
	}
}

// DeleteFramebuffer deletes the given framebuffer object.
//
// http://www.khronos.org/opengles/sdk/docs/man3/html/glDeleteFramebuffers.xhtml
func DeleteFramebuffer(v Framebuffer) {
	fmt.Printf("DeleteFramebuffer : %v\n", v)
	gl.DeleteFramebuffers(1, &v.Value)
	if err := int(GetError()); err != NO_ERROR {
		fmt.Printf("ERROR: %v\n", err)
	}
}

// DeleteProgram deletes the given program object.
//
// http://www.khronos.org/opengles/sdk/docs/man3/html/glDeleteProgram.xhtml
func DeleteProgram(p Program) {
	fmt.Printf("DeleteProgram : %v\n", p)
	gl.DeleteProgram(p.Value)
	if err := int(GetError()); err != NO_ERROR {
		fmt.Printf("ERROR: %v\n", err)
	}
}

// DeleteRenderbuffer deletes the given render buffer object.
//
// http://www.khronos.org/opengles/sdk/docs/man3/html/glDeleteRenderbuffers.xhtml
func DeleteRenderbuffer(v Renderbuffer) {
	fmt.Printf("DeleteRenderbuffer : %v\n", v)
	gl.DeleteRenderbuffers(1, &v.Value)
	if err := int(GetError()); err != NO_ERROR {
		fmt.Printf("ERROR: %v\n", err)
	}
}

// DeleteShader deletes shader s.
//
// http://www.khronos.org/opengles/sdk/docs/man3/html/glDeleteShader.xhtml
func DeleteShader(s Shader) {
	fmt.Printf("DeleteShader : %v\n", s)
	gl.DeleteShader(s.Value)
	if err := int(GetError()); err != NO_ERROR {
		fmt.Printf("ERROR: %v\n", err)
	}
}

// DeleteTexture deletes the given texture object.
//
// http://www.khronos.org/opengles/sdk/docs/man3/html/glDeleteTextures.xhtml
func DeleteTexture(v Texture) {
	fmt.Printf("DeleteTexture : %v\n", v)
	gl.DeleteTextures(1, &v.Value)
	if err := int(GetError()); err != NO_ERROR {
		fmt.Printf("ERROR: %v\n", err)
	}
}

// DeleteVertexArray deletes the given VAO.
//
// http://www.khronos.org/opengles/sdk/docs/man3/html/glDeleteVertexArrays.xhtml
func DeleteVertexArray(v VertexArray) {
	fmt.Printf("DeleteVertexArray : %v\n", v)
	gl.DeleteVertexArrays(1, &v.Value)
	if err := int(GetError()); err != NO_ERROR {
		fmt.Printf("ERROR: %v\n", err)
	}
}

// DepthFunc sets the function used for depth buffer comparisons.
//
// Valid fn values:
//	NEVER
//	LESS
//	EQUAL
//	LEQUAL
//	GREATER
//	NOTEQUAL
//	GEQUAL
//	ALWAYS
//
// http://www.khronos.org/opengles/sdk/docs/man3/html/glDepthFunc.xhtml
func DepthFunc(fn Enum) {
	fmt.Printf("DepthFunc : %v\n", fn)
	gl.DepthFunc(uint32(fn))
	if err := int(GetError()); err != NO_ERROR {
		fmt.Printf("ERROR: %v\n", err)
	}
}

// DepthMask sets the depth buffer enabled for writing.
//
// http://www.khronos.org/opengles/sdk/docs/man3/html/glDepthMask.xhtml
func DepthMask(flag bool) {
	fmt.Printf("DepthMask : %v\n", flag)
	gl.DepthMask(flag)
	if err := int(GetError()); err != NO_ERROR {
		fmt.Printf("ERROR: %v\n", err)
	}
}

// DepthRangef sets the mapping from normalized device coordinates to
// window coordinates.
//
// http://www.khronos.org/opengles/sdk/docs/man3/html/glDepthRangef.xhtml
func DepthRangef(n, f float32) {
	fmt.Printf("DepthRangef : %v, %v\n", n, f)
	gl.DepthRangef(n, f)
	if err := int(GetError()); err != NO_ERROR {
		fmt.Printf("ERROR: %v\n", err)
	}
}

// DetachShader detaches the shader s from the program p.
//
// http://www.khronos.org/opengles/sdk/docs/man3/html/glDetachShader.xhtml
func DetachShader(p Program, s Shader) {
	fmt.Printf("DetachShader : %v, %v\n", p, s)
	gl.DetachShader(p.Value, s.Value)
	if err := int(GetError()); err != NO_ERROR {
		fmt.Printf("ERROR: %v\n", err)
	}
}

// Disable disables various GL capabilities.
//
// http://www.khronos.org/opengles/sdk/docs/man3/html/glDisable.xhtml
func Disable(cap Enum) {
	fmt.Printf("Disable : %v\n", cap)
	gl.Disable(uint32(cap))
	if err := int(GetError()); err != NO_ERROR {
		fmt.Printf("ERROR: %v\n", err)
	}
}

// DisableVertexAttribArray disables a vertex attribute array.
//
// http://www.khronos.org/opengles/sdk/docs/man3/html/glDisableVertexAttribArray.xhtml
func DisableVertexAttribArray(a Attrib) {
	fmt.Printf("DisableVertexAttribArray : %v\n", a)
	gl.DisableVertexAttribArray(uint32(a.Value))
	if err := int(GetError()); err != NO_ERROR {
		fmt.Printf("ERROR: %v\n", err)
	}
}

// DrawArrays renders geometric primitives from the bound data.
//
// http://www.khronos.org/opengles/sdk/docs/man3/html/glDrawArrays.xhtml
func DrawArrays(mode Enum, first, count int) {
	fmt.Printf("DrawArrays : %v, %v, %v\n", mode, first, count)
	gl.DrawArrays(uint32(mode), int32(first), int32(count))
	if err := int(GetError()); err != NO_ERROR {
		fmt.Printf("ERROR: %v\n", err)
	}
}

// DrawElements renders primitives from a bound buffer.
//
// http://www.khronos.org/opengles/sdk/docs/man3/html/glDrawElements.xhtml
func DrawElements(mode Enum, count int, ty Enum, offset int) {
	fmt.Printf("DrawElements : %v, %v, %v, %v\n", mode, count, ty, offset)
	gl.DrawElements(uint32(mode), int32(count), uint32(ty), gl.PtrOffset(offset))
	if err := int(GetError()); err != NO_ERROR {
		fmt.Printf("ERROR: %v\n", err)
	}
}

// Enable enables various GL capabilities.
//
// http://www.khronos.org/opengles/sdk/docs/man3/html/glEnable.xhtml
func Enable(cap Enum) {
	fmt.Printf("Enable : %v\n", cap)
	gl.Enable(uint32(cap))
	if err := int(GetError()); err != NO_ERROR {
		fmt.Printf("ERROR: %v\n", err)
	}
}

// EnableVertexAttribArray enables a vertex attribute array.
//
// http://www.khronos.org/opengles/sdk/docs/man3/html/glEnableVertexAttribArray.xhtml
func EnableVertexAttribArray(a Attrib) {
	fmt.Printf("EnableVertexAttribArray : %v\n", a)
	gl.EnableVertexAttribArray(uint32(a.Value))
	if err := int(GetError()); err != NO_ERROR {
		fmt.Printf("ERROR: %v\n", err)
	}
}

// Finish blocks until the effects of all previously called GL
// commands are complete.
//
// http://www.khronos.org/opengles/sdk/docs/man3/html/glFinish.xhtml
func Finish() {
	fmt.Printf("Finish\n")
	gl.Finish()
	if err := int(GetError()); err != NO_ERROR {
		fmt.Printf("ERROR: %v\n", err)
	}
}

// Flush empties all buffers. It does not block.
//
// An OpenGL implementation may buffer network communication,
// the command stream, or data inside the graphics accelerator.
//
// http://www.khronos.org/opengles/sdk/docs/man3/html/glFlush.xhtml
func Flush() {
	fmt.Printf("Flush\n")
	gl.Flush()
	if err := int(GetError()); err != NO_ERROR {
		fmt.Printf("ERROR: %v\n", err)
	}
}

// FramebufferRenderbuffer attaches rb to the current frame buffer.
//
// http://www.khronos.org/opengles/sdk/docs/man3/html/glFramebufferRenderbuffer.xhtml
func FramebufferRenderbuffer(target, attachment, rbTarget Enum, rb Renderbuffer) {
	fmt.Printf("FramebufferRenderbuffer : %v, %v, %v, %v\n", target, attachment, rbTarget, rb)
	gl.FramebufferRenderbuffer(uint32(target), uint32(attachment), uint32(rbTarget), rb.Value)
	if err := int(GetError()); err != NO_ERROR {
		fmt.Printf("ERROR: %v\n", err)
	}
}

// FramebufferTexture2D attaches the t to the current frame buffer.
//
// http://www.khronos.org/opengles/sdk/docs/man3/html/glFramebufferTexture2D.xhtml
func FramebufferTexture2D(target, attachment, texTarget Enum, t Texture, level int) {
	fmt.Printf("FramebufferTexture2D : %v, %v, %v, %v, %v\n", target, attachment, texTarget, t, level)
	gl.FramebufferTexture2D(uint32(target), uint32(attachment), uint32(texTarget), t.Value, int32(level))
	if err := int(GetError()); err != NO_ERROR {
		fmt.Printf("ERROR: %v\n", err)
	}
}

// FrontFace defines which polygons are front-facing.
//
// Valid modes: CW, CCW.
//
// http://www.khronos.org/opengles/sdk/docs/man3/html/glFrontFace.xhtml
func FrontFace(mode Enum) {
	fmt.Printf("FrontFace : %v\n", mode)
	gl.FrontFace(uint32(mode))
	if err := int(GetError()); err != NO_ERROR {
		fmt.Printf("ERROR: %v\n", err)
	}
}

// GenerateMipmap generates mipmaps for the current texture.
//
// http://www.khronos.org/opengles/sdk/docs/man3/html/glGenerateMipmap.xhtml
func GenerateMipmap(target Enum) {
	fmt.Printf("GenerateMipmap : %v\n", target)
	gl.GenerateMipmap(uint32(target))
	if err := int(GetError()); err != NO_ERROR {
		fmt.Printf("ERROR: %v\n", err)
	}
}

// GetActiveAttrib returns details about an active attribute variable.
// A value of 0 for index selects the first active attribute variable.
// Permissible values for index range from 0 to the number of active
// attribute variables minus 1.
//
// http://www.khronos.org/opengles/sdk/docs/man3/html/glGetActiveAttrib.xhtml
func GetActiveAttrib(p Program, index uint32) (name string, size int, ty Enum) {
	var length, si int32
	var typ uint32
	name = strings.Repeat("\x00", 256)
	cname := gl.Str(name)
	gl.GetActiveAttrib(p.Value, uint32(index), int32(len(name)-1), &length, &si, &typ, cname)
	name = name[:strings.IndexRune(name, 0)]
	fmt.Printf("GetActiveAttrib : %v, %v -> %v, %v, %v\n", p, index, name, si, typ)
	if err := int(GetError()); err != NO_ERROR {
		fmt.Printf("ERROR: %v\n", err)
	}
	return name, int(si), Enum(typ)
}

// GetActiveUniform returns details about an active uniform variable.
// A value of 0 for index selects the first active uniform variable.
// Permissible values for index range from 0 to the number of active
// uniform variables minus 1.
//
// http://www.khronos.org/opengles/sdk/docs/man3/html/glGetActiveUniform.xhtml
func GetActiveUniform(p Program, index uint32) (name string, size int, ty Enum) {
	var length, si int32
	var typ uint32
	name = strings.Repeat("\x00", 256)
	cname := gl.Str(name)
	gl.GetActiveUniform(p.Value, uint32(index), int32(len(name)-1), &length, &si, &typ, cname)
	name = name[:strings.IndexRune(name, 0)]
	fmt.Printf("GetActiveUniform : %v, %v -> %v, %v, %v\n", p, index, name, si, typ)
	if err := int(GetError()); err != NO_ERROR {
		fmt.Printf("ERROR: %v\n", err)
	}
	return name, int(si), Enum(typ)
}

// GetAttachedShaders returns the shader objects attached to program p.
//
// http://www.khronos.org/opengles/sdk/docs/man3/html/glGetAttachedShaders.xhtml
func GetAttachedShaders(p Program) []Shader {
	shadersLen := GetProgrami(p, ATTACHED_SHADERS)
	var n int32
	buf := make([]uint32, shadersLen)
	gl.GetAttachedShaders(uint32(p.Value), int32(shadersLen), &n, &buf[0])
	buf = buf[:int(n)]
	shaders := make([]Shader, int(n))
	for i, s := range buf {
		shaders[i] = Shader{Value: uint32(s)}
	}
	fmt.Printf("GetAttachedShaders : %v -> %v\n", p, buf)
	if err := int(GetError()); err != NO_ERROR {
		fmt.Printf("ERROR: %v\n", err)
	}
	return shaders
}

// GetAttribLocation returns the location of an attribute variable.
//
// http://www.khronos.org/opengles/sdk/docs/man3/html/glGetAttribLocation.xhtml
func GetAttribLocation(p Program, name string) Attrib {
	fmt.Printf("GetAttribLocation : %v -> %v\n", p, name)
	return Attrib{Value: uint(gl.GetAttribLocation(p.Value, gl.Str(name+"\x00")))}
}

// GetBooleanv returns the boolean values of parameter pname.
//
// Many boolean parameters can be queried more easily using IsEnabled.
//
// http://www.khronos.org/opengles/sdk/docs/man3/html/glGet.xhtml
func GetBooleanv(dst []bool, pname Enum) {
	gl.GetBooleanv(uint32(pname), &dst[0])
	fmt.Printf("GetBooleanv : %v -> %v\n", pname, dst)
	if err := int(GetError()); err != NO_ERROR {
		fmt.Printf("ERROR: %v\n", err)
	}
}

// GetFloatv returns the float values of parameter pname.
//
// http://www.khronos.org/opengles/sdk/docs/man3/html/glGet.xhtml
func GetFloatv(dst []float32, pname Enum) {
	gl.GetFloatv(uint32(pname), &dst[0])
	fmt.Printf("GetFloatv : %v -> %v\n", pname, dst)
	if err := int(GetError()); err != NO_ERROR {
		fmt.Printf("ERROR: %v\n", err)
	}
}

// GetIntegerv returns the int values of parameter pname.
//
// Single values may be queried more easily using GetInteger.
//
// http://www.khronos.org/opengles/sdk/docs/man3/html/glGet.xhtml
func GetIntegerv(pname Enum, data []int32) {
	gl.GetIntegerv(uint32(pname), &data[0])
	fmt.Printf("GetIntegerv : %v -> %v\n", pname, data)
	if err := int(GetError()); err != NO_ERROR {
		fmt.Printf("ERROR: %v\n", err)
	}
}

// GetInteger returns the int value of parameter pname.
//
// http://www.khronos.org/opengles/sdk/docs/man3/html/glGet.xhtml
func GetInteger(pname Enum) int {
	var data int32
	gl.GetIntegerv(uint32(pname), &data)
	fmt.Printf("GetInteger : %v -> %v\n", pname, data)
	if err := int(GetError()); err != NO_ERROR {
		fmt.Printf("ERROR: %v\n", err)
	}
	return int(data)
}

// GetBufferParameteri returns a parameter for the active buffer.
//
// http://www.khronos.org/opengles/sdk/docs/man3/html/glGetBufferParameteriv.xhtml
func GetBufferParameteri(target, pname Enum) int {
	var params int32
	gl.GetBufferParameteriv(uint32(target), uint32(pname), &params)
	fmt.Printf("GetBufferParameteri : %v, %v -> %v\n", target, pname, params)
	if err := int(GetError()); err != NO_ERROR {
		fmt.Printf("ERROR: %v\n", err)
	}
	return int(params)
}

// GetError returns the next error.
//
// http://www.khronos.org/opengles/sdk/docs/man3/html/glGetError.xhtml
func GetError() Enum {
	return Enum(gl.GetError())
}

// GetBoundFramebuffer returns the currently bound framebuffer.
// Use this method instead of gl.GetInteger(gl.FRAMEBUFFER_BINDING) to
// enable support on all platforms
func GetBoundFramebuffer() Framebuffer {
	var b int32
	gl.GetIntegerv(FRAMEBUFFER_BINDING, &b)
	fmt.Printf("GetBoundFramebuffer -> %v\n", b)
	if err := int(GetError()); err != NO_ERROR {
		fmt.Printf("ERROR: %v\n", err)
	}
	return Framebuffer{Value: uint32(b)}
}

// GetFramebufferAttachmentParameteri returns attachment parameters
// for the active framebuffer object.
//
// http://www.khronos.org/opengles/sdk/docs/man3/html/glGetFramebufferAttachmentParameteriv.xhtml
func GetFramebufferAttachmentParameteri(target, attachment, pname Enum) int {
	var param int32
	gl.GetFramebufferAttachmentParameteriv(uint32(target), uint32(attachment), uint32(pname), &param)
	fmt.Printf("GetFramebufferAttachmentParameteri : %v, %v, %v -> %v\n", target, attachment, pname, param)
	if err := int(GetError()); err != NO_ERROR {
		fmt.Printf("ERROR: %v\n", err)
	}
	return int(param)
}

// GetProgrami returns a parameter value for a program.
//
// http://www.khronos.org/opengles/sdk/docs/man3/html/glGetProgramiv.xhtml
func GetProgrami(p Program, pname Enum) int {
	var result int32
	gl.GetProgramiv(p.Value, uint32(pname), &result)
	fmt.Printf("GetProgrami : %v, %v -> %v\n", p, pname, result)
	if err := int(GetError()); err != NO_ERROR {
		fmt.Printf("ERROR: %v\n", err)
	}
	return int(result)
}

// GetProgramInfoLog returns the information log for a program.
//
// http://www.khronos.org/opengles/sdk/docs/man3/html/glGetProgramInfoLog.xhtml
func GetProgramInfoLog(p Program) string {
	var logLength int32
	gl.GetProgramiv(p.Value, gl.INFO_LOG_LENGTH, &logLength)
	if logLength == 0 {
		return ""
	}

	logBuffer := make([]uint8, logLength)
	gl.GetProgramInfoLog(p.Value, logLength, nil, &logBuffer[0])
	fmt.Printf("GetProgramInfoLog : %v -> %v\n", p, gl.GoStr(&logBuffer[0]))
	if err := int(GetError()); err != NO_ERROR {
		fmt.Printf("ERROR: %v\n", err)
	}
	return gl.GoStr(&logBuffer[0])
}

// GetRenderbufferParameteri returns a parameter value for a render buffer.
//
// http://www.khronos.org/opengles/sdk/docs/man3/html/glGetRenderbufferParameteriv.xhtml
func GetRenderbufferParameteri(target, pname Enum) int {
	var result int32
	gl.GetRenderbufferParameteriv(uint32(target), uint32(pname), &result)
	fmt.Printf("GetRenderbufferParameteri : %v, %v -> %v\n", target, pname, result)
	if err := int(GetError()); err != NO_ERROR {
		fmt.Printf("ERROR: %v\n", err)
	}
	return int(result)
}

// GetRenderbufferParameteri returns a parameter value for a shader.
//
// http://www.khronos.org/opengles/sdk/docs/man3/html/glGetShaderiv.xhtml
func GetShaderi(s Shader, pname Enum) int {
	var result int32
	gl.GetShaderiv(s.Value, uint32(pname), &result)
	fmt.Printf("GetShaderi : %v, %v -> %v\n", s, pname, result)
	if err := int(GetError()); err != NO_ERROR {
		fmt.Printf("ERROR: %v\n", err)
	}
	return int(result)
}

// GetShaderInfoLog returns the information log for a shader.
//
// http://www.khronos.org/opengles/sdk/docs/man3/html/glGetShaderInfoLog.xhtml
func GetShaderInfoLog(s Shader) string {
	var logLength int32
	gl.GetShaderiv(s.Value, gl.INFO_LOG_LENGTH, &logLength)
	if logLength == 0 {
		return ""
	}

	logBuffer := make([]uint8, logLength)
	gl.GetShaderInfoLog(s.Value, logLength, nil, &logBuffer[0])
	fmt.Printf("GetShaderInfoLog : %v -> %v\n", s, gl.GoStr(&logBuffer[0]))
	if err := int(GetError()); err != NO_ERROR {
		fmt.Printf("ERROR: %v\n", err)
	}
	return gl.GoStr(&logBuffer[0])
}

// GetShaderPrecisionFormat returns range and precision limits for
// shader types.
//
// http://www.khronos.org/opengles/sdk/docs/man3/html/glGetShaderPrecisionFormat.xhtml
func GetShaderPrecisionFormat(shadertype, precisiontype Enum) (rangeLow, rangeHigh, precision int) {
	var cRange [2]int32
	var cPrecision int32

	gl.GetShaderPrecisionFormat(uint32(shadertype), uint32(precisiontype), &cRange[0], &cPrecision)
	fmt.Printf("GetShaderPrecisionFormat : %v, %v -> %v, %v, %v\n", shadertype, precisiontype, cRange[0], cRange[1], cPrecision)
	if err := int(GetError()); err != NO_ERROR {
		fmt.Printf("ERROR: %v\n", err)
	}
	return int(cRange[0]), int(cRange[1]), int(cPrecision)
}

// GetShaderSource returns source code of shader s.
//
// http://www.khronos.org/opengles/sdk/docs/man3/html/glGetShaderSource.xhtml
func GetShaderSource(s Shader) string {
	sourceLen := GetShaderi(s, gl.SHADER_SOURCE_LENGTH)
	if sourceLen == 0 {
		return ""
	}
	buf := make([]byte, sourceLen)
	gl.GetShaderSource(s.Value, int32(sourceLen), nil, &buf[0])
	fmt.Printf("GetShaderSource : %v -> %v\n", s, string(buf))
	if err := int(GetError()); err != NO_ERROR {
		fmt.Printf("ERROR: %v\n", err)
	}
	return string(buf)
}

// GetString reports current GL state.
//
// Valid name values:
//	EXTENSIONS
//	RENDERER
//	SHADING_LANGUAGE_VERSION
//	VENDOR
//	VERSION
//
// http://www.khronos.org/opengles/sdk/docs/man3/html/glGetString.xhtml
func GetString(pname Enum) string {
	str := gl.GoStr(gl.GetString(uint32(pname)))
	fmt.Printf("GetString : %v -> %v\n", pname, str)
	if err := int(GetError()); err != NO_ERROR {
		fmt.Printf("ERROR: %v\n", err)
	}
	return str
}

// GetTexParameterfv returns the float values of a texture parameter.
//
// http://www.khronos.org/opengles/sdk/docs/man3/html/glGetTexParameter.xhtml
func GetTexParameterfv(dst []float32, target, pname Enum) {
	gl.GetTexParameterfv(uint32(target), uint32(pname), &dst[0])
	fmt.Printf("GetTexParameterfv : %v, %v -> %v\n", target, pname, dst[0])
	if err := int(GetError()); err != NO_ERROR {
		fmt.Printf("ERROR: %v\n", err)
	}
}

// GetTexParameteriv returns the int values of a texture parameter.
//
// http://www.khronos.org/opengles/sdk/docs/man3/html/glGetTexParameter.xhtml
func GetTexParameteriv(dst []int32, target, pname Enum) {
	gl.GetTexParameteriv(uint32(target), uint32(pname), &dst[0])
	fmt.Printf("GetTexParameteriv : %v, %v -> %v\n", target, pname, dst[0])
	if err := int(GetError()); err != NO_ERROR {
		fmt.Printf("ERROR: %v\n", err)
	}
}

// GetUniformfv returns the float values of a uniform variable.
//
// http://www.khronos.org/opengles/sdk/docs/man3/html/glGetUniform.xhtml
func GetUniformfv(dst []float32, src Uniform, p Program) {
	gl.GetUniformfv(p.Value, src.Value, &dst[0])
	fmt.Printf("GetUniformfv : %v, %v -> %v\n", src, p, dst[0])
	if err := int(GetError()); err != NO_ERROR {
		fmt.Printf("ERROR: %v\n", err)
	}
}

// GetUniformiv returns the float values of a uniform variable.
//
// http://www.khronos.org/opengles/sdk/docs/man3/html/glGetUniform.xhtml
func GetUniformiv(dst []int32, src Uniform, p Program) {
	gl.GetUniformiv(p.Value, src.Value, &dst[0])
	fmt.Printf("GetUniformiv : %v, %v -> %v\n", src, p, dst[0])
	if err := int(GetError()); err != NO_ERROR {
		fmt.Printf("ERROR: %v\n", err)
	}
}

// GetUniformLocation returns the location of a uniform variable.
//
// http://www.khronos.org/opengles/sdk/docs/man3/html/glGetUniformLocation.xhtml
func GetUniformLocation(p Program, name string) Uniform {
	loc := Uniform{Value: gl.GetUniformLocation(p.Value, gl.Str(name+"\x00"))}
	fmt.Printf("GetUniformLocation : %v, %v -> %v\n", p, name, loc)
	if err := int(GetError()); err != NO_ERROR {
		fmt.Printf("ERROR: %v\n", err)
	}
	return loc
}

// GetVertexAttribf reads the float value of a vertex attribute.
//
// http://www.khronos.org/opengles/sdk/docs/man3/html/glGetVertexAttrib.xhtml
func GetVertexAttribf(src Attrib, pname Enum) float32 {
	var result float32
	gl.GetVertexAttribfv(uint32(src.Value), uint32(pname), &result)
	fmt.Printf("GetVertexAttribf : %v, %v -> %v\n", src, pname, result)
	if err := int(GetError()); err != NO_ERROR {
		fmt.Printf("ERROR: %v\n", err)
	}
	return result
}

// GetVertexAttribfv reads float values of a vertex attribute.
//
// http://www.khronos.org/opengles/sdk/docs/man3/html/glGetVertexAttrib.xhtml
func GetVertexAttribfv(dst []float32, src Attrib, pname Enum) {
	gl.GetVertexAttribfv(uint32(src.Value), uint32(pname), &dst[0])
	fmt.Printf("GetVertexAttribfv : %v, %v -> %v\n", src, pname, dst[0])
	if err := int(GetError()); err != NO_ERROR {
		fmt.Printf("ERROR: %v\n", err)
	}
}

// GetVertexAttribi reads the int value of a vertex attribute.
//
// http://www.khronos.org/opengles/sdk/docs/man3/html/glGetVertexAttrib.xhtml
func GetVertexAttribi(src Attrib, pname Enum) int32 {
	var result int32
	gl.GetVertexAttribiv(uint32(src.Value), uint32(pname), &result)
	fmt.Printf("GetVertexAttribi : %v, %v -> %v\n", src, pname, result)
	if err := int(GetError()); err != NO_ERROR {
		fmt.Printf("ERROR: %v\n", err)
	}
	return result
}

// GetVertexAttribiv reads int values of a vertex attribute.
//
// http://www.khronos.org/opengles/sdk/docs/man3/html/glGetVertexAttrib.xhtml
func GetVertexAttribiv(dst []int32, src Attrib, pname Enum) {
	gl.GetVertexAttribiv(uint32(src.Value), uint32(pname), &dst[0])
	fmt.Printf("GetVertexAttribiv : %v, %v -> %v\n", src, pname, dst[0])
	if err := int(GetError()); err != NO_ERROR {
		fmt.Printf("ERROR: %v\n", err)
	}
}

// Hint sets implementation-specific modes.
//
// http://www.khronos.org/opengles/sdk/docs/man3/html/glHint.xhtml
func Hint(target, mode Enum) {
	fmt.Printf("Hint : %v, %v\n", target, mode)
	gl.Hint(uint32(target), uint32(mode))
	if err := int(GetError()); err != NO_ERROR {
		fmt.Printf("ERROR: %v\n", err)
	}
}

// IsBuffer reports if b is a valid buffer.
//
// http://www.khronos.org/opengles/sdk/docs/man3/html/glIsBuffer.xhtml
func IsBuffer(b Buffer) bool {
	r := gl.IsBuffer(b.Value)
	fmt.Printf("IsBuffer : %v -> %v\n", b, r)
	if err := int(GetError()); err != NO_ERROR {
		fmt.Printf("ERROR: %v\n", err)
	}
	return r
}

// IsEnabled reports if cap is an enabled capability.
//
// http://www.khronos.org/opengles/sdk/docs/man3/html/glIsEnabled.xhtml
func IsEnabled(cap Enum) bool {
	r := gl.IsEnabled(uint32(cap))
	fmt.Printf("IsEnabled : %v -> %v\n", cap, r)
	if err := int(GetError()); err != NO_ERROR {
		fmt.Printf("ERROR: %v\n", err)
	}
	return r
}

// IsFramebuffer reports if fb is a valid frame buffer.
//
// http://www.khronos.org/opengles/sdk/docs/man3/html/glIsFramebuffer.xhtml
func IsFramebuffer(fb Framebuffer) bool {
	r := gl.IsFramebuffer(fb.Value)
	fmt.Printf("IsFramebuffer : %v -> %v\n", fb, r)
	if err := int(GetError()); err != NO_ERROR {
		fmt.Printf("ERROR: %v\n", err)
	}
	return r
}

// IsProgram reports if p is a valid program object.
//
// http://www.khronos.org/opengles/sdk/docs/man3/html/glIsProgram.xhtml
func IsProgram(p Program) bool {
	r := gl.IsFramebuffer(p.Value)
	fmt.Printf("IsProgram : %v -> %v\n", p, r)
	if err := int(GetError()); err != NO_ERROR {
		fmt.Printf("ERROR: %v\n", err)
	}
	return r
}

// IsRenderbuffer reports if rb is a valid render buffer.
//
// http://www.khronos.org/opengles/sdk/docs/man3/html/glIsRenderbuffer.xhtml
func IsRenderbuffer(rb Renderbuffer) bool {
	r := gl.IsRenderbuffer(rb.Value)
	fmt.Printf("IsProgram : %v -> %v\n", rb, r)
	if err := int(GetError()); err != NO_ERROR {
		fmt.Printf("ERROR: %v\n", err)
	}
	return r
}

// IsShader reports if s is valid shader.
//
// http://www.khronos.org/opengles/sdk/docs/man3/html/glIsShader.xhtml
func IsShader(s Shader) bool {
	r := gl.IsShader(s.Value)
	fmt.Printf("IsShader : %v -> %v\n", s, r)
	if err := int(GetError()); err != NO_ERROR {
		fmt.Printf("ERROR: %v\n", err)
	}
	return r
}

// IsTexture reports if t is a valid texture.
//
// http://www.khronos.org/opengles/sdk/docs/man3/html/glIsTexture.xhtml
func IsTexture(t Texture) bool {
	r := gl.IsTexture(t.Value)
	fmt.Printf("IsTexture : %v -> %v\n", t, r)
	if err := int(GetError()); err != NO_ERROR {
		fmt.Printf("ERROR: %v\n", err)
	}
	return r
}

// LineWidth specifies the width of lines.
//
// http://www.khronos.org/opengles/sdk/docs/man3/html/glLineWidth.xhtml
func LineWidth(width float32) {
	fmt.Printf("LineWidth : %v\n", width)
	if err := int(GetError()); err != NO_ERROR {
		fmt.Printf("ERROR: %v\n", err)
	}
	gl.LineWidth(width)
}

// LinkProgram links the specified program.
//
// http://www.khronos.org/opengles/sdk/docs/man3/html/glLinkProgram.xhtml
func LinkProgram(p Program) {
	fmt.Printf("LinkProgram : %v\n", p)
	if err := int(GetError()); err != NO_ERROR {
		fmt.Printf("ERROR: %v\n", err)
	}
	gl.LinkProgram(p.Value)
}

// PixelStorei sets pixel storage parameters.
//
// http://www.khronos.org/opengles/sdk/docs/man3/html/glPixelStorei.xhtml
func PixelStorei(pname Enum, param int32) {
	fmt.Printf("PixelStorei : %v, %v\n", pname, param)
	gl.PixelStorei(uint32(pname), param)
	if err := int(GetError()); err != NO_ERROR {
		fmt.Printf("ERROR: %v\n", err)
	}
}

// PolygonMode sets Polygon Mode.
//
// http://www.khronos.org/opengles/sdk/docs/man3/html/glPolygonMode.xhtml
func PolygonMode(face, mode Enum) {
	fmt.Printf("PolygonMode : %v, %v\n", face, mode)
	gl.PolygonMode(uint32(face), uint32(mode))
	if err := int(GetError()); err != NO_ERROR {
		fmt.Printf("ERROR: %v\n", err)
	}
}

// PolygonOffset sets the scaling factors for depth offsets.
//
// http://www.khronos.org/opengles/sdk/docs/man3/html/glPolygonOffset.xhtml
func PolygonOffset(factor, units float32) {
	fmt.Printf("PolygonOffset : %v, %v\n", factor, units)
	gl.PolygonOffset(factor, units)
	if err := int(GetError()); err != NO_ERROR {
		fmt.Printf("ERROR: %v\n", err)
	}
}

// ReadPixels returns pixel data from a buffer.
//
// http://www.khronos.org/opengles/sdk/docs/man3/html/glReadPixels.xhtml
func ReadPixels(dst []byte, x, y, width, height int, format, ty Enum) {
	gl.ReadPixels(int32(x), int32(y), int32(width), int32(height), uint32(format), uint32(ty), gl.Ptr(&dst[0]))
	fmt.Printf("ReadPixels : %v, %v, %v, %v, %v, %v -> %v\n", x, y, width, height, format, ty, len(dst))
	if err := int(GetError()); err != NO_ERROR {
		fmt.Printf("ERROR: %v\n", err)
	}
}

// ReleaseShaderCompiler frees resources allocated by the shader compiler.
//
// http://www.khronos.org/opengles/sdk/docs/man3/html/glReleaseShaderCompiler.xhtml
func ReleaseShaderCompiler() {
	fmt.Printf("ReleaseShaderCompiler\n")
	gl.ReleaseShaderCompiler()
	if err := int(GetError()); err != NO_ERROR {
		fmt.Printf("ERROR: %v\n", err)
	}
}

// RenderbufferStorage establishes the data storage, format, and
// dimensions of a renderbuffer object's image.
//
// http://www.khronos.org/opengles/sdk/docs/man3/html/glRenderbufferStorage.xhtml
func RenderbufferStorage(target, internalFormat Enum, width, height int) {
	fmt.Printf("RenderbufferStorage : %v, %v, %v, %v\n", target, internalFormat, width, height)
	gl.RenderbufferStorage(uint32(target), uint32(internalFormat), int32(width), int32(height))
	if err := int(GetError()); err != NO_ERROR {
		fmt.Printf("ERROR: %v\n", err)
	}
}

// SampleCoverage sets multisample coverage parameters.
//
// http://www.khronos.org/opengles/sdk/docs/man3/html/glSampleCoverage.xhtml
func SampleCoverage(value float32, invert bool) {
	fmt.Printf("SampleCoverage : %v, %v\n", value, invert)
	gl.SampleCoverage(value, invert)
	if err := int(GetError()); err != NO_ERROR {
		fmt.Printf("ERROR: %v\n", err)
	}
}

// Scissor defines the scissor box rectangle, in window coordinates.
//
// http://www.khronos.org/opengles/sdk/docs/man3/html/glScissor.xhtml
func Scissor(x, y, width, height int32) {
	fmt.Printf("Scissor : %v, %v, %v, %v\n", x, y, width, height)
	gl.Scissor(x, y, width, height)
	if err := int(GetError()); err != NO_ERROR {
		fmt.Printf("ERROR: %v\n", err)
	}
}

// ShaderSource sets the source code of s to the given source code.
//
// http://www.khronos.org/opengles/sdk/docs/man3/html/glShaderSource.xhtml
func ShaderSource(s Shader, src string) {
	fmt.Printf("ShaderSource : %v, %v\n", s, src)
	glsource, free := gl.Strs(src + "\x00")
	gl.ShaderSource(s.Value, 1, glsource, nil)
	if err := int(GetError()); err != NO_ERROR {
		fmt.Printf("ERROR: %v\n", err)
	}
	free()
}

// StencilFunc sets the front and back stencil test reference value.
//
// http://www.khronos.org/opengles/sdk/docs/man3/html/glStencilFunc.xhtml
func StencilFunc(fn Enum, ref int, mask uint32) {
	fmt.Printf("ActivStencilFunceTexture : %v, %v, %v\n", fn, ref, mask)
	gl.StencilFunc(uint32(fn), int32(ref), mask)
	if err := int(GetError()); err != NO_ERROR {
		fmt.Printf("ERROR: %v\n", err)
	}
}

// StencilFunc sets the front or back stencil test reference value.
//
// http://www.khronos.org/opengles/sdk/docs/man3/html/glStencilFuncSeparate.xhtml
func StencilFuncSeparate(face, fn Enum, ref int, mask uint32) {
	fmt.Printf("StencilFuncSeparate : %v, %v, %v, %v\n", face, fn, ref, mask)
	gl.StencilFuncSeparate(uint32(face), uint32(fn), int32(ref), mask)
	if err := int(GetError()); err != NO_ERROR {
		fmt.Printf("ERROR: %v\n", err)
	}
}

// StencilMask controls the writing of bits in the stencil planes.
//
// http://www.khronos.org/opengles/sdk/docs/man3/html/glStencilMask.xhtml
func StencilMask(mask uint32) {
	fmt.Printf("StencilMask : %v\n", mask)
	gl.StencilMask(mask)
	if err := int(GetError()); err != NO_ERROR {
		fmt.Printf("ERROR: %v\n", err)
	}
}

// StencilMaskSeparate controls the writing of bits in the stencil planes.
//
// http://www.khronos.org/opengles/sdk/docs/man3/html/glStencilMaskSeparate.xhtml
func StencilMaskSeparate(face Enum, mask uint32) {
	fmt.Printf("StencilMaskSeparate : %v, %v\n", face, mask)
	gl.StencilMaskSeparate(uint32(face), mask)
	if err := int(GetError()); err != NO_ERROR {
		fmt.Printf("ERROR: %v\n", err)
	}
}

// StencilOp sets front and back stencil test actions.
//
// http://www.khronos.org/opengles/sdk/docs/man3/html/glStencilOp.xhtml
func StencilOp(fail, zfail, zpass Enum) {
	fmt.Printf("StencilOp : %v, %v, %v\n", fail, zfail, zpass)
	gl.StencilOp(uint32(fail), uint32(zfail), uint32(zpass))
	if err := int(GetError()); err != NO_ERROR {
		fmt.Printf("ERROR: %v\n", err)
	}
}

// StencilOpSeparate sets front or back stencil tests.
//
// http://www.khronos.org/opengles/sdk/docs/man3/html/glStencilOpSeparate.xhtml
func StencilOpSeparate(face, sfail, dpfail, dppass Enum) {
	fmt.Printf("StencilOpSeparate : %v, %v, %v ,%v\n", face, sfail, dpfail, dppass)
	gl.StencilOpSeparate(uint32(face), uint32(sfail), uint32(dpfail), uint32(dppass))
	if err := int(GetError()); err != NO_ERROR {
		fmt.Printf("ERROR: %v\n", err)
	}
}

// TexImage2D writes a 2D texture image.
//
// http://www.khronos.org/opengles/sdk/docs/man3/html/glTexImage2D.xhtml
func TexImage2D(target Enum, level int, width, height int, format Enum, ty Enum, data []byte) {
	fmt.Printf("TexImage2D : %v, %v, %v, %v, %v, %v, %v\n", target, level, width, height, format, ty, len(data))
	p := unsafe.Pointer(nil)
	if len(data) > 0 {
		p = gl.Ptr(&data[0])
	}
	gl.TexImage2D(uint32(target), int32(level), int32(format), int32(width), int32(height), 0, uint32(format), uint32(ty), p)
	if err := int(GetError()); err != NO_ERROR {
		fmt.Printf("ERROR: %v\n", err)
	}
}

// TexSubImage2D writes a subregion of a 2D texture image.
//
// http://www.khronos.org/opengles/sdk/docs/man3/html/glTexSubImage2D.xhtml
func TexSubImage2D(target Enum, level int, x, y, width, height int, format, ty Enum, data []byte) {
	fmt.Printf("TexSubImage2D : %v, %v, %v, %v, %v, %v, %v\n", target, level, width, height, format, ty, len(data))
	gl.TexSubImage2D(uint32(target), int32(level), int32(x), int32(y), int32(width), int32(height), uint32(format), uint32(ty), gl.Ptr(&data[0]))
	if err := int(GetError()); err != NO_ERROR {
		fmt.Printf("ERROR: %v\n", err)
	}
}

// TexParameterf sets a float texture parameter.
//
// http://www.khronos.org/opengles/sdk/docs/man3/html/glTexParameter.xhtml
func TexParameterf(target, pname Enum, param float32) {
	fmt.Printf("TexParameterf : %v, %v, %v\n", target, pname, param)
	gl.TexParameterf(uint32(target), uint32(pname), param)
	if err := int(GetError()); err != NO_ERROR {
		fmt.Printf("ERROR: %v\n", err)
	}
}

// TexParameterfv sets a float texture parameter array.
//
// http://www.khronos.org/opengles/sdk/docs/man3/html/glTexParameter.xhtml
func TexParameterfv(target, pname Enum, params []float32) {
	fmt.Printf("TexParameterfv : %v, %v, %v\n", target, pname, params)
	gl.TexParameterfv(uint32(target), uint32(pname), &params[0])
	if err := int(GetError()); err != NO_ERROR {
		fmt.Printf("ERROR: %v\n", err)
	}
}

// TexParameteri sets an integer texture parameter.
//
// http://www.khronos.org/opengles/sdk/docs/man3/html/glTexParameter.xhtml
func TexParameteri(target, pname Enum, param int) {
	fmt.Printf("TexParameteri : %v, %v, %v\n", target, pname, param)
	gl.TexParameteri(uint32(target), uint32(pname), int32(param))
	if err := int(GetError()); err != NO_ERROR {
		fmt.Printf("ERROR: %v\n", err)
	}
}

// TexParameteriv sets an integer texture parameter array.
//
// http://www.khronos.org/opengles/sdk/docs/man3/html/glTexParameter.xhtml
func TexParameteriv(target, pname Enum, params []int32) {
	fmt.Printf("TexParameteriv : %v, %v, %v\n", target, pname, params)
	gl.TexParameteriv(uint32(target), uint32(pname), &params[0])
	if err := int(GetError()); err != NO_ERROR {
		fmt.Printf("ERROR: %v\n", err)
	}
}

// Uniform1f writes a float uniform variable.
//
// http://www.khronos.org/opengles/sdk/docs/man3/html/glUniform.xhtml
func Uniform1f(dst Uniform, v float32) {
	if dst.Valid() {
		fmt.Printf("Uniform1f : %v, %v\n", dst, v)
		gl.Uniform1f(dst.Value, v)
		if err := int(GetError()); err != NO_ERROR {
			fmt.Printf("ERROR: %v\n", err)
		}
	}
}

// Uniform1fv writes a [len(src)]float uniform array.
//
// http://www.khronos.org/opengles/sdk/docs/man3/html/glUniform.xhtml
func Uniform1fv(dst Uniform, src []float32) {
	if dst.Valid() {
		fmt.Printf("Uniform1fv : %v, %v\n", dst, src)
		gl.Uniform1fv(dst.Value, int32(len(src)), &src[0])
		if err := int(GetError()); err != NO_ERROR {
			fmt.Printf("ERROR: %v\n", err)
		}
	}
}

// Uniform1i writes an int uniform variable.
//
// Uniform1i and Uniform1iv are the only two functions that may be used
// to load uniform variables defined as sampler types. Loading samplers
// with any other function will result in a INVALID_OPERATION error.
//
// http://www.khronos.org/opengles/sdk/docs/man3/html/glUniform.xhtml
func Uniform1i(dst Uniform, v int) {
	if dst.Valid() {
		fmt.Printf("Uniform1i : %v, %v\n", dst, v)
		gl.Uniform1i(dst.Value, int32(v))
		if err := int(GetError()); err != NO_ERROR {
			fmt.Printf("ERROR: %v\n", err)
		}
	}
}

// Uniform1iv writes a int uniform array of len(src) elements.
//
// Uniform1i and Uniform1iv are the only two functions that may be used
// to load uniform variables defined as sampler types. Loading samplers
// with any other function will result in a INVALID_OPERATION error.
//
// http://www.khronos.org/opengles/sdk/docs/man3/html/glUniform.xhtml
func Uniform1iv(dst Uniform, src []int32) {
	if dst.Valid() {
		fmt.Printf("Uniform1iv : %v, %v\n", dst, src)
		gl.Uniform1iv(dst.Value, int32(len(src)), &src[0])
		if err := int(GetError()); err != NO_ERROR {
			fmt.Printf("ERROR: %v\n", err)
		}
	}
}

// Uniform2f writes a vec2 uniform variable.
//
// http://www.khronos.org/opengles/sdk/docs/man3/html/glUniform.xhtml
func Uniform2f(dst Uniform, v0, v1 float32) {
	if dst.Valid() {
		fmt.Printf("Uniform2f : %v, %v, %v\n", dst, v0, v1)
		gl.Uniform2f(dst.Value, v0, v1)
		if err := int(GetError()); err != NO_ERROR {
			fmt.Printf("ERROR: %v\n", err)
		}
	}
}

// Uniform2fv writes a vec2 uniform array of len(src)/2 elements.
//
// http://www.khronos.org/opengles/sdk/docs/man3/html/glUniform.xhtml
func Uniform2fv(dst Uniform, src []float32) {
	if dst.Valid() {
		fmt.Printf("Uniform2fv : %v, %v\n", dst, src)
		gl.Uniform2fv(dst.Value, int32(len(src)/2), &src[0])
		if err := int(GetError()); err != NO_ERROR {
			fmt.Printf("ERROR: %v\n", err)
		}
	}
}

// Uniform2i writes an ivec2 uniform variable.
//
// http://www.khronos.org/opengles/sdk/docs/man3/html/glUniform.xhtml
func Uniform2i(dst Uniform, v0, v1 int) {
	if dst.Valid() {
		fmt.Printf("Uniform2i : %v, %v, %v\n", dst, v0, v1)
		gl.Uniform2i(dst.Value, int32(v0), int32(v1))
		if err := int(GetError()); err != NO_ERROR {
			fmt.Printf("ERROR: %v\n", err)
		}
	}
}

// Uniform2iv writes an ivec2 uniform array of len(src)/2 elements.
//
// http://www.khronos.org/opengles/sdk/docs/man3/html/glUniform.xhtml
func Uniform2iv(dst Uniform, src []int32) {
	if dst.Valid() {
		fmt.Printf("Uniform2iv : %v, %v\n", dst, src)
		gl.Uniform2iv(dst.Value, int32(len(src)/2), &src[0])
		if err := int(GetError()); err != NO_ERROR {
			fmt.Printf("ERROR: %v\n", err)
		}
	}
}

// Uniform3f writes a vec3 uniform variable.
//
// http://www.khronos.org/opengles/sdk/docs/man3/html/glUniform.xhtml
func Uniform3f(dst Uniform, v0, v1, v2 float32) {
	if dst.Valid() {
		fmt.Printf("Uniform3f : %v, %v, %v, %v\n", dst, v0, v1, v2)
		gl.Uniform3f(dst.Value, v0, v1, v2)
		if err := int(GetError()); err != NO_ERROR {
			fmt.Printf("ERROR: %v\n", err)
		}
	}
}

// Uniform3fv writes a vec3 uniform array of len(src)/3 elements.
//
// http://www.khronos.org/opengles/sdk/docs/man3/html/glUniform.xhtml
func Uniform3fv(dst Uniform, src []float32) {
	if dst.Valid() {
		fmt.Printf("Uniform3fv : %v, %v\n", dst, src)
		gl.Uniform3fv(dst.Value, int32(len(src)), &src[0])
		if err := int(GetError()); err != NO_ERROR {
			fmt.Printf("ERROR: %v\n", err)
		}
	}
}

// Uniform3i writes an ivec3 uniform variable.
//
// http://www.khronos.org/opengles/sdk/docs/man3/html/glUniform.xhtml
func Uniform3i(dst Uniform, v0, v1, v2 int32) {
	if dst.Valid() {
		fmt.Printf("Uniform3i : %v, %v, %v, %v\n", dst, v0, v1, v2)
		gl.Uniform3i(dst.Value, v0, v1, v2)
		if err := int(GetError()); err != NO_ERROR {
			fmt.Printf("ERROR: %v\n", err)
		}
	}
}

// Uniform3iv writes an ivec3 uniform array of len(src)/3 elements.
//
// http://www.khronos.org/opengles/sdk/docs/man3/html/glUniform.xhtml
func Uniform3iv(dst Uniform, src []int32) {
	if dst.Valid() {
		fmt.Printf("Uniform3iv : %v, %v\n", dst, src)
		gl.Uniform3iv(dst.Value, int32(len(src)/3), &src[0])
		if err := int(GetError()); err != NO_ERROR {
			fmt.Printf("ERROR: %v\n", err)
		}
	}
}

// Uniform4f writes a vec4 uniform variable.
//
// http://www.khronos.org/opengles/sdk/docs/man3/html/glUniform.xhtml
func Uniform4f(dst Uniform, v0, v1, v2, v3 float32) {
	if dst.Valid() {
		fmt.Printf("Uniform3i : %v, %v, %v, %v, %v\n", dst, v0, v1, v2, v3)
		gl.Uniform4f(dst.Value, v0, v1, v2, v3)
		if err := int(GetError()); err != NO_ERROR {
			fmt.Printf("ERROR: %v\n", err)
		}
	}
}

// Uniform4fv writes a vec4 uniform array of len(src)/4 elements.
//
// http://www.khronos.org/opengles/sdk/docs/man3/html/glUniform.xhtml
func Uniform4fv(dst Uniform, src []float32) {
	if dst.Valid() {
		fmt.Printf("Uniform4fv : %v, %v\n", dst, src)
		gl.Uniform4fv(dst.Value, int32(len(src)/4), (*float32)(unsafe.Pointer(&src[0])))
		if err := int(GetError()); err != NO_ERROR {
			fmt.Printf("ERROR: %v\n", err)
		}
	}
}

// Uniform4i writes an ivec4 uniform variable.
//
// http://www.khronos.org/opengles/sdk/docs/man3/html/glUniform.xhtml
func Uniform4i(dst Uniform, v0, v1, v2, v3 int32) {
	if dst.Valid() {
		fmt.Printf("Uniform4i : %v, %v, %v, %v, %v\n", dst, v0, v1, v2, v3)
		gl.Uniform4i(dst.Value, v0, v1, v2, v3)
		if err := int(GetError()); err != NO_ERROR {
			fmt.Printf("ERROR: %v\n", err)
		}
	}
}

// Uniform4i writes an ivec4 uniform array of len(src)/4 elements.
//
// http://www.khronos.org/opengles/sdk/docs/man3/html/glUniform.xhtml
func Uniform4iv(dst Uniform, src []int32) {
	if dst.Valid() {
		fmt.Printf("Uniform4iv : %v, %v\n", dst, src)
		gl.Uniform4iv(dst.Value, int32(len(src)/4), &src[0])
		if err := int(GetError()); err != NO_ERROR {
			fmt.Printf("ERROR: %v\n", err)
		}
	}
}

// UniformMatrix2fv writes 2x2 matrices. Each matrix uses four
// float32 values, so the number of matrices written is len(src)/4.
//
// Each matrix must be supplied in column major order.
//
// http://www.khronos.org/opengles/sdk/docs/man3/html/glUniform.xhtml
func UniformMatrix2fv(dst Uniform, src []float32) {
	if dst.Valid() {
		fmt.Printf("UniformMatrix2fv : %v, %v\n", dst, src)
		gl.UniformMatrix2fv(dst.Value, int32(len(src)/(2*2)), false, &src[0])
		if err := int(GetError()); err != NO_ERROR {
			fmt.Printf("ERROR: %v\n", err)
		}
	}
}

// UniformMatrix3fv writes 3x3 matrices. Each matrix uses nine
// float32 values, so the number of matrices written is len(src)/9.
//
// Each matrix must be supplied in column major order.
//
// http://www.khronos.org/opengles/sdk/docs/man3/html/glUniform.xhtml
func UniformMatrix3fv(dst Uniform, src []float32) {
	if dst.Valid() {
		fmt.Printf("UniformMatrix3fv : %v, %v\n", dst, src)
		gl.UniformMatrix3fv(dst.Value, int32(len(src)/(3*3)), false, &src[0])
		if err := int(GetError()); err != NO_ERROR {
			fmt.Printf("ERROR: %v\n", err)
		}
	}
}

// UniformMatrix4fv writes 4x4 matrices. Each matrix uses 16
// float32 values, so the number of matrices written is len(src)/16.
//
// Each matrix must be supplied in column major order.
//
// http://www.khronos.org/opengles/sdk/docs/man3/html/glUniform.xhtml
func UniformMatrix4fv(dst Uniform, src []float32) {
	if dst.Valid() {
		fmt.Printf("UniformMatrix4fv : %v, %v\n", dst, src)
		gl.UniformMatrix4fv(dst.Value, int32(len(src)/(4*4)), false, &src[0])
		if err := int(GetError()); err != NO_ERROR {
			fmt.Printf("ERROR: %v\n", err)
		}
	}
}

// UseProgram sets the active program.
//
// http://www.khronos.org/opengles/sdk/docs/man3/html/glUseProgram.xhtml
func UseProgram(p Program) {
	fmt.Printf("UseProgram : %v\n", p)
	gl.UseProgram(p.Value)
	if err := int(GetError()); err != NO_ERROR {
		fmt.Printf("ERROR: %v\n", err)
	}
}

// ValidateProgram checks to see whether the executables contained in
// program can execute given the current OpenGL state.
//
// Typically only used for debugging.
//
// http://www.khronos.org/opengles/sdk/docs/man3/html/glValidateProgram.xhtml
func ValidateProgram(p Program) {
	fmt.Printf("ValidateProgram : %v\n", p)
	gl.ValidateProgram(uint32(p.Value))
	if err := int(GetError()); err != NO_ERROR {
		fmt.Printf("ERROR: %v\n", err)
	}
}

// VertexAttrib1f writes a float vertex attribute.
//
// http://www.khronos.org/opengles/sdk/docs/man3/html/glVertexAttrib.xhtml
func VertexAttrib1f(dst Attrib, x float32) {
	fmt.Printf("VertexAttrib1f : %v, %v\n", dst, x)
	gl.VertexAttrib1f(uint32(dst.Value), x)
	if err := int(GetError()); err != NO_ERROR {
		fmt.Printf("ERROR: %v\n", err)
	}
}

// VertexAttrib1fv writes a float vertex attribute.
//
// http://www.khronos.org/opengles/sdk/docs/man3/html/glVertexAttrib.xhtml
func VertexAttrib1fv(dst Attrib, src []float32) {
	fmt.Printf("VertexAttrib1fv : %v, %v\n", dst, src)
	gl.VertexAttrib1fv(uint32(dst.Value), &src[0])
	if err := int(GetError()); err != NO_ERROR {
		fmt.Printf("ERROR: %v\n", err)
	}
}

// VertexAttrib2f writes a vec2 vertex attribute.
//
// http://www.khronos.org/opengles/sdk/docs/man3/html/glVertexAttrib.xhtml
func VertexAttrib2f(dst Attrib, x, y float32) {
	fmt.Printf("VertexAttrib2f : %v, %v, %v\n", dst, x, y)
	gl.VertexAttrib2f(uint32(dst.Value), x, y)
	if err := int(GetError()); err != NO_ERROR {
		fmt.Printf("ERROR: %v\n", err)
	}
}

// VertexAttrib2fv writes a vec2 vertex attribute.
//
// http://www.khronos.org/opengles/sdk/docs/man3/html/glVertexAttrib.xhtml
func VertexAttrib2fv(dst Attrib, src []float32) {
	fmt.Printf("VertexAttrib2fv : %v, %v\n", dst, src)
	gl.VertexAttrib2fv(uint32(dst.Value), &src[0])
	if err := int(GetError()); err != NO_ERROR {
		fmt.Printf("ERROR: %v\n", err)
	}
}

// VertexAttrib3f writes a vec3 vertex attribute.
//
// http://www.khronos.org/opengles/sdk/docs/man3/html/glVertexAttrib.xhtml
func VertexAttrib3f(dst Attrib, x, y, z float32) {
	fmt.Printf("VertexAttrib3f : %v, %v, %v, %v\n", dst, x, y, z)
	gl.VertexAttrib3f(uint32(dst.Value), x, y, z)
	if err := int(GetError()); err != NO_ERROR {
		fmt.Printf("ERROR: %v\n", err)
	}
}

// VertexAttrib3fv writes a vec3 vertex attribute.
//
// http://www.khronos.org/opengles/sdk/docs/man3/html/glVertexAttrib.xhtml
func VertexAttrib3fv(dst Attrib, src []float32) {
	fmt.Printf("VertexAttrib3fv : %v, %v\n", dst, src)
	gl.VertexAttrib3fv(uint32(dst.Value), &src[0])
	if err := int(GetError()); err != NO_ERROR {
		fmt.Printf("ERROR: %v\n", err)
	}
}

// VertexAttrib4f writes a vec4 vertex attribute.
//
// http://www.khronos.org/opengles/sdk/docs/man3/html/glVertexAttrib.xhtml
func VertexAttrib4f(dst Attrib, x, y, z, w float32) {
	fmt.Printf("VertexAttrib4f : %v, %v, %v, %v, %v\n", dst, x, y, z, w)
	gl.VertexAttrib4f(uint32(dst.Value), x, y, z, w)
	if err := int(GetError()); err != NO_ERROR {
		fmt.Printf("ERROR: %v\n", err)
	}
}

// VertexAttrib4fv writes a vec4 vertex attribute.
//
// http://www.khronos.org/opengles/sdk/docs/man3/html/glVertexAttrib.xhtml
func VertexAttrib4fv(dst Attrib, src []float32) {
	fmt.Printf("VertexAttrib4fv : %v, %v\n", dst, src)
	gl.VertexAttrib4fv(uint32(dst.Value), &src[0])
	if err := int(GetError()); err != NO_ERROR {
		fmt.Printf("ERROR: %v\n", err)
	}
}

// VertexAttribPointer uses a bound buffer to define vertex attribute data.
//
// Direct use of VertexAttribPointer to load data into OpenGL is not
// supported via the Go bindings. Instead, use BindBuffer with an
// ARRAY_BUFFER and then fill it using BufferData.
//
// The size argument specifies the number of components per attribute,
// between 1-4. The stride argument specifies the byte offset between
// consecutive vertex attributes.
//
// http://www.khronos.org/opengles/sdk/docs/man3/html/glVertexAttribPointer.xhtml
func VertexAttribPointer(dst Attrib, size int, ty Enum, normalized bool, stride, offset int) {
	fmt.Printf("VertexAttribPointer : %v, %v, %v, %v, %v, %v\n", dst, size, ty, normalized, stride, offset)
	gl.VertexAttribPointer(uint32(dst.Value), int32(size), uint32(ty), normalized, int32(stride), gl.PtrOffset(offset))
	if err := int(GetError()); err != NO_ERROR {
		fmt.Printf("ERROR: %v\n", err)
	}
}

// Viewport sets the viewport, an affine transformation that
// normalizes device coordinates to window coordinates.
//
// http://www.khronos.org/opengles/sdk/docs/man3/html/glViewport.xhtml
func Viewport(x, y, width, height int) {
	fmt.Printf("Viewport : %v, %v, %v, %v\n", x, y, width, height)
	gl.Viewport(int32(x), int32(y), int32(width), int32(height))
	if err := int(GetError()); err != NO_ERROR {
		fmt.Printf("ERROR: %v\n", err)
	}
}
