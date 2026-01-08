package profile

import "hidtool/app/beep"

type Profile int

const (
	Off      Profile = 0
	All      Profile = 1
	Keyboard Profile = 2
	Mice     Profile = 3
)

var currentProfile Profile = All

func nextProfile(current Profile) Profile {
	switch current {
	case Off:
		return All
	case All:
		return Keyboard
	case Keyboard:
		return Mice
	case Mice:
		return Off
	}
	return Off
}

func prevProfile(current Profile) Profile {
	switch current {
	case Off:
		return Mice
	case All:
		return Off
	case Keyboard:
		return All
	case Mice:
		return Keyboard
	}
	return Off
}

func PlayProfileSound() {
	switch currentProfile {
	case Off:
		beep.Play(beep.LOW_BEEP, 1000)
	case All:
		beep.Play(beep.MEDIUM_BEEP, 500)
		beep.Play(beep.HIGH_BEEP, 500)
	case Keyboard:
		beep.Play(beep.MEDIUM_BEEP, 500)
	case Mice:
		beep.Play(beep.HIGH_BEEP, 500)
	}
}

func SetProfile(profile Profile) {
	currentProfile = profile
	// PlayProfileSound()
}

func NextProfile() {
	currentProfile = nextProfile(currentProfile)
	// PlayProfileSound()
}

func PrevProfile() {
	currentProfile = prevProfile(currentProfile)
	// PlayProfileSound()
}

func GetProfile() Profile {
	return currentProfile
}
