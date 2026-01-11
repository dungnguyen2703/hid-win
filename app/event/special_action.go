package event

import "hidtool/app/keyboard"

// -------------------------WindowLeft Action-------------------------
type WindowLeft struct{}

func (w WindowLeft) Run() {
	KeyDown(keyboard.CTRL)
	KeyDown(keyboard.WIN)
	KeyDown(keyboard.ARROW_LEFT)
	KeyUp(keyboard.ARROW_LEFT)
	KeyUp(keyboard.WIN)
	KeyUp(keyboard.CTRL)

}

// -------------------------WindowRight Action-------------------------
type WindowRight struct{}

func (w WindowRight) Run() {
	KeyDown(keyboard.CTRL)
	KeyDown(keyboard.WIN)
	KeyDown(keyboard.ARROW_RIGHT)
	KeyUp(keyboard.ARROW_RIGHT)
	KeyUp(keyboard.WIN)
	KeyUp(keyboard.CTRL)
}
