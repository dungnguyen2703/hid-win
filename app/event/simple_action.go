package event

import (
	"encoding/json"
	"fmt"
	"hidtool/app/keyboard"
	"time"
)

// -------------------------Press Action-------------------------
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

// -------------------------TapDown Action-------------------------
type TapDown struct {
	Key keyboard.KEY `json:"key"`
}

func (p *TapDown) UnmarshalJSON(data []byte) error {
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

func (k TapDown) Run() {
	KeyDown(k.Key)
}

// -------------------------TapUp Action-------------------------
type TapUp struct {
	Key keyboard.KEY `json:"key"`
}

func (p *TapUp) UnmarshalJSON(data []byte) error {
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

func (k TapUp) Run() {
	KeyUp(k.Key)
}

// -------------------------Delay Action-------------------------
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
