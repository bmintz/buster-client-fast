package mouse

import (
	"github.com/go-vgo/robotgo"
)

func MoveMouse(x, y int) {
	robotgo.MoveMouse(x, y)
}

func ClickMouse() {
	robotgo.MouseClick()
}
