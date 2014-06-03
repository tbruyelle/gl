// Copyright 2012 The go-gl Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package gl

// #include "gl.h"
import "C"

// Vertex Arrays
type VertexArray Object

func GenVertexArray() VertexArray {
	var a C.GLuint
	C.glGenVertexArraysOES(1, &a)
	return VertexArray(a)
}

func GenVertexArrays(arrays []VertexArray) {
	if len(arrays) > 0 {
		C.glGenVertexArraysOES(C.GLsizei(len(arrays)), (*C.GLuint)(&arrays[0]))
	}
}

func (array VertexArray) Delete() {
	C.glDeleteVertexArraysOES(1, (*C.GLuint)(&array))
}

func DeleteVertexArrays(arrays []VertexArray) {
	if len(arrays) > 0 {
		C.glDeleteVertexArraysOES(C.GLsizei(len(arrays)), (*C.GLuint)(&arrays[0]))
	}
}

func (array VertexArray) Bind() {
	C.glBindVertexArrayOES(C.GLuint(array))
}

func (array VertexArray) Unbind() {
	C.glBindVertexArrayOES(C.GLuint(0))
}
