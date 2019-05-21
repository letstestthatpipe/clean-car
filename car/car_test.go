package car

import "testing"

func TestCarDoors_IsLocked_shouldBeLocked(t *testing.T) {
	dooors := getLockedDoors()

	if !dooors.IsLocked() {
		t.Errorf("doors should be LOCKED, got: unlocked (false)")
	}
}

func TestCarDoors_GetToggleLocksCommand_Locked(t *testing.T) {
	dooors := getLockedDoors()

	command := dooors.GetToggleLocksCommand()
	expected := `{"command" : "UNLOCK"}`
	if command != expected {
		t.Errorf("Command generated was %s but should be %s", command, expected )
	}
}

func getLockedDoors() CarDoors {
	dooors := CarDoors{
		Doorlockstatusfrontright: Door{
			Value: "LOCKED",
		},
		Doorlockstatusfrontleft: Door{
			Value: "LOCKED",
		},
		Doorlockstatusrearright: Door{
			Value: "LOCKED",
		},
		Doorlockstatusrearleft: Door{
			Value: "LOCKED",
		},
	}
	return dooors
}

func TestCarDoors_IsLocked_shouldBeUnlocked(t *testing.T) {
	dooors := getUnlockedDoors()

	if dooors.IsLocked() {
		t.Errorf("doors should be UNLOCKED, got locked (true)")
	}
}

func TestCarDoors_GetToggleLocksCommand_Unlockd(t *testing.T) {
	dooors := getUnlockedDoors()

	command := dooors.GetToggleLocksCommand()
	expected := `{"command" : "LOCK"}`
	if command != expected {
		t.Errorf("Command generated was %s but should be %s", command, expected )
	}
}

func getUnlockedDoors() CarDoors {
	dooors := CarDoors{
		Doorlockstatusfrontright: Door{
			Value: "UNLOCKED",
		},
		Doorlockstatusfrontleft: Door{
			Value: "LOCKED",
		},
		Doorlockstatusrearright: Door{
			Value: "LOCKED",
		},
		Doorlockstatusrearleft: Door{
			Value: "LOCKED",
		},
	}
	return dooors
}

