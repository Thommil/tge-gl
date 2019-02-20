// Copyright 2018 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build js

package gl

import (
	binary "encoding/binary"
	fmt "fmt"
	math "math"
	js "syscall/js"

	tge "github.com/thommil/tge"
)

type plugin struct {
	glContext *js.Value
}

var _pluginInstance = &plugin{}

func (p *plugin) Init(runtime tge.Runtime) error {
	renderer := runtime.GetRenderer()
	switch renderer.(type) {
	case *js.Value:
		p.glContext = renderer.(*js.Value)
	default:
		return fmt.Errorf("Runtime renderer must be a *syscall/js.Value")
	}
	return nil
}

func (p *plugin) GetName() string {
	return Name
}

func (p *plugin) Dispose() {
	p.glContext = nil
}

// GetPlugin returns plugin handler
func GetPlugin() tge.Plugin {
	return _pluginInstance
}

func ActiveTexture(texture Enum) {
	_pluginInstance.glContext.Call("activeTexture", int(texture))
}

func AttachShader(p Program, s Shader) {
	_pluginInstance.glContext.Call("attachShader", p.Value, s.Value)
}

func BindAttribLocation(p Program, a Attrib, name string) {
	_pluginInstance.glContext.Call("bindAttribLocation", p.Value, a.Value, name)
}

func BindBuffer(target Enum, b Buffer) {
	_pluginInstance.glContext.Call("bindBuffer", int(target), b.Value)
}

func BindFramebuffer(target Enum, fb Framebuffer) {
	_pluginInstance.glContext.Call("bindFramebuffer", int(target), fb.Value)
}

func BindRenderbuffer(target Enum, rb Renderbuffer) {
	_pluginInstance.glContext.Call("bindRenderbuffer", int(target), rb.Value)
}

func BindTexture(target Enum, t Texture) {
	_pluginInstance.glContext.Call("bindTexture", int(target), t.Value)
}

func BlendColor(red, green, blue, alpha float32) {
	_pluginInstance.glContext.Call("blendColor", red, green, blue, alpha)
}

func BlendEquation(mode Enum) {
	_pluginInstance.glContext.Call("blendEquation", int(mode))
}

func BlendEquationSeparate(modeRGB, modeAlpha Enum) {
	_pluginInstance.glContext.Call("blendEquationSeparate", modeRGB, modeAlpha)
}

func BlendFunc(sfactor, dfactor Enum) {
	_pluginInstance.glContext.Call("blendFunc", int(sfactor), int(dfactor))
}

func BlendFuncSeparate(sfactorRGB, dfactorRGB, sfactorAlpha, dfactorAlpha Enum) {
	_pluginInstance.glContext.Call("blendFuncSeparate", int(sfactorRGB), int(dfactorRGB), int(sfactorAlpha), int(dfactorAlpha))
}

func BufferData(target Enum, src []byte, usage Enum) {
	srcTA := js.TypedArrayOf(src)
	_pluginInstance.glContext.Call("bufferData", int(target), srcTA, int(usage))
	srcTA.Release()
}

func BufferInit(target Enum, size int, usage Enum) {
	_pluginInstance.glContext.Call("bufferData", int(target), size, int(usage))
}

func BufferSubData(target Enum, offset int, data []byte) {
	_pluginInstance.glContext.Call("bufferSubData", int(target), offset, data)
}

func CheckFramebufferStatus(target Enum) Enum {
	return Enum(_pluginInstance.glContext.Call("checkFramebufferStatus", int(target)).Int())
}

func Clear(mask Enum) {
	_pluginInstance.glContext.Call("clear", int(mask))
}

func ClearColor(red, green, blue, alpha float32) {
	_pluginInstance.glContext.Call("clearColor", red, green, blue, alpha)
}

func ClearDepthf(d float32) {
	_pluginInstance.glContext.Call("clearDepth", d)
}

func ClearStencil(s int) {
	_pluginInstance.glContext.Call("clearStencil", s)
}

func ColorMask(red, green, blue, alpha bool) {
	_pluginInstance.glContext.Call("colorMask", red, green, blue, alpha)
}

func CompileShader(s Shader) {
	_pluginInstance.glContext.Call("compileShader", s.Value)
}

func CompressedTexImage2D(target Enum, level int, internalformat Enum, width, height, border int, data []byte) {
	_pluginInstance.glContext.Call("compressedTexImage2D", int(target), level, internalformat, width, height, border, data)
}

func CompressedTexSubImage2D(target Enum, level, xoffset, yoffset, width, height int, format Enum, data []byte) {
	_pluginInstance.glContext.Call("compressedTexSubImage2D", int(target), level, xoffset, yoffset, width, height, format, data)
}

func CopyTexImage2D(target Enum, level int, internalformat Enum, x, y, width, height, border int) {
	_pluginInstance.glContext.Call("copyTexImage2D", int(target), level, internalformat, x, y, width, height, border)
}

func CopyTexSubImage2D(target Enum, level, xoffset, yoffset, x, y, width, height int) {
	_pluginInstance.glContext.Call("copyTexSubImage2D", int(target), level, xoffset, yoffset, x, y, width, height)
}

func CreateBuffer() Buffer {
	return Buffer{Value: _pluginInstance.glContext.Call("createBuffer")}
}

func CreateFramebuffer() Framebuffer {
	return Framebuffer{Value: _pluginInstance.glContext.Call("createFramebuffer")}
}

func CreateProgram() Program {
	return Program{Value: _pluginInstance.glContext.Call("createProgram")}
}

func CreateRenderbuffer() Renderbuffer {
	return Renderbuffer{Value: _pluginInstance.glContext.Call("createRenderbuffer")}
}

func CreateShader(ty Enum) Shader {
	return Shader{Value: _pluginInstance.glContext.Call("createShader", int(ty))}
}

func CreateTexture() Texture {
	return Texture{Value: _pluginInstance.glContext.Call("createTexture")}
}

func CullFace(mode Enum) {
	_pluginInstance.glContext.Call("cullFace", int(mode))
}

func DeleteBuffer(v Buffer) {
	_pluginInstance.glContext.Call("deleteBuffer", v.Value)
}

func DeleteFramebuffer(v Framebuffer) {
	_pluginInstance.glContext.Call("deleteFramebuffer", v.Value)
}

func DeleteProgram(p Program) {
	_pluginInstance.glContext.Call("deleteProgram", p.Value)
}

func DeleteRenderbuffer(v Renderbuffer) {
	_pluginInstance.glContext.Call("deleteRenderbuffer", v.Value)
}

func DeleteShader(s Shader) {
	_pluginInstance.glContext.Call("deleteShader", s.Value)
}

func DeleteTexture(v Texture) {
	_pluginInstance.glContext.Call("deleteTexture", v.Value)
}

func DepthFunc(fn Enum) {
	_pluginInstance.glContext.Call("depthFunc", fn)
}

func DepthMask(flag bool) {
	_pluginInstance.glContext.Call("depthMask", flag)
}

func DepthRangef(n, f float32) {
	_pluginInstance.glContext.Call("depthRange", n, f)
}

func DetachShader(p Program, s Shader) {
	_pluginInstance.glContext.Call("detachShader", p.Value, s.Value)
}

func Disable(cap Enum) {
	_pluginInstance.glContext.Call("disable", int(cap))
}

func DisableVertexAttribArray(a Attrib) {
	_pluginInstance.glContext.Call("disableVertexAttribArray", a.Value)
}

func DrawArrays(mode Enum, first, count int) {
	_pluginInstance.glContext.Call("drawArrays", int(mode), first, count)
}

func DrawElements(mode Enum, count int, ty Enum, offset int) {
	_pluginInstance.glContext.Call("drawElements", int(mode), count, int(ty), offset)
}

func Enable(cap Enum) {
	_pluginInstance.glContext.Call("enable", int(cap))
}

func EnableVertexAttribArray(a Attrib) {
	_pluginInstance.glContext.Call("enableVertexAttribArray", a.Value)
}

func Finish() {
	_pluginInstance.glContext.Call("finish")
}

func Flush() {
	_pluginInstance.glContext.Call("flush")
}

func FramebufferRenderbuffer(target, attachment, rbTarget Enum, rb Renderbuffer) {
	_pluginInstance.glContext.Call("framebufferRenderbuffer", target, attachment, int(rbTarget), rb.Value)
}

func FramebufferTexture2D(target, attachment, texTarget Enum, t Texture, level int) {
	_pluginInstance.glContext.Call("framebufferTexture2D", target, attachment, int(texTarget), t.Value, level)
}

func FrontFace(mode Enum) {
	_pluginInstance.glContext.Call("frontFace", int(mode))
}

func GenerateMipmap(target Enum) {
	_pluginInstance.glContext.Call("generateMipmap", int(target))
}

func GetActiveAttrib(p Program, index uint32) (name string, size int, ty Enum) {
	ai := _pluginInstance.glContext.Call("getActiveAttrib", p.Value, index)
	return ai.Get("name").String(), ai.Get("size").Int(), Enum(ai.Get("type").Int())
}

func GetActiveUniform(p Program, index uint32) (name string, size int, ty Enum) {
	ai := _pluginInstance.glContext.Call("getActiveUniform", p.Value, index)
	return ai.Get("name").String(), ai.Get("size").Int(), Enum(ai.Get("type").Int())
}

func GetAttachedShaders(p Program) []Shader {
	objs := _pluginInstance.glContext.Call("getAttachedShaders", p.Value)
	shaders := make([]Shader, objs.Length())
	for i := 0; i < objs.Length(); i++ {
		shaders[i] = Shader{Value: objs.Index(i)}
	}
	return shaders
}

func GetAttribLocation(p Program, name string) Attrib {
	return Attrib{Value: _pluginInstance.glContext.Call("getAttribLocation", p.Value, name).Int()}
}

func GetBooleanv(dst []bool, pname Enum) {
	println("GetBooleanv: not yet tested (TODO: remove this after it's confirmed to work. Your feedback is welcome.)")
	result := _pluginInstance.glContext.Call("getParameter", int(pname))
	length := result.Length()
	for i := 0; i < length; i++ {
		dst[i] = result.Index(i).Bool()
	}
}

func GetFloatv(dst []float32, pname Enum) {
	println("GetFloatv: not yet tested (TODO: remove this after it's confirmed to work. Your feedback is welcome.)")
	result := _pluginInstance.glContext.Call("getParameter", int(pname))
	length := result.Length()
	for i := 0; i < length; i++ {
		dst[i] = float32(result.Index(i).Float())
	}
}

func GetIntegerv(pname Enum, data []int32) {
	result := _pluginInstance.glContext.Call("getParameter", int(pname))
	length := result.Length()
	for i := 0; i < length; i++ {
		data[i] = int32(result.Index(i).Int())
	}
}

func GetInteger(pname Enum) int {
	return _pluginInstance.glContext.Call("getParameter", int(pname)).Int()
}

func GetBufferParameteri(target, pname Enum) int {
	return _pluginInstance.glContext.Call("getBufferParameter", int(target), int(pname)).Int()
}

func GetError() Enum {
	return Enum(_pluginInstance.glContext.Call("getError").Int())
}

func GetBoundFramebuffer() Framebuffer {
	return Framebuffer{Value: _pluginInstance.glContext.Call("getParameter", FRAMEBUFFER_BINDING)}
}

func GetFramebufferAttachmentParameteri(target, attachment, pname Enum) int {
	return _pluginInstance.glContext.Call("getFramebufferAttachmentParameter", int(target), int(attachment), int(pname)).Int()
}

func GetProgrami(p Program, pname Enum) int {
	switch pname {
	case DELETE_STATUS, LINK_STATUS, VALIDATE_STATUS:
		if _pluginInstance.glContext.Call("getProgramParameter", p.Value, int(pname)).Bool() {
			return TRUE
		}
		return FALSE
	default:
		return _pluginInstance.glContext.Call("getProgramParameter", p.Value, int(pname)).Int()
	}
}

func GetProgramInfoLog(p Program) string {
	return _pluginInstance.glContext.Call("getProgramInfoLog", p.Value).String()
}

func GetRenderbufferParameteri(target, pname Enum) int {
	return _pluginInstance.glContext.Call("getRenderbufferParameter", int(target), int(pname)).Int()
}

func GetShaderi(s Shader, pname Enum) int {
	switch pname {
	case DELETE_STATUS, COMPILE_STATUS:
		if _pluginInstance.glContext.Call("getShaderParameter", s.Value, int(pname)).Bool() {
			return TRUE
		}
		return FALSE
	default:
		return _pluginInstance.glContext.Call("getShaderParameter", s.Value, int(pname)).Int()
	}
}

func GetShaderInfoLog(s Shader) string {
	return _pluginInstance.glContext.Call("getShaderInfoLog", s.Value).String()
}

func GetShaderPrecisionFormat(shadertype, precisiontype Enum) (rangeMin, rangeMax, precision int) {
	println("GetShaderPrecisionFormat: not yet tested (TODO: remove this after it's confirmed to work. Your feedback is welcome.)")
	format := _pluginInstance.glContext.Call("getShaderPrecisionFormat", shadertype, precisiontype)
	rangeMin = format.Get("rangeMin").Int()
	rangeMax = format.Get("rangeMax").Int()
	precision = format.Get("precision").Int()
	return
}

func GetShaderSource(s Shader) string {
	return _pluginInstance.glContext.Call("getShaderSource", s.Value).String()
}

func GetString(pname Enum) string {
	return _pluginInstance.glContext.Call("getParameter", int(pname)).String()
}

func GetTexParameterfv(dst []float32, target, pname Enum) {
	dst[0] = float32(_pluginInstance.glContext.Call("getTexParameter", int(pname)).Float())
}

func GetTexParameteriv(dst []int32, target, pname Enum) {
	dst[0] = int32(_pluginInstance.glContext.Call("getTexParameter", int(pname)).Int())
}

func GetUniformfv(dst []float32, src Uniform, p Program) {
	println("GetUniformfv: not yet tested (TODO: remove this after it's confirmed to work. Your feedback is welcome.)")
	result := _pluginInstance.glContext.Call("getUniform")
	length := result.Length()
	for i := 0; i < length; i++ {
		dst[i] = float32(result.Index(i).Float())
	}
}

func GetUniformiv(dst []int32, src Uniform, p Program) {
	println("GetUniformiv: not yet tested (TODO: remove this after it's confirmed to work. Your feedback is welcome.)")
	result := _pluginInstance.glContext.Call("getUniform")
	length := result.Length()
	for i := 0; i < length; i++ {
		dst[i] = int32(result.Index(i).Int())
	}
}

func GetUniformLocation(p Program, name string) Uniform {
	return Uniform{Value: _pluginInstance.glContext.Call("getUniformLocation", p.Value, name)}
}

func GetVertexAttribf(src Attrib, pname Enum) float32 {
	return float32(_pluginInstance.glContext.Call("getVertexAttrib", src.Value, int(pname)).Float())
}

func GetVertexAttribfv(dst []float32, src Attrib, pname Enum) {
	println("GetVertexAttribfv: not yet tested (TODO: remove this after it's confirmed to work. Your feedback is welcome.)")
	result := _pluginInstance.glContext.Call("getVertexAttrib")
	length := result.Length()
	for i := 0; i < length; i++ {
		dst[i] = float32(result.Index(i).Float())
	}
}

func GetVertexAttribi(src Attrib, pname Enum) int32 {
	return int32(_pluginInstance.glContext.Call("getVertexAttrib", src.Value, int(pname)).Int())
}

func GetVertexAttribiv(dst []int32, src Attrib, pname Enum) {
	println("GetVertexAttribiv: not yet tested (TODO: remove this after it's confirmed to work. Your feedback is welcome.)")
	result := _pluginInstance.glContext.Call("getVertexAttrib")
	length := result.Length()
	for i := 0; i < length; i++ {
		dst[i] = int32(result.Index(i).Int())
	}
}

func Hint(target, mode Enum) {
	_pluginInstance.glContext.Call("hint", int(target), int(mode))
}

func IsBuffer(b Buffer) bool {
	return _pluginInstance.glContext.Call("isBuffer", b.Value).Bool()
}

func IsEnabled(cap Enum) bool {
	return _pluginInstance.glContext.Call("isEnabled", int(cap)).Bool()
}

func IsFramebuffer(fb Framebuffer) bool {
	return _pluginInstance.glContext.Call("isFramebuffer", fb.Value).Bool()
}

func IsProgram(p Program) bool {
	return _pluginInstance.glContext.Call("isProgram", p.Value).Bool()
}

func IsRenderbuffer(rb Renderbuffer) bool {
	return _pluginInstance.glContext.Call("isRenderbuffer", rb.Value).Bool()
}

func IsShader(s Shader) bool {
	return _pluginInstance.glContext.Call("isShader", s.Value).Bool()
}

func IsTexture(t Texture) bool {
	return _pluginInstance.glContext.Call("isTexture", t.Value).Bool()
}

func LineWidth(width float32) {
	_pluginInstance.glContext.Call("lineWidth", width)
}

func LinkProgram(p Program) {
	_pluginInstance.glContext.Call("linkProgram", p.Value)
}

func PixelStorei(pname Enum, param int32) {
	_pluginInstance.glContext.Call("pixelStorei", int(pname), param)
}

func PolygonOffset(factor, units float32) {
	_pluginInstance.glContext.Call("polygonOffset", factor, units)
}

func ReadPixels(dst []byte, x, y, width, height int, format, ty Enum) {
	println("ReadPixels: not yet tested (TODO: remove this after it's confirmed to work. Your feedback is welcome.)")
	if ty == Enum(UNSIGNED_BYTE) {
		_pluginInstance.glContext.Call("readPixels", x, y, width, height, format, int(ty), dst)
	} else {
		tmpDst := make([]float32, len(dst)/4)
		_pluginInstance.glContext.Call("readPixels", x, y, width, height, format, int(ty), tmpDst)
		for i, f := range tmpDst {
			binary.LittleEndian.PutUint32(dst[i*4:], math.Float32bits(f))
		}
	}
}

func ReleaseShaderCompiler() {
	// do nothing
}

func RenderbufferStorage(target, internalFormat Enum, width, height int) {
	_pluginInstance.glContext.Call("renderbufferStorage", target, internalFormat, width, height)
}

func SampleCoverage(value float32, invert bool) {
	_pluginInstance.glContext.Call("sampleCoverage", value, invert)
}

func Scissor(x, y, width, height int32) {
	_pluginInstance.glContext.Call("scissor", x, y, width, height)
}

func ShaderSource(s Shader, src string) {
	_pluginInstance.glContext.Call("shaderSource", s.Value, src)
}

func StencilFunc(fn Enum, ref int, mask uint32) {
	_pluginInstance.glContext.Call("stencilFunc", fn, ref, mask)
}

func StencilFuncSeparate(face, fn Enum, ref int, mask uint32) {
	_pluginInstance.glContext.Call("stencilFuncSeparate", face, fn, ref, mask)
}

func StencilMask(mask uint32) {
	_pluginInstance.glContext.Call("stencilMask", mask)
}

func StencilMaskSeparate(face Enum, mask uint32) {
	_pluginInstance.glContext.Call("stencilMaskSeparate", face, mask)
}

func StencilOp(fail, zfail, zpass Enum) {
	_pluginInstance.glContext.Call("stencilOp", fail, zfail, zpass)
}

func StencilOpSeparate(face, sfail, dpfail, dppass Enum) {
	_pluginInstance.glContext.Call("stencilOpSeparate", face, sfail, dpfail, dppass)
}

func TexImage2D(target Enum, level int, width, height int, format Enum, ty Enum, data []byte) {
	var p interface{}
	if data != nil {
		dataTA := js.TypedArrayOf(data)
		defer dataTA.Release()
		p = dataTA
	}
	_pluginInstance.glContext.Call("texImage2D", int(target), level, int(format), width, height, 0, int(format), int(ty), p)
}

func TexSubImage2D(target Enum, level int, x, y, width, height int, format, ty Enum, data []byte) {
	_pluginInstance.glContext.Call("texSubImage2D", int(target), level, x, y, width, height, format, int(ty), data)
}

func TexParameterf(target, pname Enum, param float32) {
	_pluginInstance.glContext.Call("texParameterf", int(target), int(pname), param)
}

func TexParameterfv(target, pname Enum, params []float32) {
	println("TexParameterfv: not yet tested (TODO: remove this after it's confirmed to work. Your feedback is welcome.)")
	for _, param := range params {
		_pluginInstance.glContext.Call("texParameterf", int(target), int(pname), param)
	}
}

func TexParameteri(target, pname Enum, param int) {
	_pluginInstance.glContext.Call("texParameteri", int(target), int(pname), param)
}

func TexParameteriv(target, pname Enum, params []int32) {
	println("TexParameteriv: not yet tested (TODO: remove this after it's confirmed to work. Your feedback is welcome.)")
	for _, param := range params {
		_pluginInstance.glContext.Call("texParameteri", int(target), int(pname), param)
	}
}

func Uniform1f(dst Uniform, v float32) {
	_pluginInstance.glContext.Call("uniform1f", dst.Value, v)
}

func Uniform1fv(dst Uniform, src []float32) {
	_pluginInstance.glContext.Call("uniform1fv", dst.Value, src)
}

func Uniform1i(dst Uniform, v int) {
	_pluginInstance.glContext.Call("uniform1i", dst.Value, v)
}

func Uniform1iv(dst Uniform, src []int32) {
	_pluginInstance.glContext.Call("uniform1iv", dst.Value, src)
}

func Uniform2f(dst Uniform, v0, v1 float32) {
	_pluginInstance.glContext.Call("uniform2f", dst.Value, v0, v1)
}

func Uniform2fv(dst Uniform, src []float32) {
	_pluginInstance.glContext.Call("uniform2fv", dst.Value, src)
}

func Uniform2i(dst Uniform, v0, v1 int) {
	_pluginInstance.glContext.Call("uniform2i", dst.Value, v0, v1)
}

func Uniform2iv(dst Uniform, src []int32) {
	_pluginInstance.glContext.Call("uniform2iv", dst.Value, src)
}

func Uniform3f(dst Uniform, v0, v1, v2 float32) {
	_pluginInstance.glContext.Call("uniform3f", dst.Value, v0, v1, v2)
}

func Uniform3fv(dst Uniform, src []float32) {
	_pluginInstance.glContext.Call("uniform3fv", dst.Value, src)
}

func Uniform3i(dst Uniform, v0, v1, v2 int32) {
	_pluginInstance.glContext.Call("uniform3i", dst.Value, v0, v1, v2)
}

func Uniform3iv(dst Uniform, src []int32) {
	_pluginInstance.glContext.Call("uniform3iv", dst.Value, src)
}

func Uniform4f(dst Uniform, v0, v1, v2, v3 float32) {
	_pluginInstance.glContext.Call("uniform4f", dst.Value, v0, v1, v2, v3)
}

func Uniform4fv(dst Uniform, src []float32) {
	_pluginInstance.glContext.Call("uniform4fv", dst.Value, src)
}

func Uniform4i(dst Uniform, v0, v1, v2, v3 int32) {
	_pluginInstance.glContext.Call("uniform4i", dst.Value, v0, v1, v2, v3)
}

func Uniform4iv(dst Uniform, src []int32) {
	_pluginInstance.glContext.Call("uniform4iv", dst.Value, src)
}

func UniformMatrix2fv(dst Uniform, src []float32) {
	_pluginInstance.glContext.Call("uniformMatrix2fv", dst.Value, false, src)
}

func UniformMatrix3fv(dst Uniform, src []float32) {
	_pluginInstance.glContext.Call("uniformMatrix3fv", dst.Value, false, src)
}

func UniformMatrix4fv(dst Uniform, src []float32) {
	srcTA := js.TypedArrayOf(src)
	_pluginInstance.glContext.Call("uniformMatrix4fv", dst.Value, false, srcTA)
	srcTA.Release()
}

func UseProgram(p Program) {
	// Workaround for js.Value zero value.
	if p.Value == (js.Value{}) {
		p.Value = js.Null()
	}
	_pluginInstance.glContext.Call("useProgram", p.Value)
}

func ValidateProgram(p Program) {
	_pluginInstance.glContext.Call("validateProgram", p.Value)
}

func VertexAttrib1f(dst Attrib, x float32) {
	_pluginInstance.glContext.Call("vertexAttrib1f", dst.Value, x)
}

func VertexAttrib1fv(dst Attrib, src []float32) {
	_pluginInstance.glContext.Call("vertexAttrib1fv", dst.Value, src)
}

func VertexAttrib2f(dst Attrib, x, y float32) {
	_pluginInstance.glContext.Call("vertexAttrib2f", dst.Value, x, y)
}

func VertexAttrib2fv(dst Attrib, src []float32) {
	_pluginInstance.glContext.Call("vertexAttrib2fv", dst.Value, src)
}

func VertexAttrib3f(dst Attrib, x, y, z float32) {
	_pluginInstance.glContext.Call("vertexAttrib3f", dst.Value, x, y, z)
}

func VertexAttrib3fv(dst Attrib, src []float32) {
	_pluginInstance.glContext.Call("vertexAttrib3fv", dst.Value, src)
}

func VertexAttrib4f(dst Attrib, x, y, z, w float32) {
	_pluginInstance.glContext.Call("vertexAttrib4f", dst.Value, x, y, z, w)
}

func VertexAttrib4fv(dst Attrib, src []float32) {
	_pluginInstance.glContext.Call("vertexAttrib4fv", dst.Value, src)
}

func VertexAttribPointer(dst Attrib, size int, ty Enum, normalized bool, stride, offset int) {
	_pluginInstance.glContext.Call("vertexAttribPointer", dst.Value, size, int(ty), normalized, stride, offset)
}

func Viewport(x, y, width, height int) {
	_pluginInstance.glContext.Call("viewport", x, y, width, height)
}
