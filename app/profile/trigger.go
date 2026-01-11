package profile

import (
	"encoding/json"
	"fmt"
	"hidtool/app/keyboard"
	"hidtool/app/mice"
)

type Trigger interface {
	isTrigger(keyboard.KEY, mice.BUTTON) bool
}

func UnmarshalTrigger(raw json.RawMessage) (Trigger, error) {
	var meta struct {
		Type string `json:"type"`
	}
	if err := json.Unmarshal(raw, &meta); err != nil {
		return nil, err
	}
	// Case-insensitive check could be good, but stick to exact for now based on JSON
	switch meta.Type {
	case "key", "Key": // supporting both temporarily if needed, or strictly one. user used "key" in previous jsons.
		var t KeyTrigger
		if err := json.Unmarshal(raw, &t); err != nil {
			return nil, err
		}
		return &t, nil
	case "mouse", "Mouse":
		var t MiceTrigger
		if err := json.Unmarshal(raw, &t); err != nil {
			return nil, err
		}
		return &t, nil
	default:
		return nil, fmt.Errorf("unknown trigger type: %s", meta.Type)
	}
}

type MiceTrigger struct {
	Button mice.BUTTON `json:"button"`
}

func (m *MiceTrigger) UnmarshalJSON(data []byte) error {
	type Alias MiceTrigger
	aux := &struct {
		*Alias
	}{
		Alias: (*Alias)(m),
	}
	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}
	if !mice.IsValidButton(string(m.Button)) {
		return fmt.Errorf("invalid mouse button: %s", m.Button)
	}
	return nil
}

func (m MiceTrigger) isTrigger(key keyboard.KEY, btn mice.BUTTON) bool {
	isTrigged := mice.IsButtonClicked(m.Button)
	return m.Button == btn || isTrigged
}

type KeyTrigger struct {
	Key keyboard.KEY `json:"key"`
}

func (k *KeyTrigger) UnmarshalJSON(data []byte) error {
	var raw struct {
		Key interface{} `json:"key"`
	}
	if err := json.Unmarshal(data, &raw); err != nil {
		return err
	}
	switch v := raw.Key.(type) {
	case string:
		if val, ok := keyboard.GetKeyFromName(v); ok {
			k.Key = val
		} else {
			return fmt.Errorf("unknown key name: %s", v)
		}
	default:
		return fmt.Errorf("invalid key type: must be a string name of the key")
	}
	return nil
}

func (k KeyTrigger) isTrigger(key keyboard.KEY, btn mice.BUTTON) bool {
	isTrigged := keyboard.IsKeyPressed(k.Key)
	return k.Key == key || isTrigged
}
