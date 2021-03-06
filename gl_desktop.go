// Copyright (c) 2019 Thomas MILLET. All rights reserved.

// +build darwin freebsd linux windows
// +build !android
// +build !ios
// +build !js

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

func (p *plugin) Init(runtime tge.Runtime) error {
	if !p.isInit {
		p.isInit = true
		return gl.Init()
	}
	return fmt.Errorf("Already initialized")
}

func (p *plugin) Dispose() {
	p.isInit = false
	FlushCache()
}

// GetGLSLVersion gives the glsl version ti put in #version ${VERSION}
func GetGLSLVersion() string {
	return "330 core"
}

// FlushCache free memory cache, should be called between scenes
func FlushCache() {
	byteArrayBuffer = make([]byte, 0)
	byteArrayBufferExtendFactor = 1
}

// ActiveTexture sets the active texture unit.
//
// http://www.khronos.org/opengles/sdk/docs/man3/html/glActiveTexture.xhtml
func ActiveTexture(texture Enum) {
	gl.ActiveTexture(uint32(texture))
}

// AttachShader attaches a shader to a program.
//
// http://www.khronos.org/opengles/sdk/docs/man3/html/glAttachShader.xhtml
func AttachShader(p Program, s Shader) {
	gl.AttachShader(uint32(p), uint32(s))
}

// BindAttribLocation binds a vertex attribute index with a named
// variable.
//
// http://www.khronos.org/opengles/sdk/docs/man3/html/glBindAttribLocation.xhtml
func BindAttribLocation(p Program, a Attrib, name string) {
	gl.BindAttribLocation(uint32(p), uint32(a), gl.Str(name+"\x00"))
}

// BindBuffer binds a buffer.
//
// http://www.khronos.org/opengles/sdk/docs/man3/html/glBindBuffer.xhtml
func BindBuffer(target Enum, b Buffer) {
	gl.BindBuffer(uint32(target), uint32(b))
}

// BindFramebuffer binds a framebuffer.
//
// http://www.khronos.org/opengles/sdk/docs/man3/html/glBindFramebuffer.xhtml
func BindFramebuffer(target Enum, fb Framebuffer) {
	gl.BindFramebuffer(uint32(target), uint32(fb))
}

// BindRenderbuffer binds a render buffer.
//
// http://www.khronos.org/opengles/sdk/docs/man3/html/glBindRenderbuffer.xhtml
func BindRenderbuffer(target Enum, rb Renderbuffer) {
	gl.BindRenderbuffer(uint32(target), uint32(rb))
}

// BindTexture binds a texture.
//
// http://www.khronos.org/opengles/sdk/docs/man3/html/glBindTexture.xhtml
func BindTexture(target Enum, t Texture) {
	gl.BindTexture(uint32(target), uint32(t))
}

// BindVertexArray binds a VAO.
//
// http://www.khronos.org/opengles/sdk/docs/man3/html/glBindVertexArray.xhtml
func BindVertexArray(vao VertexArray) {
	gl.BindVertexArray(uint32(vao))
}

// BlendColor sets the blend color.
//
// http://www.khronos.org/opengles/sdk/docs/man3/html/glBlendColor.xhtml
func BlendColor(red, green, blue, alpha float32) {
	gl.BlendColor(red, green, blue, alpha)
}

// BlendEquation sets both RGB and alpha blend equations.
//
// http://www.khronos.org/opengles/sdk/docs/man3/html/glBlendEquation.xhtml
func BlendEquation(mode Enum) {
	gl.BlendEquation(uint32(mode))
}

// BlendEquationSeparate sets RGB and alpha blend equations separately.
//
// http://www.khronos.org/opengles/sdk/docs/man3/html/glBlendEquationSeparate.xhtml
func BlendEquationSeparate(modeRGB, modeAlpha Enum) {
	gl.BlendEquationSeparate(uint32(modeRGB), uint32(modeAlpha))
}

// BlendFunc sets the pixel blending factors.
//
// http://www.khronos.org/opengles/sdk/docs/man3/html/glBlendFunc.xhtml
func BlendFunc(sfactor, dfactor Enum) {
	gl.BlendFunc(uint32(sfactor), uint32(dfactor))
}

// BlendFuncSeparate sets the pixel RGB and alpha blending factors separately.
//
// http://www.khronos.org/opengles/sdk/docs/man3/html/glBlendFuncSeparate.xhtml
func BlendFuncSeparate(sfactorRGB, dfactorRGB, sfactorAlpha, dfactorAlpha Enum) {
	gl.BlendFuncSeparate(uint32(sfactorRGB), uint32(dfactorRGB), uint32(sfactorAlpha), uint32(dfactorAlpha))
}

// BufferData creates a new data store for the bound buffer object.
//
// http://www.khronos.org/opengles/sdk/docs/man3/html/glBufferData.xhtml
func BufferData(target Enum, src []byte, usage Enum) {
	gl.BufferData(uint32(target), int(len(src)), gl.Ptr(&src[0]), uint32(usage))
}

// BufferInit creates a new unitialized data store for the bound buffer object.
//
// http://www.khronos.org/opengles/sdk/docs/man3/html/glBufferData.xhtml
func BufferInit(target Enum, size int, usage Enum) {
	gl.BufferData(uint32(target), size, nil, uint32(usage))
}

// BufferSubData sets some of data in the bound buffer object.
//
// http://www.khronos.org/opengles/sdk/docs/man3/html/glBufferSubData.xhtml
func BufferSubData(target Enum, offset int, data []byte) {
	gl.BufferSubData(uint32(target), offset, int(len(data)), gl.Ptr(&data[0]))
}

// CheckFramebufferStatus reports the completeness status of the
// active framebuffer.
//
// http://www.khronos.org/opengles/sdk/docs/man3/html/glCheckFramebufferStatus.xhtml
func CheckFramebufferStatus(target Enum) Enum {
	return Enum(gl.CheckFramebufferStatus(uint32(target)))
}

// Clear clears the window.
//
// The behavior of Clear is influenced by the pixel ownership test,
// the scissor test, dithering, and the buffer writemasks.
//
// http://www.khronos.org/opengles/sdk/docs/man3/html/glClear.xhtml
func Clear(mask Enum) {
	gl.Clear(uint32(mask))
}

// ClearColor specifies the RGBA values used to clear color buffers.
//
// http://www.khronos.org/opengles/sdk/docs/man3/html/glClearColor.xhtml
func ClearColor(red, green, blue, alpha float32) {
	gl.ClearColor(red, green, blue, alpha)
}

// ClearDepthf sets the depth value used to clear the depth buffer.
//
// http://www.khronos.org/opengles/sdk/docs/man3/html/glClearDepthf.xhtml
func ClearDepthf(d float32) {
	gl.ClearDepthf(d)
}

// ClearStencil sets the index used to clear the stencil buffer.
//
// http://www.khronos.org/opengles/sdk/docs/man3/html/glClearStencil.xhtml
func ClearStencil(s int) {
	gl.ClearStencil(int32(s))
}

// ColorMask specifies whether color components in the framebuffer
// can be written.
//
// http://www.khronos.org/opengles/sdk/docs/man3/html/glColorMask.xhtml
func ColorMask(red, green, blue, alpha bool) {
	gl.ColorMask(red, green, blue, alpha)
}

// CompileShader compiles the source code of s.
//
// http://www.khronos.org/opengles/sdk/docs/man3/html/glCompileShader.xhtml
func CompileShader(s Shader) {
	gl.CompileShader(uint32(s))
}

// CompressedTexImage2D writes a compressed 2D texture.
//
// http://www.khronos.org/opengles/sdk/docs/man3/html/glCompressedTexImage2D.xhtml
func CompressedTexImage2D(target Enum, level int, internalformat Enum, width, height, border int, data []byte) {
	gl.CompressedTexImage2D(uint32(target), int32(level), uint32(internalformat), int32(width), int32(height), int32(border), int32(len(data)), gl.Ptr(data))
}

// CompressedTexSubImage2D writes a subregion of a compressed 2D texture.
//
// http://www.khronos.org/opengles/sdk/docs/man3/html/glCompressedTexSubImage2D.xhtml
func CompressedTexSubImage2D(target Enum, level, xoffset, yoffset, width, height int, format Enum, data []byte) {
	gl.CompressedTexSubImage2D(uint32(target), int32(level), int32(xoffset), int32(yoffset), int32(width), int32(height), uint32(format), int32(len(data)), gl.Ptr(data))
}

// CopyTexImage2D writes a 2D texture from the current framebuffer.
//
// http://www.khronos.org/opengles/sdk/docs/man3/html/glCopyTexImage2D.xhtml
func CopyTexImage2D(target Enum, level int, internalformat Enum, x, y, width, height, border int) {
	gl.CopyTexImage2D(uint32(target), int32(level), uint32(internalformat), int32(x), int32(y), int32(width), int32(height), int32(border))
}

// CopyTexSubImage2D writes a 2D texture subregion from the
// current framebuffer.
//
// http://www.khronos.org/opengles/sdk/docs/man3/html/glCopyTexSubImage2D.xhtml
func CopyTexSubImage2D(target Enum, level, xoffset, yoffset, x, y, width, height int) {
	gl.CopyTexSubImage2D(uint32(target), int32(level), int32(xoffset), int32(yoffset), int32(x), int32(y), int32(width), int32(height))
}

// CreateBuffer creates a buffer object.
//
// http://www.khronos.org/opengles/sdk/docs/man3/html/glGenBuffers.xhtml
func CreateBuffer() Buffer {
	var b uint32
	gl.GenBuffers(1, &b)
	return Buffer(b)
}

// CreateFramebuffer creates a framebuffer object.
//
// http://www.khronos.org/opengles/sdk/docs/man3/html/glGenFramebuffers.xhtml
func CreateFramebuffer() Framebuffer {
	var b uint32
	gl.GenFramebuffers(1, &b)
	return Framebuffer(b)
}

// CreateProgram creates a new empty program object.
//
// http://www.khronos.org/opengles/sdk/docs/man3/html/glCreateProgram.xhtml
func CreateProgram() Program {
	return Program(uint32(gl.CreateProgram()))
}

// CreateRenderbuffer create a renderbuffer object.
//
// http://www.khronos.org/opengles/sdk/docs/man3/html/glGenRenderbuffers.xhtml
func CreateRenderbuffer() Renderbuffer {
	var b uint32
	gl.GenRenderbuffers(1, &b)
	return Renderbuffer(b)
}

// CreateShader creates a new empty shader object.
//
// http://www.khronos.org/opengles/sdk/docs/man3/html/glCreateShader.xhtml
func CreateShader(ty Enum) Shader {
	return Shader(uint32(gl.CreateShader(uint32(ty))))
}

// CreateTexture creates a texture object.
//
// http://www.khronos.org/opengles/sdk/docs/man3/html/glGenTextures.xhtml
func CreateTexture() Texture {
	var t uint32
	gl.GenTextures(1, &t)
	return Texture(t)
}

// CreateVertexArray creates a VAO.
//
// http://www.khronos.org/opengles/sdk/docs/man3/html/glGenVertexArrays.xhtml
func CreateVertexArray() VertexArray {
	var vao uint32
	gl.GenVertexArrays(1, &vao)
	return VertexArray(vao)
}

// CullFace specifies which polygons are candidates for culling.
//
// Valid modes: FRONT, BACK, FRONT_AND_BACK.
//
// http://www.khronos.org/opengles/sdk/docs/man3/html/glCullFace.xhtml
func CullFace(mode Enum) {
	gl.CullFace(uint32(mode))
}

// DeleteBuffer deletes the given buffer object.
//
// http://www.khronos.org/opengles/sdk/docs/man3/html/glDeleteBuffers.xhtml
func DeleteBuffer(v Buffer) {
	u := uint32(v)
	gl.DeleteBuffers(1, &u)
}

// DeleteFramebuffer deletes the given framebuffer object.
//
// http://www.khronos.org/opengles/sdk/docs/man3/html/glDeleteFramebuffers.xhtml
func DeleteFramebuffer(v Framebuffer) {
	u := uint32(v)
	gl.DeleteFramebuffers(1, &u)
}

// DeleteProgram deletes the given program object.
//
// http://www.khronos.org/opengles/sdk/docs/man3/html/glDeleteProgram.xhtml
func DeleteProgram(p Program) {
	gl.DeleteProgram(uint32(p))
}

// DeleteRenderbuffer deletes the given render buffer object.
//
// http://www.khronos.org/opengles/sdk/docs/man3/html/glDeleteRenderbuffers.xhtml
func DeleteRenderbuffer(v Renderbuffer) {
	u := uint32(v)
	gl.DeleteRenderbuffers(1, &u)
}

// DeleteShader deletes shader s.
//
// http://www.khronos.org/opengles/sdk/docs/man3/html/glDeleteShader.xhtml
func DeleteShader(s Shader) {
	gl.DeleteShader(uint32(s))
}

// DeleteTexture deletes the given texture object.
//
// http://www.khronos.org/opengles/sdk/docs/man3/html/glDeleteTextures.xhtml
func DeleteTexture(v Texture) {
	u := uint32(v)
	gl.DeleteTextures(1, &u)
}

// DeleteVertexArray deletes the given VAO.
//
// http://www.khronos.org/opengles/sdk/docs/man3/html/glDeleteVertexArrays.xhtml
func DeleteVertexArray(v VertexArray) {
	u := uint32(v)
	gl.DeleteVertexArrays(1, &u)
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
	gl.DepthFunc(uint32(fn))
}

// DepthMask sets the depth buffer enabled for writing.
//
// http://www.khronos.org/opengles/sdk/docs/man3/html/glDepthMask.xhtml
func DepthMask(flag bool) {
	gl.DepthMask(flag)
}

// DepthRangef sets the mapping from normalized device coordinates to
// window coordinates.
//
// http://www.khronos.org/opengles/sdk/docs/man3/html/glDepthRangef.xhtml
func DepthRangef(n, f float32) {
	gl.DepthRangef(n, f)
}

// DetachShader detaches the shader s from the program p.
//
// http://www.khronos.org/opengles/sdk/docs/man3/html/glDetachShader.xhtml
func DetachShader(p Program, s Shader) {
	gl.DetachShader(uint32(p), uint32(s))
}

// Disable disables various GL capabilities.
//
// http://www.khronos.org/opengles/sdk/docs/man3/html/glDisable.xhtml
func Disable(cap Enum) {
	gl.Disable(uint32(cap))
}

// DisableVertexAttribArray disables a vertex attribute array.
//
// http://www.khronos.org/opengles/sdk/docs/man3/html/glDisableVertexAttribArray.xhtml
func DisableVertexAttribArray(a Attrib) {
	gl.DisableVertexAttribArray(uint32(a))
}

// DrawArrays renders geometric primitives from the bound data.
//
// http://www.khronos.org/opengles/sdk/docs/man3/html/glDrawArrays.xhtml
func DrawArrays(mode Enum, first, count int) {
	gl.DrawArrays(uint32(mode), int32(first), int32(count))
}

// DrawElements renders primitives from a bound buffer.
//
// http://www.khronos.org/opengles/sdk/docs/man3/html/glDrawElements.xhtml
func DrawElements(mode Enum, count int, ty Enum, offset int) {
	gl.DrawElements(uint32(mode), int32(count), uint32(ty), gl.PtrOffset(offset))
}

// Enable enables various GL capabilities.
//
// http://www.khronos.org/opengles/sdk/docs/man3/html/glEnable.xhtml
func Enable(cap Enum) {
	gl.Enable(uint32(cap))
}

// EnableVertexAttribArray enables a vertex attribute array.
//
// http://www.khronos.org/opengles/sdk/docs/man3/html/glEnableVertexAttribArray.xhtml
func EnableVertexAttribArray(a Attrib) {
	gl.EnableVertexAttribArray(uint32(a))
}

// Finish blocks until the effects of all previously called GL
// commands are complete.
//
// http://www.khronos.org/opengles/sdk/docs/man3/html/glFinish.xhtml
func Finish() {
	gl.Finish()
}

// Flush empties all buffers. It does not block.
//
// An OpenGL implementation may buffer network communication,
// the command stream, or data inside the graphics accelerator.
//
// http://www.khronos.org/opengles/sdk/docs/man3/html/glFlush.xhtml
func Flush() {
	gl.Flush()
}

// FramebufferRenderbuffer attaches rb to the current frame buffer.
//
// http://www.khronos.org/opengles/sdk/docs/man3/html/glFramebufferRenderbuffer.xhtml
func FramebufferRenderbuffer(target, attachment, rbTarget Enum, rb Renderbuffer) {
	gl.FramebufferRenderbuffer(uint32(target), uint32(attachment), uint32(rbTarget), uint32(rb))
}

// FramebufferTexture2D attaches the t to the current frame buffer.
//
// http://www.khronos.org/opengles/sdk/docs/man3/html/glFramebufferTexture2D.xhtml
func FramebufferTexture2D(target, attachment, texTarget Enum, t Texture, level int) {
	gl.FramebufferTexture2D(uint32(target), uint32(attachment), uint32(texTarget), uint32(t), int32(level))
}

// FrontFace defines which polygons are front-facing.
//
// Valid modes: CW, CCW.
//
// http://www.khronos.org/opengles/sdk/docs/man3/html/glFrontFace.xhtml
func FrontFace(mode Enum) {
	gl.FrontFace(uint32(mode))
}

// GenerateMipmap generates mipmaps for the current texture.
//
// http://www.khronos.org/opengles/sdk/docs/man3/html/glGenerateMipmap.xhtml
func GenerateMipmap(target Enum) {
	gl.GenerateMipmap(uint32(target))
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
	gl.GetActiveAttrib(uint32(p), uint32(index), int32(len(name)-1), &length, &si, &typ, cname)
	name = name[:strings.IndexRune(name, 0)]
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
	gl.GetActiveUniform(uint32(p), uint32(index), int32(len(name)-1), &length, &si, &typ, cname)
	name = name[:strings.IndexRune(name, 0)]
	return name, int(si), Enum(typ)
}

// GetAttachedShaders returns the shader objects attached to program p.
//
// http://www.khronos.org/opengles/sdk/docs/man3/html/glGetAttachedShaders.xhtml
func GetAttachedShaders(p Program) []Shader {
	shadersLen := GetProgrami(p, ATTACHED_SHADERS)
	var n int32
	buf := make([]uint32, shadersLen)
	gl.GetAttachedShaders(uint32(p), int32(shadersLen), &n, &buf[0])
	buf = buf[:int(n)]
	shaders := make([]Shader, int(n))
	for i, s := range buf {
		shaders[i] = Shader(uint32(s))
	}
	return shaders
}

// GetAttribLocation returns the location of an attribute variable.
//
// http://www.khronos.org/opengles/sdk/docs/man3/html/glGetAttribLocation.xhtml
func GetAttribLocation(p Program, name string) Attrib {
	return Attrib(gl.GetAttribLocation(uint32(p), gl.Str(name+"\x00")))
}

// GetBooleanv returns the boolean values of parameter pname.
//
// Many boolean parameters can be queried more easily using IsEnabled.
//
// http://www.khronos.org/opengles/sdk/docs/man3/html/glGet.xhtml
func GetBooleanv(dst []bool, pname Enum) {
	gl.GetBooleanv(uint32(pname), &dst[0])
}

// GetFloatv returns the float values of parameter pname.
//
// http://www.khronos.org/opengles/sdk/docs/man3/html/glGet.xhtml
func GetFloatv(dst []float32, pname Enum) {
	gl.GetFloatv(uint32(pname), &dst[0])
}

// GetIntegerv returns the int values of parameter pname.
//
// Single values may be queried more easily using GetInteger.
//
// http://www.khronos.org/opengles/sdk/docs/man3/html/glGet.xhtml
func GetIntegerv(pname Enum, data []int32) {
	gl.GetIntegerv(uint32(pname), &data[0])
}

// GetInteger returns the int value of parameter pname.
//
// http://www.khronos.org/opengles/sdk/docs/man3/html/glGet.xhtml
func GetInteger(pname Enum) int {
	var data int32
	gl.GetIntegerv(uint32(pname), &data)
	return int(data)
}

// GetBufferParameteri returns a parameter for the active buffer.
//
// http://www.khronos.org/opengles/sdk/docs/man3/html/glGetBufferParameteriv.xhtml
func GetBufferParameteri(target, pname Enum) int {
	var params int32
	gl.GetBufferParameteriv(uint32(target), uint32(pname), &params)
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
	return Framebuffer(uint32(b))
}

// GetFramebufferAttachmentParameteri returns attachment parameters
// for the active framebuffer object.
//
// http://www.khronos.org/opengles/sdk/docs/man3/html/glGetFramebufferAttachmentParameteriv.xhtml
func GetFramebufferAttachmentParameteri(target, attachment, pname Enum) int {
	var param int32
	gl.GetFramebufferAttachmentParameteriv(uint32(target), uint32(attachment), uint32(pname), &param)
	return int(param)
}

// GetProgrami returns a parameter value for a program.
//
// http://www.khronos.org/opengles/sdk/docs/man3/html/glGetProgramiv.xhtml
func GetProgrami(p Program, pname Enum) int {
	var result int32
	gl.GetProgramiv(uint32(p), uint32(pname), &result)
	return int(result)
}

// GetProgramInfoLog returns the information log for a program.
//
// http://www.khronos.org/opengles/sdk/docs/man3/html/glGetProgramInfoLog.xhtml
func GetProgramInfoLog(p Program) string {
	var logLength int32
	gl.GetProgramiv(uint32(p), gl.INFO_LOG_LENGTH, &logLength)
	if logLength == 0 {
		return ""
	}

	logBuffer := make([]uint8, logLength)
	gl.GetProgramInfoLog(uint32(p), logLength, nil, &logBuffer[0])
	return gl.GoStr(&logBuffer[0])
}

// GetRenderbufferParameteri returns a parameter value for a render buffer.
//
// http://www.khronos.org/opengles/sdk/docs/man3/html/glGetRenderbufferParameteriv.xhtml
func GetRenderbufferParameteri(target, pname Enum) int {
	var result int32
	gl.GetRenderbufferParameteriv(uint32(target), uint32(pname), &result)
	return int(result)
}

// GetShaderi returns a parameter value for a shader.
//
// http://www.khronos.org/opengles/sdk/docs/man3/html/glGetShaderiv.xhtml
func GetShaderi(s Shader, pname Enum) int {
	var result int32
	gl.GetShaderiv(uint32(s), uint32(pname), &result)
	return int(result)
}

// GetShaderInfoLog returns the information log for a shader.
//
// http://www.khronos.org/opengles/sdk/docs/man3/html/glGetShaderInfoLog.xhtml
func GetShaderInfoLog(s Shader) string {
	var logLength int32
	gl.GetShaderiv(uint32(s), gl.INFO_LOG_LENGTH, &logLength)
	if logLength == 0 {
		return ""
	}

	logBuffer := make([]uint8, logLength)
	gl.GetShaderInfoLog(uint32(s), logLength, nil, &logBuffer[0])
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
	gl.GetShaderSource(uint32(s), int32(sourceLen), nil, &buf[0])
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
	return gl.GoStr(gl.GetString(uint32(pname)))
}

// GetTexParameterfv returns the float values of a texture parameter.
//
// http://www.khronos.org/opengles/sdk/docs/man3/html/glGetTexParameter.xhtml
func GetTexParameterfv(dst []float32, target, pname Enum) {
	gl.GetTexParameterfv(uint32(target), uint32(pname), &dst[0])
}

// GetTexParameteriv returns the int values of a texture parameter.
//
// http://www.khronos.org/opengles/sdk/docs/man3/html/glGetTexParameter.xhtml
func GetTexParameteriv(dst []int32, target, pname Enum) {
	gl.GetTexParameteriv(uint32(target), uint32(pname), &dst[0])
}

// GetUniformfv returns the float values of a uniform variable.
//
// http://www.khronos.org/opengles/sdk/docs/man3/html/glGetUniform.xhtml
func GetUniformfv(dst []float32, src Uniform, p Program) {
	gl.GetUniformfv(uint32(p), int32(src), &dst[0])
}

// GetUniformiv returns the float values of a uniform variable.
//
// http://www.khronos.org/opengles/sdk/docs/man3/html/glGetUniform.xhtml
func GetUniformiv(dst []int32, src Uniform, p Program) {
	gl.GetUniformiv(uint32(p), int32(src), &dst[0])
}

// GetUniformLocation returns the location of a uniform variable.
//
// http://www.khronos.org/opengles/sdk/docs/man3/html/glGetUniformLocation.xhtml
func GetUniformLocation(p Program, name string) Uniform {
	return Uniform(gl.GetUniformLocation(uint32(p), gl.Str(name+"\x00")))
}

// GetVertexAttribf reads the float value of a vertex attribute.
//
// http://www.khronos.org/opengles/sdk/docs/man3/html/glGetVertexAttrib.xhtml
func GetVertexAttribf(src Attrib, pname Enum) float32 {
	var result float32
	gl.GetVertexAttribfv(uint32(src), uint32(pname), &result)
	return result
}

// GetVertexAttribfv reads float values of a vertex attribute.
//
// http://www.khronos.org/opengles/sdk/docs/man3/html/glGetVertexAttrib.xhtml
func GetVertexAttribfv(dst []float32, src Attrib, pname Enum) {
	gl.GetVertexAttribfv(uint32(src), uint32(pname), &dst[0])
}

// GetVertexAttribi reads the int value of a vertex attribute.
//
// http://www.khronos.org/opengles/sdk/docs/man3/html/glGetVertexAttrib.xhtml
func GetVertexAttribi(src Attrib, pname Enum) int32 {
	var result int32
	gl.GetVertexAttribiv(uint32(src), uint32(pname), &result)
	return result
}

// GetVertexAttribiv reads int values of a vertex attribute.
//
// http://www.khronos.org/opengles/sdk/docs/man3/html/glGetVertexAttrib.xhtml
func GetVertexAttribiv(dst []int32, src Attrib, pname Enum) {
	gl.GetVertexAttribiv(uint32(src), uint32(pname), &dst[0])
}

// Hint sets implementation-specific modes.
//
// http://www.khronos.org/opengles/sdk/docs/man3/html/glHint.xhtml
func Hint(target, mode Enum) {
	gl.Hint(uint32(target), uint32(mode))
}

// IsBuffer reports if b is a valid buffer.
//
// http://www.khronos.org/opengles/sdk/docs/man3/html/glIsBuffer.xhtml
func IsBuffer(b Buffer) bool {
	return gl.IsBuffer(uint32(b))
}

// IsEnabled reports if cap is an enabled capability.
//
// http://www.khronos.org/opengles/sdk/docs/man3/html/glIsEnabled.xhtml
func IsEnabled(cap Enum) bool {
	return gl.IsEnabled(uint32(cap))
}

// IsFramebuffer reports if fb is a valid frame buffer.
//
// http://www.khronos.org/opengles/sdk/docs/man3/html/glIsFramebuffer.xhtml
func IsFramebuffer(fb Framebuffer) bool {
	return gl.IsFramebuffer(uint32(fb))
}

// IsProgram reports if p is a valid program object.
//
// http://www.khronos.org/opengles/sdk/docs/man3/html/glIsProgram.xhtml
func IsProgram(p Program) bool {
	return gl.IsProgram(uint32(p))
}

// IsRenderbuffer reports if rb is a valid render buffer.
//
// http://www.khronos.org/opengles/sdk/docs/man3/html/glIsRenderbuffer.xhtml
func IsRenderbuffer(rb Renderbuffer) bool {
	return gl.IsRenderbuffer(uint32(rb))
}

// IsShader reports if s is valid shader.
//
// http://www.khronos.org/opengles/sdk/docs/man3/html/glIsShader.xhtml
func IsShader(s Shader) bool {
	return gl.IsShader(uint32(s))
}

// IsTexture reports if t is a valid texture.
//
// http://www.khronos.org/opengles/sdk/docs/man3/html/glIsTexture.xhtml
func IsTexture(t Texture) bool {
	return gl.IsTexture(uint32(t))
}

// LineWidth specifies the width of lines.
//
// http://www.khronos.org/opengles/sdk/docs/man3/html/glLineWidth.xhtml
func LineWidth(width float32) {
	gl.LineWidth(width)
}

// LinkProgram links the specified program.
//
// http://www.khronos.org/opengles/sdk/docs/man3/html/glLinkProgram.xhtml
func LinkProgram(p Program) {
	gl.LinkProgram(uint32(p))
}

// PixelStorei sets pixel storage parameters.
//
// http://www.khronos.org/opengles/sdk/docs/man3/html/glPixelStorei.xhtml
func PixelStorei(pname Enum, param int32) {
	gl.PixelStorei(uint32(pname), param)
}

// PolygonMode sets Polygon Mode.
//
// http://www.khronos.org/opengles/sdk/docs/man3/html/glPolygonMode.xhtml
func PolygonMode(face, mode Enum) {
	gl.PolygonMode(uint32(face), uint32(mode))
}

// PolygonOffset sets the scaling factors for depth offsets.
//
// http://www.khronos.org/opengles/sdk/docs/man3/html/glPolygonOffset.xhtml
func PolygonOffset(factor, units float32) {
	gl.PolygonOffset(factor, units)
}

// ReadPixels returns pixel data from a buffer.
//
// http://www.khronos.org/opengles/sdk/docs/man3/html/glReadPixels.xhtml
func ReadPixels(dst []byte, x, y, width, height int, format, ty Enum) {
	gl.ReadPixels(int32(x), int32(y), int32(width), int32(height), uint32(format), uint32(ty), gl.Ptr(&dst[0]))
}

// ReleaseShaderCompiler frees resources allocated by the shader compiler.
//
// http://www.khronos.org/opengles/sdk/docs/man3/html/glReleaseShaderCompiler.xhtml
func ReleaseShaderCompiler() {
	gl.ReleaseShaderCompiler()
}

// RenderbufferStorage establishes the data storage, format, and
// dimensions of a renderbuffer object's image.
//
// http://www.khronos.org/opengles/sdk/docs/man3/html/glRenderbufferStorage.xhtml
func RenderbufferStorage(target, internalFormat Enum, width, height int) {
	gl.RenderbufferStorage(uint32(target), uint32(internalFormat), int32(width), int32(height))
}

// SampleCoverage sets multisample coverage parameters.
//
// http://www.khronos.org/opengles/sdk/docs/man3/html/glSampleCoverage.xhtml
func SampleCoverage(value float32, invert bool) {
	gl.SampleCoverage(value, invert)
}

// Scissor defines the scissor box rectangle, in window coordinates.
//
// http://www.khronos.org/opengles/sdk/docs/man3/html/glScissor.xhtml
func Scissor(x, y, width, height int32) {
	gl.Scissor(x, y, width, height)
}

// ShaderSource sets the source code of s to the given source code.
//
// http://www.khronos.org/opengles/sdk/docs/man3/html/glShaderSource.xhtml
func ShaderSource(s Shader, src string) {
	glsource, free := gl.Strs(src + "\x00")
	gl.ShaderSource(uint32(s), 1, glsource, nil)
	free()
}

// StencilFunc sets the front and back stencil test reference value.
//
// http://www.khronos.org/opengles/sdk/docs/man3/html/glStencilFunc.xhtml
func StencilFunc(fn Enum, ref int, mask uint32) {
	gl.StencilFunc(uint32(fn), int32(ref), mask)
}

// StencilFuncSeparate sets the front or back stencil test reference value.
//
// http://www.khronos.org/opengles/sdk/docs/man3/html/glStencilFuncSeparate.xhtml
func StencilFuncSeparate(face, fn Enum, ref int, mask uint32) {
	gl.StencilFuncSeparate(uint32(face), uint32(fn), int32(ref), mask)
}

// StencilMask controls the writing of bits in the stencil planes.
//
// http://www.khronos.org/opengles/sdk/docs/man3/html/glStencilMask.xhtml
func StencilMask(mask uint32) {
	gl.StencilMask(mask)
}

// StencilMaskSeparate controls the writing of bits in the stencil planes.
//
// http://www.khronos.org/opengles/sdk/docs/man3/html/glStencilMaskSeparate.xhtml
func StencilMaskSeparate(face Enum, mask uint32) {
	gl.StencilMaskSeparate(uint32(face), mask)
}

// StencilOp sets front and back stencil test actions.
//
// http://www.khronos.org/opengles/sdk/docs/man3/html/glStencilOp.xhtml
func StencilOp(fail, zfail, zpass Enum) {
	gl.StencilOp(uint32(fail), uint32(zfail), uint32(zpass))
}

// StencilOpSeparate sets front or back stencil tests.
//
// http://www.khronos.org/opengles/sdk/docs/man3/html/glStencilOpSeparate.xhtml
func StencilOpSeparate(face, sfail, dpfail, dppass Enum) {
	gl.StencilOpSeparate(uint32(face), uint32(sfail), uint32(dpfail), uint32(dppass))
}

// TexImage2D writes a 2D texture image.
//
// http://www.khronos.org/opengles/sdk/docs/man3/html/glTexImage2D.xhtml
func TexImage2D(target Enum, level int, width, height int, format Enum, ty Enum, data []byte) {
	p := unsafe.Pointer(nil)
	if len(data) > 0 {
		p = gl.Ptr(&data[0])
	}
	gl.TexImage2D(uint32(target), int32(level), int32(format), int32(width), int32(height), 0, uint32(format), uint32(ty), p)
}

// TexSubImage2D writes a subregion of a 2D texture image.
//
// http://www.khronos.org/opengles/sdk/docs/man3/html/glTexSubImage2D.xhtml
func TexSubImage2D(target Enum, level int, x, y, width, height int, format, ty Enum, data []byte) {
	gl.TexSubImage2D(uint32(target), int32(level), int32(x), int32(y), int32(width), int32(height), uint32(format), uint32(ty), gl.Ptr(&data[0]))
}

// TexParameterf sets a float texture parameter.
//
// http://www.khronos.org/opengles/sdk/docs/man3/html/glTexParameter.xhtml
func TexParameterf(target, pname Enum, param float32) {
	gl.TexParameterf(uint32(target), uint32(pname), param)
}

// TexParameterfv sets a float texture parameter array.
//
// http://www.khronos.org/opengles/sdk/docs/man3/html/glTexParameter.xhtml
func TexParameterfv(target, pname Enum, params []float32) {
	gl.TexParameterfv(uint32(target), uint32(pname), &params[0])
}

// TexParameteri sets an integer texture parameter.
//
// http://www.khronos.org/opengles/sdk/docs/man3/html/glTexParameter.xhtml
func TexParameteri(target, pname Enum, param int) {
	gl.TexParameteri(uint32(target), uint32(pname), int32(param))
}

// TexParameteriv sets an integer texture parameter array.
//
// http://www.khronos.org/opengles/sdk/docs/man3/html/glTexParameter.xhtml
func TexParameteriv(target, pname Enum, params []int32) {
	gl.TexParameteriv(uint32(target), uint32(pname), &params[0])
}

// Uniform1f writes a float uniform variable.
//
// http://www.khronos.org/opengles/sdk/docs/man3/html/glUniform.xhtml
func Uniform1f(dst Uniform, v float32) {
	if dst.Valid() {
		gl.Uniform1f(int32(dst), v)
	}
}

// Uniform1fv writes a [len(src)]float uniform array.
//
// http://www.khronos.org/opengles/sdk/docs/man3/html/glUniform.xhtml
func Uniform1fv(dst Uniform, src []float32) {
	if dst.Valid() {
		gl.Uniform1fv(int32(dst), int32(len(src)), &src[0])
	}
}

// Uniform1fvP Pointer version of Uniform1fv (faster)
func Uniform1fvP(dst Uniform, count int32, value *float32) {
	if dst.Valid() {
		gl.Uniform1fv(int32(dst), count, value)
	}
}

// Uniform1fvUP Unsafe Pointer version of Uniform1fv (faster)
func Uniform1fvUP(dst Uniform, count int32, value unsafe.Pointer) {
	if dst.Valid() {
		gl.Uniform1fv(int32(dst), count, (*float32)(value))
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
		gl.Uniform1i(int32(dst), int32(v))
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
		gl.Uniform1iv(int32(dst), int32(len(src)), &src[0])
	}
}

// Uniform1ivP Pointer version of Uniform1iv (faster)
func Uniform1ivP(dst Uniform, count int32, value *int32) {
	if dst.Valid() {
		gl.Uniform1iv(int32(dst), count, value)
	}
}

// Uniform1ivUP Unsafe Pointer version of Uniform1iv (faster)
func Uniform1ivUP(dst Uniform, count int32, value unsafe.Pointer) {
	if dst.Valid() {
		gl.Uniform1iv(int32(dst), count, (*int32)(value))
	}
}

// Uniform2f writes a vec2 uniform variable.
//
// http://www.khronos.org/opengles/sdk/docs/man3/html/glUniform.xhtml
func Uniform2f(dst Uniform, v0, v1 float32) {
	if dst.Valid() {
		gl.Uniform2f(int32(dst), v0, v1)
	}
}

// Uniform2fv writes a vec2 uniform array of len(src)/2 elements.
//
// http://www.khronos.org/opengles/sdk/docs/man3/html/glUniform.xhtml
func Uniform2fv(dst Uniform, src []float32) {
	if dst.Valid() {
		gl.Uniform2fv(int32(dst), int32(len(src)/2), &src[0])
	}
}

// Uniform2fvP Pointer version of Uniform2fv (faster)
func Uniform2fvP(dst Uniform, count int32, value *float32) {
	if dst.Valid() {
		gl.Uniform2fv(int32(dst), count, value)
	}
}

// Uniform2fvUP Unsafe Pointer version of Uniform2fv (faster)
func Uniform2fvUP(dst Uniform, count int32, value unsafe.Pointer) {
	if dst.Valid() {
		gl.Uniform2fv(int32(dst), count, (*float32)(value))
	}
}

// Uniform2i writes an ivec2 uniform variable.
//
// http://www.khronos.org/opengles/sdk/docs/man3/html/glUniform.xhtml
func Uniform2i(dst Uniform, v0, v1 int) {
	if dst.Valid() {
		gl.Uniform2i(int32(dst), int32(v0), int32(v1))
	}
}

// Uniform2iv writes an ivec2 uniform array of len(src)/2 elements.
//
// http://www.khronos.org/opengles/sdk/docs/man3/html/glUniform.xhtml
func Uniform2iv(dst Uniform, src []int32) {
	if dst.Valid() {
		gl.Uniform2iv(int32(dst), int32(len(src)/2), &src[0])
	}
}

// Uniform2ivP Pointer version of Uniform2iv (faster)
func Uniform2ivP(dst Uniform, count int32, value *int32) {
	if dst.Valid() {
		gl.Uniform2iv(int32(dst), count, value)
	}
}

// Uniform2ivUP Unsafe Pointer version of Uniform2iv (faster)
func Uniform2ivUP(dst Uniform, count int32, value unsafe.Pointer) {
	if dst.Valid() {
		gl.Uniform2iv(int32(dst), count, (*int32)(value))
	}
}

// Uniform3f writes a vec3 uniform variable.
//
// http://www.khronos.org/opengles/sdk/docs/man3/html/glUniform.xhtml
func Uniform3f(dst Uniform, v0, v1, v2 float32) {
	if dst.Valid() {
		gl.Uniform3f(int32(dst), v0, v1, v2)
	}
}

// Uniform3fv writes a vec3 uniform array of len(src)/3 elements.
//
// http://www.khronos.org/opengles/sdk/docs/man3/html/glUniform.xhtml
func Uniform3fv(dst Uniform, src []float32) {
	if dst.Valid() {
		gl.Uniform3fv(int32(dst), int32(len(src)/3), &src[0])
	}
}

// Uniform3fvP Pointer version of Uniform3fv (faster)
func Uniform3fvP(dst Uniform, count int32, value *float32) {
	if dst.Valid() {
		gl.Uniform3fv(int32(dst), count, value)
	}
}

// Uniform3fvUP Unsafe Pointer version of Uniform3fv (faster)
func Uniform3fvUP(dst Uniform, count int32, value unsafe.Pointer) {
	if dst.Valid() {
		gl.Uniform3fv(int32(dst), count, (*float32)(value))
	}
}

// Uniform3i writes an ivec3 uniform variable.
//
// http://www.khronos.org/opengles/sdk/docs/man3/html/glUniform.xhtml
func Uniform3i(dst Uniform, v0, v1, v2 int32) {
	if dst.Valid() {
		gl.Uniform3i(int32(dst), v0, v1, v2)
	}
}

// Uniform3iv writes an ivec3 uniform array of len(src)/3 elements.
//
// http://www.khronos.org/opengles/sdk/docs/man3/html/glUniform.xhtml
func Uniform3iv(dst Uniform, src []int32) {
	if dst.Valid() {
		gl.Uniform3iv(int32(dst), int32(len(src)/3), &src[0])
	}
}

// Uniform3ivP Pointer version of Uniform3iv (faster)
func Uniform3ivP(dst Uniform, count int32, value *int32) {
	if dst.Valid() {
		gl.Uniform3iv(int32(dst), count, value)
	}
}

// Uniform3ivUP Unsafe Pointer version of Uniform3iv (faster)
func Uniform3ivUP(dst Uniform, count int32, value unsafe.Pointer) {
	if dst.Valid() {
		gl.Uniform3iv(int32(dst), count, (*int32)(value))
	}
}

// Uniform4f writes a vec4 uniform variable.
//
// http://www.khronos.org/opengles/sdk/docs/man3/html/glUniform.xhtml
func Uniform4f(dst Uniform, v0, v1, v2, v3 float32) {
	if dst.Valid() {
		gl.Uniform4f(int32(dst), v0, v1, v2, v3)
	}
}

// Uniform4fv writes a vec4 uniform array of len(src)/4 elements.
//
// http://www.khronos.org/opengles/sdk/docs/man3/html/glUniform.xhtml
func Uniform4fv(dst Uniform, src []float32) {
	if dst.Valid() {
		gl.Uniform4fv(int32(dst), int32(len(src)/4), &src[0])
	}
}

// Uniform4fvP Pointer version of Uniform4fv (faster)
func Uniform4fvP(dst Uniform, count int32, value *float32) {
	if dst.Valid() {
		gl.Uniform4fv(int32(dst), count, value)
	}
}

// Uniform4fvUP Unsafe Pointer version of Uniform4fv (faster)
func Uniform4fvUP(dst Uniform, count int32, value unsafe.Pointer) {
	if dst.Valid() {
		gl.Uniform4fv(int32(dst), count, (*float32)(value))
	}
}

// Uniform4i writes an ivec4 uniform variable.
//
// http://www.khronos.org/opengles/sdk/docs/man3/html/glUniform.xhtml
func Uniform4i(dst Uniform, v0, v1, v2, v3 int32) {
	if dst.Valid() {
		gl.Uniform4i(int32(dst), v0, v1, v2, v3)
	}
}

// Uniform4iv writes an ivec4 uniform array of len(src)/4 elements.
//
// http://www.khronos.org/opengles/sdk/docs/man3/html/glUniform.xhtml
func Uniform4iv(dst Uniform, src []int32) {
	if dst.Valid() {
		gl.Uniform4iv(int32(dst), int32(len(src)/4), &src[0])
	}
}

// Uniform4ivP Pointer version of Uniform4iv (faster)
func Uniform4ivP(dst Uniform, count int32, value *int32) {
	if dst.Valid() {
		gl.Uniform4iv(int32(dst), count, value)
	}
}

// Uniform4ivUP Unsafe Pointer version of Uniform4iv (faster)
func Uniform4ivUP(dst Uniform, count int32, value unsafe.Pointer) {
	if dst.Valid() {
		gl.Uniform4iv(int32(dst), count, (*int32)(value))
	}
}

// UniformMatrix2fv writes 2x2 matrices. Each matrix uses four
// float32 values, so the number of matrices written is len(src)/4.
//
// Each matrix must be supplied in column major order.
//
// http://www.khronos.org/opengles/sdk/docs/man3/html/glUniform.xhtml
func UniformMatrix2fv(dst Uniform, transpose bool, src []float32) {
	if dst.Valid() {
		gl.UniformMatrix2fv(int32(dst), int32(len(src)/(2*2)), transpose, &src[0])
	}
}

// UniformMatrix2fvP Pointer version of UniformMatrix2fv (faster)
func UniformMatrix2fvP(dst Uniform, count int32, transpose bool, value *float32) {
	if dst.Valid() {
		gl.UniformMatrix2fv(int32(dst), count, transpose, value)
	}
}

// UniformMatrix2fvUP Unsafe Pointer version of UniformMatrix2fv (faster)
func UniformMatrix2fvUP(dst Uniform, count int32, transpose bool, value unsafe.Pointer) {
	if dst.Valid() {
		gl.UniformMatrix2fv(int32(dst), count, transpose, (*float32)(value))
	}
}

// UniformMatrix3fv writes 3x3 matrices. Each matrix uses nine
// float32 values, so the number of matrices written is len(src)/9.
//
// Each matrix must be supplied in column major order.
//
// http://www.khronos.org/opengles/sdk/docs/man3/html/glUniform.xhtml
func UniformMatrix3fv(dst Uniform, transpose bool, src []float32) {
	if dst.Valid() {
		gl.UniformMatrix3fv(int32(dst), int32(len(src)/(3*3)), transpose, &src[0])
	}
}

// UniformMatrix3fvP Pointer version of UniformMatrix3fv (faster)
func UniformMatrix3fvP(dst Uniform, count int32, transpose bool, value *float32) {
	if dst.Valid() {
		gl.UniformMatrix3fv(int32(dst), count, transpose, value)
	}
}

// UniformMatrix3fvUP Unsafe Pointer version of UniformMatrix3fv (faster)
func UniformMatrix3fvUP(dst Uniform, count int32, transpose bool, value unsafe.Pointer) {
	if dst.Valid() {
		gl.UniformMatrix3fv(int32(dst), count, transpose, (*float32)(value))
	}
}

// UniformMatrix4fv writes 4x4 matrices. Each matrix uses 16
// float32 values, so the number of matrices written is len(src)/16.
//
// Each matrix must be supplied in column major order.
//
// http://www.khronos.org/opengles/sdk/docs/man3/html/glUniform.xhtml
func UniformMatrix4fv(dst Uniform, transpose bool, src []float32) {
	if dst.Valid() {
		gl.UniformMatrix4fv(int32(dst), int32(len(src)/(4*4)), transpose, &src[0])
	}
}

// UniformMatrix4fvP Pointer version of UniformMatrix4fv (faster)
func UniformMatrix4fvP(dst Uniform, count int32, transpose bool, value *float32) {
	if dst.Valid() {
		gl.UniformMatrix4fv(int32(dst), count, transpose, value)
	}
}

// UniformMatrix4fvUP Unsafe Pointer version of UniformMatrix4fv (faster)
func UniformMatrix4fvUP(dst Uniform, count int32, transpose bool, value unsafe.Pointer) {
	if dst.Valid() {
		gl.UniformMatrix4fv(int32(dst), count, transpose, (*float32)(value))
	}
}

// UseProgram sets the active program.
//
// http://www.khronos.org/opengles/sdk/docs/man3/html/glUseProgram.xhtml
func UseProgram(p Program) {
	gl.UseProgram(uint32(p))
}

// ValidateProgram checks to see whether the executables contained in
// program can execute given the current OpenGL state.
//
// Typically only used for debugging.
//
// http://www.khronos.org/opengles/sdk/docs/man3/html/glValidateProgram.xhtml
func ValidateProgram(p Program) {
	gl.ValidateProgram(uint32(p))
}

// VertexAttrib1f writes a float vertex attribute.
//
// http://www.khronos.org/opengles/sdk/docs/man3/html/glVertexAttrib.xhtml
func VertexAttrib1f(dst Attrib, x float32) {
	gl.VertexAttrib1f(uint32(dst), x)
}

// VertexAttrib1fv writes a float vertex attribute.
//
// http://www.khronos.org/opengles/sdk/docs/man3/html/glVertexAttrib.xhtml
func VertexAttrib1fv(dst Attrib, src []float32) {
	gl.VertexAttrib1fv(uint32(dst), &src[0])
}

// VertexAttrib2f writes a vec2 vertex attribute.
//
// http://www.khronos.org/opengles/sdk/docs/man3/html/glVertexAttrib.xhtml
func VertexAttrib2f(dst Attrib, x, y float32) {
	gl.VertexAttrib2f(uint32(dst), x, y)
}

// VertexAttrib2fv writes a vec2 vertex attribute.
//
// http://www.khronos.org/opengles/sdk/docs/man3/html/glVertexAttrib.xhtml
func VertexAttrib2fv(dst Attrib, src []float32) {
	gl.VertexAttrib2fv(uint32(dst), &src[0])
}

// VertexAttrib3f writes a vec3 vertex attribute.
//
// http://www.khronos.org/opengles/sdk/docs/man3/html/glVertexAttrib.xhtml
func VertexAttrib3f(dst Attrib, x, y, z float32) {
	gl.VertexAttrib3f(uint32(dst), x, y, z)
}

// VertexAttrib3fv writes a vec3 vertex attribute.
//
// http://www.khronos.org/opengles/sdk/docs/man3/html/glVertexAttrib.xhtml
func VertexAttrib3fv(dst Attrib, src []float32) {
	gl.VertexAttrib3fv(uint32(dst), &src[0])
}

// VertexAttrib4f writes a vec4 vertex attribute.
//
// http://www.khronos.org/opengles/sdk/docs/man3/html/glVertexAttrib.xhtml
func VertexAttrib4f(dst Attrib, x, y, z, w float32) {
	gl.VertexAttrib4f(uint32(dst), x, y, z, w)
}

// VertexAttrib4fv writes a vec4 vertex attribute.
//
// http://www.khronos.org/opengles/sdk/docs/man3/html/glVertexAttrib.xhtml
func VertexAttrib4fv(dst Attrib, src []float32) {
	gl.VertexAttrib4fv(uint32(dst), &src[0])
}

// VertexAttribPointer uses a bound buffer to define vertex attribute data.
//
// http://www.khronos.org/opengles/sdk/docs/man3/html/glVertexAttribPointer.xhtml
func VertexAttribPointer(dst Attrib, size int, ty Enum, normalized bool, stride, offset int) {
	gl.VertexAttribPointer(uint32(dst), int32(size), uint32(ty), normalized, int32(stride), gl.PtrOffset(offset))
}

// Viewport sets the viewport, an affine transformation that
// normalizes device coordinates to window coordinates.
//
// http://www.khronos.org/opengles/sdk/docs/man3/html/glViewport.xhtml
func Viewport(x, y, width, height int) {
	gl.Viewport(int32(x), int32(y), int32(width), int32(height))
}
