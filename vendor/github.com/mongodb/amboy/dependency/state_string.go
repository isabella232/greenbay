// Code generated by "stringer -type=State"; DO NOT EDIT

package dependency

import "fmt"

const stateName = "ReadyPassedBlockedUnresolved"

var stateIndex = [...]uint8{0, 5, 11, 18, 28}

func (i State) String() string {
	if i < 0 || i >= State(len(stateIndex)-1) {
		return fmt.Sprintf("State(%d)", i)
	}
	return stateName[stateIndex[i]:stateIndex[i+1]]
}
