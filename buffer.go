// Copyright 2012 The go-gl Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package gl

// #include "gl.h"
import "C"

// Buffer Objects

type Buffer Object

// Create single buffer object
func GenBuffer() Buffer {
	var b C.GLuint
	C.glGenBuffers(1, &b)
	return Buffer(b)
}

// Fill slice with new buffers
func GenBuffers(buffers []Buffer) {
	if len(buffers) > 0 {
		C.glGenBuffers(C.GLsizei(len(buffers)), (*C.GLuint)(&buffers[0]))
	}
}

// Delete buffer object
func (buffer Buffer) Delete() {
	b := C.GLuint(buffer)
	C.glDeleteBuffers(1, &b)
}

// Delete all buffers in slice
func DeleteBuffers(buffers []Buffer) {
	if len(buffers) > 0 {
		C.glDeleteBuffers(C.GLsizei(len(buffers)), (*C.GLuint)(&buffers[0]))
	}
}

// Bind this buffer as target
func (buffer Buffer) Bind(target GLenum) {
	C.glBindBuffer(C.GLenum(target), C.GLuint(buffer))
}

// Remove buffer binding
func (buffer Buffer) Unbind(target GLenum) {
	C.glBindBuffer(C.GLenum(target), C.GLuint(0))
}

// Creates and initializes a buffer object's data store
func BufferData(target GLenum, size int, data interface{}, usage GLenum) {
	C.glBufferData(C.GLenum(target), C.GLsizeiptr(size), ptr(data), C.GLenum(usage))
}

//  Update a subset of a buffer object's data store
func BufferSubData(target GLenum, offset int, size int, data interface{}) {
	C.glBufferSubData(C.GLenum(target), C.GLintptr(offset), C.GLsizeiptr(size),
		ptr(data))
}

// Return parameters of a buffer object
func GetBufferParameteriv(target GLenum, pname GLenum) int32 {
	var param C.GLint
	C.glGetBufferParameteriv(C.GLenum(target), C.GLenum(pname), &param)
	return int32(param)
}
