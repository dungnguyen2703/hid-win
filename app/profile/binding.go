package profile

import (
	"encoding/json"
	"hidtool/app/event"
	"hidtool/app/keyboard"
	"hidtool/app/mice"
)

type Binding interface {
	Action()
	DisableLatestInput() bool
	isTrigger(keyboard.KEY, mice.BUTTON) bool
}

type BindingMapping struct {
	Triggers             []Trigger      `json:"triggers"`                        // Conditions to activate the binding
	Actions              []event.Action `json:"actions"`                         // List of steps to execute
	DisabledLastestInput bool           `json:"disabled_latest_input,omitempty"` // Whether to disable the latest input event after action
}

func (b BindingMapping) Action() {
	for _, action := range b.Actions {
		action.Run()
	}
}

func (b BindingMapping) DisableLatestInput() bool {
	return b.DisabledLastestInput
}

func (b BindingMapping) isTrigger(key keyboard.KEY, button mice.BUTTON) bool {
	if len(b.Triggers) == 0 {
		return false
	}
	for _, trigger := range b.Triggers {
		if !trigger.isTrigger(key, button) {
			return false
		}
	}
	return true
}

func (b *BindingMapping) UnmarshalJSON(data []byte) error {
	type Alias BindingMapping
	aux := &struct {
		Triggers []json.RawMessage `json:"triggers"`
		Actions  []json.RawMessage `json:"actions"`
		*Alias
	}{
		Alias: (*Alias)(b),
	}
	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}

	for _, raw := range aux.Triggers {
		t, err := UnmarshalTrigger(raw)
		if err != nil {
			return err
		}
		b.Triggers = append(b.Triggers, t)
	}

	for _, raw := range aux.Actions {
		a, err := event.UnmarshalAction(raw)
		if err != nil {
			return err
		}
		b.Actions = append(b.Actions, a)
	}
	return nil
}
