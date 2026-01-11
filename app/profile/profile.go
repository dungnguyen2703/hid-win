package profile

import (
	"encoding/json"
	"fmt"
	"hidtool/app/keyboard"
	"hidtool/app/mice"
)

type Profile interface {
	GetID() string
	GetName() string
	GetDescription() string
	GetBinding(key keyboard.KEY, button mice.BUTTON) Binding
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

func (p *ProfileImpl) UnmarshalJSON(data []byte) error {
	type Alias ProfileImpl
	aux := &struct {
		Bindings []json.RawMessage `json:"bindings"`
		*Alias
	}{
		Alias: (*Alias)(p),
	}
	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}

	for _, raw := range aux.Bindings {
		var meta struct {
			Type string `json:"type"`
		}
		_ = json.Unmarshal(raw, &meta)

		var b Binding
		switch meta.Type {
		case "", "mapping":
			var bm BindingMapping
			if err := json.Unmarshal(raw, &bm); err != nil {
				return err
			}
			b = &bm
		default:
			return fmt.Errorf("unknown binding type: %s", meta.Type)
		}
		p.Bindings = append(p.Bindings, b)
	}
	return nil
}
