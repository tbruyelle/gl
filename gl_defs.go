// Copyright 2012 The go-gl Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package gl

// #include "gl.h"
import "C"

// Constants
const (
	FRAMEBUFFER  = C.GL_FRAMEBUFFER
	RENDERBUFFER = C.GL_RENDERBUFFER
	ACTIVE_UNIFORM_MAX_LENGTH = C.GL_ACTIVE_UNIFORM_MAX_LENGTH
	INFO_LOG_LENGTH=C.GL_INFO_LOG_LENGTH
)
