package event

import (
	"encoding/json"
	"fmt"
)

type Action interface {
	Run()
}

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
	case "tap_down":
		var a TapDown
		if err := json.Unmarshal(raw, &a); err != nil {
			return nil, err
		}
		return a, nil
	case "tap_up":
		var a TapUp
		if err := json.Unmarshal(raw, &a); err != nil {
			return nil, err
		}
		return a, nil
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
