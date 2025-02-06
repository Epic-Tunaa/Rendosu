package input

import (
	"github.com/go-gl/glfw/v3.3/glfw"
)

type InputManager struct {
	keys	[512]bool
	mouse	[8]bool
	mousePos	struct(X, Y float64)
	keyChan	chan glfw.Key
}

func NewInputManager(window *glfw.Window) *InputManager {
	im := &InputManager{
		keyChan: make(chan glfw.Key, 1024)
	}
	window.SetKeyCallback(im.keyCallback)
	window.SetMouseButtonCallback(im.mouseCallback)
	window.SetCursorPosCallback(im.mousePosCallback)

	return im
}

func (im *InputManager) keyCallback(w *glfw.Window, key glfw.Key, scancode int, action glfw.Action, mods glfw.ModifierKey) {
	if action == glfw.Press || action == glfw.Release {
		im.keys[key] = action == glfw.Press
		select {
		case im.keyChan <- key:
		default:
		}
	}
}

func (im *InputManager) isKeyPressed(key glfw.key) bool {
	return im.keys[key]
}