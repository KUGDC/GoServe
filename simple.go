// Ported from GLUT's samples.  Original copyright below applies.

/* Copyright (c) Mark J. Kilgard, 1996. */

/* This program is freely distributable without licensing fees
   and is provided without guarantee or warrantee expressed or
   implied. This program is -not- in the public domain. */

/* This program is a response to a question posed by Gil Colgate
   <gcolgate@sirius.com> about how lengthy a program is required using
   OpenGL compared to using  Direct3D immediate mode to "draw a
   triangle at screen coordinates 0,0, to 200,200 to 20,200, and I
   want it to be blue at the top vertex, red at the left vertex, and
   green at the right vertex".  I'm not sure how long the Direct3D
   program is; Gil has used Direct3D and his guess is "about 3000
   lines of code". */

package main

import (
	"github.com/go-gl/gl"
	"github.com/rhencke/glut"
    "fmt"
    "container/list"
)

type Entity struct {
    rot_matrix [4]float32  // row major order
    pos        [2]float32  // pos offsets
    id         string      // unique identifier for obj
    entities   *list.List  // reference to the entity list
}

type GameState struct {
    leftPos float32
    units *list.List
    main_unit *Entity
}

var t GameState = GameState{}
var currentWindow glut.Window;

func reshape(w, h int) {
	/* Because Gil specified "screen coordinates" (presumably with an
	   upper-left origin), this short bit of code sets up the coordinate
	   system to correspond to actual window coodrinates.  This code
	   wouldn't be required if you chose a (more typical in 3D) abstract
	   coordinate system. */

    /* Establish viewing area to cover entire window. */
	gl.Viewport(0, 0, w, h)
    /* Start modifying the projection matrix. */
	gl.MatrixMode(gl.PROJECTION)
    /* Reset project matrix. */
	gl.LoadIdentity()
    /* Map abstract coords directly to window coords. */
	gl.Ortho(0, float64(w), 0, float64(h), -1, 1)
    /* Invert Y axis so increasing Y goes down. */
	gl.Scalef(1, -1, 1)
    /* Shift origin up to upper-left corner. */
	gl.Translatef(0, float32(-h), 0)
}

func display() {
	gl.Clear(gl.COLOR_BUFFER_BIT)

	gl.Begin(gl.TRIANGLES)
    l := t.units
    for e := l.Front(); e != nil; e = e.Next() {
        b := e.Value.(*Entity)
        //r := b.rot_matrix
        p := b.pos
        gl.Color3f(0.0, 0.0, 1.0) /* blue */
        gl.Vertex2f(p[0], p[1])
        gl.Color3f(0.0, 1.0, 0.0) /* green */
        gl.Vertex2f(20 + p[0], 20 + p[1])
        gl.Color3f(1.0, 0.0, 0.0) /* red */
        gl.Vertex2f(p[0], 20 + p[1])
    }
	gl.End()
	gl.Flush() /* Single buffered, so needs a flush. */

    glut.SwapBuffers()
}

func keyboardIn(key byte, x, y int) {
}

func specialIn(key, x, y int) {
    // If they're pressing the left key
    if (key == glut.KEY_LEFT) {
        t.main_unit.pos[0] -= 5.0
    }
    if (key == glut.KEY_RIGHT) {
        t.main_unit.pos[0] += 5.0
    }
    if (key == glut.KEY_UP) {
        t.main_unit.pos[1] -= 5.0
    }
    if (key == glut.KEY_DOWN) {
        t.main_unit.pos[1] += 5.0
    }
}

// Abstract away our logging to change later
func Log(v ...interface{}) {
    fmt.Println(v...)
}

func initWindow() {

    // Function called to do the re-rendering
    glut.DisplayFunc(display)
    // Called when the visibility of the program changes
	glut.VisibilityFunc(visible)
    // Called when a regular ascii character is pressed
    glut.KeyboardFunc(keyboardIn)
    // Called when any non-ascii character is pressed
    glut.SpecialFunc(specialIn)
    // Called when the size of the window changes
	glut.ReshapeFunc(reshape)

    // affect our projection matrix
    gl.MatrixMode(gl.PROJECTION)
    gl.LoadIdentity()  // Load an identity matrix -> projection
    // Specify the bounds of the of our scene
    gl.Ortho(0, 40, 0, 40, 0, 40)  

    // Now affect our modelview matrix
    gl.MatrixMode(gl.MODELVIEW)
    // Make points be rendererd larger
    gl.PointSize(3.0)

    currentWindow = glut.GetWindow()
}

func idle() {
	currentWindow.PostRedisplay()
}

func visible(vis int) {
	//if vis == glut.visible {
		//if !paused {
			glut.IdleFunc(idle)
		//}
	//} else {
	//	glut.idlefunc(nil)
	//}
}

func main() {
	glut.CreateWindow("Triangle Demo")
    t.units = list.New()
    e := Entity{ [...]float32{1,0,0,1},
                 [...]float32{0, 0},
                 "test",
                 t.units }
    t.units.PushBack(&e)
    t.main_unit = &e
    initWindow()
	glut.MainLoop()
}
