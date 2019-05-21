package car

import "strings"

type Car struct {
	Id           string `json:"id"`
	Licenseplate string `json:"licenseplate"`
	Finorvin     string `json:"finorvin"`
}

type CarDoors struct {
	Doorstatusfrontleft      Door `json:"doorstatusfrontleft"`
	Doorlockstatusfrontleft  Door `json:"doorlockstatusfrontleft"`
	Doorstatusfrontright     Door `json:"doorstatusfrontright"`
	Doorlockstatusfrontright Door `json:"doorlockstatusfrontright"`
	Doorstatusrearright      Door `json:"doorstatusrearright"`
	Doorlockstatusrearright  Door `json:"doorlockstatusrearright"`
	Doorstatusrearleft       Door `json:"doorstatusrearleft"`
	Doorlockstatusrearleft   Door `json:"doorlockstatusrearleft"`
	Doorlockstatusdecklid    Door `json:"doorlockstatusdecklid"`
	Doorlockstatusgas        Door `json:"doorlockstatusgas"`
	Doorlockstatusvehicle    Door `json:"doorlockstatusvehicle"`
}

type Door struct {
	Value           string `json:"value"`
	Retrievalstatus string `json:"retrievalstatus"`
	Timestamp       int64  `json:"timestamp"`
}


func (doors *CarDoors) GetToggleLocksCommand() string {
	if doors.IsLocked() {
		return `{"command" : "UNLOCK"}`
	} else {
		return `{"command" : "LOCK"}`
	}
}

func (doors *CarDoors) IsLocked() bool {
	return isLocked(doors.Doorlockstatusrearleft) &&
		isLocked(doors.Doorlockstatusrearright) &&
		isLocked(doors.Doorlockstatusfrontleft) &&
		isLocked(doors.Doorlockstatusfrontright)

}

func isLocked(door Door) bool {
	return strings.ToUpper(door.Value) == "LOCKED"
}
