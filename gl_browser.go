// Copyright (c) 2019 Thomas MILLET. All rights reserved.

// +build js

package gl

import (
	binary "encoding/binary"
	fmt "fmt"
	math "math"
	js "syscall/js"
	unsafe "unsafe"

	tge "github.com/thommil/tge"
)

type plugin struct {
	glContext *js.Value
}

func (p *plugin) Init(runtime tge.Runtime) error {
	renderer := runtime.GetRenderer()
	switch renderer.(type) {
	case *js.Value:
		p.glContext = renderer.(*js.Value)
	default:
		return fmt.Errorf("Runtime renderer must be a *syscall/js.Value")
	}

	programMap[NONE] = js.Null()
	shaderMap[NONE] = js.Null()
	bufferMap[NONE] = js.Null()
	framebufferMap[NONE] = js.Null()
	renderbufferMap[NONE] = js.Null()
	textureMap[NONE] = js.Null()
	vertexArrayMap[NONE] = js.Null()

	return nil
}

func (p *plugin) Dispose() {
	FlushCache()
}

// GetGLSLVersion gives the glsl version ti put in #version ${VERSION}
func GetGLSLVersion() string {
	return "300 es"
}

// FlushCache free memory cache, should be called between scenes
func FlushCache() {
	for k := range float32TypedArrayCacheMap {
		delete(float32TypedArrayCacheMap, k)
	}
	for k := range int32TypedArrayCacheMap {
		delete(int32TypedArrayCacheMap, k)
	}

	for k := range programMap {
		delete(programMap, k)
	}
	programMapIndex = Program(1)

	for k := range shaderMap {
		delete(shaderMap, k)
	}
	shaderMapIndex = Shader(1)

	for k := range bufferMap {
		delete(bufferMap, k)
	}
	bufferMapIndex = Buffer(1)

	for k := range framebufferMap {
		delete(framebufferMap, k)
	}
	framebufferMapIndex = Framebuffer(1)

	for k := range renderbufferMap {
		delete(renderbufferMap, k)
	}
	renderbufferMapIndex = Renderbuffer(1)

	for k := range textureMap {
		delete(textureMap, k)
	}
	textureMapIndex = Texture(1)

	for k := range uniformMap {
		delete(uniformMap, k)
	}
	uniformMapIndex = Uniform(0)

	for k := range vertexArrayMap {
		delete(vertexArrayMap, k)
	}
	vertexArrayMapIndex = VertexArray(1)

	int32ArrayBuffer = make([]int32, 0)
	int32ArrayBufferExtendFactor = 1

	float32ArrayBuffer = make([]float32, 0)
	float32ArrayBufferExtendFactor = 1

	byteArrayBuffer = make([]byte, 0)
	byteArrayBufferExtendFactor = 1
}

var programMap = make(map[Program]js.Value)
var programMapIndex = Program(1)

var shaderMap = make(map[Shader]js.Value)
var shaderMapIndex = Shader(1)

var bufferMap = make(map[Buffer]js.Value)
var bufferMapIndex = Buffer(1)

var framebufferMap = make(map[Framebuffer]js.Value)
var framebufferMapIndex = Framebuffer(1)

var renderbufferMap = make(map[Renderbuffer]js.Value)
var renderbufferMapIndex = Renderbuffer(1)

var textureMap = make(map[Texture]js.Value)
var textureMapIndex = Texture(1)

var uniformMap = make(map[Uniform]js.Value)
var uniformMapIndex = Uniform(0)

var vertexArrayMap = make(map[VertexArray]js.Value)
var vertexArrayMapIndex = VertexArray(1)

func ActiveTexture(texture Enum) {
	_pluginInstance.glContext.Call("activeTexture", int(texture))
}

func AttachShader(p Program, s Shader) {
	_pluginInstance.glContext.Call("attachShader", programMap[p], shaderMap[s])
}

func BindAttribLocation(p Program, a Attrib, name string) {
	_pluginInstance.glContext.Call("bindAttribLocation", programMap[p], int32(a), name)
}

func BindBuffer(target Enum, b Buffer) {
	_pluginInstance.glContext.Call("bindBuffer", int(target), bufferMap[b])
}

func BindFramebuffer(target Enum, fb Framebuffer) {
	_pluginInstance.glContext.Call("bindFramebuffer", int(target), framebufferMap[fb])
}

func BindRenderbuffer(target Enum, rb Renderbuffer) {
	_pluginInstance.glContext.Call("bindRenderbuffer", int(target), renderbufferMap[rb])
}

func BindTexture(target Enum, t Texture) {
	_pluginInstance.glContext.Call("bindTexture", int(target), textureMap[t])
}

func BindVertexArray(vao VertexArray) {
	_pluginInstance.glContext.Call("bindVertexArray", vertexArrayMap[vao])
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

func BufferInit(target Enum, size int, usage Enum) {
	js.TypedArrayOf(getByteArrayBuffer(size))
	_pluginInstance.glContext.Call("bufferData", int(target), size, int(usage))
}

func BufferData(target Enum, src []byte, usage Enum) {
	srcTA := js.TypedArrayOf(src)
	_pluginInstance.glContext.Call("bufferData", int(target), srcTA, int(usage))
	srcTA.Release()
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
	_pluginInstance.glContext.Call("compileShader", shaderMap[s])
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
	bufferMap[bufferMapIndex] = _pluginInstance.glContext.Call("createBuffer")
	buffer := Buffer(bufferMapIndex)
	bufferMapIndex++
	return buffer
}

func CreateFramebuffer() Framebuffer {
	framebufferMap[framebufferMapIndex] = _pluginInstance.glContext.Call("createFramebuffer")
	framebuffer := Framebuffer(framebufferMapIndex)
	framebufferMapIndex++
	return framebuffer
}

func CreateProgram() Program {
	programMap[programMapIndex] = _pluginInstance.glContext.Call("createProgram")
	program := Program(programMapIndex)
	programMapIndex++
	return program
}

func CreateRenderbuffer() Renderbuffer {
	renderbufferMap[renderbufferMapIndex] = _pluginInstance.glContext.Call("createRenderbuffer")
	renderbuffer := Renderbuffer(renderbufferMapIndex)
	renderbufferMapIndex++
	return renderbuffer
}

func CreateShader(ty Enum) Shader {
	shaderMap[shaderMapIndex] = _pluginInstance.glContext.Call("createShader", int(ty))
	shader := Shader(shaderMapIndex)
	shaderMapIndex++
	return shader
}

func CreateTexture() Texture {
	textureMap[textureMapIndex] = _pluginInstance.glContext.Call("createTexture")
	texture := Texture(textureMapIndex)
	textureMapIndex++
	return texture
}

func CreateVertexArray() VertexArray {
	vertexArrayMap[vertexArrayMapIndex] = _pluginInstance.glContext.Call("createVertexArray")
	vao := VertexArray(vertexArrayMapIndex)
	vertexArrayMapIndex++
	return vao
}

func CullFace(mode Enum) {
	_pluginInstance.glContext.Call("cullFace", int(mode))
}

func DeleteBuffer(v Buffer) {
	_pluginInstance.glContext.Call("deleteBuffer", bufferMap[v])
	delete(bufferMap, v)
}

func DeleteFramebuffer(v Framebuffer) {
	_pluginInstance.glContext.Call("deleteFramebuffer", framebufferMap[v])
	delete(framebufferMap, v)
}

func DeleteProgram(p Program) {
	_pluginInstance.glContext.Call("deleteProgram", programMap[p])
	delete(programMap, p)
}

func DeleteRenderbuffer(v Renderbuffer) {
	_pluginInstance.glContext.Call("deleteRenderbuffer", renderbufferMap[v])
	delete(renderbufferMap, v)
}

func DeleteShader(s Shader) {
	_pluginInstance.glContext.Call("deleteShader", shaderMap[s])
	delete(shaderMap, s)
}

func DeleteTexture(v Texture) {
	_pluginInstance.glContext.Call("deleteTexture", textureMap[v])
	delete(textureMap, v)
}

func DeleteVertexArray(v VertexArray) {
	_pluginInstance.glContext.Call("DeleteVertexArray", vertexArrayMap[v])
	delete(vertexArrayMap, v)
}

func DepthFunc(fn Enum) {
	_pluginInstance.glContext.Call("depthFunc", uint32(fn))
}

func DepthMask(flag bool) {
	_pluginInstance.glContext.Call("depthMask", flag)
}

func DepthRangef(n, f float32) {
	_pluginInstance.glContext.Call("depthRange", n, f)
}

func DetachShader(p Program, s Shader) {
	_pluginInstance.glContext.Call("detachShader", programMap[p], shaderMap[s])
}

func Disable(cap Enum) {
	_pluginInstance.glContext.Call("disable", int(cap))
}

func DisableVertexAttribArray(a Attrib) {
	_pluginInstance.glContext.Call("disableVertexAttribArray", int32(a))
}

func DrawArrays(mode Enum, first, count int) {
	_pluginInstance.glContext.Call("drawArrays", int(mode), first, count)
}

func DrawElements(mode Enum, count int, ty Enum, offset int) {
	_pluginInstance.glContext.Call("drawElements", int(mode), count, int(ty), offset)
}

func Enable(cap Enum) {
	_pluginInstance.glContext.Call("enable", uint32(cap))
}

func EnableVertexAttribArray(a Attrib) {
	_pluginInstance.glContext.Call("enableVertexAttribArray", int32(a))
}

func Finish() {
	_pluginInstance.glContext.Call("finish")
}

func Flush() {
	_pluginInstance.glContext.Call("flush")
}

func FramebufferRenderbuffer(target, attachment, rbTarget Enum, rb Renderbuffer) {
	_pluginInstance.glContext.Call("framebufferRenderbuffer", target, attachment, int(rbTarget), renderbufferMap[rb])
}

func FramebufferTexture2D(target, attachment, texTarget Enum, t Texture, level int) {
	_pluginInstance.glContext.Call("framebufferTexture2D", target, attachment, int(texTarget), textureMap[t], level)
}

func FrontFace(mode Enum) {
	_pluginInstance.glContext.Call("frontFace", int(mode))
}

func GenerateMipmap(target Enum) {
	_pluginInstance.glContext.Call("generateMipmap", int(target))
}

func GetActiveAttrib(p Program, index uint32) (name string, size int, ty Enum) {
	ai := _pluginInstance.glContext.Call("getActiveAttrib", programMap[p], index)
	return ai.Get("name").String(), ai.Get("size").Int(), Enum(ai.Get("type").Int())
}

func GetActiveUniform(p Program, index uint32) (name string, size int, ty Enum) {
	ai := _pluginInstance.glContext.Call("getActiveUniform", programMap[p], index)
	return ai.Get("name").String(), ai.Get("size").Int(), Enum(ai.Get("type").Int())
}

func GetAttachedShaders(p Program) []Shader {
	fmt.Printf("WARNING: GetAttachedShaders not implemented\n")
	return []Shader{}
}

func GetAttribLocation(p Program, name string) Attrib {
	return Attrib(int32(_pluginInstance.glContext.Call("getAttribLocation", programMap[p], name).Int()))
}

func GetBooleanv(dst []bool, pname Enum) {
	result := _pluginInstance.glContext.Call("getParameter", int(pname))
	length := result.Length()
	for i := 0; i < length; i++ {
		dst[i] = result.Index(i).Bool()
	}
}

func GetFloatv(dst []float32, pname Enum) {
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
	fmt.Printf("WARNING: GetAttachedShaders not implemented\n")
	return NONE
}

func GetFramebufferAttachmentParameteri(target, attachment, pname Enum) int {
	return _pluginInstance.glContext.Call("getFramebufferAttachmentParameter", int(target), int(attachment), int(pname)).Int()
}

func GetProgrami(p Program, pname Enum) int {
	switch pname {
	case DELETE_STATUS, LINK_STATUS, VALIDATE_STATUS:
		if _pluginInstance.glContext.Call("getProgramParameter", programMap[p], int(pname)).Bool() {
			return TRUE
		}
		return FALSE
	default:
		return _pluginInstance.glContext.Call("getProgramParameter", programMap[p], int(pname)).Int()
	}
}

func GetProgramInfoLog(p Program) string {
	return _pluginInstance.glContext.Call("getProgramInfoLog", programMap[p]).String()
}

func GetRenderbufferParameteri(target, pname Enum) int {
	return _pluginInstance.glContext.Call("getRenderbufferParameter", int(target), int(pname)).Int()
}

func GetShaderi(s Shader, pname Enum) int {
	switch pname {
	case DELETE_STATUS, COMPILE_STATUS:
		if _pluginInstance.glContext.Call("getShaderParameter", shaderMap[s], int(pname)).Bool() {
			return TRUE
		}
		return FALSE
	default:
		return _pluginInstance.glContext.Call("getShaderParameter", shaderMap[s], int(pname)).Int()
	}
}

func GetShaderInfoLog(s Shader) string {
	return _pluginInstance.glContext.Call("getShaderInfoLog", shaderMap[s]).String()
}

func GetShaderPrecisionFormat(shadertype, precisiontype Enum) (rangeMin, rangeMax, precision int) {
	format := _pluginInstance.glContext.Call("getShaderPrecisionFormat", uint32(shadertype), uint32(precisiontype))
	rangeMin = format.Get("rangeMin").Int()
	rangeMax = format.Get("rangeMax").Int()
	precision = format.Get("precision").Int()
	return
}

func GetShaderSource(s Shader) string {
	return _pluginInstance.glContext.Call("getShaderSource", shaderMap[s]).String()
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
	result := _pluginInstance.glContext.Call("getUniform")
	length := result.Length()
	for i := 0; i < length; i++ {
		dst[i] = float32(result.Index(i).Float())
	}
}

func GetUniformiv(dst []int32, src Uniform, p Program) {
	result := _pluginInstance.glContext.Call("getUniform")
	length := result.Length()
	for i := 0; i < length; i++ {
		dst[i] = int32(result.Index(i).Int())
	}
}

func GetUniformLocation(p Program, name string) Uniform {
	uniform := _pluginInstance.glContext.Call("getUniformLocation", programMap[p], name)
	uniformIndex := *(*Uniform)(unsafe.Pointer(&uniform))
	uniformMap[uniformIndex] = uniform
	return Uniform(uniformIndex)

}

func GetVertexAttribf(src Attrib, pname Enum) float32 {
	return float32(_pluginInstance.glContext.Call("getVertexAttrib", int32(src), int(pname)).Float())
}

func GetVertexAttribfv(dst []float32, src Attrib, pname Enum) {
	result := _pluginInstance.glContext.Call("getVertexAttrib")
	length := result.Length()
	for i := 0; i < length; i++ {
		dst[i] = float32(result.Index(i).Float())
	}
}

func GetVertexAttribi(src Attrib, pname Enum) int32 {
	return int32(_pluginInstance.glContext.Call("getVertexAttrib", int32(src), int(pname)).Int())
}

func GetVertexAttribiv(dst []int32, src Attrib, pname Enum) {
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
	if buffer, found := bufferMap[b]; found {
		return _pluginInstance.glContext.Call("isBuffer", buffer).Bool()
	}
	return false
}

func IsEnabled(cap Enum) bool {
	return _pluginInstance.glContext.Call("isEnabled", int(cap)).Bool()
}

func IsFramebuffer(fb Framebuffer) bool {
	if framebuffer, found := framebufferMap[fb]; found {
		return _pluginInstance.glContext.Call("isFramebuffer", framebuffer).Bool()
	}
	return false
}

func IsProgram(p Program) bool {
	if program, found := programMap[p]; found {
		return _pluginInstance.glContext.Call("isProgram", program).Bool()
	}
	return false
}

func IsRenderbuffer(rb Renderbuffer) bool {
	if renderbuffer, found := renderbufferMap[rb]; found {
		return _pluginInstance.glContext.Call("isRenderbuffer", renderbuffer).Bool()
	}
	return false
}

func IsShader(s Shader) bool {
	if shader, found := shaderMap[s]; found {
		return _pluginInstance.glContext.Call("isShader", shader).Bool()
	}
	return false
}

func IsTexture(t Texture) bool {
	if texture, found := textureMap[t]; found {
		return _pluginInstance.glContext.Call("isTexture", texture).Bool()
	}
	return false
}

func LineWidth(width float32) {
	_pluginInstance.glContext.Call("lineWidth", width)
}

func LinkProgram(p Program) {
	_pluginInstance.glContext.Call("linkProgram", programMap[p])
}

func PixelStorei(pname Enum, param int32) {
	_pluginInstance.glContext.Call("pixelStorei", int(pname), param)
}

func PolygonOffset(factor, units float32) {
	_pluginInstance.glContext.Call("polygonOffset", factor, units)
}

func PolygonMode(face, mode Enum) {
	fmt.Printf("WARNING: PolygonMode not implemented\n")
}

func ReadPixels(dst []byte, x, y, width, height int, format, ty Enum) {
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
	fmt.Printf("WARNING: ReleaseShaderCompiler not implemented\n")
}

func RenderbufferStorage(target, internalFormat Enum, width, height int) {
	_pluginInstance.glContext.Call("renderbufferStorage", target, uint32(internalFormat), width, height)
}

func SampleCoverage(value float32, invert bool) {
	_pluginInstance.glContext.Call("sampleCoverage", value, invert)
}

func Scissor(x, y, width, height int32) {
	_pluginInstance.glContext.Call("scissor", x, y, width, height)
}

func ShaderSource(s Shader, src string) {
	_pluginInstance.glContext.Call("shaderSource", shaderMap[s], src)
}

func StencilFunc(fn Enum, ref int, mask uint32) {
	_pluginInstance.glContext.Call("stencilFunc", uint32(fn), ref, mask)
}

func StencilFuncSeparate(face, fn Enum, ref int, mask uint32) {
	_pluginInstance.glContext.Call("stencilFuncSeparate", uint32(face), uint32(fn), ref, mask)
}

func StencilMask(mask uint32) {
	_pluginInstance.glContext.Call("stencilMask", mask)
}

func StencilMaskSeparate(face Enum, mask uint32) {
	_pluginInstance.glContext.Call("stencilMaskSeparate", uint32(face), mask)
}

func StencilOp(fail, zfail, zpass Enum) {
	_pluginInstance.glContext.Call("stencilOp", uint32(fail), uint32(zfail), uint32(zpass))
}

func StencilOpSeparate(face, sfail, dpfail, dppass Enum) {
	_pluginInstance.glContext.Call("stencilOpSeparate", uint32(face), uint32(sfail), uint32(dpfail), uint32(dppass))
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
	for _, param := range params {
		_pluginInstance.glContext.Call("texParameterf", int(target), int(pname), param)
	}
}

func TexParameteri(target, pname Enum, param int) {
	_pluginInstance.glContext.Call("texParameteri", int(target), int(pname), param)
}

func TexParameteriv(target, pname Enum, params []int32) {
	for _, param := range params {
		_pluginInstance.glContext.Call("texParameteri", int(target), int(pname), param)
	}
}

//go:linkname memmove runtime.memmove
func memmove(to, from unsafe.Pointer, n uintptr)

// int32 array singleton, allocate 1KB at startup
var int32ArrayBuffer = make([]int32, 0)
var int32ArrayBufferExtendFactor = 1

const int32Offset = unsafe.Sizeof(int32(0))

func getInt32ArrayBuffer(size int) []int32 {
	if size > len(int32ArrayBuffer) {
		for (1024 * int32ArrayBufferExtendFactor) < size {
			int32ArrayBufferExtendFactor++
		}
		int32ArrayBuffer = make([]int32, (1024 * int32ArrayBufferExtendFactor))
	}
	return int32ArrayBuffer[:size]
}

var int32TypedArrayCacheMap = make(map[uintptr]*js.TypedArray)

// No affectation if done due to magic coincidence between this cache and syscall.js one \o/
func getInt32TypedArrayFromCache(src []int32) *js.TypedArray {
	key := uintptr(unsafe.Pointer(&src[0])) + uintptr(len(src))
	if b, found := int32TypedArrayCacheMap[key]; found {
		return b
	} else {
		b := js.TypedArrayOf(src)
		int32TypedArrayCacheMap[key] = &b
		return &b
	}
}

func getInt32TypedArrayFromCacheP(size int, src *int32) *js.TypedArray {
	b := getInt32ArrayBuffer(size)
	memmove(unsafe.Pointer(&b[0]), unsafe.Pointer(src), uintptr(size)*int32Offset)
	return getInt32TypedArrayFromCache(b)
}

func getInt32TypedArrayFromCacheUP(size int, src unsafe.Pointer) *js.TypedArray {
	b := getInt32ArrayBuffer(size)
	memmove(unsafe.Pointer(&b[0]), src, uintptr(size)*int32Offset)
	return getInt32TypedArrayFromCache(b)
}

// float32 array singleton, allocate 1KB at startup
var float32ArrayBuffer = make([]float32, 0)
var float32ArrayBufferExtendFactor = 1

const float32Offset = unsafe.Sizeof(float32(0))

func getFloat32ArrayBuffer(size int) []float32 {
	if size > len(float32ArrayBuffer) {
		for (1024 * float32ArrayBufferExtendFactor) < size {
			float32ArrayBufferExtendFactor++
		}
		float32ArrayBuffer = make([]float32, (1024 * float32ArrayBufferExtendFactor))
	}
	return float32ArrayBuffer[:size]
}

var float32TypedArrayCacheMap = make(map[uintptr]*js.TypedArray)

// Hack using backed array of js.TypeArray and internal cache of syscall/js
func getFloat32TypedArrayFromCache(src []float32) *js.TypedArray {
	key := uintptr(unsafe.Pointer(&src[0])) + uintptr(len(src))
	if b, found := float32TypedArrayCacheMap[key]; found {
		return b
	} else {
		b := js.TypedArrayOf(src)
		float32TypedArrayCacheMap[key] = &b
		return &b
	}
}

func getFloat32TypedArrayFromCacheP(size int, src *float32) *js.TypedArray {
	b := getFloat32ArrayBuffer(size)
	memmove(unsafe.Pointer(&b[0]), unsafe.Pointer(src), uintptr(size)*float32Offset)
	return getFloat32TypedArrayFromCache(b)
}

func getFloat32TypedArrayFromCacheUP(size int, src unsafe.Pointer) *js.TypedArray {
	b := getFloat32ArrayBuffer(size)
	memmove(unsafe.Pointer(&b[0]), src, uintptr(size)*float32Offset)
	return getFloat32TypedArrayFromCache(b)
}

func Uniform1f(dst Uniform, v float32) {
	_pluginInstance.glContext.Call("uniform1f", uniformMap[dst], v)
}

func Uniform1fv(dst Uniform, src []float32) {
	_pluginInstance.glContext.Call("uniform1fv", uniformMap[dst], *getFloat32TypedArrayFromCache(src))
}

func Uniform1fvP(dst Uniform, count int32, value *float32) {
	_pluginInstance.glContext.Call("uniform1fv", uniformMap[dst], *getFloat32TypedArrayFromCacheP(int(count), value))
}

func Uniform1fvUP(dst Uniform, count int32, value unsafe.Pointer) {
	_pluginInstance.glContext.Call("uniform1fv", uniformMap[dst], *getFloat32TypedArrayFromCacheUP(int(count), value))
}

func Uniform1i(dst Uniform, v int) {
	_pluginInstance.glContext.Call("uniform1i", uniformMap[dst], v)
}

func Uniform1iv(dst Uniform, src []int32) {
	_pluginInstance.glContext.Call("uniform1iv", uniformMap[dst], *getInt32TypedArrayFromCache(src))
}

func Uniform1ivP(dst Uniform, count int32, value *int32) {
	_pluginInstance.glContext.Call("uniform1iv", uniformMap[dst], *getInt32TypedArrayFromCacheP(int(count), value))
}

func Uniform1ivUP(dst Uniform, count int32, value unsafe.Pointer) {
	_pluginInstance.glContext.Call("uniform1iv", uniformMap[dst], *getInt32TypedArrayFromCacheUP(int(count), value))
}

func Uniform2f(dst Uniform, v0, v1 float32) {
	_pluginInstance.glContext.Call("uniform2f", uniformMap[dst], v0, v1)
}

func Uniform2fv(dst Uniform, src []float32) {
	_pluginInstance.glContext.Call("uniform2fv", uniformMap[dst], *getFloat32TypedArrayFromCache(src))
}

func Uniform2fvP(dst Uniform, count int32, value *float32) {
	_pluginInstance.glContext.Call("uniform2fv", uniformMap[dst], *getFloat32TypedArrayFromCacheP(int(count*2), value))
}

func Uniform2fvUP(dst Uniform, count int32, value unsafe.Pointer) {
	_pluginInstance.glContext.Call("uniform2fv", uniformMap[dst], *getFloat32TypedArrayFromCacheUP(int(count*2), value))
}

func Uniform2i(dst Uniform, v0, v1 int) {
	_pluginInstance.glContext.Call("uniform2i", uniformMap[dst], v0, v1)
}

func Uniform2iv(dst Uniform, src []int32) {
	_pluginInstance.glContext.Call("uniform2iv", uniformMap[dst], *getInt32TypedArrayFromCache(src))
}

func Uniform2ivP(dst Uniform, count int32, value *int32) {
	_pluginInstance.glContext.Call("uniform2iv", uniformMap[dst], *getInt32TypedArrayFromCacheP(int(count*2), value))
}

func Uniform2ivUP(dst Uniform, count int32, value unsafe.Pointer) {
	_pluginInstance.glContext.Call("uniform2iv", uniformMap[dst], *getInt32TypedArrayFromCacheUP(int(count*2), value))
}

func Uniform3f(dst Uniform, v0, v1, v2 float32) {
	_pluginInstance.glContext.Call("uniform3f", uniformMap[dst], v0, v1, v2)
}

func Uniform3fv(dst Uniform, src []float32) {
	_pluginInstance.glContext.Call("uniform3fv", uniformMap[dst], *getFloat32TypedArrayFromCache(src))
}

func Uniform3fvP(dst Uniform, count int32, value *float32) {
	_pluginInstance.glContext.Call("uniform3fv", uniformMap[dst], *getFloat32TypedArrayFromCacheP(int(count*3), value))
}

func Uniform3fvUP(dst Uniform, count int32, value unsafe.Pointer) {
	_pluginInstance.glContext.Call("uniform3fv", uniformMap[dst], *getFloat32TypedArrayFromCacheUP(int(count*3), value))
}

func Uniform3i(dst Uniform, v0, v1, v2 int32) {
	_pluginInstance.glContext.Call("uniform3i", uniformMap[dst], v0, v1, v2)
}

func Uniform3iv(dst Uniform, src []int32) {
	_pluginInstance.glContext.Call("uniform3iv", uniformMap[dst], *getInt32TypedArrayFromCache(src))
}

func Uniform3ivP(dst Uniform, count int32, value *int32) {
	_pluginInstance.glContext.Call("uniform3iv", uniformMap[dst], *getInt32TypedArrayFromCacheP(int(count*3), value))
}

func Uniform3ivUP(dst Uniform, count int32, value unsafe.Pointer) {
	_pluginInstance.glContext.Call("uniform3iv", uniformMap[dst], *getInt32TypedArrayFromCacheUP(int(count*3), value))
}

func Uniform4f(dst Uniform, v0, v1, v2, v3 float32) {
	_pluginInstance.glContext.Call("uniform4f", uniformMap[dst], v0, v1, v2, v3)
}

func Uniform4fv(dst Uniform, src []float32) {
	_pluginInstance.glContext.Call("uniform4fv", uniformMap[dst], *getFloat32TypedArrayFromCache(src))
}

func Uniform4fvP(dst Uniform, count int32, value *float32) {
	_pluginInstance.glContext.Call("uniform4fv", uniformMap[dst], *getFloat32TypedArrayFromCacheP(int(count*4), value))
}

func Uniform4fvUP(dst Uniform, count int32, value unsafe.Pointer) {
	_pluginInstance.glContext.Call("uniform4fv", uniformMap[dst], *getFloat32TypedArrayFromCacheUP(int(count*4), value))
}

func Uniform4i(dst Uniform, v0, v1, v2, v3 int32) {
	_pluginInstance.glContext.Call("uniform4i", uniformMap[dst], v0, v1, v2, v3)
}

func Uniform4iv(dst Uniform, src []int32) {
	_pluginInstance.glContext.Call("uniform4iv", uniformMap[dst], *getInt32TypedArrayFromCache(src))
}

func Uniform4ivP(dst Uniform, count int32, value *int32) {
	_pluginInstance.glContext.Call("uniform4iv", uniformMap[dst], *getInt32TypedArrayFromCacheP(int(count*4), value))
}

func Uniform4ivUP(dst Uniform, count int32, value unsafe.Pointer) {
	_pluginInstance.glContext.Call("uniform4iv", uniformMap[dst], *getInt32TypedArrayFromCacheUP(int(count*4), value))
}

func UniformMatrix2fv(dst Uniform, transpose bool, src []float32) {
	_pluginInstance.glContext.Call("uniformMatrix2fv", uniformMap[dst], transpose, *getFloat32TypedArrayFromCache(src))
}

func UniformMatrix2fvP(dst Uniform, count int32, transpose bool, value *float32) {
	_pluginInstance.glContext.Call("uniformMatrix2fv", uniformMap[dst], transpose, *getFloat32TypedArrayFromCacheP(int(count*4), value))
}

func UniformMatrix2fvUP(dst Uniform, count int32, transpose bool, value unsafe.Pointer) {
	_pluginInstance.glContext.Call("uniformMatrix2fv", uniformMap[dst], transpose, *getFloat32TypedArrayFromCacheUP(int(count*4), value))
}

func UniformMatrix3fv(dst Uniform, transpose bool, src []float32) {
	_pluginInstance.glContext.Call("uniformMatrix3fv", uniformMap[dst], transpose, *getFloat32TypedArrayFromCache(src))
}

func UniformMatrix3fvP(dst Uniform, count int32, transpose bool, value *float32) {
	_pluginInstance.glContext.Call("uniformMatrix3fv", uniformMap[dst], transpose, *getFloat32TypedArrayFromCacheP(int(count*9), value))
}

func UniformMatrix3fvUP(dst Uniform, count int32, transpose bool, value unsafe.Pointer) {
	_pluginInstance.glContext.Call("uniformMatrix3fv", uniformMap[dst], transpose, *getFloat32TypedArrayFromCacheUP(int(count*9), value))
}

func UniformMatrix4fv(dst Uniform, transpose bool, src []float32) {
	_pluginInstance.glContext.Call("uniformMatrix4fv", uniformMap[dst], transpose, *getFloat32TypedArrayFromCache(src))
}

func UniformMatrix4fvP(dst Uniform, count int32, transpose bool, value *float32) {
	_pluginInstance.glContext.Call("uniformMatrix4fv", uniformMap[dst], transpose, *getFloat32TypedArrayFromCacheP(int(count*16), value))
}

func UniformMatrix4fvUP(dst Uniform, count int32, transpose bool, value unsafe.Pointer) {
	_pluginInstance.glContext.Call("uniformMatrix4fv", uniformMap[dst], transpose, *getFloat32TypedArrayFromCacheUP(int(count*16), value))
}

func UseProgram(p Program) {
	_pluginInstance.glContext.Call("useProgram", programMap[p])
}

func ValidateProgram(p Program) {
	_pluginInstance.glContext.Call("validateProgram", programMap[p])
}

func VertexAttrib1f(dst Attrib, x float32) {
	_pluginInstance.glContext.Call("vertexAttrib1f", int32(dst), x)
}

func VertexAttrib1fv(dst Attrib, src []float32) {
	_pluginInstance.glContext.Call("vertexAttrib1fv", int32(dst), src)
}

func VertexAttrib2f(dst Attrib, x, y float32) {
	_pluginInstance.glContext.Call("vertexAttrib2f", int32(dst), x, y)
}

func VertexAttrib2fv(dst Attrib, src []float32) {
	_pluginInstance.glContext.Call("vertexAttrib2fv", int32(dst), src)
}

func VertexAttrib3f(dst Attrib, x, y, z float32) {
	_pluginInstance.glContext.Call("vertexAttrib3f", int32(dst), x, y, z)
}

func VertexAttrib3fv(dst Attrib, src []float32) {
	_pluginInstance.glContext.Call("vertexAttrib3fv", int32(dst), src)
}

func VertexAttrib4f(dst Attrib, x, y, z, w float32) {
	_pluginInstance.glContext.Call("vertexAttrib4f", int32(dst), x, y, z, w)
}

func VertexAttrib4fv(dst Attrib, src []float32) {
	_pluginInstance.glContext.Call("vertexAttrib4fv", int32(dst), src)
}

func VertexAttribPointer(dst Attrib, size int, ty Enum, normalized bool, stride, offset int) {
	_pluginInstance.glContext.Call("vertexAttribPointer", int32(dst), size, int(ty), normalized, stride, offset)
}

func Viewport(x, y, width, height int) {
	_pluginInstance.glContext.Call("viewport", x, y, width, height)
}
