package profile

import (
	"hidtool/app/event"
	"hidtool/app/keyboard"
	"hidtool/app/mice"
)

type Profile interface {
	GetID() string
	GetName() string
	GetDescription() string
	GetBinding(key keyboard.KEY, button mice.BUTTON) Binding
}

type Trigger interface {
	isTrigger(keyboard.KEY, mice.BUTTON) bool
}

type MiceTrigger struct {
	Button mice.BUTTON `json:"button"`
}

func (m MiceTrigger) isTrigger(key keyboard.KEY, btn mice.BUTTON) bool {
	isTrigged := mice.IsButtonClicked(m.Button)
	return m.Button == btn || isTrigged
}

type KeyTrigger struct {
	Key keyboard.KEY `json:"key"`
}

func (k KeyTrigger) isTrigger(key keyboard.KEY, btn mice.BUTTON) bool {
	isTrigged := keyboard.IsKeyPressed(k.Key)
	return k.Key == key || isTrigged
}

type Binding interface {
	Action()
	DisableLatestInput() bool
	isTrigger(keyboard.KEY, mice.BUTTON) bool
}

// Binding represents a mapping rule
type BindingImpl struct {
	Triggers             []Trigger      `json:"triggers"`                        // Conditions to activate the binding
	Actions              []event.Action `json:"actions"`                         // List of steps to execute
	DisabledLastestInput bool           `json:"disabled_latest_input,omitempty"` // Whether to disable the latest input event after action
}

func (b BindingImpl) Action() {
	for _, action := range b.Actions {
		action.Run()
	}
}

func (b BindingImpl) DisableLatestInput() bool {
	return b.DisabledLastestInput
}

func (b BindingImpl) isTrigger(key keyboard.KEY, button mice.BUTTON) bool {
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

// Profile contains a collection of Bindings
type ProfileImpl struct {
	ID          string    `json:"id,omitempty"`
	Name        string    `json:"name"`
	Description string    `json:"description,omitempty"`
	Bindings    []Binding `json:"bindings"`
}

func (m *ProfileImpl) GetID() string {
	return m.ID
}

func (m *ProfileImpl) GetName() string {
	return m.Name
}

func (m *ProfileImpl) GetDescription() string {
	return m.Description
}

// Action checks and performs actions based on the input key and button
// Returns isPerformed, isDisabledLatestAction
func (m *ProfileImpl) GetBinding(key keyboard.KEY, button mice.BUTTON) Binding {
	if key == 0 && button == "" {
		return nil
	}
	for _, binding := range m.Bindings {
		if binding.isTrigger(key, button) {
			return binding
		}
	}
	return nil
}
