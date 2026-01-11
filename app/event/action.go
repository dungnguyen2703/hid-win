package event

import (
	"encoding/json"
	"fmt"
	"hidtool/app/keyboard"
	"time"
)

func UnmarshalAction(raw json.RawMessage) (Action, error) {
	var meta struct {
		Type string `json:"type"`
	}
	if err := json.Unmarshal(raw, &meta); err != nil {
		return nil, err
	}
	switch meta.Type {
	case "window_left":
		return WindowLeft{}, nil
	case "window_right":
		return WindowRight{}, nil
	case "press":
		var a Press
		if err := json.Unmarshal(raw, &a); err != nil {
			return nil, err
		}
		return a, nil
	case "delay":
		var a Delay
		if err := json.Unmarshal(raw, &a); err != nil {
			return nil, err
		}
		return a, nil
	default:
		return nil, fmt.Errorf("unknown action type: %s", meta.Type)
	}
}

type Action interface {
	Run()
}

type Press struct {
	Key keyboard.KEY `json:"key"`
}

func (p *Press) UnmarshalJSON(data []byte) error {
	var raw struct {
		Key interface{} `json:"key"`
	}
	if err := json.Unmarshal(data, &raw); err != nil {
		return err
	}
	switch v := raw.Key.(type) {
	case string:
		if val, ok := keyboard.GetKeyFromName(v); ok {
			p.Key = val
		} else {
			return fmt.Errorf("unknown key name: %s", v)
		}
	default:
		return fmt.Errorf("invalid key type: must be a string name of the key")
	}
	return nil
}

func (k Press) Run() {
	KeyPressed(k.Key)
}

type Delay struct {
	DurationMs int `json:"duration"`
}

func (d *Delay) UnmarshalJSON(data []byte) error {
	type Alias Delay
	aux := &struct {
		*Alias
	}{
		Alias: (*Alias)(d),
	}
	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}
	if d.DurationMs <= 0 {
		return fmt.Errorf("delay duration must be greater than 0")
	}
	return nil
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
