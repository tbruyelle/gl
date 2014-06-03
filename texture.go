// Copyright 2012 The go-gl Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package gl

// #include "gl.h"
import "C"

func ActiveTexture(texture GLenum) { C.glActiveTexture(C.GLenum(texture)) }

// Texture

type Texture Object

// Create single texture object
func GenTexture() Texture {
	var b C.GLuint
	C.glGenTextures(1, &b)
	return Texture(b)
}

// Fill slice with new textures
func GenTextures(textures []Texture) {
	if len(textures) > 0 {
		C.glGenTextures(C.GLsizei(len(textures)), (*C.GLuint)(&textures[0]))
	}
}

// Delete texture object
func (texture Texture) Delete() {
	b := C.GLuint(texture)
	C.glDeleteTextures(1, &b)
}

// Delete all textures in slice
func DeleteTextures(textures []Texture) {
	if len(textures) > 0 {
		C.glDeleteTextures(C.GLsizei(len(textures)), (*C.GLuint)(&textures[0]))
	}
}

// Bind this texture as target
func (texture Texture) Bind(target GLenum) {
	C.glBindTexture(C.GLenum(target), C.GLuint(texture))
}

// Unbind this texture
func (texture Texture) Unbind(target GLenum) {
	C.glBindTexture(C.GLenum(target), 0)
}
