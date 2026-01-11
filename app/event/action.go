package event

import (
	"hidtool/app/keyboard"
	"time"
)

type Action interface {
	Run()
}

type Press struct {
	key keyboard.KEY
}

func (k Press) Run() {
	KeyPressed(k.key)
}

type Delay struct {
	DurationMs int
}

func (d Delay) Run() {
	time.Sleep(time.Duration(d.DurationMs) * time.Millisecond)
}

type WindowLeft struct{}

func (w WindowLeft) Run() {
	KeyDown(keyboard.CTRL)
	KeyDown(keyboard.WIN)
	KeyDown(keyboard.ARROW_LEFT)
	KeyUp(keyboard.ARROW_LEFT)
	KeyUp(keyboard.WIN)
	KeyUp(keyboard.CTRL)

}

type WindowRight struct{}

func (w WindowRight) Run() {
	KeyDown(keyboard.CTRL)
	KeyDown(keyboard.WIN)
	KeyDown(keyboard.ARROW_RIGHT)
	KeyUp(keyboard.ARROW_RIGHT)
	KeyUp(keyboard.WIN)
	KeyUp(keyboard.CTRL)
}
