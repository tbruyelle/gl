// Copyright 2012 The go-gl Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package gl

// #include "gl.h"
import "C"

// Framebuffer Objects
// TODO: implement GLsync stuff
type Framebuffer Object

// void glBindFramebuffer(GLenum target, GLuint framebuffer);
//
// Binds fb to target FRAMEBUFFER. To bind to a specific target, see BindTarget.
func (fb Framebuffer) Bind() {
	C.glBindFramebuffer(C.GLenum(FRAMEBUFFER), C.GLuint(fb))
}

// Binds fb to the specified target.
//
// See issue at github for why this function exists:
// http://github.com/go-gl/gl/issues/113
func (fb Framebuffer) BindTarget(target GLenum) {
	C.glBindFramebuffer(C.GLenum(target), C.GLuint(fb))
}

// Unbinds target FRAMEBUFFER. To unbind a a specific target, see UnbindTarget.
func (fb Framebuffer) Unbind() {
	C.glBindFramebuffer(C.GLenum(FRAMEBUFFER), 0)
}

// Unbinds the specified target.
//
// See issue at github for why this function exists:
// http://github.com/go-gl/gl/issues/113
func (fb Framebuffer) UnbindTarget(target GLenum) {
	C.glBindFramebuffer(C.GLenum(target), 0)
}

// GLenum glCheckFramebufferStatus(GLenum target);
func CheckFramebufferStatus(target GLenum) GLenum {
	return (GLenum)(C.glCheckFramebufferStatus(C.GLenum(target)))
}

// void glDeleteFramebuffers(GLsizei n, GLuint* framebuffers);
func (fb Framebuffer) Delete() {
	C.glDeleteFramebuffers(1, (*C.GLuint)(&fb))
}

func DeleteFramebuffers(bufs []Framebuffer) {
	if len(bufs) > 0 {
		C.glDeleteFramebuffers(C.GLsizei(len(bufs)), (*C.GLuint)(&bufs[0]))
	}
}

// void glFramebufferTexture2D(GLenum target, GLenum attachment, GLenum textarget, GLuint texture, GLint level);
func FramebufferTexture2D(target, attachment, textarget GLenum, texture Texture, level int) {
	C.glFramebufferTexture2D(C.GLenum(target), C.GLenum(attachment), C.GLenum(textarget), C.GLuint(texture), C.GLint(level))
}

// void glGenFramebuffers(GLsizei n, GLuint* ids);
func GenFramebuffer() Framebuffer {
	var b C.GLuint
	C.glGenFramebuffers(1, &b)
	return Framebuffer(b)
}

func GenFramebuffers(bufs []Framebuffer) {
	if len(bufs) > 0 {
		C.glGenFramebuffers(C.GLsizei(len(bufs)), (*C.GLuint)(&bufs[0]))
	}
}
