// Copyright (c) 2019 Thomas MILLET. All rights reserved.

// +build android ios

package gl

import (
	fmt "fmt"
	unsafe "unsafe"

	tge "github.com/thommil/tge"
	gl "github.com/thommil/tge-mobile/gl"
)

type plugin struct {
	glContext gl.Context
}

var _pluginInstance = &plugin{}

// GetInstance returns plugin handler
func GetInstance() tge.Plugin {
	return _pluginInstance
}

func (p *plugin) Init(runtime tge.Runtime) error {
	renderer := runtime.GetRenderer()
	switch renderer.(type) {
	case gl.Context:
		p.glContext = renderer.(gl.Context)
	default:
		return fmt.Errorf("Runtime renderer must be a golang.org/x/mobile/gl.Context")
	}
	return nil
}

func (p *plugin) GetName() string {
	return Name
}

func (p *plugin) Dispose() {
	p.glContext = nil
	FlushCache()
}

// GetGLSLVersion gives the glsl version ti put in #version ${VERSION}
func GetGLSLVersion() string {
	return "300 es"
}

// FlushCache free memory cache, should be called between scenes
func FlushCache() {

}

// ActiveTexture sets the active texture unit.
//
// http://www.khronos.org/opengles/sdk/docs/man3/html/glActiveTexture.xhtml
func ActiveTexture(texture Enum) {
	_pluginInstance.glContext.ActiveTexture(gl.Enum(texture))

}

// AttachShader attaches a shader to a program.
//
// http://www.khronos.org/opengles/sdk/docs/man3/html/glAttachShader.xhtml
func AttachShader(p Program, s Shader) {
	_pluginInstance.glContext.AttachShader(gl.Program{Init: true, Value: p.Value}, gl.Shader{s.Value})
}

// BindAttribLocation binds a vertex attribute index with a named
// variable.
//
// http://www.khronos.org/opengles/sdk/docs/man3/html/glBindAttribLocation.xhtml
func BindAttribLocation(p Program, a Attrib, name string) {
	_pluginInstance.glContext.BindAttribLocation(gl.Program{Init: true, Value: p.Value}, gl.Attrib{a.Value}, name+"\x00")
}

// BindBuffer binds a buffer.
//
// http://www.khronos.org/opengles/sdk/docs/man3/html/glBindBuffer.xhtml
func BindBuffer(target Enum, b Buffer) {
	_pluginInstance.glContext.BindBuffer(gl.Enum(target), gl.Buffer{b.Value})
}

// BindFramebuffer binds a framebuffer.
//
// http://www.khronos.org/opengles/sdk/docs/man3/html/glBindFramebuffer.xhtml
func BindFramebuffer(target Enum, fb Framebuffer) {
	_pluginInstance.glContext.BindFramebuffer(gl.Enum(target), gl.Framebuffer{fb.Value})
}

// BindRenderbuffer binds a render buffer.
//
// http://www.khronos.org/opengles/sdk/docs/man3/html/glBindRenderbuffer.xhtml
func BindRenderbuffer(target Enum, rb Renderbuffer) {
	_pluginInstance.glContext.BindRenderbuffer(gl.Enum(target), gl.Renderbuffer{rb.Value})
}

// BindTexture binds a texture.
//
// http://www.khronos.org/opengles/sdk/docs/man3/html/glBindTexture.xhtml
func BindTexture(target Enum, t Texture) {
	_pluginInstance.glContext.BindTexture(gl.Enum(target), gl.Texture{t.Value})
}

// BindVertexArray binds a vertex array.
//
// http://www.khronos.org/opengles/sdk/docs/man3/html/glBindVertexArray.xhtml
func BindVertexArray(rb VertexArray) {
	_pluginInstance.glContext.BindVertexArray(gl.VertexArray{rb.Value})
}

// BlendColor sets the blend color.
//
// http://www.khronos.org/opengles/sdk/docs/man3/html/glBlendColor.xhtml
func BlendColor(red, green, blue, alpha float32) {
	_pluginInstance.glContext.BlendColor(red, green, blue, alpha)
}

// BlendEquation sets both RGB and alpha blend equations.
//
// http://www.khronos.org/opengles/sdk/docs/man3/html/glBlendEquation.xhtml
func BlendEquation(mode Enum) {
	_pluginInstance.glContext.BlendEquation(gl.Enum(mode))
}

// BlendEquationSeparate sets RGB and alpha blend equations separately.
//
// http://www.khronos.org/opengles/sdk/docs/man3/html/glBlendEquationSeparate.xhtml
func BlendEquationSeparate(modeRGB, modeAlpha Enum) {
	_pluginInstance.glContext.BlendEquationSeparate(gl.Enum(modeRGB), gl.Enum(modeAlpha))
}

// BlendFunc sets the pixel blending factors.
//
// http://www.khronos.org/opengles/sdk/docs/man3/html/glBlendFunc.xhtml
func BlendFunc(sfactor, dfactor Enum) {
	_pluginInstance.glContext.BlendFunc(gl.Enum(sfactor), gl.Enum(dfactor))
}

// BlendFunc sets the pixel RGB and alpha blending factors separately.
//
// http://www.khronos.org/opengles/sdk/docs/man3/html/glBlendFuncSeparate.xhtml
func BlendFuncSeparate(sfactorRGB, dfactorRGB, sfactorAlpha, dfactorAlpha Enum) {
	_pluginInstance.glContext.BlendFuncSeparate(gl.Enum(sfactorRGB), gl.Enum(dfactorRGB), gl.Enum(sfactorAlpha), gl.Enum(dfactorAlpha))
}

// BufferData creates a new data store for the bound buffer object.
//
// http://www.khronos.org/opengles/sdk/docs/man3/html/glBufferData.xhtml
func BufferData(target Enum, src []byte, usage Enum) {
	_pluginInstance.glContext.BufferData(gl.Enum(target), src, gl.Enum(usage))
}

// BufferInit creates a new unitialized data store for the bound buffer object.
//
// http://www.khronos.org/opengles/sdk/docs/man3/html/glBufferData.xhtml
func BufferInit(target Enum, size int, usage Enum) {
	_pluginInstance.glContext.BufferInit(gl.Enum(target), size, gl.Enum(usage))
}

// BufferSubData sets some of data in the bound buffer object.
//
// http://www.khronos.org/opengles/sdk/docs/man3/html/glBufferSubData.xhtml
func BufferSubData(target Enum, offset int, data []byte) {
	_pluginInstance.glContext.BufferSubData(gl.Enum(target), offset, data)
}

// CheckFramebufferStatus reports the completeness status of the
// active framebuffer.
//
// http://www.khronos.org/opengles/sdk/docs/man3/html/glCheckFramebufferStatus.xhtml
func CheckFramebufferStatus(target Enum) Enum {
	return Enum(_pluginInstance.glContext.CheckFramebufferStatus(gl.Enum(target)))
}

// Clear clears the window.
//
// The behavior of Clear is influenced by the pixel ownership test,
// the scissor test, dithering, and the buffer writemasks.
//
// http://www.khronos.org/opengles/sdk/docs/man3/html/glClear.xhtml
func Clear(mask Enum) {
	_pluginInstance.glContext.Clear(gl.Enum(mask))
}

// ClearColor specifies the RGBA values used to clear color buffers.
//
// http://www.khronos.org/opengles/sdk/docs/man3/html/glClearColor.xhtml
func ClearColor(red, green, blue, alpha float32) {
	_pluginInstance.glContext.ClearColor(red, green, blue, alpha)
}

// ClearDepthf sets the depth value used to clear the depth buffer.
//
// http://www.khronos.org/opengles/sdk/docs/man3/html/glClearDepthf.xhtml
func ClearDepthf(d float32) {
	_pluginInstance.glContext.ClearDepthf(d)
}

// ClearStencil sets the index used to clear the stencil buffer.
//
// http://www.khronos.org/opengles/sdk/docs/man3/html/glClearStencil.xhtml
func ClearStencil(s int) {
	_pluginInstance.glContext.ClearStencil(s)
}

// ColorMask specifies whether color components in the framebuffer
// can be written.
//
// http://www.khronos.org/opengles/sdk/docs/man3/html/glColorMask.xhtml
func ColorMask(red, green, blue, alpha bool) {
	_pluginInstance.glContext.ColorMask(red, green, blue, alpha)
}

// CompileShader compiles the source code of s.
//
// http://www.khronos.org/opengles/sdk/docs/man3/html/glCompileShader.xhtml
func CompileShader(s Shader) {
	_pluginInstance.glContext.CompileShader(gl.Shader{s.Value})
}

// CompressedTexImage2D writes a compressed 2D texture.
//
// http://www.khronos.org/opengles/sdk/docs/man3/html/glCompressedTexImage2D.xhtml
func CompressedTexImage2D(target Enum, level int, internalformat Enum, width, height, border int, data []byte) {
	_pluginInstance.glContext.CompressedTexImage2D(gl.Enum(target), level, gl.Enum(internalformat), width, height, border, data)
}

// CompressedTexSubImage2D writes a subregion of a compressed 2D texture.
//
// http://www.khronos.org/opengles/sdk/docs/man3/html/glCompressedTexSubImage2D.xhtml
func CompressedTexSubImage2D(target Enum, level, xoffset, yoffset, width, height int, format Enum, data []byte) {
	_pluginInstance.glContext.CompressedTexSubImage2D(gl.Enum(target), level, xoffset, yoffset, width, height, gl.Enum(format), data)
}

// CopyTexImage2D writes a 2D texture from the current framebuffer.
//
// http://www.khronos.org/opengles/sdk/docs/man3/html/glCopyTexImage2D.xhtml
func CopyTexImage2D(target Enum, level int, internalformat Enum, x, y, width, height, border int) {
	_pluginInstance.glContext.CopyTexImage2D(gl.Enum(target), level, gl.Enum(internalformat), x, y, width, height, border)
}

// CopyTexSubImage2D writes a 2D texture subregion from the
// current framebuffer.
//
// http://www.khronos.org/opengles/sdk/docs/man3/html/glCopyTexSubImage2D.xhtml
func CopyTexSubImage2D(target Enum, level, xoffset, yoffset, x, y, width, height int) {
	_pluginInstance.glContext.CopyTexSubImage2D(gl.Enum(target), level, xoffset, yoffset, x, y, width, height)
}

// CreateBuffer creates a buffer object.
//
// http://www.khronos.org/opengles/sdk/docs/man3/html/glGenBuffers.xhtml
func CreateBuffer() Buffer {
	return Buffer{_pluginInstance.glContext.CreateBuffer().Value}
}

// CreateFramebuffer creates a framebuffer object.
//
// http://www.khronos.org/opengles/sdk/docs/man3/html/glGenFramebuffers.xhtml
func CreateFramebuffer() Framebuffer {
	return Framebuffer{_pluginInstance.glContext.CreateFramebuffer().Value}
}

// CreateProgram creates a new empty program object.
//
// http://www.khronos.org/opengles/sdk/docs/man3/html/glCreateProgram.xhtml
func CreateProgram() Program {
	return Program{Value: _pluginInstance.glContext.CreateProgram().Value}
}

// CreateRenderbuffer create a renderbuffer object.
//
// http://www.khronos.org/opengles/sdk/docs/man3/html/glGenRenderbuffers.xhtml
func CreateRenderbuffer() Renderbuffer {
	return Renderbuffer{_pluginInstance.glContext.CreateRenderbuffer().Value}
}

// CreateShader creates a new empty shader object.
//
// http://www.khronos.org/opengles/sdk/docs/man3/html/glCreateShader.xhtml
func CreateShader(ty Enum) Shader {
	return Shader{_pluginInstance.glContext.CreateShader(gl.Enum(ty)).Value}
}

// CreateTexture creates a texture object.
//
// http://www.khronos.org/opengles/sdk/docs/man3/html/glGenTextures.xhtml
func CreateTexture() Texture {
	return Texture{_pluginInstance.glContext.CreateTexture().Value}
}

// CreateTVertexArray creates a vertex array.
//
// http://www.khronos.org/opengles/sdk/docs/man3/html/glGenVertexArrays.xhtml
func CreateVertexArray() VertexArray {
	return VertexArray{_pluginInstance.glContext.CreateVertexArray().Value}
}

// CullFace specifies which polygons are candidates for culling.
//
// Valid modes: FRONT, BACK, FRONT_AND_BACK.
//
// http://www.khronos.org/opengles/sdk/docs/man3/html/glCullFace.xhtml
func CullFace(mode Enum) {
	_pluginInstance.glContext.CullFace(gl.Enum(mode))
}

// DeleteBuffer deletes the given buffer object.
//
// http://www.khronos.org/opengles/sdk/docs/man3/html/glDeleteBuffers.xhtml
func DeleteBuffer(v Buffer) {
	_pluginInstance.glContext.DeleteBuffer(gl.Buffer{v.Value})
}

// DeleteFramebuffer deletes the given framebuffer object.
//
// http://www.khronos.org/opengles/sdk/docs/man3/html/glDeleteFramebuffers.xhtml
func DeleteFramebuffer(v Framebuffer) {
	_pluginInstance.glContext.DeleteFramebuffer(gl.Framebuffer{v.Value})
}

// DeleteProgram deletes the given program object.
//
// http://www.khronos.org/opengles/sdk/docs/man3/html/glDeleteProgram.xhtml
func DeleteProgram(p Program) {
	_pluginInstance.glContext.DeleteProgram(gl.Program{Init: true, Value: p.Value})
}

// DeleteRenderbuffer deletes the given render buffer object.
//
// http://www.khronos.org/opengles/sdk/docs/man3/html/glDeleteRenderbuffers.xhtml
func DeleteRenderbuffer(v Renderbuffer) {
	_pluginInstance.glContext.DeleteRenderbuffer(gl.Renderbuffer{v.Value})
}

// DeleteShader deletes shader s.
//
// http://www.khronos.org/opengles/sdk/docs/man3/html/glDeleteShader.xhtml
func DeleteShader(s Shader) {
	_pluginInstance.glContext.DeleteShader(gl.Shader{s.Value})
}

// DeleteTexture deletes the given texture object.
//
// http://www.khronos.org/opengles/sdk/docs/man3/html/glDeleteTextures.xhtml
func DeleteTexture(v Texture) {
	_pluginInstance.glContext.DeleteTexture(gl.Texture{v.Value})
}

// DeleteVertexArray deletes the given render buffer object.
//
// http://www.khronos.org/opengles/sdk/docs/man3/html/glDeleteVertexArrays.xhtml
func DeleteVertexArray(v VertexArray) {
	_pluginInstance.glContext.DeleteVertexArray(gl.VertexArray{v.Value})
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
	_pluginInstance.glContext.DepthFunc(gl.Enum(fn))
}

// DepthMask sets the depth buffer enabled for writing.
//
// http://www.khronos.org/opengles/sdk/docs/man3/html/glDepthMask.xhtml
func DepthMask(flag bool) {
	_pluginInstance.glContext.DepthMask(flag)
}

// DepthRangef sets the mapping from normalized device coordinates to
// window coordinates.
//
// http://www.khronos.org/opengles/sdk/docs/man3/html/glDepthRangef.xhtml
func DepthRangef(n, f float32) {
	_pluginInstance.glContext.DepthRangef(n, f)
}

// DetachShader detaches the shader s from the program p.
//
// http://www.khronos.org/opengles/sdk/docs/man3/html/glDetachShader.xhtml
func DetachShader(p Program, s Shader) {
	_pluginInstance.glContext.DetachShader(gl.Program{Init: true, Value: p.Value}, gl.Shader{s.Value})
}

// Disable disables various GL capabilities.
//
// http://www.khronos.org/opengles/sdk/docs/man3/html/glDisable.xhtml
func Disable(cap Enum) {
	_pluginInstance.glContext.Disable(gl.Enum(cap))
}

// DisableVertexAttribArray disables a vertex attribute array.
//
// http://www.khronos.org/opengles/sdk/docs/man3/html/glDisableVertexAttribArray.xhtml
func DisableVertexAttribArray(a Attrib) {
	_pluginInstance.glContext.DisableVertexAttribArray(gl.Attrib{a.Value})
}

// DrawArrays renders geometric primitives from the bound data.
//
// http://www.khronos.org/opengles/sdk/docs/man3/html/glDrawArrays.xhtml
func DrawArrays(mode Enum, first, count int) {
	_pluginInstance.glContext.DrawArrays(gl.Enum(mode), first, count)
}

// DrawElements renders primitives from a bound buffer.
//
// http://www.khronos.org/opengles/sdk/docs/man3/html/glDrawElements.xhtml
func DrawElements(mode Enum, count int, ty Enum, offset int) {
	_pluginInstance.glContext.DrawElements(gl.Enum(mode), count, gl.Enum(ty), offset)
}

// Enable enables various GL capabilities.
//
// http://www.khronos.org/opengles/sdk/docs/man3/html/glEnable.xhtml
func Enable(cap Enum) {
	_pluginInstance.glContext.Enable(gl.Enum(cap))
}

// EnableVertexAttribArray enables a vertex attribute array.
//
// http://www.khronos.org/opengles/sdk/docs/man3/html/glEnableVertexAttribArray.xhtml
func EnableVertexAttribArray(a Attrib) {
	_pluginInstance.glContext.EnableVertexAttribArray(gl.Attrib{a.Value})
}

// Finish blocks until the effects of all previously called GL
// commands are complete.
//
// http://www.khronos.org/opengles/sdk/docs/man3/html/glFinish.xhtml
func Finish() {
	_pluginInstance.glContext.Finish()
}

// Flush empties all buffers. It does not block.
//
// An OpenGL implementation may buffer network communication,
// the command stream, or data inside the graphics accelerator.
//
// http://www.khronos.org/opengles/sdk/docs/man3/html/glFlush.xhtml
func Flush() {
	_pluginInstance.glContext.Flush()
}

// FramebufferRenderbuffer attaches rb to the current frame buffer.
//
// http://www.khronos.org/opengles/sdk/docs/man3/html/glFramebufferRenderbuffer.xhtml
func FramebufferRenderbuffer(target, attachment, rbTarget Enum, rb Renderbuffer) {
	_pluginInstance.glContext.FramebufferRenderbuffer(gl.Enum(target), gl.Enum(attachment), gl.Enum(rbTarget), gl.Renderbuffer{rb.Value})
}

// FramebufferTexture2D attaches the t to the current frame buffer.
//
// http://www.khronos.org/opengles/sdk/docs/man3/html/glFramebufferTexture2D.xhtml
func FramebufferTexture2D(target, attachment, texTarget Enum, t Texture, level int) {
	_pluginInstance.glContext.FramebufferTexture2D(gl.Enum(target), gl.Enum(attachment), gl.Enum(texTarget), gl.Texture{t.Value}, level)
}

// FrontFace defines which polygons are front-facing.
//
// Valid modes: CW, CCW.
//
// http://www.khronos.org/opengles/sdk/docs/man3/html/glFrontFace.xhtml
func FrontFace(mode Enum) {
	_pluginInstance.glContext.FrontFace(gl.Enum(mode))
}

// GenerateMipmap generates mipmaps for the current texture.
//
// http://www.khronos.org/opengles/sdk/docs/man3/html/glGenerateMipmap.xhtml
func GenerateMipmap(target Enum) {
	_pluginInstance.glContext.GenerateMipmap(gl.Enum(target))
}

// GetActiveAttrib returns details about an active attribute variable.
// A value of 0 for index selects the first active attribute variable.
// Permissible values for index range from 0 to the number of active
// attribute variables minus 1.
//
// http://www.khronos.org/opengles/sdk/docs/man3/html/glGetActiveAttrib.xhtml
func GetActiveAttrib(p Program, index uint32) (name string, size int, ty Enum) {
	n, s, t := _pluginInstance.glContext.GetActiveAttrib(gl.Program{Init: true, Value: p.Value}, index)
	return n, s, Enum(t)
}

// GetActiveUniform returns details about an active uniform variable.
// A value of 0 for index selects the first active uniform variable.
// Permissible values for index range from 0 to the number of active
// uniform variables minus 1.
//
// http://www.khronos.org/opengles/sdk/docs/man3/html/glGetActiveUniform.xhtml
func GetActiveUniform(p Program, index uint32) (name string, size int, ty Enum) {
	n, s, t := _pluginInstance.glContext.GetActiveUniform(gl.Program{Init: true, Value: p.Value}, index)
	return n, s, Enum(t)
}

// GetAttachedShaders returns the shader objects attached to program p.
//
// http://www.khronos.org/opengles/sdk/docs/man3/html/glGetAttachedShaders.xhtml
func GetAttachedShaders(p Program) []Shader {
	shaders := _pluginInstance.glContext.GetAttachedShaders(gl.Program{Init: true, Value: p.Value})
	s := make([]Shader, len(shaders))
	for i, el := range shaders {
		s[i] = Shader{el.Value}
	}
	return s
}

// GetAttribLocation returns the location of an attribute variable.
//
// http://www.khronos.org/opengles/sdk/docs/man3/html/glGetAttribLocation.xhtml
func GetAttribLocation(p Program, name string) Attrib {
	return Attrib{_pluginInstance.glContext.GetAttribLocation(gl.Program{Init: true, Value: p.Value}, name).Value}
}

// GetBooleanv returns the boolean values of parameter pname.
//
// Many boolean parameters can be queried more easily using IsEnabled.
//
// http://www.khronos.org/opengles/sdk/docs/man3/html/glGet.xhtml
func GetBooleanv(dst []bool, pname Enum) {
	_pluginInstance.glContext.GetBooleanv(dst, gl.Enum(pname))
}

// GetFloatv returns the float values of parameter pname.
//
// http://www.khronos.org/opengles/sdk/docs/man3/html/glGet.xhtml
func GetFloatv(dst []float32, pname Enum) {
	_pluginInstance.glContext.GetFloatv(dst, gl.Enum(pname))
}

// GetIntegerv returns the int values of parameter pname.
//
// Single values may be queried more easily using GetInteger.
//
// http://www.khronos.org/opengles/sdk/docs/man3/html/glGet.xhtml
func GetIntegerv(pname Enum, data []int32) {
	_pluginInstance.glContext.GetIntegerv(data, gl.Enum(pname))
}

// GetInteger returns the int value of parameter pname.
//
// http://www.khronos.org/opengles/sdk/docs/man3/html/glGet.xhtml
func GetInteger(pname Enum) int {
	return _pluginInstance.glContext.GetInteger(gl.Enum(pname))
}

// GetBufferParameteri returns a parameter for the active buffer.
//
// http://www.khronos.org/opengles/sdk/docs/man3/html/glGetBufferParameteriv.xhtml
func GetBufferParameteri(target, pname Enum) int {
	return _pluginInstance.glContext.GetBufferParameteri(gl.Enum(target), gl.Enum(pname))
}

// GetError returns the next error.
//
// http://www.khronos.org/opengles/sdk/docs/man3/html/glGetError.xhtml
func GetError() Enum {
	return Enum(_pluginInstance.glContext.GetError())
}

// GetBoundFramebuffer returns the currently bound framebuffer.
// Use this method instead of _pluginInstance.glContext.GetInteger(_pluginInstance.glContext.FRAMEBUFFER_BINDING) to
// enable support on all platforms
func GetBoundFramebuffer() Framebuffer {
	b := make([]int32, 1)
	_pluginInstance.glContext.GetIntegerv(b, gl.FRAMEBUFFER_BINDING)
	return Framebuffer{Value: uint32(b[0])}
}

// GetFramebufferAttachmentParameteri returns attachment parameters
// for the active framebuffer object.
//
// http://www.khronos.org/opengles/sdk/docs/man3/html/glGetFramebufferAttachmentParameteriv.xhtml
func GetFramebufferAttachmentParameteri(target, attachment, pname Enum) int {
	return _pluginInstance.glContext.GetFramebufferAttachmentParameteri(gl.Enum(target), gl.Enum(attachment), gl.Enum(pname))
}

// GetProgrami returns a parameter value for a program.
//
// http://www.khronos.org/opengles/sdk/docs/man3/html/glGetProgramiv.xhtml
func GetProgrami(p Program, pname Enum) int {
	return _pluginInstance.glContext.GetProgrami(gl.Program{Init: true, Value: p.Value}, gl.Enum(pname))
}

// GetProgramInfoLog returns the information log for a program.
//
// http://www.khronos.org/opengles/sdk/docs/man3/html/glGetProgramInfoLog.xhtml
func GetProgramInfoLog(p Program) string {
	return _pluginInstance.glContext.GetProgramInfoLog(gl.Program{Init: true, Value: p.Value})
}

// GetRenderbufferParameteri returns a parameter value for a render buffer.
//
// http://www.khronos.org/opengles/sdk/docs/man3/html/glGetRenderbufferParameteriv.xhtml
func GetRenderbufferParameteri(target, pname Enum) int {
	return _pluginInstance.glContext.GetRenderbufferParameteri(gl.Enum(target), gl.Enum(pname))
}

// GetRenderbufferParameteri returns a parameter value for a shader.
//
// http://www.khronos.org/opengles/sdk/docs/man3/html/glGetShaderiv.xhtml
func GetShaderi(s Shader, pname Enum) int {
	return _pluginInstance.glContext.GetShaderi(gl.Shader{s.Value}, gl.Enum(pname))
}

// GetShaderInfoLog returns the information log for a shader.
//
// http://www.khronos.org/opengles/sdk/docs/man3/html/glGetShaderInfoLog.xhtml
func GetShaderInfoLog(s Shader) string {
	return _pluginInstance.glContext.GetShaderInfoLog(gl.Shader{s.Value})
}

// GetShaderPrecisionFormat returns range and precision limits for
// shader types.
//
// http://www.khronos.org/opengles/sdk/docs/man3/html/glGetShaderPrecisionFormat.xhtml
func GetShaderPrecisionFormat(shadertype, precisiontype Enum) (rangeLow, rangeHigh, precision int) {
	return _pluginInstance.glContext.GetShaderPrecisionFormat(gl.Enum(shadertype), gl.Enum(precisiontype))
}

// GetShaderSource returns source code of shader s.
//
// http://www.khronos.org/opengles/sdk/docs/man3/html/glGetShaderSource.xhtml
func GetShaderSource(s Shader) string {
	return _pluginInstance.glContext.GetShaderSource(gl.Shader{s.Value})
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
	return _pluginInstance.glContext.GetString(gl.Enum(pname))
}

// GetTexParameterfv returns the float values of a texture parameter.
//
// http://www.khronos.org/opengles/sdk/docs/man3/html/glGetTexParameter.xhtml
func GetTexParameterfv(dst []float32, target, pname Enum) {
	_pluginInstance.glContext.GetTexParameterfv(dst, gl.Enum(target), gl.Enum(pname))
}

// GetTexParameteriv returns the int values of a texture parameter.
//
// http://www.khronos.org/opengles/sdk/docs/man3/html/glGetTexParameter.xhtml
func GetTexParameteriv(dst []int32, target, pname Enum) {
	_pluginInstance.glContext.GetTexParameteriv(dst, gl.Enum(target), gl.Enum(pname))
}

// GetUniformfv returns the float values of a uniform variable.
//
// http://www.khronos.org/opengles/sdk/docs/man3/html/glGetUniform.xhtml
func GetUniformfv(dst []float32, src Uniform, p Program) {
	_pluginInstance.glContext.GetUniformfv(dst, gl.Uniform{src.Value}, gl.Program{Init: true, Value: p.Value})
}

// GetUniformiv returns the float values of a uniform variable.
//
// http://www.khronos.org/opengles/sdk/docs/man3/html/glGetUniform.xhtml
func GetUniformiv(dst []int32, src Uniform, p Program) {
	_pluginInstance.glContext.GetUniformiv(dst, gl.Uniform{src.Value}, gl.Program{Init: true, Value: p.Value})
}

// GetUniformLocation returns the location of a uniform variable.
//
// http://www.khronos.org/opengles/sdk/docs/man3/html/glGetUniformLocation.xhtml
func GetUniformLocation(p Program, name string) Uniform {
	return Uniform{_pluginInstance.glContext.GetUniformLocation(gl.Program{Init: true, Value: p.Value}, name).Value}
}

// GetVertexAttribf reads the float value of a vertex attribute.
//
// http://www.khronos.org/opengles/sdk/docs/man3/html/glGetVertexAttrib.xhtml
func GetVertexAttribf(src Attrib, pname Enum) float32 {
	return _pluginInstance.glContext.GetVertexAttribf(gl.Attrib{src.Value}, gl.Enum(pname))
}

// GetVertexAttribfv reads float values of a vertex attribute.
//
// http://www.khronos.org/opengles/sdk/docs/man3/html/glGetVertexAttrib.xhtml
func GetVertexAttribfv(dst []float32, src Attrib, pname Enum) {
	_pluginInstance.glContext.GetVertexAttribfv(dst, gl.Attrib{src.Value}, gl.Enum(pname))
}

// GetVertexAttribi reads the int value of a vertex attribute.
//
// http://www.khronos.org/opengles/sdk/docs/man3/html/glGetVertexAttrib.xhtml
func GetVertexAttribi(src Attrib, pname Enum) int32 {
	return _pluginInstance.glContext.GetVertexAttribi(gl.Attrib{src.Value}, gl.Enum(pname))
}

// GetVertexAttribiv reads int values of a vertex attribute.
//
// http://www.khronos.org/opengles/sdk/docs/man3/html/glGetVertexAttrib.xhtml
func GetVertexAttribiv(dst []int32, src Attrib, pname Enum) {
	_pluginInstance.glContext.GetVertexAttribiv(dst, gl.Attrib{src.Value}, gl.Enum(pname))
}

// Hint sets implementation-specific modes.
//
// http://www.khronos.org/opengles/sdk/docs/man3/html/glHint.xhtml
func Hint(target, mode Enum) {
	_pluginInstance.glContext.Hint(gl.Enum(target), gl.Enum(mode))
}

// IsBuffer reports if b is a valid buffer.
//
// http://www.khronos.org/opengles/sdk/docs/man3/html/glIsBuffer.xhtml
func IsBuffer(b Buffer) bool {
	return _pluginInstance.glContext.IsBuffer(gl.Buffer{b.Value})
}

// IsEnabled reports if cap is an enabled capability.
//
// http://www.khronos.org/opengles/sdk/docs/man3/html/glIsEnabled.xhtml
func IsEnabled(cap Enum) bool {
	return _pluginInstance.glContext.IsEnabled(gl.Enum(cap))
}

// IsFramebuffer reports if fb is a valid frame buffer.
//
// http://www.khronos.org/opengles/sdk/docs/man3/html/glIsFramebuffer.xhtml
func IsFramebuffer(fb Framebuffer) bool {
	return _pluginInstance.glContext.IsFramebuffer(gl.Framebuffer{fb.Value})
}

// IsProgram reports if p is a valid program object.
//
// http://www.khronos.org/opengles/sdk/docs/man3/html/glIsProgram.xhtml
func IsProgram(p Program) bool {
	return _pluginInstance.glContext.IsProgram(gl.Program{Init: true, Value: p.Value})
}

// IsRenderbuffer reports if rb is a valid render buffer.
//
// http://www.khronos.org/opengles/sdk/docs/man3/html/glIsRenderbuffer.xhtml
func IsRenderbuffer(rb Renderbuffer) bool {
	return _pluginInstance.glContext.IsRenderbuffer(gl.Renderbuffer{rb.Value})
}

// IsShader reports if s is valid shader.
//
// http://www.khronos.org/opengles/sdk/docs/man3/html/glIsShader.xhtml
func IsShader(s Shader) bool {
	return _pluginInstance.glContext.IsShader(gl.Shader{s.Value})
}

// IsTexture reports if t is a valid texture.
//
// http://www.khronos.org/opengles/sdk/docs/man3/html/glIsTexture.xhtml
func IsTexture(t Texture) bool {
	return _pluginInstance.glContext.IsTexture(gl.Texture{t.Value})
}

// LineWidth specifies the width of lines.
//
// http://www.khronos.org/opengles/sdk/docs/man3/html/glLineWidth.xhtml
func LineWidth(width float32) {
	_pluginInstance.glContext.LineWidth(width)
}

// LinkProgram links the specified program.
//
// http://www.khronos.org/opengles/sdk/docs/man3/html/glLinkProgram.xhtml
func LinkProgram(p Program) {
	_pluginInstance.glContext.LinkProgram(gl.Program{Init: true, Value: p.Value})
}

// PixelStorei sets pixel storage parameters.
//
// http://www.khronos.org/opengles/sdk/docs/man3/html/glPixelStorei.xhtml
func PixelStorei(pname Enum, param int32) {
	_pluginInstance.glContext.PixelStorei(gl.Enum(pname), param)
}

// PolygonOffset sets the scaling factors for depth offsets.
//
// http://www.khronos.org/opengles/sdk/docs/man3/html/glPolygonOffset.xhtml
func PolygonOffset(factor, units float32) {
	_pluginInstance.glContext.PolygonOffset(factor, units)
}

// PolygonMode sets Polygon Mode.
//
// http://www.khronos.org/opengles/sdk/docs/man3/html/glPolygonMode.xhtml
func PolygonMode(face, mode Enum) {
	fmt.Printf("WARNING: PolygonMode not implemented\n")
}

// ReadPixels returns pixel data from a buffer.
//
// http://www.khronos.org/opengles/sdk/docs/man3/html/glReadPixels.xhtml
func ReadPixels(dst []byte, x, y, width, height int, format, ty Enum) {
	_pluginInstance.glContext.ReadPixels(dst, x, y, width, height, gl.Enum(format), gl.Enum(ty))
}

// ReleaseShaderCompiler frees resources allocated by the shader compiler.
//
// http://www.khronos.org/opengles/sdk/docs/man3/html/glReleaseShaderCompiler.xhtml
func ReleaseShaderCompiler() {
	_pluginInstance.glContext.ReleaseShaderCompiler()
}

// RenderbufferStorage establishes the data storage, format, and
// dimensions of a renderbuffer object's image.
//
// http://www.khronos.org/opengles/sdk/docs/man3/html/glRenderbufferStorage.xhtml
func RenderbufferStorage(target, internalFormat Enum, width, height int) {
	_pluginInstance.glContext.RenderbufferStorage(gl.Enum(target), gl.Enum(internalFormat), width, height)
}

// SampleCoverage sets multisample coverage parameters.
//
// http://www.khronos.org/opengles/sdk/docs/man3/html/glSampleCoverage.xhtml
func SampleCoverage(value float32, invert bool) {
	_pluginInstance.glContext.SampleCoverage(value, invert)
}

// Scissor defines the scissor box rectangle, in window coordinates.
//
// http://www.khronos.org/opengles/sdk/docs/man3/html/glScissor.xhtml
func Scissor(x, y, width, height int32) {
	_pluginInstance.glContext.Scissor(x, y, width, height)
}

// ShaderSource sets the source code of s to the given source code.
//
// http://www.khronos.org/opengles/sdk/docs/man3/html/glShaderSource.xhtml
func ShaderSource(s Shader, src string) {
	_pluginInstance.glContext.ShaderSource(gl.Shader{s.Value}, src)
}

// StencilFunc sets the front and back stencil test reference value.
//
// http://www.khronos.org/opengles/sdk/docs/man3/html/glStencilFunc.xhtml
func StencilFunc(fn Enum, ref int, mask uint32) {
	_pluginInstance.glContext.StencilFunc(gl.Enum(fn), ref, mask)
}

// StencilFunc sets the front or back stencil test reference value.
//
// http://www.khronos.org/opengles/sdk/docs/man3/html/glStencilFuncSeparate.xhtml
func StencilFuncSeparate(face, fn Enum, ref int, mask uint32) {
	_pluginInstance.glContext.StencilFuncSeparate(gl.Enum(face), gl.Enum(fn), ref, mask)
}

// StencilMask controls the writing of bits in the stencil planes.
//
// http://www.khronos.org/opengles/sdk/docs/man3/html/glStencilMask.xhtml
func StencilMask(mask uint32) {
	_pluginInstance.glContext.StencilMask(mask)
}

// StencilMaskSeparate controls the writing of bits in the stencil planes.
//
// http://www.khronos.org/opengles/sdk/docs/man3/html/glStencilMaskSeparate.xhtml
func StencilMaskSeparate(face Enum, mask uint32) {
	_pluginInstance.glContext.StencilMaskSeparate(gl.Enum(face), mask)
}

// StencilOp sets front and back stencil test actions.
//
// http://www.khronos.org/opengles/sdk/docs/man3/html/glStencilOp.xhtml
func StencilOp(fail, zfail, zpass Enum) {
	_pluginInstance.glContext.StencilOp(gl.Enum(fail), gl.Enum(zfail), gl.Enum(zpass))
}

// StencilOpSeparate sets front or back stencil tests.
//
// http://www.khronos.org/opengles/sdk/docs/man3/html/glStencilOpSeparate.xhtml
func StencilOpSeparate(face, sfail, dpfail, dppass Enum) {
	_pluginInstance.glContext.StencilOpSeparate(gl.Enum(face), gl.Enum(sfail), gl.Enum(dpfail), gl.Enum(dppass))
}

// TexImage2D writes a 2D texture image.
//
// http://www.khronos.org/opengles/sdk/docs/man3/html/glTexImage2D.xhtml
func TexImage2D(target Enum, level int, width, height int, format Enum, ty Enum, data []byte) {
	_pluginInstance.glContext.TexImage2D(gl.Enum(target), level, int(format), width, height, gl.Enum(format), gl.Enum(ty), data)
}

// TexSubImage2D writes a subregion of a 2D texture image.
//
// http://www.khronos.org/opengles/sdk/docs/man3/html/glTexSubImage2D.xhtml
func TexSubImage2D(target Enum, level int, x, y, width, height int, format, ty Enum, data []byte) {
	_pluginInstance.glContext.TexSubImage2D(gl.Enum(target), level, x, y, width, height, gl.Enum(format), gl.Enum(ty), data)
}

// TexParameterf sets a float texture parameter.
//
// http://www.khronos.org/opengles/sdk/docs/man3/html/glTexParameter.xhtml
func TexParameterf(target, pname Enum, param float32) {
	_pluginInstance.glContext.TexParameterf(gl.Enum(target), gl.Enum(pname), param)
}

// TexParameterfv sets a float texture parameter array.
//
// http://www.khronos.org/opengles/sdk/docs/man3/html/glTexParameter.xhtml
func TexParameterfv(target, pname Enum, params []float32) {
	_pluginInstance.glContext.TexParameterfv(gl.Enum(target), gl.Enum(pname), params)
}

// TexParameteri sets an integer texture parameter.
//
// http://www.khronos.org/opengles/sdk/docs/man3/html/glTexParameter.xhtml
func TexParameteri(target, pname Enum, param int) {
	_pluginInstance.glContext.TexParameteri(gl.Enum(target), gl.Enum(pname), param)
}

// TexParameteriv sets an integer texture parameter array.
//
// http://www.khronos.org/opengles/sdk/docs/man3/html/glTexParameter.xhtml
func TexParameteriv(target, pname Enum, params []int32) {
	_pluginInstance.glContext.TexParameteriv(gl.Enum(target), gl.Enum(pname), params)
}

// Uniform1f writes a float uniform variable.
//
// http://www.khronos.org/opengles/sdk/docs/man3/html/glUniform.xhtml
func Uniform1f(dst Uniform, v float32) {
	_pluginInstance.glContext.Uniform1f(gl.Uniform{dst.Value}, v)
}

// Uniform1fv writes a [len(src)]float uniform array.
//
// http://www.khronos.org/opengles/sdk/docs/man3/html/glUniform.xhtml
func Uniform1fv(dst Uniform, src []float32) {
	_pluginInstance.glContext.Uniform1fv(gl.Uniform{dst.Value}, src)
}

func Uniform1fvP(dst Uniform, count int32, value *float32) {
	fmt.Println("Uniform1fvP")
}

func Uniform1fvUP(dst Uniform, count int32, value unsafe.Pointer) {
	fmt.Println("Uniform1fvUP")
}

// Uniform1i writes an int uniform variable.
//
// Uniform1i and Uniform1iv are the only two functions that may be used
// to load uniform variables defined as sampler types. Loading samplers
// with any other function will result in a INVALID_OPERATION error.
//
// http://www.khronos.org/opengles/sdk/docs/man3/html/glUniform.xhtml
func Uniform1i(dst Uniform, v int) {
	_pluginInstance.glContext.Uniform1i(gl.Uniform{dst.Value}, v)
}

// Uniform1iv writes a int uniform array of len(src) elements.
//
// Uniform1i and Uniform1iv are the only two functions that may be used
// to load uniform variables defined as sampler types. Loading samplers
// with any other function will result in a INVALID_OPERATION error.
//
// http://www.khronos.org/opengles/sdk/docs/man3/html/glUniform.xhtml
func Uniform1iv(dst Uniform, src []int32) {
	_pluginInstance.glContext.Uniform1iv(gl.Uniform{dst.Value}, src)
}

func Uniform1ivP(dst Uniform, count int32, value *int32) {
	_pluginInstance.glContext.Uniform1ivP(gl.Uniform{dst.Value}, count, value)
}

func Uniform1ivUP(dst Uniform, count int32, value unsafe.Pointer) {
	_pluginInstance.glContext.Uniform1ivUP(gl.Uniform{dst.Value}, count, value)
}

// Uniform2f writes a vec2 uniform variable.
//
// http://www.khronos.org/opengles/sdk/docs/man3/html/glUniform.xhtml
func Uniform2f(dst Uniform, v0, v1 float32) {
	_pluginInstance.glContext.Uniform2f(gl.Uniform{dst.Value}, v0, v1)
}

// Uniform2fv writes a vec2 uniform array of len(src)/2 elements.
//
// http://www.khronos.org/opengles/sdk/docs/man3/html/glUniform.xhtml
func Uniform2fv(dst Uniform, src []float32) {
	_pluginInstance.glContext.Uniform2fv(gl.Uniform{dst.Value}, src)
}

func Uniform2fvP(dst Uniform, count int32, value *float32) {
	_pluginInstance.glContext.Uniform2fvP(gl.Uniform{dst.Value}, count, value)
}

func Uniform2fvUP(dst Uniform, count int32, value unsafe.Pointer) {
	_pluginInstance.glContext.Uniform2fvUP(gl.Uniform{dst.Value}, count, value)
}

// Uniform2i writes an ivec2 uniform variable.
//
// http://www.khronos.org/opengles/sdk/docs/man3/html/glUniform.xhtml
func Uniform2i(dst Uniform, v0, v1 int) {
	_pluginInstance.glContext.Uniform2i(gl.Uniform{dst.Value}, v0, v1)
}

// Uniform2iv writes an ivec2 uniform array of len(src)/2 elements.
//
// http://www.khronos.org/opengles/sdk/docs/man3/html/glUniform.xhtml
func Uniform2iv(dst Uniform, src []int32) {
	_pluginInstance.glContext.Uniform2iv(gl.Uniform{dst.Value}, src)
}

func Uniform2ivP(dst Uniform, count int32, value *int32) {
	_pluginInstance.glContext.Uniform2ivP(gl.Uniform{dst.Value}, count, value)
}

func Uniform2ivUP(dst Uniform, count int32, value unsafe.Pointer) {
	_pluginInstance.glContext.Uniform2ivUP(gl.Uniform{dst.Value}, count, value)
}

// Uniform3f writes a vec3 uniform variable.
//
// http://www.khronos.org/opengles/sdk/docs/man3/html/glUniform.xhtml
func Uniform3f(dst Uniform, v0, v1, v2 float32) {
	_pluginInstance.glContext.Uniform3f(gl.Uniform{dst.Value}, v0, v1, v2)
}

// Uniform3fv writes a vec3 uniform array of len(src)/3 elements.
//
// http://www.khronos.org/opengles/sdk/docs/man3/html/glUniform.xhtml
func Uniform3fv(dst Uniform, src []float32) {
	_pluginInstance.glContext.Uniform3fv(gl.Uniform{dst.Value}, src)
}

func Uniform3fvP(dst Uniform, count int32, value *float32) {
	_pluginInstance.glContext.Uniform3fvP(gl.Uniform{dst.Value}, count, value)
}

func Uniform3fvUP(dst Uniform, count int32, value unsafe.Pointer) {
	_pluginInstance.glContext.Uniform3fvUP(gl.Uniform{dst.Value}, count, value)
}

// Uniform3i writes an ivec3 uniform variable.
//
// http://www.khronos.org/opengles/sdk/docs/man3/html/glUniform.xhtml
func Uniform3i(dst Uniform, v0, v1, v2 int32) {
	_pluginInstance.glContext.Uniform3i(gl.Uniform{dst.Value}, v0, v1, v2)
}

// Uniform3iv writes an ivec3 uniform array of len(src)/3 elements.
//
// http://www.khronos.org/opengles/sdk/docs/man3/html/glUniform.xhtml
func Uniform3iv(dst Uniform, src []int32) {
	_pluginInstance.glContext.Uniform3iv(gl.Uniform{dst.Value}, src)
}

func Uniform3ivP(dst Uniform, count int32, value *int32) {
	_pluginInstance.glContext.Uniform3ivP(gl.Uniform{dst.Value}, count, value)
}

func Uniform3ivUP(dst Uniform, count int32, value unsafe.Pointer) {
	_pluginInstance.glContext.Uniform3ivUP(gl.Uniform{dst.Value}, count, value)
}

// Uniform4f writes a vec4 uniform variable.
//
// http://www.khronos.org/opengles/sdk/docs/man3/html/glUniform.xhtml
func Uniform4f(dst Uniform, v0, v1, v2, v3 float32) {
	_pluginInstance.glContext.Uniform4f(gl.Uniform{dst.Value}, v0, v1, v2, v3)
}

// Uniform4fv writes a vec4 uniform array of len(src)/4 elements.
//
// http://www.khronos.org/opengles/sdk/docs/man3/html/glUniform.xhtml
func Uniform4fv(dst Uniform, src []float32) {
	_pluginInstance.glContext.Uniform4fv(gl.Uniform{dst.Value}, src)
}

func Uniform4fvP(dst Uniform, count int32, value *float32) {
	_pluginInstance.glContext.Uniform4fvP(gl.Uniform{dst.Value}, count, value)
}

func Uniform4fvUP(dst Uniform, count int32, value unsafe.Pointer) {
	_pluginInstance.glContext.Uniform4fvUP(gl.Uniform{dst.Value}, count, value)
}

// Uniform4i writes an ivec4 uniform variable.
//
// http://www.khronos.org/opengles/sdk/docs/man3/html/glUniform.xhtml
func Uniform4i(dst Uniform, v0, v1, v2, v3 int32) {
	_pluginInstance.glContext.Uniform4i(gl.Uniform{dst.Value}, v0, v1, v2, v3)
}

// Uniform4i writes an ivec4 uniform array of len(src)/4 elements.
//
// http://www.khronos.org/opengles/sdk/docs/man3/html/glUniform.xhtml
func Uniform4iv(dst Uniform, src []int32) {
	_pluginInstance.glContext.Uniform4iv(gl.Uniform{dst.Value}, src)
}

func Uniform4ivP(dst Uniform, count int32, value *int32) {
	_pluginInstance.glContext.Uniform4ivP(gl.Uniform{dst.Value}, count, value)
}

func Uniform4ivUP(dst Uniform, count int32, value unsafe.Pointer) {
	_pluginInstance.glContext.Uniform4ivUP(gl.Uniform{dst.Value}, count, value)
}

// UniformMatrix2fv writes 2x2 matrices. Each matrix uses four
// float32 values, so the number of matrices written is len(src)/4.
//
// Each matrix must be supplied in column major order.
//
// http://www.khronos.org/opengles/sdk/docs/man3/html/glUniform.xhtml
func UniformMatrix2fv(dst Uniform, src []float32) {
	_pluginInstance.glContext.UniformMatrix2fv(gl.Uniform{dst.Value}, src)
}

func UniformMatrix2fvP(dst Uniform, count int32, transpose bool, value *float32) {
	_pluginInstance.glContext.UniformMatrix2fvP(gl.Uniform{dst.Value}, count, value)
}

func UniformMatrix2fvUP(dst Uniform, count int32, transpose bool, value unsafe.Pointer) {
	_pluginInstance.glContext.UniformMatrix2fvUP(gl.Uniform{dst.Value}, count, value)
}

// UniformMatrix3fv writes 3x3 matrices. Each matrix uses nine
// float32 values, so the number of matrices written is len(src)/9.
//
// Each matrix must be supplied in column major order.
//
// http://www.khronos.org/opengles/sdk/docs/man3/html/glUniform.xhtml
func UniformMatrix3fv(dst Uniform, src []float32) {
	_pluginInstance.glContext.UniformMatrix3fv(gl.Uniform{dst.Value}, src)
}

func UniformMatrix3fvP(dst Uniform, count int32, transpose bool, value *float32) {
	_pluginInstance.glContext.UniformMatrix3fvP(gl.Uniform{dst.Value}, count, value)
}

func UniformMatrix3fvUP(dst Uniform, count int32, transpose bool, value unsafe.Pointer) {
	_pluginInstance.glContext.UniformMatrix3fvUP(gl.Uniform{dst.Value}, count, value)
}

// UniformMatrix4fv writes 4x4 matrices. Each matrix uses 16
// float32 values, so the number of matrices written is len(src)/16.
//
// Each matrix must be supplied in column major order.
//
// http://www.khronos.org/opengles/sdk/docs/man3/html/glUniform.xhtml
func UniformMatrix4fv(dst Uniform, src []float32) {
	_pluginInstance.glContext.UniformMatrix4fv(gl.Uniform{dst.Value}, src)
}

func UniformMatrix4fvP(dst Uniform, count int32, transpose bool, value *float32) {
	_pluginInstance.glContext.UniformMatrix4fvP(gl.Uniform{dst.Value}, count, value)
}

func UniformMatrix4fvUP(dst Uniform, count int32, transpose bool, value unsafe.Pointer) {
	_pluginInstance.glContext.UniformMatrix4fvUP(gl.Uniform{dst.Value}, count, value)
}

// UseProgram sets the active program.
//
// http://www.khronos.org/opengles/sdk/docs/man3/html/glUseProgram.xhtml
func UseProgram(p Program) {
	_pluginInstance.glContext.UseProgram(gl.Program{Init: true, Value: p.Value})
}

// ValidateProgram checks to see whether the executables contained in
// program can execute given the current OpenGL state.
//
// Typically only used for debugging.
//
// http://www.khronos.org/opengles/sdk/docs/man3/html/glValidateProgram.xhtml
func ValidateProgram(p Program) {
	_pluginInstance.glContext.ValidateProgram(gl.Program{Init: true, Value: p.Value})
}

// VertexAttrib1f writes a float vertex attribute.
//
// http://www.khronos.org/opengles/sdk/docs/man3/html/glVertexAttrib.xhtml
func VertexAttrib1f(dst Attrib, x float32) {
	_pluginInstance.glContext.VertexAttrib1f(gl.Attrib{dst.Value}, x)
}

// VertexAttrib1fv writes a float vertex attribute.
//
// http://www.khronos.org/opengles/sdk/docs/man3/html/glVertexAttrib.xhtml
func VertexAttrib1fv(dst Attrib, src []float32) {
	_pluginInstance.glContext.VertexAttrib1fv(gl.Attrib{dst.Value}, src)
}

// VertexAttrib2f writes a vec2 vertex attribute.
//
// http://www.khronos.org/opengles/sdk/docs/man3/html/glVertexAttrib.xhtml
func VertexAttrib2f(dst Attrib, x, y float32) {
	_pluginInstance.glContext.VertexAttrib2f(gl.Attrib{dst.Value}, x, y)
}

// VertexAttrib2fv writes a vec2 vertex attribute.
//
// http://www.khronos.org/opengles/sdk/docs/man3/html/glVertexAttrib.xhtml
func VertexAttrib2fv(dst Attrib, src []float32) {
	_pluginInstance.glContext.VertexAttrib2fv(gl.Attrib{dst.Value}, src)
}

// VertexAttrib3f writes a vec3 vertex attribute.
//
// http://www.khronos.org/opengles/sdk/docs/man3/html/glVertexAttrib.xhtml
func VertexAttrib3f(dst Attrib, x, y, z float32) {
	_pluginInstance.glContext.VertexAttrib3f(gl.Attrib{dst.Value}, x, y, z)
}

// VertexAttrib3fv writes a vec3 vertex attribute.
//
// http://www.khronos.org/opengles/sdk/docs/man3/html/glVertexAttrib.xhtml
func VertexAttrib3fv(dst Attrib, src []float32) {
	_pluginInstance.glContext.VertexAttrib3fv(gl.Attrib{dst.Value}, src)
}

// VertexAttrib4f writes a vec4 vertex attribute.
//
// http://www.khronos.org/opengles/sdk/docs/man3/html/glVertexAttrib.xhtml
func VertexAttrib4f(dst Attrib, x, y, z, w float32) {
	_pluginInstance.glContext.VertexAttrib4f(gl.Attrib{dst.Value}, x, y, z, w)
}

// VertexAttrib4fv writes a vec4 vertex attribute.
//
// http://www.khronos.org/opengles/sdk/docs/man3/html/glVertexAttrib.xhtml
func VertexAttrib4fv(dst Attrib, src []float32) {
	_pluginInstance.glContext.VertexAttrib4fv(gl.Attrib{dst.Value}, src)
}

// VertexAttribPointer uses a bound buffer to define vertex attribute data.
//
// http://www.khronos.org/opengles/sdk/docs/man3/html/glVertexAttribPointer.xhtml
func VertexAttribPointer(dst Attrib, size int, ty Enum, normalized bool, stride, offset int) {
	_pluginInstance.glContext.VertexAttribPointer(gl.Attrib{dst.Value}, size, gl.Enum(ty), normalized, stride, offset)
}

// Viewport sets the viewport, an affine transformation that
// normalizes device coordinates to window coordinates.
//
// http://www.khronos.org/opengles/sdk/docs/man3/html/glViewport.xhtml
func Viewport(x, y, width, height int) {
	_pluginInstance.glContext.Viewport(x, y, width, height)
}
