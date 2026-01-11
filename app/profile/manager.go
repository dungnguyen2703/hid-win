package profile

import (
	"hidtool/app/event"
	"hidtool/app/keyboard"
	"hidtool/app/mice"
)

var profile1 = &ProfileImpl{
	ID:          "1",
	Name:        "Keyboard and Mice",
	Description: "Move window on F1, F2, Back, Forward",
	Bindings: []Binding{
		BindingImpl{
			DisabledLastestInput: true,
			Triggers: []Trigger{
				KeyTrigger{Key: keyboard.F1},
			},
			Actions: []event.Action{
				event.WindowLeft{},
			},
		},
		BindingImpl{
			DisabledLastestInput: true,
			Triggers: []Trigger{
				KeyTrigger{Key: keyboard.F2},
			},
			Actions: []event.Action{
				event.WindowRight{},
			},
		},
		BindingImpl{
			DisabledLastestInput: true,
			Triggers: []Trigger{
				MiceTrigger{Button: mice.BACK_BUTTON},
			},
			Actions: []event.Action{
				event.WindowLeft{},
			},
		},
		BindingImpl{
			DisabledLastestInput: true,
			Triggers: []Trigger{
				MiceTrigger{Button: mice.FORWARD_BUTTON},
			},
			Actions: []event.Action{
				event.WindowRight{},
			},
		},
	},
}

var profile2 = &ProfileImpl{
	ID:          "2",
	Name:        "Keyboard",
	Description: "Move window on F1, F2",
	Bindings: []Binding{
		BindingImpl{
			DisabledLastestInput: true,
			Triggers: []Trigger{
				KeyTrigger{Key: keyboard.F1},
			},
			Actions: []event.Action{
				event.WindowLeft{},
			},
		},
		BindingImpl{
			DisabledLastestInput: true,
			Triggers: []Trigger{
				KeyTrigger{Key: keyboard.F2},
			},
			Actions: []event.Action{
				event.WindowRight{},
			},
		},
	},
}

var profile3 = &ProfileImpl{
	ID:          "3",
	Name:        "Mice",
	Description: "Move window on Back, Forward",
	Bindings: []Binding{
		BindingImpl{
			DisabledLastestInput: true,
			Triggers: []Trigger{
				MiceTrigger{Button: mice.BACK_BUTTON},
			},
			Actions: []event.Action{
				event.WindowLeft{},
			},
		},
		BindingImpl{
			DisabledLastestInput: true,
			Triggers: []Trigger{
				MiceTrigger{Button: mice.FORWARD_BUTTON},
			},
			Actions: []event.Action{
				event.WindowRight{},
			},
		},
	},
}

var List = []Profile{
	profile1,
	profile2,
	profile3,
}

var currentProfile Profile = profile1

func SetCurrentProfile(profile Profile) {
	currentProfile = profile
}

func GetCurrentProfile() Profile {
	return currentProfile
}
